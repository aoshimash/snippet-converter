package util

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

const fixturesDir = "../../data/fixtures"

func TestGetFilePathes(t *testing.T) {
	testCases := []struct {
		name        string
		dir         string
		expectedErr error
	}{
		{
			name:        "success",
			dir:         fixturesDir + "/cpp",
			expectedErr: nil,
		},
		{
			name:        "no such file or direcotry",
			dir:         "./wrong_directory",
			expectedErr: fs.ErrNotExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GetFilePathes(tc.dir)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
