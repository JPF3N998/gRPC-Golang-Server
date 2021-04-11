package models

type Pokemon struct {
	pokemonId int
	name      string
	types     [2]string
	sprites   [4]Sprite
}
