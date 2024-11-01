package models

type Athlete struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
	SportType string `json:"sport_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

type CreateAthleteRequest struct {
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
	SportType string `json:"sport_type"`
}

type GetAthleteRequest struct {
	ID string `json:"id"`
}

type GetAthleteResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CountryName string `json:"country_name"`
	CountryID   string `json:"country_id"`
	SportType   string `json:"sport_type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

type ListOfAthleteRequest struct{}

type ListOfAthleteResponse struct {
	Athletes []GetAthleteResponse `json:"athletes"`
}

type UpdateAthleteRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
	SportType string `json:"sport_type"`
}

type DeleteAthleteRequest struct {
	ID string `json:"id"`
}

type DeleteAthleteResponse struct {
	Status string `json:"status"`
}
