package param

type AddProgram struct {
	TypeName int    `json:"type"`
	Name     string `json:"name"`
	Point    string `json:"point"`
	Actors   string `json:"actors"`
}

type UpdateProgram struct {
	Id       int    `json:"id"`
	TypeName string `json:"type"`
	Name     string `json:"name"`
	Point    string `json:"point"`
	Actors   string `json:"actors"`
}
