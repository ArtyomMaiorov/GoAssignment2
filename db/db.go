package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

type Database struct {
	Conn *sql.DB
}

func Connect(psqlInfo string) (*Database, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// connection pool settings
	db.SetMaxOpenConns(10)   // max open connections
	db.SetMaxIdleConns(5)    // maximum idle connections

	return &Database{Conn: db}, nil
}

func (db *Database) Close() error {
	return db.Conn.Close()
}

func (db *Database) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		age INT NOT NULL
	);
	`
	_, err := db.Conn.Exec(query)
	return err
}

func (db *Database) InsertUser(name string, age int) error {
	query := "INSERT INTO users (name, age) VALUES ($1, $2)"
	_, err := db.Conn.Exec(query, name, age)
	return err
}

func (db *Database) QueryUsers(ageFilter *int, sortBy string, page, pageSize int) ([]User, error) {
	query := "SELECT id, name, age FROM users WHERE 1=1"
	args := []interface{}{}

	if ageFilter != nil {
		query += " AND age = $1"
		args = append(args, *ageFilter)
	}

	if sortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", sortBy)
	}

	query += " LIMIT $2 OFFSET $3"
	args = append(args, pageSize, (page-1)*pageSize)

	rows, err := db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (db *Database) UpdateUser(id uint, name string, age int) error {
	query := "UPDATE users SET name = $1, age = $2 WHERE id = $3"
	_, err := db.Conn.Exec(query, name, age, id)
	return err
}

func (db *Database) DeleteUser(id uint) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Conn.Exec(query, id)
	return err
}
