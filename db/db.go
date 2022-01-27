package db

import (
	"fmt"
	"os"

	"github.com/etcd-io/bbolt"
	"github.com/pkg/errors"
)

const (
	relativeDBPath = "/timebox/db"
	dbFilename     = "timebox.db"
	configBucket   = "tbconfig"
)

func fullDBPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "getting home directory path")
	}

	return fmt.Sprintf("%s/%s", homeDir, relativeDBPath), nil
}

func fullDBFilePath() (string, error) {
	dbPath, err := fullDBPath()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", dbPath, dbFilename), nil
}

func openDB() (*bbolt.DB, error) {
	filePath, err := fullDBFilePath()
	if err != nil {
		return nil, err
	}

	return bbolt.Open(filePath, 0777, nil)
}

// Init will setup the config bucket.
func Init() error {
	filePath, err := fullDBPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filePath, 0777); err != nil {
		return errors.Wrap(err, "making db directory")
	}

	db, err := openDB()
	if err != nil {
		return errors.Wrap(err, "creating database")
	}
	defer db.Close()

	return db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(configBucket))
		if err != nil {
			return errors.Wrap(err, "creating config bucket")
		}

		err = b.Put([]byte("beep"), []byte("off"))
		return errors.Wrap(err, "setting beep config")
	})
}

// NewBucket will create a new database bucket with the given name.
func NewBucket(name string) error {
	var errMsg = "creating new bucket"
	db, err := openDB()
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	defer db.Close()

	return db.Update(func(tx *bbolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(name)); err != nil {
			return errors.Wrap(err, errMsg)
		}
		return err
	})
}
