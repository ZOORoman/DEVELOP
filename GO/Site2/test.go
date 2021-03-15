package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"

	"./models"
	"./utils"
	"log"
)

var data 		map[string]models.Session
var db 			*sql.DB
var p 			[]*models.Post
var err 		error
var databaseId 	     	int
var databaseUsername 	string
var databasePassword 	string
var res_id		int
var res_title		string
var res_content		string
var res_id_usr		int
var idForEdit		int
var idFromUser		int
var cookies 		string
var Username 		string


const COOKIE_NAME  = "sessionId"


func MainHandler(w http.ResponseWriter, r *http.Request) {
	p = []*models.Post{}

	t, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	rows, err := db.Query("SELECT id, title, content, id_user FROM posts")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		//err = rows.Scan(&post.ID, &post.Title, &post.Content, post.UserID)
		err = rows.Scan(&res_id, &res_title, &res_content, &res_id_usr)
		//log.Println("CHECKING STRUCT:", models.Post{res_id, res_title, res_content, databaseId})
		if err != nil {
			log.Println(err)
			continue
		}
		data := &models.Post{ID: res_id, Title: res_title, Content: res_content, UserID: res_id_usr}
		p = append(p, data)
	}

	t.ExecuteTemplate(w, "index", p)

}

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/log.html", "templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "login", nil)
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	err := db.QueryRow("SELECT id, username, password FROM users WHERE username=?", username).Scan(&databaseId, &databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(w, r, "/signup", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/signup", 301)
		return
	}

	http.Redirect(w, r, "/index", 302)

}


func GetSignupHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/log.html", "templates/index.html", "templates/signup.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "signup", nil)

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		cookies := ensureCookie(r, w) // generation cookies
		if cookies == "" {
			fmt.Println("COOKIES DONT WRITE")
		}
		fmt.Println("COOKIES:", cookies)

		if err != nil {
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO users(username, password, cookies) VALUES(?, ?, ?)", username, hashedPassword, cookies)
		if err != nil {
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}

		fmt.Println("User created: ", username, password)
		http.Redirect(w, r, "/index", 302)
		return
	case err != nil:
		http.Error(w, "Server error, unable to create your account.", 500)
		return
	default:
		http.Redirect(w, r, "/", 301)
	}
}

func WriteHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "write", nil)
}

func SavepostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")


	if id != "" {
		_, err := db.Exec("UPDATE posts SET title=?, content=? WHERE id_user=?", title, content, databaseId)
		if err != nil {
			http.Error(w, "Server error, unable to add your post.", 500)
			return
		}

	} else {
		_, err = db.Exec("INSERT INTO posts(title, content, id_user) VALUES (?, ?, ?)", title, content, databaseId)
		if err != nil {
			http.Error(w, "Server error, unable to add your post.", 500)
			return
		}
	}

	http.Redirect(w, r, "/index", 302)
}


func EditHandler(w http.ResponseWriter, r *http.Request) {
	//res := make([]int, 0)
	data = make(map[string]models.Session)


	cookie := getCookie()
	user := getUsername()


	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	rows, err := db.Query("SELECT id,id_user FROM posts WHERE id_user=?", &databaseId)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&idForEdit, &idFromUser)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	data[cookie] = models.Session{Id: idForEdit, Username: user, IdPost: idFromUser}
	log.Println(data[cookie].IdPost)

	// Нужно получать еще ид поста через БД и сравнивать со значением IdPost, если всё ок, то разрешать редачить пост
	// Если нет, шлём нахуй

	//
	//for _, br := range res {
	//	if br == databaseId {
	//		log.Println(true)
	//		log.Println("br:", br)
	//		log.Println("databaseID", databaseId)
	//		//t.ExecuteTemplate(w, "write", nil)
	//	} else {
	//		log.Println("br:", br)
	//		log.Println("databaseID", databaseId)
	//		log.Println(false)
	//	}
	//	//http.Redirect(w, r, "/index", 302)
	//
	//}

	t.ExecuteTemplate(w, "write", nil)

	//http.Redirect(w, r, "/index", 302)

}


func getCookie() string {
	rows, err := db.Query("SELECT cookies FROM users WHERE id=?", &databaseId)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&cookies)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return cookies

}

func getUsername() string {
	rows, err := db.Query("SELECT username FROM users WHERE id=?", &databaseId)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&Username)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return Username
}

//func DeleteHandler(w http.ResponseWriter, r *http.Request) {
//	id := r.FormValue("id")
//	if id == "" {
//		http.NotFound(w, r)
//		return
//	}
//	delete(posts, id)
//
//	http.Redirect(w, r, "/index", 302)
//}


func ensureCookie(r *http.Request, w http.ResponseWriter) string {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		return  cookie.Value
	}
	sessionId := utils.GenerateId()

	cookie = &http.Cookie{
		Name: COOKIE_NAME,
		Value: sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}
	http.SetCookie(w, cookie)

	return sessionId
}

func main () {
	fmt.Println("[+]Listen on port :3000")

	db, err = sql.Open("mysql", "root@/test")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	// Check connect db
	err = db.Ping()
	if err != nil {
		fmt.Println("[-]Database is not connect")
		panic(err.Error())
	} else {
		fmt.Println("[+]Database is connect")
	}

	http.Handle("./assets", http.StripPrefix("/assets/", http.FileServer(http.Dir("/assets"))))
	http.HandleFunc("/", GetLoginHandler)
	http.HandleFunc("/index", MainHandler)
	http.HandleFunc("/login", PostLoginHandler)
	http.HandleFunc("/write", WriteHandler)
	http.HandleFunc("/SavePost", SavepostHandler)
	//http.HandleFunc("/delete", DeleteHandler)
	http.HandleFunc("/edit", EditHandler)
	http.HandleFunc("/signup", GetSignupHandler)
	http.HandleFunc("/Signup", SignupHandler)

	http.ListenAndServe(":3000", nil)
}
