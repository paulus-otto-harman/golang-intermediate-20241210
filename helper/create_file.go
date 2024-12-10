package helper

import "os"

// Fungsi untuk membuat folder jika belum ada
func EnsureFolder(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return os.MkdirAll(folderPath, os.ModePerm)
	}
	return nil
}
