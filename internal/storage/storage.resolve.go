package storage

import (
	"RewriteProject/internal/assets"
	"path/filepath"
)

// Resolve Path Option end Return valid path
func (l *LocalStorage) resolvePath(dir string, fileName string) (string, error) {
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
		return "", assets.ErrStorageUnavailable
	}
}
