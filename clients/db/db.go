package db

type Database map[string]struct{}

func Init() Database {
	return make(Database)
}
