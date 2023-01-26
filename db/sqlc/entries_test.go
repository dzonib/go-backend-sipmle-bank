package db

import (
	"context"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

//	type CreateEntryParams struct {
//		AccountID int64
//		Amount    int64
//	}
func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)

	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.Equal(t, entry.AccountID, entry1.AccountID)
	require.Equal(t, entry.Amount, entry1.Amount)
	require.Equal(t, entry.ID, entry1.ID)
	require.WithinDuration(t, entry.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	account := createRandomAccount(t)
	args := UpdateEntryParams{
		ID:        entry.ID,
		Amount:    666,
		AccountID: account.ID,
	}

	err := testQueries.UpdateEntry(context.Background(), args)

	require.NoError(t, err)

	updatedEntry, _ := testQueries.GetEntry(context.Background(), entry.ID)

	require.Equal(t, updatedEntry.Amount, args.Amount)
	require.Equal(t, updatedEntry.AccountID, args.AccountID)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	deletedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.Error(t, err)
	require.Empty(t, deletedEntry)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entryList, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entryList)
	require.Len(t, entryList, 5)
}
