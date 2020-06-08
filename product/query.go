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

func createProductItem(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DBConn.ExecContext(ctx, `INSERT INTO products 
		(name, price, created_at, modified_at) VALUES (?, ?, ?, ?)`,
		product.Name, product.Price, product.CreatedAt, product.ModifiedAt,
	)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertID), nil
}
