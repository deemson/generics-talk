package dao

type Model interface {
	Id() string
	SetId(id string)
}
