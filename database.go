package datacon

//Database : contract of Database for guardian
type Database interface {
	GetConnection() (interface{}, error)
}

