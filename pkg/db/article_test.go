package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePriceLog(t *testing.T) {
	dbc, err := connectForTesting(t)
	require.NoError(t, err)
	defer dbc.DropTableIfExists(&Article{})
	require.NoError(t, dbc.AutoMigrate(&Article{}).Error)

	plID, err := AddArticle(dbc)
	require.NoError(t, err)
	require.NotZero(t, plID)
}
