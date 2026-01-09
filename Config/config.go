package Config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configaration Config

type Config struct {
	Version            string
	ServiceName        string
	HttpPort           int64
	JwtSecretKey       string
	AllowedOrigins     string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	FrontendURL        string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("requaird version")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("requaird service name")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("requaird http port")
		os.Exit(1)
	}
	Port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("requaird http port")
		return
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("required jwt secret key")
		os.Exit(1)
	}
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}
	//OAUTH GOOGLE CONFIG
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	if googleClientID == "" {
		fmt.Println("required Google Client ID")
		os.Exit(1)
	}

	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if googleClientSecret == "" {
		fmt.Println("required Google Client Secret")
		os.Exit(1)
	}

	googleRedirectURL := os.Getenv("GOOGLE_REDIRECT_URL")
	if googleRedirectURL == "" {
		googleRedirectURL = "http://localhost:3000/auth/google/callback"
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	configaration = Config{
		Version:            version,
		ServiceName:        serviceName,
		HttpPort:           Port,
		JwtSecretKey:       jwtSecretKey,
		AllowedOrigins:     allowedOrigins,
		GoogleClientID:     googleClientID,
		GoogleClientSecret: googleClientSecret,
		GoogleRedirectURL:  googleRedirectURL,
		FrontendURL:        frontendURL,
	}
}
func GetConfig() Config {
	if configaration == (Config{}) {
		loadConfig()
	}
	return configaration
}
