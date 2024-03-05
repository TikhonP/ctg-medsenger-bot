package db

var schema = `
	CREATE TABLE IF NOT EXISTS contracts (
	    id INTEGER PRIMARY KEY NOT NULL,
	    is_active BOOLEAN NOT NULL,
	    agent_token STRING,
	    patient_name STRING,
	    patient_email STRING,
	    patient_sex STRING,
	    patient_phone STRING,
	    timezone_offset INTEGER
	);
`

// Contract represents Medsenger contract.
// Create on agent /init and persist during agent lifecycle.
type Contract struct {
	Id             int
	IsActive       bool
	AgentToken     string
	PatientName    string
	PatientEmail   string
	PatientSex     string
	PatientPhone   string
	TimezoneOffset int
}

func UpsetContract(id int) error {
	query := "INSERT INTO contracts (id, is_active) VALUES (?, ?) ON CONFLICT conflict_target conflict_action"
	_, err := db.Exec(query, id, true)
	return err
}
