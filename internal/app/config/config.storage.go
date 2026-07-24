package config

type StorageConfig struct {
	StorageRoot         string
	StoragePathUpload   string
	StoragePathTemp     string
	StoragePathAvatar   string
	StoragePathDocument string
}

func LoadStorageConfig() StorageConfig {
	return StorageConfig{
		StorageRoot:         getEnv("STORAGE_ROOT", "./storage"),
		StoragePathUpload:   getEnv("STORAGE_PATH_UPLOAD", "uploads"),
		StoragePathTemp:     getEnv("STORAGE_PATH_TEMP", "tmp"),
		StoragePathAvatar:   getEnv("STORAGE_PATH_AVATAR", "avatars"),
		StoragePathDocument: getEnv("STORAGE_PATH_DOCUMENT", "documents"),
	}
}
