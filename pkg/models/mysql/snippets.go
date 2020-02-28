package mysql

import (
	"database/sql"
	"errors"
	"time"

	"adamhartleb/gists/pkg/models"
)

var layout = time.Now().String()

// DB.Query() for SELECT statements that return multiple rows.
// DB.QueryRow() for SELECT statements that return a single row.
// DB.Exec() for statements that don't return any rows (INSERT, DELETE)
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(
		stmt,
		title,
		content,
		expires,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	cursor := m.DB.QueryRow(stmt, id)

	snippet := models.Snippet{}

	if err := cursor.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return &snippet, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}

