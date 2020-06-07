package product

import (
	"fmt"
	"net/http"
	"simple-crud/cors"
)

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	productsHandler := http.HandlerFunc(HandlerProductListCreate)
	productHandler := http.HandlerFunc(HandleProductDetailUpdateDelete)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsPath), cors.Middleware(productsHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsPath), cors.Middleware(productHandler))
}
