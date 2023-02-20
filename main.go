package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jamf/go-mysqldump"
)

func main() {
	config := mysql.NewConfig()

	config.User = os.Getenv("DB_USER")
	config.DBName = os.Getenv("DB_NAME")
	config.Passwd = os.Getenv("DB_PASS")
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%s", os.Getenv("DB_URL"), os.Getenv("DB_PORT"))

	currentTime := time.Now()
	dumpDir := "/opt/backup/dumps"

	dumpFilenameFormat := fmt.Sprintf("%d-%d-%d", currentTime.Year(), currentTime.Day(), currentTime.Hour())

	fmt.Printf(dumpDir, dumpFilenameFormat)

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println("Error opening database", err)
		return
	}

	//Register db
	dumper, err := mysqldump.Register(db, dumpDir, dumpFilenameFormat)
	if err != nil {
		fmt.Println("Error registering database", err)
		return
	}

	// dump database to a file
	err := dumper.Dump()
	if err != nil {
		fmt.Println(/golang.org/x/tools/internal/typesinternal"error dumping,", err)
		return
	}
	fmt.Printf("file is saved to %s", dumpFilenameFormat)
	// close dumper
	dumper.Close()
}
