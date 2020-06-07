package product

import (
	// "encoding/json"
	// "fmt"
	// "log"
	"net/http"
	// "strconv"
	// "strings"
)

const productsPath = "products"

// HandlerProductListCreate :
func HandlerProductListCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("In List"))
	return
}

// HandleProductDetailUpdateDelete :
func HandleProductDetailUpdateDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In Detail"))
	return
}