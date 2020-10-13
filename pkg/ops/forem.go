package ops

import (
	"context"
	"log"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
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

		id, err := db.AddArticle(s.DB, p)
		if err != nil {
			log.Printf("error saving article: %v", err, a)
			continue
		}
		log.Println("saving article with id ", id)
	}
}
