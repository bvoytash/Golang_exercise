package models

import (
	"app/db"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (event *Event) SaveEvent() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return err
	}
	event.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `UPDATE events
	SET name=?, description=?, location=?, datetime=?, user_id=?
	WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return err
	}
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return err
	}
	return err
}

func (e Event) Register(userID int64) error {
	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userID)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return err
	}
	return err
}

func (e Event) CancelRegister(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id=? AND user_id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return err
	}
	return err

}
