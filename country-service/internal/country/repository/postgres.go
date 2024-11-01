package repository

import (
	"country-service/logger"
	"database/sql"
	"errors"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
	"github.com/sirupsen/logrus"
)

type PostgresCountryRepository struct {
	DB *sql.DB
}

func NewPostgresCountryRepository(db *sql.DB) CountryRepository {
	return &PostgresCountryRepository{
		DB: db,
	}
}

func (db *PostgresCountryRepository) CreateCountry(req *pb.CreateCountryRequest) (*pb.Country, error) {

	resp := pb.Country{}
	query := `
	INSERT INTO countries(name, flag, region) 
	VALUES($1, $2, $3)
	RETURNING id, name, flag, region, created_at, updated_at, deleted_at`

	err := db.DB.QueryRow(query, req.Name, req.Flag, req.Region).Scan(
		&resp.Id,
		&resp.Name,
		&resp.Flag,
		&resp.Region,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)
	if err != nil {
		logger.Error("Creating country failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	logger.Info("Country created successfully", logrus.Fields{
		"country_id": resp.Id,
		"name":       resp.Name,
	})

	return &resp, nil
}

func (db *PostgresCountryRepository) GetCountry(req *pb.GetCountryRequest) (*pb.Country, error) {

	resp := pb.Country{}
	query := `
	SELECT id, name, flag, region, created_at, updated_at, deleted_at 
	FROM countries 
	WHERE id=$1 AND deleted_at=0`

	err := db.DB.QueryRow(query, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.Flag,
		&resp.Region,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)
	if err != nil {
		logger.Error("Retrieving country failed", logrus.Fields{
			"error":      err,
			"country_id": req.Id,
		})
		return nil, err
	}

	logger.Info("Country retrieved successfully", logrus.Fields{
		"country_id": resp.Id,
		"name":       resp.Name,
	})

	return &resp, nil
}

func (db *PostgresCountryRepository) ListOfCountry(req *pb.ListOfCountryRequest) (*pb.ListOfCountryResponse, error) {

	resp := pb.ListOfCountryResponse{}
	rows, err := db.DB.Query(`
	SELECT id, name, flag, region, created_at, updated_at, deleted_at 
	FROM countries
	WHERE deleted_at=0`)
	if err != nil {
		logger.Error("Listing countries failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := pb.Country{}
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Flag,
			&item.Region,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt,
		)
		if err != nil {
			logger.Error("Decoding country failed", logrus.Fields{
				"error": err,
			})
			return nil, err
		}
		resp.Countries = append(resp.Countries, &item)
	}

	logger.Info("Countries listed successfully", logrus.Fields{
		"countries_count": len(resp.Countries),
	})

	return &resp, nil
}

func (db *PostgresCountryRepository) UpdateCountry(req *pb.UpdateCountryRequest) (*pb.Country, error) {

	resp := pb.Country{}
	query := `
	UPDATE countries 
	SET name=$1, flag=$2, region=$3, updated_at=NOW() 
	WHERE id=$4 AND deleted_at=0
	RETURNING id, name, flag, region, created_at, updated_at, deleted_at`

	err := db.DB.QueryRow(query, req.Name, req.Flag, req.Region, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.Flag,
		&resp.Region,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)
	if err != nil {
		logger.Error("Updating country failed", logrus.Fields{
			"error":      err,
			"country_id": req.Id,
		})
		return nil, err
	}

	logger.Info("Country updated successfully", logrus.Fields{
		"country_id": resp.Id,
		"name":       resp.Name,
	})

	return &resp, nil
}

func (db *PostgresCountryRepository) DeleteCountry(req *pb.DeleteCountryRequest) (*pb.DeleteCountryResponse, error) {

	resp := pb.DeleteCountryResponse{}
	query := `
	UPDATE countries 
	SET deleted_at=DATE_PART('epoch', CURRENT_TIMESTAMP)::INT  
	WHERE id=$1`

	result, err := db.DB.Exec(query, req.Id)
	if err != nil {
		logger.Error("Deleting country failed", logrus.Fields{
			"error":      err,
			"country_id": req.Id,
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
			"country_id": req.Id,
		})
		return nil, errors.New("no rows affected")
	}

	resp.Status = "deleted successfully"
	logger.Info("Country deleted successfully", logrus.Fields{
		"country_id": req.Id,
	})

	return &resp, nil
}
