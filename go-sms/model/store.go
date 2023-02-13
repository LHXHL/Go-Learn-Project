package model

import (
	"go.etcd.io/bbolt"
	"time"
)

type Store interface {
	Set(key string, value []byte) error
	Get(Key string) ([]byte, error)
	Del(key string) error
	Expire(key string, seconds int) error
	Len() (int, error)
	Close() error
	FLush() error
}

type DB struct {
	db *bbolt.DB
}

func (D *DB) Set(key string, value []byte) error {
	return D.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tableName)
		return b.Put([]byte(key), value)
	})
}

func (D *DB) Get(Key string) ([]byte, error) {
	var value []byte
	err := D.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tableName)
		value = b.Get([]byte(Key))
		return nil
	})
	return value, err
}

func (D *DB) Del(key string) error {
	return D.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tableName)
		return b.Delete([]byte(key))
	})
}

func (D *DB) Expire(key string, seconds int) error {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	go func() {
		<-timer.C
		D.Del(key)
	}()
	return nil
}

func (D *DB) Len() (int, error) {
	var count int
	err := D.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tableName)
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			count++
		}
		return nil
	})
	return count, err
}

func (D *DB) Close() error {
	return D.db.Close()
}

func (D *DB) FLush() error {
	return D.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tableName)
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			b.Delete(k)
		}
		return nil
	})
}

var tableName = []byte("sms")
var DbNow *DB

func InitDB(dbName string) {
	DbNow = &DB{
		db: openDatabase(dbName),
	}
}

func openDatabase(dbName string) *bbolt.DB {
	db, err := bbolt.Open(dbName, 0600, nil)
	if err != nil {
		return nil
	}

	tables := [...][]byte{
		tableName,
	}

	db.Update(func(tx *bbolt.Tx) error {
		for _, table := range tables {
			_, err2 := tx.CreateBucketIfNotExists(table)
			if err2 != nil {
				return err2
			}
		}
		return nil
	})
	return db
}

// var _ Store = (*DB)(nil) 可以一键生成方法
