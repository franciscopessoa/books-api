package migrations

import (
	"github.com/franciscopessoa/books-api/database/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
