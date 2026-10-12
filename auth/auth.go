package auth

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/crypto/bcrypt"
)

func hashFilePath() string {
	_, thisFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(thisFile))
	return filepath.Join(projectRoot, "database", "auth.hash")
}

func IsConfigured() bool {
	_, err := os.Stat(hashFilePath())
	return err == nil
}

func SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	path := hashFilePath()
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	return os.WriteFile(path, hash, 0600)
}

func Check(password string) (bool, error) {
	hash, err := os.ReadFile(hashFilePath())
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
		return false, nil
	default:
		return false, err
	}
}
