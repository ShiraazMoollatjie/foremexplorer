package db

import (
	"strconv"
	"time"

	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
	"github.com/jinzhu/gorm"
)

// Article is a snapshot of the price for a particular product.
type Article struct {
	gorm.Model
	TypeOf               string
	DevToID              int
	Title                string
	Description          string `gorm:"type:varchar(4000)"`
	ReadablePublishDate  string
	Tags                 string
	URL                  string
	CommentsCount        int
	PublicReactionsCount int
	CrosspostedAt        time.Time `gorm:"default:current_timestamp"`
	PublishedAt          time.Time `gorm:"default:current_timestamp"`
	PublishedTimestamp   time.Time `gorm:"default:current_timestamp"`
}

// AddArticle will create a price log entry for the provided product and price.
func AddArticle(db *gorm.DB, article gophorem.Article) (uint, error) {
	p := &Article{
		TypeOf:               article.TypeOf,
		DevToID:              article.ID,
		Title:                article.Title,
		Description:          article.Description,
		ReadablePublishDate:  article.ReadablePublishDate,
		Tags:                 article.Tags,
		URL:                  article.URL,
		CommentsCount:        article.CommentsCount,
		PublicReactionsCount: article.PublicReactionsCount,
		CrosspostedAt:        article.CreatedAt.Time,
		PublishedAt:          article.PublishedAt.Time,
		PublishedTimestamp:   article.PublishedTimestamp.Time,
	}
	err := db.Create(p).Error

	return p.ID, err
}

func LookupByDevToID(db *gorm.DB, devToID int) (*Article, error) {
	var a Article
	r := db.Where(map[string]interface{}{"dev_to_id": strconv.Itoa(devToID)}).First(&a)

	return &a, r.Error
}
