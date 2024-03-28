package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simple-bank/util"
	"testing"
	"time"
)

func CreateRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func CreateRandomEntryForAccount(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := CreateRandomEntry(t)
	actual, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actual)

	require.Equal(t, actual.AccountID, entry.AccountID)
	require.Equal(t, actual.Amount, entry.Amount)

	require.NotZero(t, actual.ID)
	require.NotZero(t, actual.CreatedAt)
	require.WithinDuration(t, entry.CreatedAt, actual.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	account := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	for i := 0; i < 3; i++ {
		CreateRandomEntryForAccount(t, account)
	}

	arg := ListEntriesParams{
		Limit:     5,
		Offset:    0,
		AccountID: account.ID,
	}

	actual, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, actual, 3)
	for _, entry := range actual {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, account.ID)
	}
}
