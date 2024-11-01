package repository

import (
	"testing"
	"time"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/medalspb"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateMedal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPostgresMedalRepo(db)

	mock.ExpectQuery("SELECT 1 FROM countries").WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	mock.ExpectQuery("SELECT 1 FROM events").WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	mock.ExpectQuery("SELECT 1 FROM athletes").WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	mock.ExpectQuery("INSERT INTO medals").WithArgs("1", "GOLD", "1", "1").WillReturnRows(sqlmock.NewRows([]string{"id", "country_id", "type", "event_id", "athlete_id", "created_at", "updated_at", "deleted_at"}).AddRow(1, "1", "GOLD", "1", "1", time.Now(), time.Now(), nil))

	req := &pb.CreateMedalRequest{
		CountryId: "1",
		Type:      1,
		EventId:   "1",
		AthleteId: "1",
	}

	resp, err := repo.CreateMedal(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "1", resp.CountryId)
	assert.Equal(t, "GOLD", resp.Type)
}

func TestUpdateMedal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPostgresMedalRepo(db)

	mock.ExpectQuery("UPDATE medals").WithArgs("1", "SILVER", "1", "1", sqlmock.AnyArg(), "1").WillReturnRows(sqlmock.NewRows([]string{"id", "country_id", "type", "event_id", "athlete_id", "created_at", "updated_at", "deleted_at"}).AddRow(1, "1", "SILVER", "1", "1", time.Now(), time.Now(), nil))

	req := &pb.UpdateMedalRequest{
		Id:        "1",
		CountryId: "1",
		Type:      2,
		EventId:   "1",
		AthleteId: "1",
	}

	resp, err := repo.UpdateMedal(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "1", resp.CountryId)
	assert.Equal(t, "SILVER", resp.Type)
}

func TestDeleteMedal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPostgresMedalRepo(db)

	mock.ExpectExec("UPDATE medals SET deleted_at").WithArgs(sqlmock.AnyArg(), "1").WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.DeleteMedalRequest{
		Id: "1",
	}

	resp, err := repo.DeleteMedal(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

func TestGetMedalById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPostgresMedalRepo(db)

	mock.ExpectQuery("SELECT (.+) FROM medals WHERE id").WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"id", "country_id", "type", "event_id", "athlete_id", "created_at", "updated_at", "deleted_at"}).AddRow("1", "1", "GOLD", "1", "1", time.Now(), time.Now(), nil))

	req := &pb.GetMedalByIdRequest{
		Id: "1",
	}

	resp, err := repo.GetMedalById(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, "GOLD", resp.Type)
}

func TestGetMedals(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPostgresMedalRepo(db)

	rows := sqlmock.NewRows([]string{"id", "country_id", "type", "event_id", "athlete_id", "created_at", "updated_at", "deleted_at"}).
		AddRow("1", "1", "GOLD", "1", "1", time.Now(), time.Now(), nil).
		AddRow("2", "2", "SILVER", "2", "2", time.Now(), time.Now(), nil)

	mock.ExpectQuery("SELECT (.+) FROM medals").WillReturnRows(rows)

	resp, err := repo.GetMedals(&pb.VoidMedal{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Medals, 2)
}

func TestGetMedalByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPostgresMedalRepo(db)

	rows := sqlmock.NewRows([]string{"id", "country_id", "type", "event_id", "athlete_id", "created_at", "updated_at", "deleted_at"}).
		AddRow("1", "1", "GOLD", "1", "1", time.Now(), time.Now(), nil)

	mock.ExpectQuery("SELECT (.+) FROM medals WHERE country_id").WithArgs("1", "GOLD", "1", "1").WillReturnRows(rows)

	req := &pb.GetMedalByFilterRequest{
		CountryId: "1",
		Type:      1,
		EventId:   "1",
		AthleteId: "1",
	}

	resp, err := repo.GetMedalByFilter(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Medals, 1)
	assert.Equal(t, "GOLD", resp.Medals[0].Type)
}
