package data

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Storage interface {
	CreateEvent(e *Event) (int32, error)
	GetAllEvents() ([]Event, error)
	UpdateEvent(id int32, b *EventBase) (*Event, error)
	DeleteEvent(id int32) error
	CheckVersion() (string, error)
}

type postgreStorage struct {
	db *sql.DB
}

func NewPostgreStorage(user string, password string, dbname string) (*postgreStorage, error) {
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=postgres sslmode=disable", user, dbname, password)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	q := `
		CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		title TEXT,
		likes NUMERIC,
		media VARCHAR(100)[],
		author NUMERIC,
		created_at TIMESTAMP,
		date TIMESTAMP,
		description TEXT,
		members NUMERIC[]
	)`
	_, err = db.Exec(q)
	if err != nil {
		return nil, err
	}

	return &postgreStorage{
		db: db,
	}, nil
}

func (s *postgreStorage) CheckVersion() (string, error) {
	rows, err := s.db.Query("SELECT version()")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var vers string
	for rows.Next() {
		if err := rows.Scan(&vers); err != nil {
			return "", err
		}
	}
	return vers, nil
}

func (s *postgreStorage) CreateEvent(e *Event) (int32, error) {
	q := `
	INSERT INTO events
	(title, likes, media, author, created_at, date, description, members)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id
	`
	var id int32
	rows, err := s.db.Query(q,
		e.Title,
		e.Likes,
		pq.Array(e.Media),
		e.Author,
		e.CreatedAt,
		e.Date,
		e.Description,
		pq.Array(e.Members))
	if err != nil {
		return -1, err
	}
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (s *postgreStorage) GetAllEvents() ([]Event, error) {
	events := make([]Event, 0)

	q := `SELECT * FROM events ORDER BY id`
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		event := Event{}
		var members pq.Int32Array
		if err := rows.Scan(&event.Id,
			&event.Title,
			&event.Likes,
			pq.Array(&event.Media),
			&event.Author,
			&event.CreatedAt,
			&event.Date,
			&event.Description,
			&members); err != nil {
			return nil, err
		}
		event.Members = []int32(members)
		events = append(events, event)
	}
	return events, nil
}

func (s *postgreStorage) UpdateEvent(id int32, b *EventBase) (*Event, error) {
	q := `
	UPDATE events SET
	title = $1,
	media = $2,
	date = $3,
	description = $4,
	members = $5
	WHERE id = $6
	RETURNING *
	`
	e := new(Event)
	rows, err := s.db.Query(q,
		b.Title,
		pq.Array(b.Media),
		b.Date,
		b.Description,
		pq.Array(b.Members),
		id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var members pq.Int32Array
		if err := rows.Scan(&e.Id,
			&e.Title,
			&e.Likes,
			pq.Array(&e.Media),
			&e.Author,
			&e.CreatedAt,
			&e.Date,
			&e.Description,
			&members); err != nil {
			return nil, err
		}
		e.Members = []int32(members)
	}
	return e, nil
}

func (s *postgreStorage) DeleteEvent(id int32) error {
	q := `DELETE FROM events WHERE id = $1`
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}
