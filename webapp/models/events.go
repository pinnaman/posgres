package models

import (
	"database/sql"
)

type Event struct {
	Id        string
	Created   string
	Start     string
	End       sql.NullString
	Title     string
	Completed string
}

func AllEvents() ([]*Event, error) {
	//fmt.Println("HERE 0....")
	//InitDB("postgres://postgres:secret@192.168.99.100/realtime sslmode=disable")
	InitDB("host=192.168.99.100 user=postgres password=secret dbname=realtime sslmode=disable")
	rows, err := db.Query("select * from events ")
	//fmt.Println("HERE 1....")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	evts := make([]*Event, 0)
	for rows.Next() {
		evt := new(Event)
		err := rows.Scan(&evt.Id, &evt.Created, &evt.Start, &evt.End, &evt.Title, &evt.Completed)
		if err != nil {
			return nil, err
		}
		evts = append(evts, evt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return evts, nil
}
