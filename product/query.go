package product

import (
	"context"
	"log"
	"time"

	"simple-crud/database"
)

func getProductList() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, (price/100) as price, created_at, modified_at, is_deleted, deleted_at FROM products`
	results, err := database.DBConn.QueryContext(ctx, query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()

	products := make([]Product, 0)
	for results.Next() {
		var p Product
		results.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt, &p.ModifiedAt, &p.IsDeleted, &p.DeletedAt)
		products = append(products, p)
	}
	return products, nil
}
