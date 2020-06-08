package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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
		var product Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		product.CreatedAt = currentTime
		product.ModifiedAt = currentTime

		productID, err := createProductItem(product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"productId":%d}`, productID)))
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
