package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/EdlanioJ/nlwpacetime/server/configs"
	"github.com/EdlanioJ/nlwpacetime/server/database"
	"github.com/EdlanioJ/nlwpacetime/server/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthResponse struct {
	Token string `json:"token"`
}

func Auth(code string) (AuthResponse, error) {
	githubAccessToken, err := getGithubToken(code)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("GithubTokenError: %s", err.Error())
	}

	userInfo, err := getGithubUser(githubAccessToken)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("GithubUserError: %s", err.Error())
	}

	db, err := database.OpenConnection()
	if err != nil {
		return AuthResponse{}, fmt.Errorf("DBConnectionError: %s", err.Error())
	}
	defer db.Close()

	user := models.User{}
	err = db.Get(&user, "SELECT * FROM users WHERE github_id=$1", fmt.Sprintf("%v", userInfo.ID))
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return AuthResponse{}, fmt.Errorf("DatabaseError: %s", err.Error())
	}

	if user.ID == "" {
		user.ID = uuid.New().String()
		user.AvatarUrl = userInfo.AvatarUrl
		user.GithubID = fmt.Sprintf("%v", userInfo.ID)
		user.Login = userInfo.Login
		user.Name = userInfo.Name

		_, err = db.NamedExec("INSERT INTO users (id, github_id, name, login, avatar_url) VALUES (:id, :github_id, :name, :login, :avatar_url)", user)
		if err != nil {
			return AuthResponse{}, fmt.Errorf("DatabaseError: %s", err.Error())
		}
	}

	token, err := generateToken(user.ID, user.Name, user.AvatarUrl)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("TokenError: %s", err.Error())
	}

	return AuthResponse{
		Token: token,
	}, nil
}

func generateToken(sub string, name string, avatar string) (string, error) {
	claims := jwt.MapClaims{
		"sub":       sub,
		"name":      name,
		"avatarUrl": avatar,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cfg := configs.GetApi()

	signingKey := []byte(cfg.JwtSecret)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getGithubUser(token string) (models.GithubUserResponse, error) {
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if err != nil {
		return models.GithubUserResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.GithubUserResponse{}, err
	}

	userBody, _ := io.ReadAll(res.Body)

	var ghUserRes models.GithubUserResponse
	_ = json.Unmarshal(userBody, &ghUserRes)

	return ghUserRes, nil
}

func getGithubToken(code string) (string, error) {
	cfg := configs.GetGithub()
	reqBodyMap := map[string]string{
		"client_id":     cfg.ClientID,
		"client_secret": cfg.ClientSecret,
		"code":          code,
	}

	reqJson, _ := json.Marshal(reqBodyMap)

	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(reqJson),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	resBody, _ := io.ReadAll(res.Body)

	var ghResp models.GithubTokenResponse
	_ = json.Unmarshal(resBody, &ghResp)

	return ghResp.AccessToken, nil
}
