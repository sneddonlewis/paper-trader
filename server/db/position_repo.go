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

func (repo *PositionRepo) GetPositions() ([]*model.Position, error) {
	rows, err := repo.db.Query("select ticker, price, quantity from positions")
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.Ticker, &position.Price, &position.Quantity)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	log.Println(positions)
	return positions, nil
}

func (repo *PositionRepo) GetPositionsByTicker(ticker string) ([]*model.Position, error) {
	rows, err := repo.db.Query("SELECT ticker, price, quantity FROM positions WHERE ticker = ?", ticker)
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.Ticker, &position.Price, &position.Quantity)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	log.Println(positions)
	return positions, nil
}

func (repo *PositionRepo) OpenPosition(p *model.Position) (*model.Position, error) {
	if p.Quantity == 0.0 {
		return nil, errors.New("cannot open position of quantity 0.0")
	}
	_, err := repo.db.Exec("INSERT INTO positions VALUES (?, ?, ?)", p.Ticker, p.Price, p.Quantity)
	if err != nil {
		return nil, err
	}
	return p, nil
}
