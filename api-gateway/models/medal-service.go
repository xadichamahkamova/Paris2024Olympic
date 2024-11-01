package models

type MedalType int32

const (
	GOLD   MedalType = 0
	SILVER MedalType = 1
	BRONZE MedalType = 2
)

type Medal struct {
	ID        string    `json:"id"`
	CountryID string    `json:"country_id"`
	Type      MedalType `json:"type"`
	EventID   string    `json:"event_id"`
	AthleteID string    `json:"athlete_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt int64     `json:"deleted_at"`
}

type CreateMedalRequest struct {
	CountryID string    `json:"country_id"`
	Type      MedalType `json:"type"`
	EventID   string    `json:"event_id"`
	AthleteID string    `json:"athlete_id"`
}

type CreateMedalResponse struct {
	ID        string    `json:"id"`
	CountryID string    `json:"country_id"`
	Type      MedalType `json:"type"`
	EventID   string    `json:"event_id"`
	AthleteID string    `json:"athlete_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt int64     `json:"deleted_at"`
}

type GetMedalByIdRequest struct {
	ID string `json:"id"`
}

type GetMedalByIdResponse struct {
	ID        string    `json:"id"`
	CountryID string    `json:"country_id"`
	Type      MedalType `json:"type"`
	EventID   string    `json:"event_id"`
	AthleteID string    `json:"athlete_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt int64     `json:"deleted_at"`
}

type GetMedalsResponse struct {
	Medals []Medal `json:"medals"`
}

type GetMedalByFilterRequest struct {
	CountryID string    `json:"country_id,omitempty"`
	Type      MedalType `json:"type,omitempty"`
	EventID   string    `json:"event_id,omitempty"`
	AthleteID string    `json:"athlete_id,omitempty"`
}

type GetMedalByFilterResponse struct {
	Medals []Medal `json:"medals"`
}

type UpdateMedalRequest struct {
	ID        string    `json:"id"`
	CountryID string    `json:"country_id"`
	Type      MedalType `json:"type"`
	EventID   string    `json:"event_id"`
	AthleteID string    `json:"athlete_id"`
}

type UpdateMedalResponse struct {
	ID        string    `json:"id"`
	CountryID string    `json:"country_id"`
	Type      MedalType `json:"type"`
	EventID   string    `json:"event_id"`
	AthleteID string    `json:"athlete_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt int64     `json:"deleted_at"`
}

type DeleteMedalRequest struct {
	ID string `json:"id"`
}

type DeleteMedalResponse struct {
	Success bool `json:"success"`
}

type VoidMedal struct{}
