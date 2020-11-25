package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var mainBkt = []byte("allTasks")
var doneBkt = []byte("completedTasks")

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
		_, err = tx.CreateBucketIfNotExists(mainBkt)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(doneBkt)
		return err
	})
}

//AddTask inserts a new task into the database
func AddTask(task string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(mainBkt)
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
		bucket := tx.Bucket(mainBkt)
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

//RemoveTask removes completed task from the database
func RemoveTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(mainBkt)
		if bucket.Get(itob(key)) == nil {
			return fmt.Errorf("task with id %d does not exist", key)
		}
		return bucket.Delete(itob(key))
	})
}

//DoTask marks a task as complete and puts it
//in the done bucket of the database
func DoTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		mainBucket := tx.Bucket(mainBkt)
		doneBucket := tx.Bucket(doneBkt)
		if mainBucket.Get(itob(key)) == nil {
			return fmt.Errorf("task with id %d does not exist", key)
		}
		doneBucket.Put(mainBucket.Get(itob(key)), timetob(time.Now()))
		return mainBucket.Delete(itob(key))
	})
}

//Size returns the number of entries in the main bucket
func Size() int {
	var size int
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(mainBkt)
		size = bucket.Stats().KeyN
		return nil
	})
	return size
}

//converts integer to byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//converts byte slice to integer
func btoi(b []byte) int {
	val := binary.BigEndian.Uint64(b)
	return int(val)
}

//encodes time to byte array
func timetob(t time.Time) []byte {
	val, _ := t.GobEncode()
	return val
}
