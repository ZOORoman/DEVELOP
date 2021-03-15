package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type Cont struct {
	Name string "name"
	Desc string "desc"

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
//
// 	//Добавление данных
// 	// insert, err := db.Query("INSERT INTO `themas` (`name`, `desc`, `reating`) VALUES('Fl','Прогэ','4')")
// 	// if err != nil{
// 	// 	panic(err)
// 	// }
// 	// defer insert.Close()
//
// 	// fmt.Printf("Обля!")
// // err := db.QueryRow("SELECT data->>'id', data->>'type', data->>'title' FROM message WHERE data->>'id'=$1", id).
// // Scan(m.Id, m.Type, m.Title)
//
	res, err := db.Query("SELECT `name`, `desc` FROM `themas`")
	if err != nil{
		panic(err)
	}

	for res.Next() {
		 var cont Cont
		 // res.Scan(&content.Name, &content.Desc)
		err = res.Scan(&cont.Name, &cont.Desc)
		if err != nil {
			panic(err)
		}
		// err = res.Scan(&content.Desc)
		// if err != nil {
		// 	panic(err)
		// }

	 	fmt.Printf(fmt.Sprintf("%s - это Обля! %s \n", cont.Name, cont.Desc))
	}
}



// package main
//
// import (
//   "fmt"
//   "net/http"
//   "html/template"
// )
//
// func index(w http.ResponseWriter, r *http.Request)  {
//   // fmt.Fprintf(w, "Go!")
//   t, err := template.ParseFiles("templates/index.html")
//   // t, err := template.ParseFiles( "templates/index.html", "templates/header.html", "templates/footer.html")
//
//   if err != nil {
//     fmt.Fprintf(w, err.Error())
//   }
//
//   t.ExecuteTemplate(w, "index", nil) // Чтобы рвботали шаблоны
// }
//
// func handleFunc() {
//   http.HandleFunc("/", index)
//   http.ListenAndServe(":3333", nil)
// }
//
// func main () {
//   handleFunc()
// }
