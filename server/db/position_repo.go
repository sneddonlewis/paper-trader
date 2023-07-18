package db

import (
	"database/sql"
	"paper-trader/model"
)

type PositionRepo struct {
	db *sql.DB
}

func NewPositionRepo(db *sql.DB) PositionRepo {
	return PositionRepo{db: db}
}

func (repo *PositionRepo) GetPositions() ([]*model.Position, error) {
	rows, err := repo.db.Query("select ticker, direction, price, quantity from positions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var positions []*model.Position
	for rows.Next() {
		position := new(model.Position)
		err = rows.Scan(&position.Ticker, &position.Direction, &position.Price, &position.Quantity)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}
