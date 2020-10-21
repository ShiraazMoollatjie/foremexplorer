package db

import (
	"testing"

	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
	"github.com/stretchr/testify/require"
)

func TestCreateArticle(t *testing.T) {
	dbc, err := connectForTesting(t)
	require.NoError(t, err)
	defer dbc.DropTableIfExists(&Article{})
	require.NoError(t, dbc.AutoMigrate(&Article{}).Error)

	aID, err := AddArticle(dbc, gophorem.Article{
		TypeOf:               "something",
		ID:                   123,
		Title:                "tlaksjd",
		Description:          "alsjkd",
		ReadablePublishDate:  "asdkj",
		TagList:              []string{"sde", "asd"},
		Tags:                 "askdlj, asdlkj, asda",
		URL:                  "url",
		CommentsCount:        1,
		PublicReactionsCount: 1,
	})
	require.NoError(t, err)
	require.NotZero(t, aID)

	res, err := LookupByDevToID(dbc, 123)
	require.NoError(t, err)
	require.Equal(t, aID, res.ID)
}

func TestListAllArticles(t *testing.T) {
	dbc, err := connectForTesting(t)
	require.NoError(t, err)
	defer dbc.DropTableIfExists(&Article{})
	require.NoError(t, dbc.AutoMigrate(&Article{}).Error)

	aID, err := AddArticle(dbc, gophorem.Article{
		TypeOf:               "something",
		ID:                   123,
		Title:                "tlaksjd",
		Description:          "alsjkd",
		ReadablePublishDate:  "asdkj",
		TagList:              []string{"sde", "asd"},
		Tags:                 "askdlj, asdlkj, asda",
		URL:                  "url",
		CommentsCount:        1,
		PublicReactionsCount: 1,
	})
	require.NoError(t, err)
	require.NotZero(t, aID)

	res, err := ListArticles(dbc)
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
}
