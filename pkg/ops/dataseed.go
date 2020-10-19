package ops

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
)

// SeedDevData is a once off operation that will seed all the existing data into the database.
func SeedDevData(s *state.State) {
	pageSize := 1000
	ctx := context.Background()

	pg := 1
	for {
		log.Printf("importing page number %d", pg)
		al, err := s.ForemClient.Articles(ctx, gophorem.Arguments{
			"page":     strconv.Itoa(pg),
			"per_page": strconv.Itoa(pageSize),
		})

		if err != nil {
			log.Printf("cannot fetch forem articles: %v", err)
			continue // retry
		}

		err = importPage(ctx, s, al)
		if err != nil {
			log.Printf("cannot import forem articles: %v", err)
			continue // retry
		}

		// Assume all pages equal 1000 until you reach the end.
		if len(al) < pageSize {
			break
		}
		pg++
	}
}

func importPage(ctx context.Context, s *state.State, al gophorem.Articles) error {
	t0 := time.Now()
	defer func() { log.Printf("time taken to import page %s seconds", time.Since(t0)) }()

	for _, a := range al {
		at := a.Article
		at.TagList = a.TagList
		at.Tags = a.Tags

		_, err := AddArticle(ctx, s, at)
		if err != nil {
			return err
		}
	}

	return nil
}
