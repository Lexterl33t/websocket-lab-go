package models

const (
	UPGRADER_CONN = iota
)

type Error struct {
	Status  int
	Message string
}
