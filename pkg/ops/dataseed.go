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

		err = importPages(ctx, s, al)
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

// SeedDevDataIncrementally is a once off operation that will seed all the existing data into the database.
func SeedDevDataIncrementally(s *state.State) {
	ctx := context.Background()
	const limit = 492198

	for i := int32(1); i <= limit; i++ {
		log.Printf("importing article %d", i)
		a, err := s.ForemClient.PublishedArticle(ctx, i)

		if err != nil {
			log.Printf("cannot fetch forem article: %v", err)
			continue // skip
		}

		err = importArticle(ctx, s, a)
		if err != nil {
			log.Printf("cannot import forem article: %v", err)
			continue // skip
		}

		time.Sleep(200 * time.Millisecond) // Rate limit at most 5 requests per second.
	}
}

func importArticle(ctx context.Context, s *state.State, a *gophorem.Article) error {
	t0 := time.Now()
	defer func() { log.Printf("time taken to import article %s", time.Since(t0)) }()

	_, err := AddArticle(ctx, s, *a)
	if err != nil {
		return err
	}

	return nil
}

func importPages(ctx context.Context, s *state.State, al gophorem.Articles) error {
	t0 := time.Now()
	defer func() { log.Printf("time taken to import page %s", time.Since(t0)) }()

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
