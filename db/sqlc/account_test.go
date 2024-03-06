package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"simple-bank/util"
	"testing"
	"time"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	actualAccount, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)

	require.Equal(t, actualAccount.Owner, account.Owner)
	require.Equal(t, actualAccount.Balance, account.Balance)
	require.Equal(t, actualAccount.Currency, account.Currency)

	require.NotZero(t, actualAccount.ID)
	require.NotZero(t, actualAccount.CreatedAt)
	require.WithinDuration(t, actualAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	actualAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)

	require.Equal(t, actualAccount.Owner, account.Owner)
	require.Equal(t, actualAccount.Balance, arg.Balance)
	require.Equal(t, actualAccount.Currency, account.Currency)
	require.WithinDuration(t, actualAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	actualAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Empty(t, actualAccount)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}

	actualAccounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, actualAccounts, 5)
	for _, account := range actualAccounts {
		require.NotEmpty(t, account)
	}
}
