package models

type Event struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SportType string `json:"sport_type"`
	Location  string `json:"location"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at,omitempty"` // Optional field
}

type CreateEventRequest struct {
	Name      string `json:"name"`
	SportType string `json:"sport_type"`
	Location  string `json:"location"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type GetEventRequest struct {
	ID string `json:"id"`
}

type ListOfEventRequest struct{}

type ListOfEventResponse struct {
	Events []Event `json:"events"`
}

type UpdateEventRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SportType string `json:"sport_type"`
	Location  string `json:"location"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type DeleteEventRequest struct {
	ID string `json:"id"`
}

type DeleteEventResponse struct {
	Status string `json:"status"`
}
