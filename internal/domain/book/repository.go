package book

import (
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migration() error {
	return r.db.AutoMigrate(&Book{})
}

func (r *Repository) FindAll(ctx context.Context, includeDeleted bool) ([]Book, error) {
	var books []Book
	var result *gorm.DB

	tx := r.db.WithContext(ctx)
	if includeDeleted {
		result = tx.Unscoped().Find(&books)
	} else {
		result = tx.Find(&books)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *Repository) FindAllWithAuthor(ctx context.Context) ([]Book, error) {
	var books []Book
	tx := r.db.WithContext(ctx)
	result := tx.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *Repository) FindById(ctx context.Context, id int) (Book, error) {
	var b Book
	tx := r.db.WithContext(ctx)
	result := tx.First(&b, id)
	if result.Error != nil {
		return Book{}, result.Error
	}
	return b, nil
}

func (r *Repository) Search(ctx context.Context, text string) ([]Book, error) {
	var books []Book
	text = "%" + text + "%"
	tx := r.db.WithContext(ctx)
	result := tx.Preload("Author").Joins("join bakyazi_authors aut on aut.id = bakyazi_books.author_id").Where("bakyazi_books.Name ILIKE ? OR bakyazi_books.ISBN ILIKE ? OR aut.Name ILIKE ?", text, text, text).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *Repository) Update(ctx context.Context, b Book) error {
	tx := r.db.WithContext(ctx)
	result := tx.Save(b)
	return result.Error
}

func (r *Repository) Delete(ctx context.Context, b Book) error {
	tx := r.db.WithContext(ctx)
	result := tx.Delete(&b)
	return result.Error
}

func (r *Repository) Insert(book Book) error {
	return r.db.Create(&book).Error
}

func (r *Repository) BulkInsert(books []Book) error {
	var err error
	for _, b := range books {
		err = r.Insert(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) Clear() error {
	return r.db.Unscoped().Where("id > ?", 0).Delete(&Book{}).Error
}
