package db

import (
	"encoding/binary"

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

//AddTask inserts a new task into the database
func AddTask(task string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bkt)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)
		return bucket.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}
	return id, nil
}

//ListTasks returns a slice of Task struct representing the
//incomplete tasks
func ListTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bkt)
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				btoi(k),
				string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

//DeleteTask removes completed task from the database
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bkt)
		return bucket.Delete(itob(key))
	})
}

//converts integer to byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//converst byte slice to integer
func btoi(b []byte) int {
	val := binary.BigEndian.Uint64(b)
	return int(val)
}
