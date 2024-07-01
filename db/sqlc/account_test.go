package dbbank

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/k2madureira/gobank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:       util.RandomOwner(),
		Balance:     util.RandomMoney(),
		Currency:    util.RandomCurrency(),
		CountryCode: sql.NullInt32{Int32: int32(util.RandomInt(10, 60)), Valid: true},
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreateAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	findAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, findAccount)

	require.Equal(t, account.ID, findAccount.ID)
	require.Equal(t, account.Owner, findAccount.Owner)
	require.Equal(t, account.Currency, findAccount.Currency)
	require.Equal(t, account.Balance, findAccount.Balance)
	require.Equal(t, account.CountryCode, findAccount.CountryCode)
	require.WithinDuration(t, account.CreateAt, findAccount.CreateAt, time.Second)
}
