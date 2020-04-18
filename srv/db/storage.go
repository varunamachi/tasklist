package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/varunamachi/tasklist/srv/todo"
)

var updateQuery = `
	UPDATE tasks SET 
	heading = :heading, 
	description = :description, 
	status = :status, 
	created = :created, 
	deadline = :deadline, 
	modified = :modified, 
	WHERE id = :id
`

var createQuery = `INSERT INTO tasks(
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

var deleteQuery = `DELETE FROM tasks WHERE id = $1`

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
	_, err := db.NamedExec(createQuery, ti)
	return err
}

// Remove -
func (pg *PostgresStorage) Remove(id int) error {
	_, err := db.Exec(deleteQuery, id)
	return err
}

// Name -
func (pg *PostgresStorage) Name() string {
	return "pg"
}

// Update -
func (pg *PostgresStorage) Update(item *todo.TaskItem) error {
	_, err := db.Exec(updateQuery, item)
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
func (pg *PostgresStorage) Bulk(op todo.BulkOp) (err error) {
	query := ""
	switch op.Op {
	case todo.Update:
		query = updateQuery
	case todo.Create:
		query = createQuery
	case todo.Delete:
		query = deleteQuery
	}
	ctx, done := context.WithTimeout(context.Background(), 1*time.Minute)
	defer done()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		err = tx.Commit()
	}()

	for _, ti := range op.Items {
		tx.Exec(query, ti)
	}
	return err
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
