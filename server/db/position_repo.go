package db

import (
	"database/sql"
	"errors"
	"log"
	"paper-trader/model"
	"strings"
)

type PositionRepo struct {
	db *sql.DB
}

func NewPositionRepo(db *sql.DB) PositionRepo {
	return PositionRepo{db: db}
}

func (r *PositionRepo) GetOpenPositions() ([]*model.Position, error) {
	rows, err := r.db.Query("SELECT id, portfolio_id, ticker, price, quantity, opened_at FROM positions WHERE close_price IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.ID, &position.PortfolioID, &position.Ticker, &position.Price, &position.Quantity, &position.OpenedAt)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func (r *PositionRepo) GetClosedPositions(portfolioId int32) ([]*model.ClosedPosition, error) {
	rows, err := r.db.Query("SELECT id, portfolio_id, ticker, price, quantity, opened_at, close_price, closed_at, profit FROM positions WHERE close_price IS NOT NULL AND portfolio_id = $1",
		portfolioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.ClosedPosition
	for rows.Next() {
		position := new(model.ClosedPosition)
		err = rows.Scan(&position.ID,
			&position.PortfolioID,
			&position.Ticker,
			&position.Price,
			&position.Quantity,
			&position.OpenedAt,
			&position.ClosePrice,
			&position.ClosedAt,
			&position.Profit,
		)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func (r *PositionRepo) GetPositions() ([]*model.Position, error) {
	query := `
select id, portfolio_id, ticker, price, quantity, opened_at from positions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.ID,
			&position.PortfolioID,
			&position.Ticker, &position.Price, &position.Quantity, &position.OpenedAt)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func (r *PositionRepo) GetPositionById(id int32) (*model.Position, error) {
	row := r.db.QueryRow("select id, portfolio_id, ticker, price, quantity, opened_at from positions WHERE id = $1", id)
	position := new(model.Position)
	err := row.Scan(&position.ID,
		&position.PortfolioID,
		&position.Ticker, &position.Price, &position.Quantity, &position.OpenedAt)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (r *PositionRepo) GetPositionsByTicker(ticker string) ([]*model.Position, error) {
	rows, err := r.db.Query("SELECT id, portfolio_id, ticker, price, quantity FROM positions WHERE ticker = ?", ticker)
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.ID,
			&position.PortfolioID,
			&position.Ticker, &position.Price, &position.Quantity)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	log.Println(positions)
	return positions, nil
}

func (r *PositionRepo) OpenPosition(p *model.Position) (*model.Position, error) {
	if p.Quantity == 0.0 {
		return nil, errors.New("cannot open position of quantity 0.0")
	}
	p.Ticker = strings.ToUpper(p.Ticker)
	row := r.db.QueryRow(
		"INSERT INTO positions (ticker, price, quantity) VALUES ($1, $2, $3) RETURNING id, opened_at",
		p.Ticker,
		p.Price,
		p.Quantity,
	)
	err := row.Scan(&p.ID, &p.OpenedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PositionRepo) ClosePosition(id int32, closePrice float64) (*model.ClosedPosition, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	closePositionQuery := `
		UPDATE positions
		SET close_price = $1, profit = ($1 - price) * quantity, closed_at = CURRENT_TIMESTAMP
		WHERE id = $2
		RETURNING id, portfolio_id, ticker, price, quantity, opened_at, close_price, closed_at, profit`

	updatePorfolioAmountQuery := `
		UPDATE money
    	SET amount = amount + $1 
    	WHERE portfolio_id = $2`

	positionRow := tx.QueryRow(
		closePositionQuery,
		closePrice,
		id,
	)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	closedPosition := new(model.ClosedPosition)
	err = positionRow.Scan(
		&closedPosition.ID,
		&closedPosition.PortfolioID,
		&closedPosition.Ticker,
		&closedPosition.Price,
		&closedPosition.Quantity,
		&closedPosition.OpenedAt,
		&closedPosition.ClosePrice,
		&closedPosition.ClosedAt,
		&closedPosition.Profit,
	)
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(
		updatePorfolioAmountQuery,
		closedPosition.Profit,
		closedPosition.PortfolioID,
	)
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	return closedPosition, nil
}
