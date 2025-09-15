package usecase

import (
	"os"
	"path/filepath"

	"github.com/tiagoncardoso/go-pdf/config"
)

type DeleteTempFile struct {
	env *config.EnvConfig
}

func NewDeleteTempFile(env *config.EnvConfig) *DeleteTempFile {
	return &DeleteTempFile{
		env: env,
	}
}

func (d *DeleteTempFile) Execute(fileName string) error {
	file := filepath.Join(d.env.OutputPath, fileName)

	if err := os.Remove(file); err != nil {
		return err
	}

	return nil
}
