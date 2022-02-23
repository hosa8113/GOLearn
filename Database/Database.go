package mainx

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Datastructure struct {
	Id        int       `json:"id"`
	Data1     string    `json:"data_1"`
	Data2     string    `json:"data_2"`
	MutnTmstp time.Time `json:"mutn_tmstp"`
}

func main() {

	db := ConnectToDb()
	defer db.Close()

	sql := "SELECT id, data_1, data_2, mutn_tmstp FROM new_test_table"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var o Datastructure
	rowCount := 0
	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Data1, &o.Data2, &o.MutnTmstp)
		if err != nil {
			panic(err)
		}
		rowCount++
		fmt.Println(o)
	}
	fmt.Println(strconv.Itoa(rowCount) + " Rows gelesen")

}

func ConnectToDb() *sql.DB {
	dbHost := "localhost"
	dbPort := 3306
	dbSchema := "klf_ch_auth"
	dbUser := "development_user"
	dbUserPassword := "H@nn1bal"

	tcpConnection := "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbSchema
	db, err := sql.Open("mysql", dbUser+":"+dbUserPassword+tcpConnection+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to " + tcpConnection + " was successful")
	return db
}
