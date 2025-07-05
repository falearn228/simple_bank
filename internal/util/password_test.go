package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(10)

	hshPassw1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hshPassw1)

	err = CheckPassword(password, hshPassw1)
	require.NoError(t, err)

	wrongPassw := RandomString(10)
	err = CheckPassword(wrongPassw, hshPassw1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hshPassw2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hshPassw2)

	require.NotEqual(t, hshPassw1, hshPassw2, "two hashes for the same password should be not equal")
}
