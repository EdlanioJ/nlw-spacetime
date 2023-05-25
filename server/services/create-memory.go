package services

import (
	"fmt"
	"time"

	"github.com/EdlanioJ/nlwpacetime/server/database"
	"github.com/EdlanioJ/nlwpacetime/server/models"
	"github.com/google/uuid"
)

type CreateMemoryDto struct {
	Content  string `json:"content" form:"content"`
	CoverUrl string `json:"coverUrl" form:"coverUrl"`
	IsPublic bool   `json:"isPublic" form:"isPublic"`
	UserID   string `json:"-" form:"-"`
}

func CreateMemory(dto CreateMemoryDto) (models.Memory, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return models.Memory{}, fmt.Errorf("DBConnectionError: %s", err.Error())
	}
	defer db.Close()

	memory := models.Memory{
		ID:        uuid.New().String(),
		UserID:    dto.UserID,
		CoverUrl:  dto.CoverUrl,
		Content:   dto.Content,
		IsPublic:  dto.IsPublic,
		CreatedAt: time.Now(),
	}
	_, err = db.NamedExec("INSERT INTO memories (id, cover_url, content, is_public, user_id, created_at) VALUES (:id, :cover_url, :content, :is_public, :user_id, :created_at)", memory)
	if err != nil {
		return models.Memory{}, fmt.Errorf("DatabaseError: %s", err.Error())
	}

	return memory, nil
}
