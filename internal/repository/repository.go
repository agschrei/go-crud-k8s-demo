package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/agschrei/go-crud-k8s-demo/internal/models"
)

type Repository interface {
	GetAllAircraft() ([]models.Aircraft, error)
	GetAllAirports() ([]models.Airport, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (repo *PostgresRepository) GetAllAircraft() ([]models.Aircraft, error) {
	db := repo.db
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT aircraft_code, model, range FROM aircrafts"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aircraftSlice []models.Aircraft
	for rows.Next() {
		var aircraft models.Aircraft
		if err := rows.Scan(&aircraft.Code, &aircraft.Model, &aircraft.Range); err != nil {
			return aircraftSlice, err
		}
		aircraftSlice = append(aircraftSlice, aircraft)
	}
	return aircraftSlice, nil
}

func (repo *PostgresRepository) GetAllAirports() ([]models.Airport, error) {
	db := repo.db
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT airport_code, airport_name, city,  coordinates[0] as longitude, coordinates[1] as latitude, timezone FROM airports"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var airports []models.Airport
	for rows.Next() {
		airport := models.Airport{
			Coordinates: &models.LatLonCoordinate{},
		}
		if err := rows.Scan(&airport.Code, &airport.Name, &airport.City, &airport.Coordinates.Longitude, &airport.Coordinates.Latitude, &airport.Timezone); err != nil {
			return airports, err
		}
		airports = append(airports, airport)
	}
	return airports, nil
}
