package entity

type Review struct {
	ReviewId   int    `json:"review_id" db:"review_id"`
	RentalId   int    `json:"rental_id" db:"rental_id"`
	ReviewText string `json:"review_text" db:"review_text"`
	Rating     int    `json:"rating" db:"rating"`
}
