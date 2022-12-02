package db

type Database map[string]struct{}
type URLDatabase map[string]string

func Init() Database {
	return make(Database)
}

func InitURLDatabase() URLDatabase {
	return make(URLDatabase)
}
