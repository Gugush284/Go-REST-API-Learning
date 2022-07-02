package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Stop struct {
	ID    int
	Title string
}

func getid(w http.ResponseWriter, r *http.Request) {

	stop := Read_db()

	for i := 0; i < len(stop); i++ {
		fmt.Fprintf(w, "%d - %s\n", stop[i].ID, stop[i].Title)
	}
}

func Read_db() []Stop {
	sql_database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/csv_db 6")
	if err != nil {
		panic(err)
	}

	defer sql_database.Close()
	fmt.Println("Подключение создано")

	/*statement, err := sql_database.Prepare("CREATE TABLE IF NOT EXIST people (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		panic(err)
	}
	statement.Exec()*/

	var stop []Stop

	rows, err := sql_database.Query("SELECT _id, name FROM stopsker")
	defer rows.Close()
	if err != nil {
		fmt.Println("Error in getting DB")
		//log.Error("Problem", err)
	} else {
		var id int
		var name string

		for rows.Next() {
			rows.Scan(&id, &name)
			stop = append(stop, Stop{id, name})
		}
	}

	return stop
}

func main() {

	http.HandleFunc("/", getid)
	http.ListenAndServe(":85", nil)
}
