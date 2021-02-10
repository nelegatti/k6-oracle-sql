package query

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/godror/godror"

	"github.com/loadimpact/k6/js/modules"
)

func init() {
	modules.Register("k6/x/query", new(QUERY))
}

// OracleConnector is the k6 SQL plugin.
type QUERY struct{}

func (*QUERY) Query() {
	currentTime := time.Now()
	fmt.Println("Starting at : ", currentTime.Format("03:04:05:06 PM"))

	fmt.Println("... Setting up Database Connection")

	db, err := sql.Open("godror", `user="dummy" password="dummy" connectString="localhost:1521/ORCLCDB.localdomain"
    libDir="/mnt/f/Linux/instantclient_19_6"`)
	if err != nil {
		fmt.Println("... DB Setup Failed")
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Println("... Opening Database Connection")
	// if err = db.Ping(); err != nil {
	// 	fmt.Println("Error connecting to the database: %s\n", err)
	// 	return
	// }
	fmt.Println("... Connected to Database")

	// dbQuery := "select table_name from user_tables where table_name not like 'DM$%' and table_name not like 'ODMR$%'"
	// rows, err := db.Query(dbQuery)
	// if err != nil {
	// 	fmt.Println(".....Error processing query")
	// 	fmt.Println(err)
	// 	return
	// }
	// defer rows.Close()

	// fmt.Println("... Parsing query results")
	// var tableName string
	// for rows.Next() {
	// 	rows.Scan(&tableName)
	// 	fmt.Println(tableName)
	// }

	fmt.Println("... Closing connection")
	finishTime := time.Now()
	fmt.Println("Finished at ", finishTime.Format("03:04:05:06 PM"))
}
