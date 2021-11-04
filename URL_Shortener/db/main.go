package main

import (
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

// type shortURL struct {
// 	Path string `yaml: path`
// 	Url  string `yaml: url`
// }
const dbLocation string = "urlshortener.db"
const dbName string = "urlDB"

func main() {
	//db, err := setupDB()
	db, err := db()
	if err != nil {
		fmt.Printf("could not setup db, %v", err)
	}
	// insertURL(db, "rickroll", "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley")
	// insertURL(db, "valorant", "https://www.youtube.com/watch?v=FUoqAn5T4h4&ab_channel=VALORANT")

	view(db)
	s := string(viewKey(db, "rickrolasdl"))
	if s != "" {
		fmt.Println(s)
	} else {
		fmt.Println("No url found")
	}

	db.Close()
}

func db() (*bolt.DB, error) {
	db, err := bolt.Open(dbLocation, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not create connection: %v", err)
	}
	fmt.Println("connected to db")
	return db, err
}

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open(dbLocation, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(dbName))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up bucket, %v", err)
	}
	fmt.Println("DB setup complete")
	return db, nil
}

func insertURL(db *bolt.DB, path, url string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(dbName)).Put([]byte(path), []byte(url))
		if err != nil {
			return fmt.Errorf("could not insert shortURL: %v", err)
		}
		return nil
	})
	fmt.Println("shortURL added")
	return err
}

func view(db *bolt.DB) error {
	err := db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(dbName))
		data.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	return err
}

func viewKey(db *bolt.DB, key string) []byte {
	var url []byte
	db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(dbName))
		url = data.Get([]byte(key))
		return nil
	})
	return url
}
