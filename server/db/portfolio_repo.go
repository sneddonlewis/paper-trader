package db

import (
	"database/sql"
	"paper-trader/model"
	"time"
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
	p := new(model.Portfolio)

	query := `
        SELECT 
            p.id AS portfolio_id, 
            p.user_id, 
            p.name,
            pos.id AS position_id,
            pos.ticker,
            pos.price,
            pos.quantity,
            pos.opened_at,
            pos.close_price,
            pos.closed_at,
            pos.profit
        FROM 
            portfolios p
        LEFT JOIN 
            positions pos ON p.id = pos.portfolio_id
        WHERE 
            p.id = $1
    `

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var portfolioID, positionID int32
		var ticker string
		var price, quantity float64
		var closePrice, profit sql.NullFloat64
		var openedAt, closedAt *time.Time

		err := rows.Scan(
			&portfolioID,
			&p.UserID,
			&p.Name,
			&positionID,
			&ticker,
			&price,
			&quantity,
			&openedAt,
			&closePrice,
			&closedAt,
			&profit,
		)
		if err != nil {
			return nil, err
		}

		if closedAt != nil {
			closedPosition := &model.ClosedPosition{
				ID:         positionID,
				Ticker:     ticker,
				Price:      price,
				Quantity:   quantity,
				OpenedAt:   openedAt,
				ClosePrice: closePrice,
				ClosedAt:   closedAt,
				Profit:     profit,
			}
			p.ClosedPositions = append(p.ClosedPositions, closedPosition)
		} else {
			openPosition := &model.Position{
				ID:       positionID,
				Ticker:   ticker,
				Price:    price,
				Quantity: quantity,
				OpenedAt: openedAt,
			}
			p.OpenPositions = append(p.OpenPositions, openPosition)
		}

	}

	return p, nil
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
