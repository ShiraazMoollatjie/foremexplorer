package ops

import (
	"context"
	"log"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
	"github.com/jinzhu/gorm"
)

func QueryForemArticlesForever(s *state.State) {
	ctx := context.Background()
	al, err := s.ForemClient.Articles(ctx, gophorem.Defaults())
	if err != nil {
		log.Printf("error querying dev.to: %+v", err)
	}

	for _, a := range al {
		p := a.Article
		p.TagList = a.TagList
		p.Tags = a.Tags

		id, err := AddArticle(ctx, s, p)
		if err != nil {
			log.Printf("error saving article: %+v", err)
			continue
		}
		log.Println("saving article with id ", id)
	}
}

func AddArticle(ctx context.Context, s *state.State, article gophorem.Article) (uint, error) {
	a, err := db.LookupByDevToID(s.DB, article.ID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return 0, err
	} else if err != nil && gorm.IsRecordNotFoundError(err) {
		return db.AddArticle(s.DB, article)
	}

	// Found article, so can just return the id.
	return a.ID, nil
}
