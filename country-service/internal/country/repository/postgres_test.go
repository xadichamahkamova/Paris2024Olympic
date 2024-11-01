package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
)

// Helper function to set up the test database and repository
func setupTestDB(t *testing.T) (CountryRepository, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	repo := NewPostgresCountryRepository(db)
	return repo, mock
}

func TestCreateCountry(t *testing.T) {
	repo, mock := setupTestDB(t)

	req := &pb.CreateCountryRequest{
		Name:   "CountryName",
		Flag:   "FlagURL",
		Region: "RegionName",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "flag", "region", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, req.Name, req.Flag, req.Region, "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0)

	mock.ExpectQuery(`INSERT INTO countries\(name, flag, region\) VALUES\(\$1, \$2, \$3\) RETURNING id, name, flag, region, created_at, updated_at, deleted_at`).
		WithArgs(req.Name, req.Flag, req.Region).
		WillReturnRows(rows)

	country, err := repo.CreateCountry(req)
	assert.NoError(t, err)
	assert.Equal(t, "1", country.Id)
	assert.Equal(t, req.Name, country.Name)
	assert.Equal(t, req.Flag, country.Flag)
	assert.Equal(t, req.Region, country.Region)
}

func TestGetCountry(t *testing.T) {
	repo, mock := setupTestDB(t)

	req := &pb.GetCountryRequest{Id: "1"}
	rows := sqlmock.NewRows([]string{"id", "name", "flag", "region", "created_at", "updated_at", "deleted_at"}).
		AddRow(req.Id, "CountryName", "FlagURL", "RegionName", "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0)

	mock.ExpectQuery(`SELECT id, name, flag, region, created_at, updated_at, deleted_at FROM countries WHERE id=\$1 AND deleted_at=0`).
		WithArgs(req.Id).
		WillReturnRows(rows)

	country, err := repo.GetCountry(req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, country.Id)
	assert.Equal(t, "CountryName", country.Name)
	assert.Equal(t, "FlagURL", country.Flag)
	assert.Equal(t, "RegionName", country.Region)
}

func TestListOfCountry(t *testing.T) {
	repo, mock := setupTestDB(t)

	rows := sqlmock.NewRows([]string{"id", "name", "flag", "region", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Country1", "FlagURL1", "Region1", "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0).
		AddRow(2, "Country2", "FlagURL2", "Region2", "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0)

	mock.ExpectQuery(`SELECT id, name, flag, region, created_at, updated_at, deleted_at FROM countries WHERE deleted_at=0`).
		WillReturnRows(rows)

	resp, err := repo.ListOfCountry(&pb.ListOfCountryRequest{})
	assert.NoError(t, err)
	assert.Equal(t, int(2), len(resp.Countries))
	assert.Equal(t, "1", resp.Countries[0].Id)
	assert.Equal(t, "Country1", resp.Countries[0].Name)
	assert.Equal(t, "FlagURL1", resp.Countries[0].Flag)
	assert.Equal(t, "Region1", resp.Countries[0].Region)
}

func TestUpdateCountry(t *testing.T) {
	repo, mock := setupTestDB(t)

	req := &pb.UpdateCountryRequest{
		Id:     "1",
		Name:   "UpdatedCountryName",
		Flag:   "UpdatedFlagURL",
		Region: "UpdatedRegionName",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "flag", "region", "created_at", "updated_at", "deleted_at"}).
		AddRow(req.Id, req.Name, req.Flag, req.Region, "2024-08-07T00:00:00Z", "2024-08-07T00:00:00Z", 0)

	mock.ExpectQuery(`UPDATE countries SET name=\$1, flag=\$2, region=\$3, updated_at=NOW\(\) WHERE id=\$4 AND deleted_at=0 RETURNING id, name, flag, region, created_at, updated_at, deleted_at`).
		WithArgs(req.Name, req.Flag, req.Region, req.Id).
		WillReturnRows(rows)

	country, err := repo.UpdateCountry(req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, country.Id)
	assert.Equal(t, req.Name, country.Name)
	assert.Equal(t, req.Flag, country.Flag)
	assert.Equal(t, req.Region, country.Region)
}

func TestDeleteCountry(t *testing.T) {
	repo, mock := setupTestDB(t)

	req := &pb.DeleteCountryRequest{Id: "1"}
	mock.ExpectExec(`UPDATE countries SET deleted_at=DATE_PART\('epoch', CURRENT_TIMESTAMP\)::INT WHERE id=\$1`).
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.DeleteCountry(req)
	assert.NoError(t, err)
	assert.Equal(t, "deleted successfully", resp.Status)
}
