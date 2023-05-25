package services

import (
	"fmt"

	"github.com/EdlanioJ/nlwpacetime/server/database"
	"github.com/EdlanioJ/nlwpacetime/server/models"
)

type DeleteMemoryDto struct {
	ID     string
	UserID string
}

func DeleteMemory(dto DeleteMemoryDto) error {
	memory := models.Memory{}

	db, err := database.OpenConnection()
	if err != nil {
		return fmt.Errorf("DBConnectionError: %s", err.Error())
	}
	defer db.Close()

	err = db.Get(memory, "SELECT * FROM memories WHERE id=$1", dto.ID)
	if err != nil {
		return fmt.Errorf("DatabaseError: %s", err.Error())
	}

	if memory.UserID != dto.UserID {
		return fmt.Errorf("InvalidUserError: invalid user")
	}

	_, err = db.NamedExec("DELETE FROM memories WHERE id=$1", dto.ID)
	if err != nil {
		return fmt.Errorf("DatabaseError: %s", err.Error())
	}

	return nil
}
