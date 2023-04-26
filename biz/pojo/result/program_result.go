package result

type Program struct {
	Id        int32  `json:"id"`
	TypeName  string `json:"typeName"`
	Name      string `json:"name"`
	View      string `json:"view"`
	ActorList string `json:"actorList"`
}

type ProgramResults struct {
	ProgramResults []*Program `json:"programResults"`
}
