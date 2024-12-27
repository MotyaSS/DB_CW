package storage

import (
	"fmt"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type RepairPostgres struct {
	db *sqlx.DB
}

func NewRepairPostgres(db *sqlx.DB) *RepairPostgres {
	return &RepairPostgres{db: db}
}

func (r *RepairPostgres) GetRepair(id int) (entity.Repair, error) {
	var repair entity.Repair
	query := `SELECT * FROM repairs WHERE repair_id = $1`
	err := r.db.Get(&repair, query, id)
	if err != nil {
		return entity.Repair{}, fmt.Errorf("failed to get repair: %w", err)
	}
	return repair, nil
}

func (r *RepairPostgres) CreateRepair(callerId int, repair entity.Repair) (id int, err error) {
	query := `INSERT INTO repairs (instrument_id, repair_start_date, repair_end_date, repair_cost, description) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING repair_id`

	err = r.db.QueryRow(
		query,
		repair.InstrumentId,
		repair.RepairStartDate,
		repair.RepairEndDate,
		repair.RepairCost,
		repair.Description,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create repair: %w", err)
	}
	return id, nil
}

func (r *RepairPostgres) GetInstrumentRepairs(instrumentId int) ([]entity.Repair, error) {
	var repairs []entity.Repair
	query := `SELECT * FROM repairs WHERE instrument_id = $1`
	err := r.db.Select(&repairs, query, instrumentId)
	if err != nil {
		return nil, fmt.Errorf("failed to get instrument repairs: %w", err)
	}
	return repairs, nil
}
