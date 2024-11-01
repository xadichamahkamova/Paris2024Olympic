package repository

import (
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
	"athlete-service/logger"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

type PostgresAthleteRepository struct {
	DB *sql.DB
}

func NewPostgresAthleteRepository(db *sql.DB) AthleteRepository {
	return &PostgresAthleteRepository{
		DB: db,
	}
}

func (db *PostgresAthleteRepository) CreateAthlete(req *pb.CreateAthleteRequest) (*pb.Athlete, error) {

	resp := pb.Athlete{}
	query := `
	INSERT INTO athletes(name, country_id, sport_type) 
	VALUES($1, $2, $3)
	RETURNING id, name, country_id, sport_type, created_at, updated_at, deleted_at`

	err := db.DB.QueryRow(query, req.Name, req.CountryId, req.SportType).Scan(
		&resp.Id,
		&resp.Name,
		&resp.CountryId,
		&resp.SportType,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)

	if err != nil {
		logger.Error("Creating athlete failed", logrus.Fields{
            "error": err,
        })
		return nil, err
	}

	logger.Info("Athlete created successfully", logrus.Fields{
        "athlete_id": resp.Id, 
        "name": resp.Name,
    })
	return &resp, nil
}

func (db *PostgresAthleteRepository) GetAthlete(req *pb.GetAthleteRequest) (*pb.GetAthleteResponse, error) {

	resp := pb.GetAthleteResponse{}
	query := `
	SELECT * 
	FROM athletes 
	WHERE id=$1 AND deleted_at=0`

	err := db.DB.QueryRow(query, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.CountryId,
		&resp.SportType,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)

	if err != nil {
		logger.Error("Retrieving athlete failed", logrus.Fields{
            "error": err, "athlete_id": req.Id,
        })
		return nil, err
	}

	logger.Info("Athlete retrieved successfully", logrus.Fields{
        "athlete_id": resp.Id, 
        "name": resp.Name,
    })
	return &resp, nil
}

func (db *PostgresAthleteRepository) ListAthletes(req *pb.ListOfAthleteRequest) (*pb.ListOfAthleteResponse, error) {

	resp := pb.ListOfAthleteResponse{}
	query := 
	`SELECT * 
	FROM athletes 
	WHERE deleted_at=0`
	rows, err := db.DB.Query(query)
	if err != nil {
		logger.Error("Listing athletes failed", logrus.Fields{"error": err})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := pb.GetAthleteResponse{}
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.CountryId,
			&item.SportType,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt,
		)
		if err != nil {
			logger.Error("Decoding athlete failed", logrus.Fields{
                "error": err,
            })
			return nil, err
		}
		resp.Athletes = append(resp.Athletes, &item)
	}

	logger.Info("Athletes listed successfully", logrus.Fields{
        "athletes_count": len(resp.Athletes),
    })
	return &resp, nil
}

func (db *PostgresAthleteRepository) UpdateAthlete(req *pb.UpdateAthleteRequest) (*pb.Athlete, error) {

	resp := pb.Athlete{}
	query := `
	UPDATE athletes 
	SET name=$1, country_id=$2, sport_type=$3, updated_at=NOW() 
	WHERE id=$4 AND deleted_at=0
	RETURNING id, name, country_id, sport_type, created_at, updated_at, deleted_at`

	err := db.DB.QueryRow(query, req.Name, req.CountryId, req.SportType, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.CountryId,
		&resp.SportType,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)

	if err != nil {
		logger.Error("Updating athlete failed", logrus.Fields{
            "error": err, 
            "athlete_id": req.Id,
        })
		return nil, err
	}

	logger.Info("Athlete updated successfully", logrus.Fields{
        "athlete_id": resp.Id, 
        "name": resp.Name,
    })
	return &resp, nil
}

func (db *PostgresAthleteRepository) DeleteAthlete(req *pb.DeleteAthleteRequest) (*pb.DeleteAthleteResponse, error) {

	resp := pb.DeleteAthleteResponse{}
	query := `
	UPDATE athletes 
	SET deleted_at=DATE_PART('epoch', CURRENT_TIMESTAMP)::INT 
	WHERE id=$1`

	result, err := db.DB.Exec(query, req.Id)
	if err != nil {
		logger.Error("Deleting athlete failed", logrus.Fields{
            "error": err, 
            "athlete_id": req.Id,
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
            "athlete_id": req.Id,
        })
		return nil, errors.New("no rows affected")
	}

	resp.Status = "deleted successfully"
	logger.Info("Athlete deleted successfully", logrus.Fields{
        "athlete_id": req.Id},
    )
	return &resp, nil
}
