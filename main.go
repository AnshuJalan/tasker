package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/anshujalan/tasker/cmd"
	"github.com/anshujalan/tasker/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	homePath, _ := homedir.Dir()
	dbPath := filepath.Join(homePath, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
