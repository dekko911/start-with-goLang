package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/dekko911/start-with-goLang/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (s *Store) GetProductsByID(ids []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(ids)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	args := make([]any, len(ids))
	copy(args, args) // <- IMPORTANT YOU KNOW !!!!!!!!!!!!

	// if error, use this down below
	// for i, v := range args {
	// 	args[i] = v
	// }

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (s *Store) CreateProduct(p types.CreateProductPayload) error {
	product := []any{
		p.Name,
		p.Price,
		p.Image,
		p.Description,
		p.Quantity,
	}
	_, err := s.db.Exec("INSERT INTO products (name, price, image, description, quantity) VALUES (?,?,?,?,?)", product...)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateProduct(p types.Product) error {
	product := []any{
		p.Name,
		p.Price,
		p.Image,
		p.Description,
		p.Quantity,
		p.ID,
	}
	_, err := s.db.Exec("UPDATE products SET name = ?, price = ?, image = ?, description = ?, quantity = ? WHERE id = ?", product...)
	if err != nil {
		return err
	}

	return nil
}
