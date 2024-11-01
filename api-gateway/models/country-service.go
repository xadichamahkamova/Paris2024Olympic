package models

type Country struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Flag      string `json:"flag"`
	Region    string `json:"region"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at,omitempty"`
}

type CreateCountryRequest struct {
	Name   string `json:"name"`
	Flag   string `json:"flag"`
	Region string `json:"region"`
}

type GetCountryRequest struct {
	ID string `json:"id"`
}

type ListOfCountryRequest struct{}

type ListOfCountryResponse struct {
	Countries []Country `json:"countries"`
}

type UpdateCountryRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Flag   string `json:"flag"`
	Region string `json:"region"`
}

type DeleteCountryRequest struct {
	ID string `json:"id"`
}

type DeleteCountryResponse struct {
	Status string `json:"status"`
}
