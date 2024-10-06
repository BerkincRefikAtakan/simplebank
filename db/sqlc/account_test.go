package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func radomCreateAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		owner:    util.RandomOwner(),
		balance:  util.RandomMoney(),
		currency: util.RandomCurrency(),
	}

	account, err := testQueries.createAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.owner, account.owner)
	require.Equal(t, arg.balance, account.balance)
	require.Equal(t, arg.currency, account.currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account

}

func testCreateAccount(t *testing.T) {
	randomCreateAccount(t)
}

func testGetaccount(t *testing.T) {
	account1 := randomCreateAccount(t)

	account2, err := testQueries.getAccount(context.Background(), account1)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.owner, account2.owner)
	require.Equal(t, account1.balance, account2.balance)
	require.Equal(t, account1.currency, account2.currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}
