package db

func (d *Database) Add(s string) {
	(*d)[s] = struct{}{}
}

func (d *Database) Contains(s string) bool {
	_, ok := (*d)[s]
	return ok
}

func (d *Database) Delete(s string) {
	delete(*d, s)
}

func (d *URLDatabase) Add(key string, val string) {
	(*d)[key] = val
}

func (d *URLDatabase) Contains(key string) bool {
	_, ok := (*d)[key]
	return ok
}

func (d *URLDatabase) Get(key string) string {
	return (*d)[key]
}
