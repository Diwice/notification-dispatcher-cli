package state

import (
	"log"
	"database/sql"
	"not-dis/internal/db"

	_ "github.com/mattn/go-sqlite3"
)

type State struct {
	Queries *db.Queries
}

func NewState(path string) *State, error {
	log.Println("Trying to open the database at path", path)

	localDB, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	newQueries := db.New(localDB)

	newState := State{
		Queries: db,
	}

	return &newState, nil
}
