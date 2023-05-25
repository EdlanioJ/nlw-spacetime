package services

import (
	"fmt"
	"time"

	"github.com/EdlanioJ/nlwpacetime/server/database"
	"github.com/EdlanioJ/nlwpacetime/server/models"
)

type ListMemoriesByUserResponse struct {
	ID        string    `json:"id"`
	CoverUrl  string    `json:"coverUrl"`
	Excerpt   string    `json:"excerpt"`
	CreatedAt time.Time `json:"createdAt"`
}

func ListMemoriesByUser(userId string) ([]ListMemoriesByUserResponse, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return nil, fmt.Errorf("DBConnectionError: %s", err.Error())
	}
	defer db.Close()

	memories := []models.Memory{}
	err = db.Select(&memories, "SELECT * FROM memories WHERE user_id=$1 ORDER BY created_at ASC", userId)
	if err != nil {
		return nil, fmt.Errorf("DatabaseError: %s", err.Error())
	}

	return mapCollection(memories), nil
}

func mapCollection(memories []models.Memory) []ListMemoriesByUserResponse {
	res := []ListMemoriesByUserResponse{}

	for _, source := range memories {
		var excerpt string

		if len(source.Content) > 129 {
			excerpt = fmt.Sprintf("%s...", source.Content[1:129])
		} else {
			excerpt = source.Content
		}
		memory := ListMemoriesByUserResponse{
			ID:        source.ID,
			CoverUrl:  source.CoverUrl,
			Excerpt:   excerpt,
			CreatedAt: source.CreatedAt,
		}
		res = append(res, memory)
	}
	return res
}
