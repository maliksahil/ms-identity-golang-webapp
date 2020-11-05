package app

import (
	"encoding/gob"
	"log"
	"math"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	Store = sessions.NewFilesystemStore(os.TempDir(), []byte("secret"))
	Store.MaxLength(math.MaxInt64)
	gob.Register(map[string]interface{}{})
	return nil
}
