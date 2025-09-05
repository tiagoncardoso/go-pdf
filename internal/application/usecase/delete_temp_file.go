package usecase

import (
	"os"
	"path/filepath"
)

type DeleteTempFile struct{}

func NewDeleteTempFile() *DeleteTempFile {
	return &DeleteTempFile{}
}

func (d *DeleteTempFile) Execute(fileName string) error {
	file := filepath.Join(fileName)

	if err := os.Remove(file); err != nil {
		return err
	}

	return nil
}
