package file

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

// Copies files
func Copy(fromPath string, toPath string, offset int64, limit int) error {
	soughtFile, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("file %s not found", fromPath))
		}

		return err
	}
	defer soughtFile.Close()

	fileInfo, err := soughtFile.Stat()
	if err != nil {
		return err
	}

	if int64(limit) > fileInfo.Size() {
		limit = int(fileInfo.Size())
	}

	_, err = soughtFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	newLimit := int64(limit)
	bar := pb.Full.Start64(newLimit - offset)
	barReader := bar.NewProxyReader(soughtFile)
	defer bar.Finish()

	_, err = io.CopyN(newFile, barReader, newLimit)
	if err == nil || err == io.EOF {
		return nil
	}

	return err
}
