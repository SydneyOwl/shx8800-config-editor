package filetools

import (
	_ "embed"
	"errors"
	"os"
)

//go:embed SHX8800
var SHX8800_executable []byte

func ReleaseFile() error {
	targetPath := GetCSPath()
	if targetPath == "" {
		return errors.New("")
	} else {
		handler, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		_, err = handler.Write(SHX8800_executable)
		if err != nil {
			return err
		}
		return nil
	}
}
