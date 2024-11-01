package repository

import (
	"database/sql"
	"errors"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"
	"event-service/logger" 
	"github.com/sirupsen/logrus"
)

type PostgresEventRepository struct {
	DB *sql.DB
}

func NewPostgresEventRepository(db *sql.DB) EventRepository {
	return &PostgresEventRepository{
		DB: db,
	}
}

func (db *PostgresEventRepository) CreateEvent(req *pb.CreateEventRequest) (*pb.Event, error) {

	resp := pb.Event{}
	query := `
	INSERT INTO events(name, sport_type, location, date, start_time, end_time) 
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING id, name, sport_type, location, date, start_time, end_time, created_at, updated_at, deleted_at`
	err := db.DB.QueryRow(query,
		req.Name,
		req.SportType,
		req.Location,
		req.Date,
		req.StartTime,
		req.EndTime).Scan(
		&resp.Id,
		&resp.Name,
		&resp.SportType,
		&resp.Location,
		&resp.Date,
		&resp.StartTime,
		&resp.EndTime,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)
	if err != nil {
		logger.Error("Creating event failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	logger.Info("Event created successfully", logrus.Fields{
		"event_id": resp.Id,
		"name":     resp.Name,
	})

	return &resp, nil
}

func (db *PostgresEventRepository) GetEvent(req *pb.GetEventRequest) (*pb.Event, error) {

	resp := pb.Event{}
	query := `
	SELECT id, name, sport_type, location, date, start_time, end_time, created_at, updated_at, deleted_at 
	FROM events 
	WHERE id=$1 AND deleted_at=0`
	err := db.DB.QueryRow(query, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.SportType,
		&resp.Location,
		&resp.Date,
		&resp.StartTime,
		&resp.EndTime,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)
	if err != nil {
		logger.Error("Retrieving event failed", logrus.Fields{
			"error":    err,
			"event_id": req.Id,
		})
		return nil, err
	}

	logger.Info("Event retrieved successfully", logrus.Fields{
		"event_id": resp.Id,
		"name":     resp.Name,
	})

	return &resp, nil
}

func (db *PostgresEventRepository) ListOfEvent(req *pb.ListOfEventRequest) (*pb.ListOfEventResponse, error) {

	resp := pb.ListOfEventResponse{}
	rows, err := db.DB.Query(`
	SELECT id, name, sport_type, location, date, start_time, end_time, created_at, updated_at, deleted_at 
	FROM events
	WHERE deleted_at=0`)
	if err != nil {
		logger.Error("Listing events failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}
	defer rows.Close() 

	for rows.Next() {
		item := pb.Event{}
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.SportType,
			&item.Location,
				&item.Date,
			&item.StartTime,
			&item.EndTime,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt,
		)
		if err != nil {
			logger.Error("Decoding event failed", logrus.Fields{
				"error": err,
			})
			return nil, err
		}
		resp.Events = append(resp.Events, &item)
	}

	logger.Info("Events listed successfully", logrus.Fields{
		"events_count": len(resp.Events),
	})

	return &resp, nil
}

func (db *PostgresEventRepository) UpdateEvent(req *pb.UpdateEventRequest) (*pb.Event, error) {

	resp := pb.Event{}
	query := `
	UPDATE events 
	SET name=$1, sport_type=$2, location=$3, date=$4, start_time=$5, end_time=$6, updated_at=NOW() 
	WHERE id=$7 AND deleted_at=0
	RETURNING id, name, sport_type, location, date, start_time, end_time, created_at, updated_at, deleted_at`
	err := db.DB.QueryRow(query,
		req.Name,
		req.SportType,
		req.Location,
		req.Date,
		req.StartTime,
		req.EndTime,
		req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.SportType,
		&resp.Location,
		&resp.Date,
		&resp.StartTime,
		&resp.EndTime,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)
	if err != nil {
		logger.Error("Updating event failed", logrus.Fields{
			"error":    err,
			"event_id": req.Id,
		})
		return nil, err
	}

	logger.Info("Event updated successfully", logrus.Fields{
		"event_id": resp.Id,
		"name":     resp.Name,
	})

	return &resp, nil
}

func (db *PostgresEventRepository) DeleteEvent(req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {

	resp := pb.DeleteEventResponse{}
	query := `
	UPDATE events 
	SET deleted_at=DATE_PART('epoch', CURRENT_TIMESTAMP)::INT  
	WHERE id=$1`
	result, err := db.DB.Exec(query, req.Id)
	if err != nil {
		logger.Error("Deleting event failed", logrus.Fields{
			"error":    err,
			"event_id": req.Id,
		})
		return nil, err
	}

	num, err := result.RowsAffected()
	if err != nil {
		logger.Error("Getting affected rows failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	if num == 0 {
		logger.Warn("No rows affected for deletion", logrus.Fields{
			"event_id": req.Id,
		})
		return nil, errors.New("no rows affected")
	}

	resp.Status = "deleted successfully"
	logger.Info("Event deleted successfully", logrus.Fields{
		"event_id": req.Id,
	})

	return &resp, nil
}
