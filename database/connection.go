package database

var connection string

func init() {
	connection = "mongoDb"
}

func C_database() string {
	return connection
}
