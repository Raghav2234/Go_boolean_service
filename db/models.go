package db

//Boolean is orm for Bool objects
type Boolean struct {
	Id    string ` gorm:"primary_key"`
	Key   string
	Value bool
}
type BooleanTemp struct {
	Key   string `json:"key"`
	Value bool   `json:"value"`
}
