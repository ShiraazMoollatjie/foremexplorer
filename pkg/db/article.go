package db

import "github.com/jinzhu/gorm"

// Article is a snapshot of the price for a particular product.
type Article struct {
	gorm.Model
}

// AddArticle will create a price log entry for the provided product and price.
func AddArticle(db *gorm.DB) (uint, error) {
	p := &Article{}
	err := db.Create(p).Error

	return p.ID, err
}
