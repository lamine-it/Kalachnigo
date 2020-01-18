package main

import(
	"fmt"
	"log"
	"bufio"
	"os"
	"time"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

	
func check(e error) {

	if e != nil {

		log.Fatal(e)
	}
}

func main() {

	var DatabaseName, tableName, user, pass, csvFile string

	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Database name: ")
  	sc.Scan()
	DatabaseName = sc.Text()

	fmt.Print("User: ")
	sc.Scan()
	user = sc.Text()

	fmt.Print("Password: ")
	sc.Scan()
	pass = sc.Text()

	fmt.Print("Table Name : ")
	sc.Scan()
	tableName = sc.Text()

	fmt.Print("Csv file : ")
	sc.Scan()
	csvFile = sc.Text()

	fmt.Println()
    fmt.Println("Start of operation ...")

	var start, end int64
	start = time.Now().UnixNano() // starting the counter

	// -----

	file, err := os.Open(csvFile) // the working file
	
	check(err)

	defer file.Close()

	// -----

	sqlqr := user + ":" + pass + "@tcp(127.0.0.1:3306)/" + DatabaseName

	db, err := sql.Open("mysql", sqlqr) // connection to the database

	check(err)

	defer db.Close()

	// -

	mysql.RegisterLocalFile("buffer.csv") // the buffer file

	// -----

	scanner := bufio.NewScanner(file)
	
	var content string

	i := 1 // index

	for scanner.Scan() {

		content = content + scanner.Text() + "\n"

		// reading 2000 lines
		if i == 2000 {

			f, err := os.Create("buffer.csv") // create the buffer
			check(err)

			_, err = f.WriteString(content) // write in the buffer
			check(err)
    	
    		f.Close()

    		_ ,err = db.Exec("LOAD DATA LOCAL INFILE 'buffer.csv' INTO TABLE " + tableName + " FIELDS TERMINATED BY '\"'") // write in database
    		check(err)

			content = "" // clean content
			i = 0 // restart index
		}
		
		i++ // increment index
	}

	// add the rest of contents to the database
	if content != "" {

		f, err := os.Create("buffer.csv")
		check(err)

		_, err = f.WriteString(content)
		check(err)

		_ ,err = db.Exec("LOAD DATA LOCAL INFILE 'buffer.csv' INTO TABLE " + tableName + " FIELDS TERMINATED BY '\" \"' LINES TERMINATED BY '\"\n' STARTING BY '\"'") // write in database
    	check(err)
	}

	// if no errors in scanner
	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}

	fmt.Println("Completed operation")

	end = time.Now().UnixNano()
	fmt.Println("Time in seconds :", (float64(end) - float64(start)) / 1000000000)
}
