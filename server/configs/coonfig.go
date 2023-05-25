package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API        APIConfig
	DB         DBConfig
	Uploadcare UploadcareConfig
	Github     GithubConfig
}

type APIConfig struct {
	Port      string
	JwtSecret string
}

type GithubConfig struct {
	ClientSecret string
	ClientID     string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type UploadcareConfig struct {
	SecretKey string
	PublicKey string
}

func init() {
	viper.SetDefault("PORT", "3000")

}

func Load() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port:      viper.GetString("PORT"),
		JwtSecret: viper.GetString("JWT_SECRET"),
	}
	cfg.DB = DBConfig{
		Host: viper.GetString("DB_HOST"),
		Port: viper.GetString("DB_PORT"),
		User: viper.GetString("DB_USER"),
		Pass: viper.GetString("DB_PASS"),
		Name: viper.GetString("DB_NAME"),
	}
	cfg.Uploadcare = UploadcareConfig{
		SecretKey: viper.GetString("UPLOADCARE_SECRET_KEY"),
		PublicKey: viper.GetString("UPLOADCARE_PUBLIC_KEY"),
	}

	cfg.Github = GithubConfig{
		ClientID:     viper.GetString("GITHUB_CLIENT_ID"),
		ClientSecret: viper.GetString("GITHUB_CLIENT_SECRET"),
	}
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetApi() APIConfig {
	return cfg.API
}

func GetUploadcare() UploadcareConfig {
	return cfg.Uploadcare
}

func GetGithub() GithubConfig {
	return cfg.Github
}
