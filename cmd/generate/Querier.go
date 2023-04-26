package main

import "gorm.io/gen"

type Querier interface {
	// ProgramSearch SELECT * FROM @@table
	//{{where}}
	//	{{if category_id != ""}}
	//		category_id = @categoryId
	//	{{end}}
	//	{{if name !=""}}
	//		AND name = @name
	//	{{end}}
	//	{{if num != -1}}
	//		AND num = @num
	//	{{end}}
	//{{end}}
	ProgramSearch(categoryId, name, string, num int) ([]gen.T, error)
}
