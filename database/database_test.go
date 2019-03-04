package database

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/voteright/voteright/config"
)

func TestNew(t *testing.T) {
	d, err := New(&config.Config{
		DatabaseFile: "./test.db",
	})
	if err != nil {
		fmt.Println("error connecting to test db")
		t.FailNow()
	}
	_, err = d.ExecStatement("DROP TABLE IF EXISTS mytest")
	if err != nil {
		fmt.Println("error executing command 1 on test db")
		t.FailNow()
	}
	_, err = d.ExecStatement("CREATE TABLE mytest (value INTEGER)")
	if err != nil {
		fmt.Println("error executing command 2 on test db")
		t.FailNow()
	}
	_, err = d.ExecStatement("INSERT INTO mytest VALUES (1)")
	if err != nil {
		fmt.Println("error executing command 3 on test db")
		t.FailNow()
	}
}
