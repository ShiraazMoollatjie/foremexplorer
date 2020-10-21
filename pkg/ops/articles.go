package ops

import (
	"context"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
	"github.com/jinzhu/gorm"
)

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
