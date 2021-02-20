package main

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Cont struct {
	Name string `json:"name"`
	Desc string `json:"desc"`

	// Name string `json:"name"`
	// Desc string `json:"desc"`
	// Reating uint16 `json:"reating"`
}

func main(){

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8009)/go")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	//Добавление данных
	// insert, err := db.Query("INSERT INTO `content` (`name`, `desc`, `reating`) VALUES('Fl','Прогэ','4')")
	// if err != nil{
	// 	panic(err)
	// }
	// defer insert.Close()

	// fmt.Printf("Куй")


	res, err := db.Query("SELECT `name` `desc` FROM `themas`")
	if err != nil{
		panic(err)
	}

	for res.Next() {
		 var content Cont
		 res.Scan(&content.Name, &content.Desc)
		// err := res.Scan(&content.Name, &content.Desc)
		// if err != nil {
		// 	panic(err)
		// }
	 	fmt.Printf(fmt.Sprintf("%s контент %s \n", content.Name, content.Desc))
	}
}