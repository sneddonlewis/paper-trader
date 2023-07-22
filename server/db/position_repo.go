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
	rows, err := r.db.Query("SELECT id, ticker, price, quantity FROM positions WHERE close_price IS NULL")
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.ID, &position.Ticker, &position.Price, &position.Quantity)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	log.Println(positions)
	return positions, nil
}

func (r *PositionRepo) GetPositions() ([]*model.Position, error) {
	rows, err := r.db.Query("select id, ticker, price, quantity, opened_at from positions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.ID, &position.Ticker, &position.Price, &position.Quantity, &position.OpenedAt)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func (r *PositionRepo) GetPositionById(id int32) (*model.Position, error) {
	row := r.db.QueryRow("select id, ticker, price, quantity from positions WHERE id = $1", id)
	position := new(model.Position)
	err := row.Scan(&position.ID, &position.Ticker, &position.Price, &position.Quantity)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (r *PositionRepo) GetPositionsByTicker(ticker string) ([]*model.Position, error) {
	rows, err := r.db.Query("SELECT id, ticker, price, quantity FROM positions WHERE ticker = ?", ticker)
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.ID, &position.Ticker, &position.Price, &position.Quantity)
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
	row := r.db.QueryRow(
		"UPDATE positions SET close_price = $1 WHERE id = $2 RETURNING id, ticker, price, quantity, close_price",
		closePrice,
		id,
	)
	closedPosition := new(model.ClosedPosition)
	err := row.Scan(
		&closedPosition.ID,
		&closedPosition.Ticker,
		&closedPosition.Price,
		&closedPosition.Quantity,
		&closedPosition.ClosePrice,
	)
	if err != nil {
		return nil, err
	}
	return closedPosition, nil
}
