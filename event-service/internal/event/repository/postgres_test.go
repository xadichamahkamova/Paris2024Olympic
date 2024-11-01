package repository

import (
	"testing"
	"time"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*PostgresEventRepository, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := NewPostgresEventRepository(db).(*PostgresEventRepository)

	return repo, mock, func() {
		db.Close()
	}
}

func TestCreateEvent(t *testing.T) {
	repo, mock, teardown := setupTest(t)
	defer teardown()

	req := &pb.CreateEventRequest{
		Name:      "Football Match",
		SportType: "Football",
		Location:  "Stadium",
		Date:      "2024-09-01",
		StartTime: "15:00",
		EndTime:   "17:00",
	}

	mock.ExpectQuery("INSERT INTO events").
		WithArgs(req.Name, req.SportType, req.Location, req.Date, req.StartTime, req.EndTime).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "sport_type", "location", "date", "start_time", "end_time", "created_at", "updated_at", "deleted_at"}).
			AddRow("1", req.Name, req.SportType, req.Location, req.Date, req.StartTime, req.EndTime, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), 0))

	resp, err := repo.CreateEvent(req)

	assert.NoError(t, err)
	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, req.Name, resp.Name)
}

func TestGetEvent(t *testing.T) {
	repo, mock, teardown := setupTest(t)
	defer teardown()

	req := &pb.GetEventRequest{Id: "1"}

	mock.ExpectQuery("SELECT (.+) FROM events").
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "sport_type", "location", "date", "start_time", "end_time", "created_at", "updated_at", "deleted_at"}).
			AddRow("1", "Football Match", "Football", "Stadium", "2024-09-01", "15:00", "17:00", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), 0))

	resp, err := repo.GetEvent(req)

	assert.NoError(t, err)
	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, "Football Match", resp.Name)
}

func TestListOfEvent(t *testing.T) {
	repo, mock, teardown := setupTest(t)
	defer teardown()

	mock.ExpectQuery("SELECT (.+) FROM events").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "sport_type", "location", "date", "start_time", "end_time", "created_at", "updated_at", "deleted_at"}).
			AddRow("1", "Football Match", "Football", "Stadium", "2024-09-01", "15:00", "17:00", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), 0).
			AddRow("2", "Basketball Game", "Basketball", "Arena", "2024-09-02", "18:00", "20:00", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), 0))

	resp, err := repo.ListOfEvent(&pb.ListOfEventRequest{})

	assert.NoError(t, err)
	assert.Len(t, resp.Events, 2)
	assert.Equal(t, "Football Match", resp.Events[0].Name)
	assert.Equal(t, "Basketball Game", resp.Events[1].Name)
}

func TestUpdateEvent(t *testing.T) {
	repo, mock, teardown := setupTest(t)
	defer teardown()

	req := &pb.UpdateEventRequest{
		Id:        "1",
		Name:      "Updated Football Match",
		SportType: "Football",
		Location:  "Updated Stadium",
		Date:      "2024-09-02",
		StartTime: "16:00",
		EndTime:   "18:00",
	}

	mock.ExpectQuery("UPDATE events SET").
		WithArgs(req.Name, req.SportType, req.Location, req.Date, req.StartTime, req.EndTime, req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "sport_type", "location", "date", "start_time", "end_time", "created_at", "updated_at", "deleted_at"}).
			AddRow("1", req.Name, req.SportType, req.Location, req.Date, req.StartTime, req.EndTime, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), 0))

	resp, err := repo.UpdateEvent(req)

	assert.NoError(t, err)
	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, req.Name, resp.Name)
}

func TestDeleteEvent(t *testing.T) {
	repo, mock, teardown := setupTest(t)
	defer teardown()

	req := &pb.DeleteEventRequest{Id: "1"}

	// To'g'ri SQL so'rovini aniqlang
	mock.ExpectExec(`UPDATE events SET deleted_at=DATE_PART\('epoch', CURRENT_TIMESTAMP\)::INT WHERE id=\$1`).
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.DeleteEvent(req)

	assert.NoError(t, err)
	assert.Equal(t, "deleted successfully", resp.Status)
}
