package pkg

import (
	"encoding/json"
	"go.etcd.io/bbolt"
)

const mainBucket = "main"

type database struct {
	db *bbolt.DB
}

func (d *database) GetObject(key string, out any) error {
	return d.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(mainBucket))
		if b == nil {
			return nil
		}
		val := b.Get([]byte(key))
		if val == nil {
			return nil
		}
		err := json.Unmarshal(val, out)
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *database) SaveObject(key string, data interface{}) error {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return d.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(mainBucket))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), dataJson)
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *database) Close() error {
	return d.db.Close()
}

type Database interface {
	Close() error
	GetObject(key string, out interface{}) error
	SaveObject(key string, data interface{}) error
}

func NewDB(path string) (Database, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &database{
		db: db,
	}, nil
}
