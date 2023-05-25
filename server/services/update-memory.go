package services

import (
	"fmt"

	"github.com/EdlanioJ/nlwpacetime/server/database"
	"github.com/EdlanioJ/nlwpacetime/server/models"
)

type UpdateMemoryDto struct {
	ID       string `json:"-" form:"-"`
	Content  string `json:"content" form:"content"`
	CoverUrl string `json:"coverUrl" form:"coverUrl"`
	IsPublic bool   `json:"isPublic" form:"isPublic"`
	UserID   string `json:"-" form:"-"`
}

func UpdateMemory(dto UpdateMemoryDto) (models.Memory, error) {
	memory := models.Memory{}

	db, err := database.OpenConnection()
	if err != nil {
		return models.Memory{}, fmt.Errorf("DBConnectionError: %s", err.Error())
	}
	defer db.Close()

	err = db.Get(memory, "SELECT * FROM memories WHERE id=$1", dto.ID)
	if err != nil {
		return models.Memory{}, fmt.Errorf("DatabaseError: %s", err.Error())
	}

	if memory.UserID != dto.UserID {
		return models.Memory{}, fmt.Errorf("InvalidUserError: invalid user")
	}

	memory.Content = dto.Content
	memory.CoverUrl = dto.CoverUrl
	memory.IsPublic = dto.IsPublic

	_, err = db.NamedExec("UPDATE memories SET content=:content, cover_url=:cover_url, is_public=:is_public WHERE id=:id", memory)
	if err != nil {
		return models.Memory{}, fmt.Errorf("DatabaseError: %s", err.Error())
	}

	return memory, nil
}
