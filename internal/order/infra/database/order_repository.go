package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wenceslau/sysm-product-calculator/internal/order/entity"
)

type OriderRepository struct {
	Db *sql.DB
}

// GetTotal implements entity.OrderRepositoryInterface.
func (r *OriderRepository) GetTotal() (int, error) {
	panic("unimplemented")
}

func NewOrderRepository(db *sql.DB) *OriderRepository {
	return &OriderRepository{Db: db}
}

func (r *OriderRepository) Save(order entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
