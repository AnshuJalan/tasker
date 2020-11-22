package db

import (
	"github.com/boltdb/bolt"
)

var bkt = []byte("allTasks")

var db *bolt.DB

//Task defines an item on the TODO list
type Task struct {
	Key   int
	Value string
}

//Init opens up a connection the database
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bkt)
		return err
	})
}

// //AddTask adds a new task to the boltDB database
// func AddTask(task string) error {
// 	db, err := bolt.Open("taskDB.db", 0600, nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	//add task
// 	err = db.Update(func(tx *bolt.Tx) error {
// 		bucket, err := tx.CreateBucketIfNotExists(bkt)
// 		if err != nil {
// 			return err
// 		}
// 		id := bucket.Stats().KeyN
// 		return bucket.Put([]byte(strconv.Itoa(id+1)), []byte(task))
// 	})

// 	return err
// }

// //ListTasks lists all incomplete tasks
// func ListTasks() []string {
// 	db, _ := bolt.Open("taskDB.db", 0600, nil)
// 	defer db.Close()

// 	var allTasks []string

// 	db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket(bkt)
// 		c := b.Cursor()

// 		for k, v := c.First(); k != nil; k, v = c.Next() {
// 			allTasks = append(allTasks, string(v))
// 		}
// 		return nil
// 	})

// 	return allTasks
// }
