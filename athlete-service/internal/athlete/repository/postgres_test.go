package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
)

func setupTestDB(t *testing.T) (AthleteRepository, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	repo := NewPostgresAthleteRepository(db)
	return repo, mock
}

func TestCreateAthlete(t *testing.T) {
	repo, mock := setupTestDB(t)

	req := &pb.CreateAthleteRequest{
		Name:      "AthleteName",
		CountryId: "1",
		SportType: "SportType",
	}

	// Mock the country check
	mock.ExpectQuery(`SELECT id, name FROM countries WHERE id=\$1 AND deleted_at=0`).
		WithArgs(req.CountryId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(req.CountryId, "CountryName"))

	// Mock the athlete creation
	rows := sqlmock.NewRows([]string{"id", "name", "country_id", "sport_type", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, req.Name, req.CountryId, req.SportType, "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0)

	mock.ExpectQuery(`INSERT INTO athletes\(name, country_id, sport_type\) VALUES\(\$1, \$2, \$3\) RETURNING id, name, country_id, sport_type, created_at, updated_at, deleted_at`).
		WithArgs(req.Name, req.CountryId, req.SportType).
		WillReturnRows(rows)

	athlete, err := repo.CreateAthlete(req)
	assert.NoError(t, err)
	assert.Equal(t, "1", athlete.Id)
	assert.Equal(t, req.Name, athlete.Name)
	assert.Equal(t, req.CountryId, athlete.CountryId)
	assert.Equal(t, req.SportType, athlete.SportType)
}

func TestGetAthlete(t *testing.T) {
	repo, mock := setupTestDB(t)

	req := &pb.GetAthleteRequest{Id: "1"}
	rows := sqlmock.NewRows([]string{"id", "name", "country_name", "country_id", "sport_type", "created_at", "updated_at", "deleted_at"}).
		AddRow(req.Id, "AthleteName", "CountryName", 1, "SportType", "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0)

	mock.ExpectQuery(`SELECT a.id, a.name, c.name AS country_name, a.country_id, a.sport_type, a.created_at, a.updated_at, a.deleted_at FROM athletes AS a INNER JOIN countries AS c ON a.country_id = c.id WHERE a.id=\$1 AND a.deleted_at=0 AND c.deleted_at=0`).
		WithArgs(req.Id).
		WillReturnRows(rows)

	athlete, err := repo.GetAthlete(req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, athlete.Id)
	assert.Equal(t, "AthleteName", athlete.Name)
	assert.Equal(t, "CountryName", athlete.CountryName)
	assert.Equal(t, "1", athlete.CountryId)
	assert.Equal(t, "SportType", athlete.SportType)
}

func TestListAthletes(t *testing.T) {
	repo, mock := setupTestDB(t)

	rows := sqlmock.NewRows([]string{"id", "name", "country_name", "country_id", "sport_type", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Athlete1", "Country1", 1, "SportType1", "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0).
		AddRow(2, "Athlete2", "Country2", 2, "SportType2", "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0)

	mock.ExpectQuery(`SELECT a.id, a.name, c.name AS country_name, a.country_id, a.sport_type, a.created_at, a.updated_at, a.deleted_at FROM athletes AS a INNER JOIN countries AS c ON a.country_id = c.id WHERE a.deleted_at=0 AND c.deleted_at=0`).
		WillReturnRows(rows)

	resp, err := repo.ListAthletes(&pb.ListOfAthleteRequest{})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(resp.Athletes))
	assert.Equal(t, "1", resp.Athletes[0].Id)
	assert.Equal(t, "Athlete1", resp.Athletes[0].Name)
	assert.Equal(t, "Country1", resp.Athletes[0].CountryName)
	assert.Equal(t, "SportType1", resp.Athletes[0].SportType)
}

func TestUpdateAthlete(t *testing.T) {
	repo, mock := setupTestDB(t)

	req := &pb.UpdateAthleteRequest{
		Id:        "1",
		Name:      "UpdatedAthleteName",
		CountryId: "2",
		SportType: "UpdatedSportType",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "country_id", "sport_type", "created_at", "updated_at", "deleted_at"}).
		AddRow(req.Id, req.Name, req.CountryId, req.SportType, "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0)

	mock.ExpectQuery(`UPDATE athletes SET name=\$1, country_id=\$2, sport_type=\$3, updated_at=NOW\(\) WHERE id=\$4 AND deleted_at=0 RETURNING id, name, country_id, sport_type, created_at, updated_at, deleted_at`).
		WithArgs(req.Name, req.CountryId, req.SportType, req.Id).
		WillReturnRows(rows)

	athlete, err := repo.UpdateAthlete(req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, athlete.Id)
	assert.Equal(t, req.Name, athlete.Name)
	assert.Equal(t, req.CountryId, athlete.CountryId)
	assert.Equal(t, req.SportType, athlete.SportType)
}

func TestDeleteAthlete(t *testing.T) {
	repo, mock := setupTestDB(t)

	req := &pb.DeleteAthleteRequest{Id: "1"}
	mock.ExpectExec(`UPDATE athletes SET deleted_at=DATE_PART\('epoch', CURRENT_TIMESTAMP\)::INT WHERE id=\$1`).
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.DeleteAthlete(req)
	assert.NoError(t, err)
	assert.Equal(t, "deleted successfully", resp.Status)
}
