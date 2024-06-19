package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simple-bank/util"
	"testing"
	"time"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)
	actualUser, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, actualUser)

	require.Equal(t, actualUser.Username, user.Username)
	require.Equal(t, actualUser.HashedPassword, user.HashedPassword)
	require.Equal(t, actualUser.FullName, user.FullName)
	require.Equal(t, actualUser.Email, user.Email)

	require.True(t, actualUser.PasswordChangedAt.IsZero())
	require.NotZero(t, actualUser.CreatedAt)
	require.WithinDuration(t, actualUser.CreatedAt, actualUser.CreatedAt, time.Second)
}
