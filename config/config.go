package config

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}

func LoadConfig() *Config {
    return &Config{
        DBHost:     "localhost",
        DBPort:     "5432",
        DBUser:     "fleet_manager",
        DBPassword: "fleet5427+",
        DBName:     "fleet",
    }
}
