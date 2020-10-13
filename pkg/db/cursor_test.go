package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCursor(t *testing.T) {
	dbc, err := connectForTesting(t)
	require.NoError(t, err)
	defer dbc.DropTableIfExists(&Cursor{})
	require.NoError(t, dbc.AutoMigrate(&Cursor{}).Error)

	_, err = GetCursor(dbc, "page_number")
	require.Error(t, err)

	err = SetCursor(dbc, "page_number", 1)
	require.NoError(t, err)

	res, err := GetCursor(dbc, "page_number")
	require.NoError(t, err)
	require.Equal(t, 1, res.Value)

	err = SetCursor(dbc, "page_number", 2)
	require.NoError(t, err)

	res, err = GetCursor(dbc, "page_number")
	require.NoError(t, err)
	require.Equal(t, 2, res.Value)
}
