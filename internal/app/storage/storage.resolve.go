package storage

import (
	"RewriteProject/internal/app/assets"
	"path/filepath"
)

// Resolve Path Option end Return valid path
func (l *LocalStorage) resolvePath(dir string) (string, error) {
	switch dir {
	case "upload":
		return filepath.Join(l.conf.StoragePathUpload), nil
	case "temp":
		return filepath.Join(l.conf.StoragePathTemp), nil
	case "avatar":
		return filepath.Join(l.conf.StoragePathAvatar), nil
	case "document":
		return filepath.Join(l.conf.StoragePathDocument), nil
	default:
		return "", assets.ErrStorageUnavailable
	}
}
