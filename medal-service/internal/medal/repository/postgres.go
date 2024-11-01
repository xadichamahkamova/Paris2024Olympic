package repository

import (
	"database/sql"
	"fmt"
	"medal-service/logger"
	"time"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/medalspb"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type MedalRepo struct {
	db *sql.DB
}

func NewPostgresMedalRepo(db *sql.DB) MedalRepository {
	return &MedalRepo{db: db}
}

func (r *MedalRepo) CreateMedal(req *pb.CreateMedalRequest) (*pb.CreateMedalResponse, error) {
	query := `
		INSERT INTO medals (country_id, type, event_id, athlete_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, country_id, type, event_id, athlete_id, created_at, updated_at, deleted_at`
	var medal pb.Medal
	err := r.db.QueryRow(query, req.CountryId, req.Type, req.EventId, req.AthleteId).Scan(
		&medal.Id, &medal.CountryId, &medal.Type, &medal.EventId, &medal.AthleteId, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt)
	if err != nil {
		logger.Error("Failed to create medal", logrus.Fields{
			"error": err,
		})
		return nil, fmt.Errorf("failed to create medal: %v", err)
	}

	logger.Info("Medal created successfully", logrus.Fields{
		"id": medal.Id,
	})

	return &pb.CreateMedalResponse{
		Id:        medal.Id,
		CountryId: medal.CountryId,
		Type:      medal.Type,
		EventId:   medal.EventId,
		AthleteId: medal.AthleteId,
		CreatedAt: medal.CreatedAt,
		UpdatedAt: medal.UpdatedAt,
		DeletedAt: medal.DeletedAt,
	}, nil
}

func (r *MedalRepo) UpdateMedal(req *pb.UpdateMedalRequest) (*pb.UpdateMedalResponse, error) {
	query := `
		UPDATE medals
		SET country_id = $1, type = $2, event_id = $3, athlete_id = $4, updated_at = $5
		WHERE id = $6 AND deleted_at=0
		RETURNING id, country_id, type, event_id, athlete_id, created_at, updated_at, deleted_at`
	var medal pb.Medal
	err := r.db.QueryRow(query, req.CountryId, req.Type, req.EventId, req.AthleteId, time.Now().Format(time.RFC3339), req.Id).Scan(
		&medal.Id, &medal.CountryId, &medal.Type, &medal.EventId, &medal.AthleteId, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt)
	if err != nil {
		logger.Error("Failed to update medal", logrus.Fields{
			"error": err,
			"id":    req.Id,
		})
		return nil, fmt.Errorf("failed to update medal: %v", err)
	}

	logger.Info("Medal updated successfully", logrus.Fields{
		"id": medal.Id,
	})

	return &pb.UpdateMedalResponse{
		Id:        medal.Id,
		CountryId: medal.CountryId,
		Type:      medal.Type,
		EventId:   medal.EventId,
		AthleteId: medal.AthleteId,
		CreatedAt: medal.CreatedAt,
		UpdatedAt: medal.UpdatedAt,
		DeletedAt: medal.DeletedAt,
	}, nil
}

func (r *MedalRepo) DeleteMedal(req *pb.DeleteMedalRequest) (*pb.DeleteMedalResponse, error) {
	query := `UPDATE medals SET deleted_at = $1 WHERE id = $2`
	_, err := r.db.Exec(query, time.Now().Unix(), req.Id)
	if err != nil {
		logger.Error("Failed to delete medal", logrus.Fields{
			"error": err,
			"id":    req.Id,
		})
		return nil, fmt.Errorf("failed to delete medal: %v", err)
	}

	logger.Info("Medal deleted successfully", logrus.Fields{
		"id": req.Id,
	})
	return &pb.DeleteMedalResponse{Success: true}, nil
}

func (r *MedalRepo) GetMedalById(req *pb.GetMedalByIdRequest) (*pb.GetMedalByIdResponse, error) {
	query := `SELECT id, country_id, type, event_id, athlete_id, created_at, updated_at, deleted_at FROM medals WHERE id = $1`
	var medal pb.Medal
	err := r.db.QueryRow(query, req.Id).Scan(
		&medal.Id, &medal.CountryId, &medal.Type, &medal.EventId, &medal.AthleteId, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt)
	if err != nil {
		logger.Error("Failed to get medal by id", logrus.Fields{
			"error": err,
			"id":    req.Id,
		})
		return nil, fmt.Errorf("failed to get medal by id: %v", err)
	}

	logger.Info("Medal retrieved successfully", logrus.Fields{
		"id": medal.Id,
	})
	return &pb.GetMedalByIdResponse{
		Id:        medal.Id,
		CountryId: medal.CountryId,
		Type:      medal.Type,
		EventId:   medal.EventId,
		AthleteId: medal.AthleteId,
		CreatedAt: medal.CreatedAt,
		UpdatedAt: medal.UpdatedAt,
		DeletedAt: medal.DeletedAt,
	}, nil
}

func (r *MedalRepo) GetMedals(req *pb.VoidMedal) (*pb.GetMedalsResponse, error) {
	query := `SELECT id, country_id, type, event_id, athlete_id, created_at, updated_at, deleted_at FROM medals`
	rows, err := r.db.Query(query)
	if err != nil {
		logger.Error("Failed to get medals", logrus.Fields{
			"error": err,
		})
		return nil, fmt.Errorf("failed to get medals: %v", err)
	}
	defer rows.Close()

	var medals []*pb.Medal
	for rows.Next() {
		var medal pb.Medal
		err := rows.Scan(&medal.Id, &medal.CountryId, &medal.Type, &medal.EventId, &medal.AthleteId, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt)
		if err != nil {
			logger.Error("Failed to scan medal", logrus.Fields{
				"error": err,
			})
			return nil, fmt.Errorf("failed to scan medal: %v", err)
		}
		medals = append(medals, &medal)
	}

	logger.Info("Medals retrieved successfully", logrus.Fields{
		"count": len(medals),
	})

	return &pb.GetMedalsResponse{Medals: medals}, nil
}

func (r *MedalRepo) GetMedalByFilter(req *pb.GetMedalByFilterRequest) (*pb.GetMedalByFilterResponse, error) {
	query := `SELECT id, country_id, type, event_id, athlete_id, created_at, updated_at, deleted_at FROM medals WHERE country_id = $1 AND type = $2 AND event_id = $3 AND athlete_id = $4`
	rows, err := r.db.Query(query, req.CountryId, req.Type, req.EventId, req.AthleteId)
	if err != nil {
		logger.Error("Failed to get medals by filter", logrus.Fields{
			"error": err,
		})
		return nil, fmt.Errorf("failed to get medals by filter: %v", err)
	}
	defer rows.Close()

	var medals []*pb.Medal
	for rows.Next() {
		var medal pb.Medal
		err := rows.Scan(&medal.Id, &medal.CountryId, &medal.Type, &medal.EventId, &medal.AthleteId, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt)
		if err != nil {
			logger.Error("Failed to scan medal", logrus.Fields{
				"error": err,
			})
			return nil, fmt.Errorf("failed to scan medal: %v", err)
		}
		medals = append(medals, &medal)
	}
	logger.Info("Medals retrieved by filter successfully", logrus.Fields{
		"count": len(medals),
	})

	return &pb.GetMedalByFilterResponse{Medals: medals}, nil
}

