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
	query := `INSERT INTO tasks(
		heading,
		description,
		status,
		created,
		deadline,
		modified
	) VALUES(
		:heading,
		:description,
		:status,
		:created,
		:deadline,
		:modified
	);`
	_, err := db.NamedExec(query, ti)
	return err
}

// Remove -
func (pg *PostgresStorage) Remove(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

// Name -
func (pg *PostgresStorage) Name() string {
	return "pg"
}

// Update -
func (pg *PostgresStorage) Update(item *todo.TaskItem) error {
	query := `UPDATE tasks SET 
		heading = :heading, 
		description = :description, 
		status = :status, 
		created = :created, 
		deadline = :deadline, 
		modified = :modified, 
		WHERE id = :id
	`
	_, err := db.Exec(query, item)
	return err
}

// Retrieve -
func (pg *PostgresStorage) Retrieve(id int) (*todo.TaskItem, error) {
	query := `SELECT * FROM tasks WHERE id = $1`
	var ti todo.TaskItem
	err := db.Get(ti, query, id)
	return &ti, err
}

// Bulk -
func (pg *PostgresStorage) Bulk(op todo.BulkOp) error {
	return nil
}

// RetrieveAll -
func (pg *PostgresStorage) RetrieveAll(offset, limit int) (
	[]*todo.TaskItem, error) {
	query := `SELECT * FROM tasks 
		ORDER BY modified DESC OFFSET $1 LIMIT $2;`
	items := make([]*todo.TaskItem, 0, limit)
	err := db.Select(&items, query, offset, limit)
	return items, err
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
