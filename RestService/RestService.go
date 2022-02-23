package main

import (
	mainx "GOLearn/Database"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"klf-ch-auth/router"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//these vars are not thread-safe!
var a int
var b int
var db *sql.DB

func main() {

	systemId := strings.ToLower(uuid.New().String())
	db = mainx.ConnectToDb()
	defer db.Close()

	e := echo.New()
	e.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "server is running")
	})
	e.GET("/wait/:seconds", handleWait)

	e.PUT("/db", putDataToDb)
	e.PUT("/db2", putDataToDb2)
	e.GET("/db/list", getAllData)

	fmt.Println("System-Id:", systemId)
	ShowExposedEndpoints(e)
	err := e.Start("localhost:8000")
	if err != nil {
		panic(err)
	}
}

func getAllData(context echo.Context) error {
	var objArray []*mainx.Datastructure
	sql := "SELECT id, data_1, data_2, mutn_tmstp FROM new_test_table"
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var o mainx.Datastructure
	rowCount := 0
	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Data1, &o.Data2, &o.MutnTmstp)
		if err != nil {
			return err
		}
		rowCount++
		objArray = append(objArray, &o)
	}
	fmt.Println(strconv.Itoa(rowCount) + " Rows gelesen")
	return router.Json(context, http.StatusOK, objArray)
}

func putDataToDb(context echo.Context) error {
	sqlStatement := "INSERT INTO new_test_table (data_1, data_2) VALUES (?,?)"
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	tx, err = db.Begin()

	for i := 0; i < 5; i++ {
		if err != nil {
			return err
		}
		res, err := tx.Exec(sqlStatement, "Hallo Data 1("+strconv.Itoa(i)+")", "Hallo Data 2("+strconv.Itoa(i)+")")
		if err != nil {
			return err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		fmt.Println("Data (" + strconv.Itoa(i) + ") inserted id=" + strconv.Itoa(int(id)))
		//TODO: check concurrent update/insert
		wait(10)
	}

	/*
		err = tx.Commit()
		if err != nil {
			return err
		}
	*/

	//TODO: rollback funktioniert noch nicht...
	_, err = tx.Exec(sqlStatement, "Rolledback Hallo Data 1()", "Rolledback Hallo Data 2()")
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, "Data inserted and rolled back")
}

func putDataToDb2(context echo.Context) error {
	sqlStatement := "INSERT INTO new_test_table (data_1, data_2) VALUES (?,?)"
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	tx, err = db.Begin()

	for i := 0; i < 5; i++ {
		if err != nil {
			return err
		}
		res, err := tx.Exec(sqlStatement, "Hallo Data 1("+strconv.Itoa(i)+")", "Hallo Data 2("+strconv.Itoa(i)+")")
		if err != nil {
			return err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		fmt.Println("Data (" + strconv.Itoa(i) + ") inserted id=" + strconv.Itoa(int(id)))
		//TODO: check concurrent update/insert
		wait(2)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, "Data inserted")
}

func handleWait(context echo.Context) error {
	c := rand.Intn(10)
	d := rand.Intn(10)
	a = c
	b = d
	seconds, _ := strconv.Atoi(context.Param("seconds"))
	math1 := strconv.Itoa(a) + " + " + strconv.Itoa(b) + " = "
	math2 := strconv.Itoa(c) + " + " + strconv.Itoa(d) + " = "
	wait(seconds)
	return context.JSON(http.StatusOK, "Request waited for "+strconv.Itoa(seconds)+" second(s) [ (a + b =) "+math1+strconv.Itoa(a+b)+" (c + d =) "+math2+strconv.Itoa(c+d)+"]")
}

func wait(seconds int) {
	currentTime := time.Now()
	fmt.Println("it's now", currentTime.Format("2006-01-02 15:04:05"))
	time.Sleep(time.Duration(seconds) * time.Second)
}

func ShowExposedEndpoints(e *echo.Echo) {
	//spew.Dump(e.Routes())
	data := ""
	for index, element := range e.Routes() {
		if index > 0 {
			data += ", "
		}
		data += element.Method + ":" + element.Path
	}
	fmt.Println("the following endpoints are exposed[" + data + "]")
}
