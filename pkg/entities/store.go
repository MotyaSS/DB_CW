package entity

type Store struct {
	StoreId      int    `json:"store_id" db:"store_id"`
	StoreName    string `json:"store_name" db:"store_name"`
	StoreAddress string `json:"store_address" db:"store_address"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
}
