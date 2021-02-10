package main

import ("fmt" 
	"net/http")

type Content struct {
						id 	uint16
						name string
						reating, avg_data float64

}

func (c Content) getAllInfo() string{
	return fmt.Sprintf("#%d Название вкладки %s заполнение контента %f," +
		"рейтинг: %f", c.id, c.name, c.avg_data, c.reating)
}

func (c *Content) setNewName(newName string){ //Функция передачи параметра ( * - ссылочный тип )
	c.name = newName
}
 
 func home_page(w http.ResponseWriter, r *http.Request){    // w & r - это параметры, в данном случае один принимает ввод, а другой передачу
 	// p := Content{id: 1, name: "ABAP", reating: 3.9, avg_data: 1.8}
 	datas := Content{1, "C#", 5.9, 0.8} // Запишем в структуру данные
 	datas.setNewName("C-+")             // Передаем в функцию по адресу в ссылочный тип
	fmt.Fprintf(w, "Салям!\n" + datas.getAllInfo())		// Fprintf - форматированная строка
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


 func main(){ 

 	handleRequest()
 }  