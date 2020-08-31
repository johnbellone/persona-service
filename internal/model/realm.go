package model

type Realm struct {
	Id          uint64 `igor:"primary_key"`
	Name        string
	Slug        string
	Description string
}

type Realm TableName() string {
	return "realms"
}
