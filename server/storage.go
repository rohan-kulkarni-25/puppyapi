package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresStore() (*PostgresStore, error) {
	// connStr := "user=postgres dbname=postgres password=rohan sslmode=disable"
	connStr := "user=postgres password=rohan dbname=postgres host=go_db sslmode=disable "
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return &PostgresStore{
		db: *db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createDogsTable()
}

func (s *PostgresStore) createDogsTable() error {
	query := `create table if not exists dogs (
		name varchar(100) primary key,
		path varchar(100)
	)`
	_, err := s.db.Exec(query)
	// Here we can use MustExec ALso but it will not work well with Errors
	return err
}

func (s *PostgresStore) GetDogImageByName(name string) (*Dog, error) {
	rows, _ := s.db.Queryx("select * from dogs where name = $1", name)
	for rows.Next() {
		return scanIntoDogs(rows)
	}

	return nil, fmt.Errorf("Dog with name [%s] not found", name)

}

func scanIntoDogs(rows *sqlx.Rows) (*Dog, error) {
	Dog := new(Dog)
	err := rows.Scan(
		&Dog.Name,
		&Dog.Path)

	return Dog, err
}

func (s *PostgresStore) AddDog(dog *Dog) error {
	query := `insert into dogs 
	(name,path)
	values ($1, $2)`

	_, err := s.db.Query(
		query,
		dog.Name,
		dog.Path)

	if err != nil {
		return err
	}

	return nil
}
