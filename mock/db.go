package mock

//go:generate  mockgen -source=db.go -destination=db_mock.go -package=mock

type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	value, err := db.Get(key)
	if err == nil {
		return value
	}
	return -1
}
