package models

type User struct {
	ID        string `json:"id" db:"id"`
	GithubID  string `json:"githubId" db:"github_id"`
	Name      string `json:"name" db:"name"`
	Login     string `json:"login" db:"login"`
	AvatarUrl string `json:"avatarUrl" db:"avatar_url"`
}
