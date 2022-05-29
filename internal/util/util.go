package util

import (
	"os"
	"path/filepath"
)

// ファイルパス一覧を取得する
func GetFilePathes(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			subdir_path := filepath.Join(dir, file.Name())
			subdir_file_paths, err := GetFilePathes(subdir_path)
			if err != nil {
				return nil, err
			}
			paths = append(paths, subdir_file_paths...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths, err
}
