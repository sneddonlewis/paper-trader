package db

import (
	"database/sql"
	"paper-trader/model"
)

type PortfolioRepo struct {
	db *sql.DB
}

func NewPortfolioRepo(db *sql.DB) PortfolioRepo {
	return PortfolioRepo{db: db}
}

func (r *PortfolioRepo) CreatePortfolio(userId int32, name string) (*model.Portfolio, error) {
	portfolio := new(model.Portfolio)
	err := r.db.QueryRow("INSERT INTO portfolios (user_id, name) VALUES ($1, $2) RETURNING id", userId, name).Scan(&portfolio.ID)
	if err != nil {
		return nil, err
	}
	portfolio.UserID = userId
	portfolio.Name = name
	return portfolio, nil
}

func (r *PortfolioRepo) GetPortfolioById(id int32) (*model.Portfolio, error) {
	portfolio := new(model.Portfolio)
	err := r.db.QueryRow("SELECT id, user_id, name FROM portfolios WHERE id = $1", id).Scan(&portfolio.ID, &portfolio.UserID, &portfolio.Name)
	if err != nil {
		return nil, err
	}
	return portfolio, nil
}

func (r *PortfolioRepo) GetPortfoliosByUserId(userId int32) ([]*model.Portfolio, error) {
	rows, err := r.db.Query("SELECT id, name FROM portfolios WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var portfolios []*model.Portfolio
	for rows.Next() {
		portfolio := new(model.Portfolio)
		err = rows.Scan(&portfolio.ID, &portfolio.Name)
		if err != nil {
			return nil, err
		}
		portfolio.UserID = userId
		portfolios = append(portfolios, portfolio)
	}
	return portfolios, nil
}

func (r *PortfolioRepo) UpdatePortfolio(portfolio *model.Portfolio) error {
	_, err := r.db.Exec("UPDATE portfolios SET name = $1 WHERE id = $2", portfolio.Name, portfolio.ID)
	return err
}

func (r *PortfolioRepo) DeletePortfolio(id int32) error {
	_, err := r.db.Exec("DELETE FROM portfolios WHERE id = $1", id)
	return err
}
