package main

import ("fmt" 			//биб работа с вводом\выводом
	"net/http"        	//библия работы с сервером
	"html/template")	//работа с шаблонами


// Структура
type Content struct {
						Id 	uint16
						Name string
						Reating, Avg_data float64  // Названия даем с большой буквы иначе регистр в шаблоне не будет рабкать
						Sections[] string
}
//

// Работа
func (c Content) getAllInfo() string {
	return fmt.Sprintf("#%d Название вкладки %s заполнение контента %f," +
		"рейтинг: %f", c.Id, c.Name, c.Avg_data, c.Reating)
}

func (c *Content) setNewName(newName string) { //Функция передачи параметра ( * - ссылочный тип )
	c.Name = newName
}
 
func home_page(w http.ResponseWriter, r *http.Request){    // w & r - это параметры, в данном случае один принимает ввод, а другой передачу
	datas := Content{1, "C#", 5.9, 0.8, []string{"Dev", "Work", "Learn"}}
	// fmt.Fprintf(w, "<b>Сайтец</b>")		// Fprintf - форматированная строка
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, datas)

}

  func contacts_page(w http.ResponseWriter, r *http.Request){    // w & r - это параметры, в данном случае один принимает ввод, а другой передачу
 	fmt.Fprintf(w, "Контакты")											// Fprintf - форматированная строка
 }

func handleRequest(){ 					// Чтобы не спамить главную функцию(main) создал свою (HandleRequest)
 	// fmt.Println("Hello GO!")
 	http.HandleFunc("/", home_page) 	// HandleFunc - Функция отслеживания URL адреса 
 										//	home_page - метод, его название может быть каким угодно	
 	http.HandleFunc("/contacts/", contacts_page)
 	http.ListenAndServe(":3030", nil)   // ListenAndServe - Функция передачи параметров на сервер											
 										// 	nil - пустота, означает, что никаких настроек мы не передаем
}
// 

// Начало проги
 func main(){ 

 	handleRequest()
 }
 //