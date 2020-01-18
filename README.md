## Kalachnigo
Is a fast program to import a large csv file to Mysql database.
- Video [example](https://www.youtube.com/watch?v=GLI2g12iIqo)

In the previous video we can notice :
- Execution time < 90 Seconds
- Memory usage < 30 Mo

## Requirements
- Go MySQL [Driver](https://github.com/go-sql-driver/mysql)

## Installation
- Compile the main.go file.

```text
go run main.go
```

Or

- Build an executable of the program.

```text
go build main.go
```

**Important:** Make sure that the csv file is in the same folder as the program.

## Usage

```text
Database name: name_of_databse
user : user_of_mysql_server
pass : password_of_mysql_server
Table name : table_name

Start of operation ...
Completed operation
Execution time : 70 seconds
```

- Database name: the name of your database in MySql server.
- User: the username in MySql server.
- Password: the password in MySql server.

```go
sqlqr := user + ":" + pass + "@tcp(127.0.0.1:3306)/" + DatabaseName
db, err := sql.Open("mysql", sqlqr) // connection to the database
```

- Table name: the name of the table in MySql server.

```go
_ ,err = db.Exec("LOAD DATA LOCAL INFILE 'buffer.csv' INTO TABLE " + tableName + " FIELDS TERMINATED BY '\"'") // write in database
```
## Links

Project [link](https://drive.google.com/open?id=134GWlmF4X09SwLCEwbkT3Wnx7peD2drM)
