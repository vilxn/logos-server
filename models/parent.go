package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ParentRepository struct {
	db        *sql.DB
	parent_id int64
}

func NewParentRepository(db *sql.DB) *ParentRepository {
	return &ParentRepository{db: db}
}

func (r *ParentRepository) SetParentID(id int64) {
	r.parent_id = id
}

// AddChild creates a new child and links them to the parent in one transaction.
func (r *ParentRepository) AddChild(child Child) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	res, err := tx.Exec(
		`INSERT INTO children (first_name, last_name, birth_date, notes)
         VALUES (?, ?, ?, ?)`,
		child.FirstName, child.LastName, child.BirthDate, child.Notes,
	)
	if err != nil {
		return 0, err
	}

	childID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(
		`INSERT INTO child_parents (child_id, parent_id) VALUES (?, ?)`,
		childID, r.parent_id,
	)
	if err != nil {
		return 0, err
	}

	return childID, tx.Commit()
}

func (r *ParentRepository) GetChildFromID(childID int64) (*Child, error) {
	row := r.db.QueryRow(
		`SELECT c.id, c.first_name, c.last_name, c.birth_date, c.notes
         FROM children c
         JOIN child_parents cp ON cp.child_id = c.id
         WHERE c.id = ? AND cp.parent_id = ?`,
		childID, r.parent_id,
	)

	var child Child
	err := row.Scan(&child.ID, &child.FirstName, &child.LastName, &child.BirthDate, &child.Notes)
	if err == sql.ErrNoRows {
		return nil, nil // child not found or doesn't belong to this parent
	}
	if err != nil {
		return nil, err
	}

	return &child, nil
}

// Can be troubles with performance due to GetChildFromID
// A single JOIN query would be more efficient
func (r *ParentRepository) GetChildren() ([]*Child, error) {
	rows, err := r.db.Query(
		`SELECT child_id FROM child_parents WHERE parent_id = ?`,
		r.parent_id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var children []*Child
	for rows.Next() {
		var childID int64
		if err := rows.Scan(&childID); err != nil {
			return nil, err
		}

		child, err := r.GetChildFromID(childID)
		if err != nil {
			return nil, err
		}
		if child != nil {
			children = append(children, child)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return children, nil
}
