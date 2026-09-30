package models

// GetModels returns a slice of all database models for migration
func GetModels() []any {
	return []any{
		&User{},
	}
}
