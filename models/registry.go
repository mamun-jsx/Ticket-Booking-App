package models

var migrationRegistry []any // an empty slice to hold data

// RegisterModel allows individual files to register their own structs

func RegisterModel(model any) {
	migrationRegistry = append(migrationRegistry, model)
}

// func for return everything that called by init func()
func GetRegisterModels() []any {
	return migrationRegistry
}
