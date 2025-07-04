package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"bobbabank/internal/util"

	"github.com/stretchr/testify/require"
)

// Valid - значение не является null
func createRandomTransfer(t *testing.T, account1 Account, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: sql.NullInt64{
			Int64: account1.ID,
			Valid: true,
		},
		ToAccountID: sql.NullInt64{
			Int64: account2.ID,
			Valid: true,
		},
		Amount: util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID.Int64, transfer.FromAccountID.Int64)
	require.Equal(t, arg.ToAccountID.Int64, transfer.ToAccountID.Int64)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {

		createRandomTransfer(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountID: sql.NullInt64{
			Int64: account1.ID,
			Valid: true,
		},
		ToAccountID: sql.NullInt64{
			Int64: account2.ID,
			Valid: true,
		},
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID.Int64 == account1.ID || transfer.ToAccountID.Int64 == account1.ID)
	}
}
