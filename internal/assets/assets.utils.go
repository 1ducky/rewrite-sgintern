package assets

import (
	"io"
	"path/filepath"
)

func (l *Usecase) resolvePath(dir AssetCategory, fileName string) (string, error) {
	switch dir {
	case "upload":
		return filepath.Join(l.conf.StoragePathUpload, fileName), nil
	case "temp":
		return filepath.Join(l.conf.StoragePathTemp, fileName), nil
	case "avatar":
		return filepath.Join(l.conf.StoragePathAvatar, fileName), nil
	case "document":
		return filepath.Join(l.conf.StoragePathDocument, fileName), nil
	default:
		return "", ErrStorageUnavailable
	}
}

func (l *Usecase) limitReader(reader io.Reader, maxSize int64) io.Reader {
	if maxSize == 0 {
		return reader
	}
	return io.LimitReader(reader, maxSize)
}
