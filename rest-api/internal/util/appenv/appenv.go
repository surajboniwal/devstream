package appenv

func AppEnv() string {
	return Getenv("ENV", "development")
}
