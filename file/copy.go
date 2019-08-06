package file

import (
	"errors"
	"os"
)

func Copy(fromPath string, toPath string, offset int64, limit int) (bool, error) {
	soughtFile, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, errors.New("file not found")
		}

		return false, err
	}
	defer soughtFile.Close()

	fileInfo, err := soughtFile.Stat()
	if err != nil {
		return false, err
	}

	if int64(limit) > fileInfo.Size() {
		limit = int(fileInfo.Size())
	}

	buffer := make([]byte, limit)
	_, err = soughtFile.ReadAt(buffer, offset)
	if err != nil {
		return false, err
	}

	newFile, err := os.Create(toPath)
	if err != nil {
		return false, err
	}
	defer newFile.Close()

	_, err = newFile.Write(buffer)
	if err != nil {
		return false, err
	}

	return true, nil
}
