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