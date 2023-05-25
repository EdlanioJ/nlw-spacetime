package services

import (
	"fmt"

	"github.com/EdlanioJ/nlwpacetime/server/database"
	"github.com/EdlanioJ/nlwpacetime/server/models"
)

func GetMemory(id string) (models.Memory, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return models.Memory{}, fmt.Errorf("DBConnectionError: %s", err.Error())
	}
	defer db.Close()

	memory := models.Memory{}
	err = db.Get(&memory, "SELECT * FROM memories WHERE id=$1", id)
	if err != nil {
		return models.Memory{}, fmt.Errorf("DatabaseError: %s", err.Error())
	}

	return memory, nil
}
