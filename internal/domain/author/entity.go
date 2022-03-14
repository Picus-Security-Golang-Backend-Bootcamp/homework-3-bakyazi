package author

import "gorm.io/gorm"

type Author struct {
	ID      int
	Name    string
	Deleted gorm.DeletedAt
}
