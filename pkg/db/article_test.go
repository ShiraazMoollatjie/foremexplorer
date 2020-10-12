package db

import (
	"testing"

	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
	"github.com/stretchr/testify/require"
)

func TestCreatePriceLog(t *testing.T) {
	dbc, err := connectForTesting(t)
	require.NoError(t, err)
	defer dbc.DropTableIfExists(&Article{})
	require.NoError(t, dbc.AutoMigrate(&Article{}).Error)

	plID, err := AddArticle(dbc, gophorem.Article{
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
	require.NotZero(t, plID)
}
