package models

type Child struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"` // "YYYY-MM-DD"
	Notes     string `json:"notes"`
}
