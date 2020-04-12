package db

import "github.com/varunamachi/tasklist/srv/todo"

// PostgresStorage -
type PostgresStorage struct {
}

// Init -
func (pg *PostgresStorage) Init() error {
	if has, err := hasTable("tasks"); has || err != nil {
		return err
	}
	query := `
	CREATE TABLE tasks(
		id SERIAL PRIMARY KEY,
		heading VARCHAR(256),
		description TEXT,
		status CHAR(64),
		created TIMESTAMPTZ,
		deadline TIMESTAMPTZ,
		modified TIMESTAMPTZ
	);
	`
	_, err := db.Exec(query)
	return err
}

// Add -
func (pg *PostgresStorage) Add(ti *todo.TaskItem) error {
	// query := ``
	return nil
}

// Remove -
func (pg *PostgresStorage) Remove(id int) error {
	return nil
}

// Name -
func (pg *PostgresStorage) Name() string {
	return "pg"
}

// Update -
func (pg *PostgresStorage) Update(item *todo.TaskItem) error {
	return nil
}

// Retrieve -
func (pg *PostgresStorage) Retrieve(id int) error {
	return nil
}

// Bulk -
func (pg *PostgresStorage) Bulk(op todo.BulkOp) error {
	return nil
}

// RetrieveAll -
func (pg *PostgresStorage) RetrieveAll() []*todo.TaskItem {
	return nil
}

func hasTable(tableName string) (bool, error) {
	yes := false
	err := db.Get(&yes,
		`SELECT EXISTS (
			SELECT 1 FROM information_schema.tables 
			WHERE table_schema = 'public' AND table_name = $1)`,
		tableName)
	return yes, err
}
