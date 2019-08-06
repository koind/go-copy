package file

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

// Copies files
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

	_, err = soughtFile.Seek(offset, io.SeekStart)
	if err != nil {
		return false, err
	}

	newFile, err := os.Create(toPath)
	if err != nil {
		return false, err
	}
	defer newFile.Close()

	newLimit := int64(limit)
	bar := pb.Full.Start64(newLimit - offset)
	barReader := bar.NewProxyReader(soughtFile)

	_, err = io.CopyN(newFile, barReader, newLimit)
	if err == nil || err == io.EOF {
		bar.Finish()

		return true, nil
	}

	bar.Finish()

	return false, err
}
