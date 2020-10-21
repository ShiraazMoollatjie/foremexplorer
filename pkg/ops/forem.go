package ops

import (
	"context"
	"log"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
)

func QueryForemArticlesForever(s *state.State) {
	ctx := context.Background()
	al, err := s.ForemClient.Articles(ctx, gophorem.Arguments{
		"per_page": "1000",
	})
	if err != nil {
		log.Printf("error querying forem site: %+v", err)
	}

	for _, a := range al {
		p := a.Article
		p.TagList = a.TagList
		p.Tags = a.Tags

		_, err := AddArticle(ctx, s, p)
		if err != nil {
			log.Printf("error saving article: %+v", err)
			continue
		}
	}
}
