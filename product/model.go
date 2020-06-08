package product

// Product :
type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CreatedAt  string  `json:"created_at"`
	ModifiedAt string  `json:"modified_at"`
	IsDeleted  bool    `json:"is_deleted"`
	DeletedAt  string  `json:"deleted_at"`
}
