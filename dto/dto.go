package dto

type TematicDto struct {
	Name string `json:"name"`	
}

type MovieDto struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Year int `json:"year"`
	TematicID int64 `json:"tematic_id"`
}

type UserDto struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Telephone string `json:"telephone"`
	Password string `json:"password"`
	PerfilID int64 `json:"perfil_id"`
}

type PerfilDto struct {
	Name string `json:"name"`
}

type LoginDto struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginAnswerDto struct {
	Name string `json:"name"`
	Token string `json:"token"`
}

