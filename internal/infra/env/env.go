package env

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
)

type Env struct {
	AppPort              int                      `env:"APP_PORT"`
	DBUsername           string                   `env:"DB_USERNAME"`
	DBPassword           string                   `env:"DB_PASSWORD"`
	DBHost               string                   `env:"DB_HOST"`
	DBPort               int                      `env:"DB_PORT"`
	DBName               string                   `env:"DB_NAME"`
	JWTSecretKey         string                   `env:"JWT_SECRET_KEY"`
	JWTExpired           int                      `env:"JWT_EXPIRED"`
	SerivceCost          float64                  `env:"SERVICE_COST"`
	DepositePercentage   float64                  `env:"DEPOSITE_PERCENTAGE"`
	MidtransServerKey    string                   `env:"MIDTRANS_SERVER_KEY"`
	MidtransEnvironment  midtrans.EnvironmentType `env:"MIDTRANS_ENVIRONMENT"`
	AWSS3Bucket          string                   `env:"AWS_S3_BUCKET"`
	AWSS3Endpoint        string                   `env:"AWS_S3_ENDPOINT"`
	AWSS3Region          string                   `env:"AWS_S3_REGION"`
	AWSS3AccessKey       string                   `env:"AWS_S3_ACCESS_KEY"`
	AWSS3SecretAccessKey string                   `env:"AWS_S3_SECRET_ACCESS_KEY"`
	SupabaseBukcetName   string                   `env:"SUPABASE_BUCKET_NAME"`
	SupabaseProjectURL   string                   `env:"SUPABASE_PROJECT_URL"`
	SupabaseAPIKey       string                   `env:"SUPABASE_API_KEY"`
	SupabaseAccessToken  string                   `env:"SUPABASE_ACCESS_TOKEN"`
}

func New() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	_env := new(Env)
	err = env.Parse(_env)
	if err != nil {
		return nil, err
	}

	return _env, nil
}
