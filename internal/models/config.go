package models

type Config struct {
	SqliteConfig SqliteConfig
}

type SqliteConfig struct {
	DataSourceName string
	DriverName     string
}
