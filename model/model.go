package model

import (
	"github.com/uptrace/bun"
)

type TematicModel struct {
	bun.BaseModel `bun:"table:tematic,alias:t"`

	//ID IS AUNTOINCREMENT
	ID  int64 `bun:"id,pk,autoincrement" json:"id"`
	Name string `bun:"name,notnull" json:"name"`
	Slug string `bun:"slug,notnull" json:"slug"`

}

type MovieModel struct {
	bun.BaseModel `bun:"table:movie,alias:p"`

	//ID IS AUNTOINCREMENT
	ID  int64 `bun:"id,pk,autoincrement" json:"id"`
	Name string `bun:"name,notnull" json:"name"`
	Slug string `bun:"slug" json:"slug"`
	Description string `bun:"description" json:"description"`
	Year int `bun:"year" json:"year"`
	TematicID int64 `bun:"tematic_id" json:"tematic_id"`
	Tematic *TematicModel `bun:"rel:belongs-to,join:tematic_id=id" json:"tematic"`

}

type MoviePictureModel struct {
	bun.BaseModel `bun:"table:movie_picture,alias:mp"`

	//ID IS AUNTOINCREMENT
	ID  int64 `bun:"id,pk,autoincrement" json:"id"`
	Name string `bun:"name,notnull" json:"name"`
	MovieID int64 `bun:"movie_id" json:"movie_id"`
	Picture MovieModel `bun:"rel:belongs-to,join:movie_id=id" json:"picture"`
}

type PerfilModel struct {
	bun.BaseModel `bun:"table:perfil,alias:p"`

	//ID IS AUNTOINCREMENT
	ID  int64 `bun:"id,pk,autoincrement" json:"id"`
	Name string `bun:"name,notnull" json:"name"`
}

type UserModel struct {
	bun.BaseModel `bun:"table:user,alias:u"`

	//ID IS AUNTOINCREMENT
	ID  int64 `bun:"id,pk,autoincrement" json:"id"`
	Name string `bun:"name,notnull" json:"name"`
	Email string `bun:"email,notnull" json:"email"`
	Telephone string `bun:"telephone" json:"telephone"`
	Password string `bun:"password,notnull" json:"password"`
	PerfilID int64 `bun:"perfil_id" json:"perfil_id"`
	Perfil *PerfilModel `bun:"rel:belongs-to,join:perfil_id=id" json:"perfil"`
}
