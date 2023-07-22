package db

import (
	"database/sql"
	"errors"
	"log"
	"paper-trader/model"
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
	rows, err := r.db.Query("select id, ticker, price, quantity from positions")
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

func (r *PositionRepo) OpenPosition(request *model.Position) (*model.Position, error) {
	if request.Quantity == 0.0 {
		return nil, errors.New("cannot open position of quantity 0.0")
	}
	row := r.db.QueryRow("INSERT INTO positions VALUES (?, ?, ?) RETURNING id, ticker, price, quantity", p.Ticker, p.Price, p.Quantity)
	p := new(model.Position)
	err := row.Scan(&p.ID, &p.Ticker, &p.Price, &p.Quantity)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PositionRepo) ClosePosition(id int32, closePrice float64) (*model.ClosedPosition, error) {
	row := r.db.QueryRow("UPDATE positions SET close_price = $1 WHERE id = $2 RETURNING id, ticker, price, quantity, close_price", id, closePrice)
	closedPosition := new(model.ClosedPosition)
	err := row.Scan(&closedPosition.ID, &closedPosition.Ticker, &closedPosition.Price, &closedPosition.Quantity, &closedPosition.ClosePrice)
	if err != nil {
		return nil, err
	}
	return closedPosition, nil
}
