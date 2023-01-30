package db

import (
	"context"
	"database/sql"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)

	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, args.FromAccountID)
	require.Equal(t, transfer.ToAccountID, args.ToAccountID)
	require.Equal(t, transfer.Amount, args.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)

	transfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, createdTransfer.ID, transfer.ID)
	require.Equal(t, createdTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, createdTransfer.Amount, transfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)

	account := createRandomAccount(t)

	args := UpdateTransferParams{
		ID:            createdTransfer.ID,
		FromAccountID: account.ID,
		ToAccountID:   createdTransfer.ToAccountID,
		Amount:        666,
	}

	err := testQueries.UpdateTransfer(context.Background(), args)

	require.NoError(t, err)

	transfer, _ := testQueries.GetTransfer(context.Background(), args.ID)

	require.Equal(t, transfer.Amount, args.Amount)
	require.Equal(t, transfer.FromAccountID, args.FromAccountID)
	require.Equal(t, transfer.ToAccountID, createdTransfer.ToAccountID)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)

	deletedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedTransfer)
}

func TestGetTransferList(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	args := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfersList, err := testQueries.ListTransfers(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfersList)

	require.Len(t, transfersList, 5)
}
