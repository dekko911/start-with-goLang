package order

import (
	"database/sql"

	"github.com/dekko911/start-with-goLang/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(o types.Order) (int, error) {
	order := []any{
		o.UserID,
		o.Total,
		o.Status,
		o.Address,
	}

	req, err := s.db.Exec("INSERT INTO orders (userID, total, status, address) VALUES (?,?,?,?)", order...)
	if err != nil {
		return 0, nil
	}

	id, err := req.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(o types.OrderItem) error {
	orderItem := []any{
		o.OrderID,
		o.ProductID,
		o.Quantity,
		o.Price,
	}

	_, err := s.db.Exec("INSERT INTO order_items", orderItem...)
	if err != nil {
		return err
	}

	return nil
}
