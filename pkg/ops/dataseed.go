package ops

import (
	"context"
	"log"
	"strconv"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
)

// SeedDevData is a once off operation that will seed all the existing data into the database.
func SeedDevData(s *state.State) {
	pageSize := 1000
	ctx := context.Background()

	pg := 1
	for {

		al, err := s.ForemClient.Articles(ctx, gophorem.Arguments{
			"page":     strconv.Itoa(pg),
			"per_page": strconv.Itoa(pageSize),
		})

		if err != nil {
			log.Printf("cannot fetch forem articles: %v", err)
		}

		for _, a := range al {
			at := a.Article
			at.TagList = a.TagList
			at.Tags = a.Tags

			AddArticle(ctx, s, at)
		}

		// Assume all pages equal 1000 until you reach the end.
		if len(al) < pageSize {
			break
		}
		pg++
		break
	}

}