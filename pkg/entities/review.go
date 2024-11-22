package entity

type Review struct {
	ReviewId   int    `json:"review_id"`
	RentalId   int    `json:"rental_id"`
	ReviewText string `json:"review_text"`
	Rating     int    `json:"rating"`
}
