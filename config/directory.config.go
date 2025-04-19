package config

import (
	"fmt"
	"os"
)

const (
	UploadDirectory = "./uploads"
)

func CreateDirectory() error {
	fileInfo, err := os.Stat(UploadDirectory)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(UploadDirectory, 0755); err != nil {
			return fmt.Errorf("error creando directorio: %v", err)
		}
		return nil
	}

	if err != nil {
		return fmt.Errorf("error verificando directorio: %v", err)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("la ruta %s existe pero no es un directorio", UploadDirectory)
	}

	if fileInfo.Mode().Perm()&0200 == 0 {
		return fmt.Errorf("no hay permisos de escritura en el directorio %s", UploadDirectory)
	}

	return nil
}
