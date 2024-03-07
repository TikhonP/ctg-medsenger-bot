package db

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

const schema = `
	CREATE TABLE IF NOT EXISTS contracts (
	    id INTEGER PRIMARY KEY NOT NULL,
	    is_active BOOLEAN NOT NULL,
	    agent_token VARCHAR(254),
	    patient_name VARCHAR(254),
	    patient_email VARCHAR(254),
	    patient_sex VARCHAR(20),
	    patient_phone VARCHAR(254),
	    timezone_offset INTEGER
	);
`

// Contract represents Medsenger contract.
// Create on agent /init and persist during agent lifecycle.
type Contract struct {
	Id             int     `db:"id"`
	IsActive       bool    `db:"is_active"`
	AgentToken     *string `db:"agent_token"`
	PatientName    *string `db:"patient_name"`
	PatientEmail   *string `db:"patient_email"`
	PatientSex     *string `db:"patient_sex"`
	PatientPhone   *string `db:"patient_phone"`
	TimezoneOffset *int    `db:"timezone_offset"`
}

// UpsetContract creates contract on database or sets it to active if already exists.
func UpsetContract(id int) error {
	query := "INSERT INTO contracts (id, is_active) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET is_active = EXCLUDED.is_active;"
	_, err := db.Exec(query, id, true)
	return err
}

// GetActiveContractIds returns all active contracts ids.
// Use it for medsenger status endpoint.
func GetActiveContractIds() ([]int, error) {
	var contractIds = make([]int, 0)
	err := db.Select(&contractIds, "SELECT id FROM contracts WHERE is_active = true")
	return contractIds, err
}

// MarkInactiveContractWithId sets contract with id to inactive.
// Use it for medsenger remove endpoint.
// Equivalent to DELETE FROM contracts WHERE id = ?.
func MarkInactiveContractWithId(id int) error {
	_, err := db.Exec("UPDATE contracts SET is_active = false WHERE id = $1", id)
	return err
}

// GetContractById returns contract with specified id.
func GetContractById(id int) (*Contract, error) {
	contract := new(Contract)
	err := db.Get(contract, "SELECT * FROM contracts WHERE id = $1", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Contract not found")
	}
	return contract, err
}
