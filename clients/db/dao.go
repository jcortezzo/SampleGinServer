package db

func (d *Database) Add(s string) {
	(*d)[s] = struct{}{}
}

func (d *Database) Contains(s string) bool {
	_, ok := (*d)[s]
	return ok
}
