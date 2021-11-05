package db

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

// type shortURL struct {
// 	Path string `yaml: path`
// 	Url  string `yaml: url`
// }
const dbLocation string = "urlshortener.db"
const dbName string = "urlDB"

// func main() {
// 	//db, err := setupDB()
// 	db, err := db()
// 	if err != nil {
// 		fmt.Printf("could not setup db, %v", err)
// 	}
// insertURL(db, "rickroll", "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley")
// insertURL(db, "valorant", "https://www.youtube.com/watch?v=FUoqAn5T4h4&ab_channel=VALORANT")

// 	view(db)
// 	s := string(viewKey(db, "rickroll"))
// 	if s != "" {
// 		fmt.Println(s)
// 	} else {
// 		fmt.Println("No url found")
// 	}

// 	db.Close()
// }

func DB() (*bolt.DB, error) {
	db, err := bolt.Open(dbLocation, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not create connection: %v", err)
	}
	return db, err
}

//setupDB sets up a db if it does not exist
func SetupDB() error {
	db, err := bolt.Open(dbLocation, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return fmt.Errorf("could not open DB: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(dbName))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not set up bucket, %v", err)
	}
	InsertURL(db, "rickroll", "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley")
	fmt.Println("DB setup complete")
	db.Close()
	return nil
}

//insertURL adds a new path and url in the database
func InsertURL(db *bolt.DB, path, url string) error {
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

//view prints all the paths and urls stored in the db
func View(db *bolt.DB) error {
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

//viewKey returns the url stores on given key
func ViewKey(db *bolt.DB, key string) []byte {
	var url []byte
	db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(dbName))
		url = data.Get([]byte(key))
		return nil
	})
	return url
}
