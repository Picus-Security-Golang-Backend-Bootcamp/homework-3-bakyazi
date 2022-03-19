package author

import (
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migration() error {
	return r.db.AutoMigrate(&Author{})
}

func (r *Repository) FindAll(ctx context.Context, includeDeleted bool) ([]Author, error) {
	var authors []Author
	var result *gorm.DB

	tx := r.db.WithContext(ctx)
	if includeDeleted {
		result = tx.Unscoped().Find(&authors)
	} else {
		result = tx.Find(&authors)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

func (r *Repository) Insert(author Author) error {
	return r.db.Create(&author).Error
}

func (r *Repository) BulkInsert(authors []Author) error {
	var err error
	for _, a := range authors {
		err = r.Insert(a)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) Clear() error {
	return r.db.Unscoped().Where("id > ?", 0).Delete(&Author{}).Error
}
