package entity

type Store struct {
	StoreId     int    `json:"store_id"`
	StoreName   string `json:"store_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
