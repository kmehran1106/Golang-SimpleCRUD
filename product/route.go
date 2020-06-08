package product

import (
	"fmt"
	"net/http"
	"simple-crud/cors"
)

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	productListCreateHandler := http.HandlerFunc(handlerProductListCreate)
	productDetailUpdateDeleteHandler := http.HandlerFunc(handleProductDetailUpdateDelete)

	http.Handle(
		fmt.Sprintf("%s/%s", apiBasePath, productsPath),
		cors.Middleware(productListCreateHandler),
	)
	http.Handle(
		fmt.Sprintf("%s/%s/", apiBasePath, productsPath),
		cors.Middleware(productDetailUpdateDeleteHandler),
	)
}
