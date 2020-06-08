package product

import (
	"encoding/json"
	"log"
	"net/http"
)

const productsPath = "products"

func handlerProductListCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList, err := getProductList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(productList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("In Detail"))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleProductDetailUpdateDelete :
func handleProductDetailUpdateDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In Detail"))
	return
}
