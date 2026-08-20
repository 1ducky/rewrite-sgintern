package localstorage_test

import (
	"RewriteProject/internal/assets"
	"RewriteProject/internal/config"
	"RewriteProject/internal/localstorage"
	"io"
	"path"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/google/uuid"
)

type UsecaseContract string

const (
	Write  UsecaseContract = "WRITE"
	Delete UsecaseContract = "DELETE"
	Read   UsecaseContract = "READ"
)

type WriteTestCase struct {
	Name        string
	Usecase     UsecaseContract
	Expeted     bool
	Destination string
	Reader      io.Reader
}

func createWriteTestUsecase(conf config.StorageConfig) []WriteTestCase {
	return []WriteTestCase{
		{Name: "VALID_WRITE_TEST", Usecase: Write, Expeted: true, Destination: path.Join(conf.StoragePathTemp, uuid.NewString()+"-"+".txt"), Reader: strings.NewReader("hello world")},
		{Name: "EMPTY_CONTENT_WRITE_TEST", Usecase: Write, Expeted: false, Destination: path.Join(conf.StoragePathTemp, uuid.NewString()+"-"+".txt"), Reader: strings.NewReader("")},
		{Name: "EMPTY_DESTINATION_WRITE_TEST", Usecase: Write, Expeted: false, Destination: "", Reader: strings.NewReader("hello world")},
		{Name: "INVALID_DESTINATION_WRITE_TEST", Usecase: Write, Expeted: false, Destination: path.Join("adas", uuid.NewString()+"-"+".txt"), Reader: strings.NewReader("hello world")},
		{Name: "ERROR_READER_TEST", Usecase: Write, Expeted: false, Destination: path.Join(conf.StoragePathTemp, uuid.NewString()+"-"+".txt"), Reader: iotest.ErrReader(assets.ErrAssetFailedCreate)},
		{Name: "MULTIPLE_INVALID_DESTINATION_WRITE_TEST", Usecase: Write, Expeted: false, Destination: path.Join("adas", "asdsad", uuid.NewString()+"-"+".txt"), Reader: strings.NewReader("hello world")},
	}
}

func newLocalStorage(t *testing.T, conf config.StorageConfig) assets.RepositoryContract {
	storge, err := localstorage.NewLocalStorage(conf)

	if err != nil {
		t.Fatal("Storage Failed Init")
	}
	return storge

}

func createExpectedDir(conf config.StorageConfig) []string {
	return []string{
		path.Join(conf.StorageRoot, conf.StoragePathTemp),
		path.Join(conf.StorageRoot, conf.StoragePathUpload),
		path.Join(conf.StorageRoot, conf.StoragePathDocument),
		path.Join(conf.StorageRoot, conf.StoragePathAvatar),
		path.Join(conf.StorageRoot, conf.StoragePathVideo),
	}
}

func TestWriteReadDelete(t *testing.T) {
	conf := config.StorageConfig{
		StorageRoot:         t.TempDir(),
		StoragePathTemp:     "temp",
		StoragePathUpload:   "upload",
		StoragePathDocument: "document",
		StoragePathAvatar:   "avatar",
		StoragePathVideo:    "video",
	}
	localStorage := newLocalStorage(t, conf)
	uc := createWriteTestUsecase(conf)

	for _, u := range uc {
		t.Run(u.Name+"_"+string(Write), func(t *testing.T) {
			_, writeErr := localStorage.Write(t.Context(), u.Reader, u.Destination)
			gotOK := writeErr == nil
			if gotOK != u.Expeted {
				t.Fatalf("Write: expected ok=%v, got err=%v", u.Expeted, writeErr)
			}
		})

		t.Run(u.Name+"_"+string(Read), func(t *testing.T) {
			f, readErr := localStorage.Read(t.Context(), u.Destination)
			gotOK := readErr == nil
			if gotOK != u.Expeted {
				t.Fatalf("Read: expected ok=%v, got err=%v", u.Expeted, readErr)
			}
			if f != nil {
				defer f.Close()
			}
		})
		t.Run(u.Name+"_"+string(Delete), func(t *testing.T) {
			delErr := localStorage.Delete(t.Context(), u.Destination)
			gotOK := delErr == nil
			if gotOK != u.Expeted {
				t.Fatalf("Delete: expected ok=%v, got err=%v", u.Expeted, delErr)
			}
		})
	}

}
