package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"rsc.io/quote"
)

type Film struct {
	Title    string
	Director string
}
type Tests struct {
	IdName   string
	ModeName string
	Id       string
	Base     string
	Dt       string
	Mode     string
	Start    string
	End      string
	FullUrl  string
}

func constructUrl(r *http.Request, params string) string {
	//lsChangesScriptName := "lsChanges"
	fullUrl := "http://" + strings.Split(r.Host, ":")[0] + "/lsChanges/lsChanges.php" + "?" + params
	/* 	url := "http://" + r.Host + r.RequestURI
	   	fullPath := r.URL.Path
	   	Port = r.URL.Port()
	   	fileName := path.Base(fullPath) */
	//fullUrl := strings.Replace(url, fileName, lsChangesScriptName, 1) + "?" + params
	return fullUrl
}
func main() {
	fmt.Println(quote.Go())
	handleTests := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("testLsChanges.html"))
		tests := map[string][]Tests{ //ключ string - название таблицы, []Tests - набор данных
			"Tests": {
				//id=201000003125&base=04&dt=19.09.2024&mode=changes&start=1&end=100
				{FullUrl: "", IdName: "Экоград Азов", ModeName: "Статус изменений ЛС", Id: "201000003125", Base: "04", Dt: "19.09.2024", Mode: "status", Start: "1", End: "100"},
				{FullUrl: "", IdName: "Экоград Азов", ModeName: "Изменения ЛС", Id: "201000003125", Base: "04", Dt: "19.09.2024", Mode: "changes", Start: "1", End: "100"},
				{FullUrl: "", IdName: "Экоград Азов", ModeName: "Статус Оплат ЛС", Id: "201000003125", Base: "04", Dt: "19.09.2024", Mode: "status_pay", Start: "1", End: "100"},
				{FullUrl: "", IdName: "Экоград Азов", ModeName: "Оплаты ЛС", Id: "201000003125", Base: "04", Dt: "19.09.2024", Mode: "changes_pay", Start: "1", End: "100"},
				{FullUrl: "", IdName: "Экоград Новочеркасск", ModeName: "Статус изменений ЛС", Id: "201000003592", Base: "04", Dt: "19.09.2024", Mode: "status", Start: "1", End: "100"},
				{FullUrl: "", IdName: "Экоград Новочеркасск", ModeName: "Изменения ЛС", Id: "201000003592", Base: "04", Dt: "19.09.2024", Mode: "changes", Start: "1", End: "100"},
				{FullUrl: "", IdName: "Экоград Новочеркасск", ModeName: "Статус Оплат ЛС", Id: "201000003592", Base: "04", Dt: "19.09.2024", Mode: "status_pay", Start: "1", End: "100"},
				{FullUrl: "", IdName: "Экоград Новочеркасск", ModeName: "Оплаты ЛС", Id: "201000003592", Base: "04", Dt: "19.09.2024", Mode: "changes_pay", Start: "1", End: "100"},
			},
		}
		for i, test := range tests["Tests"] {
			params := fmt.Sprintf("id=%s&base=%s&dt=%s&mode=%s&start=%s&end=%s",
				test.Id, test.Base, test.Dt, test.Mode, test.Start, test.End)
			tests["Tests"][i].FullUrl = constructUrl(r, params)

		}
		tmpl.Execute(w, tests)
		/* 	io.WriteString(w, "Hello, World!")
		io.WriteString(w, r.Method) */
	}

	/*	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("testLsChanges.html"))
		films := map[string][]Film{ //ключ string - название таблицы, []Film - набор данных
			"Films": {
				{Title: "The Matrix", Director: "The Wachowskis"},
				{Title: "The Matrix Reloaded", Director: "The Wachowskis"},
				{Title: "The Matrix Revolutions", Director: "The Wachowskis"},
			},
		}
		tmpl.Execute(w, films)
		// 	io.WriteString(w, "Hello, World!")
		//io.WriteString(w, r.Method)
	} */
	handleAddTest := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		log.Print("HTMX request recieved")
		log.Print(r.Header.Get("HX-Request"))

		idName := r.PostFormValue("idName")
		modeName := r.PostFormValue("modeName")
		id := r.PostFormValue("id")
		base := r.PostFormValue("base")
		dt := r.PostFormValue("dt")
		mode := r.PostFormValue("mode")
		start := r.PostFormValue("start")
		end := r.PostFormValue("end")
		params := fmt.Sprintf("id=%s&base=%s&dt=%s&mode=%s&start=%s&end=%s",
			id, base, dt, mode, start, end)
		fullUrl := constructUrl(r, params)

		htmlStr := "<li class='list-group-item bg-primary text-white'>" +
			fmt.Sprintf("<h3>%s - %s</h3>"+
				"<a class='text-white' href='%s'>%s </a>", idName, modeName, fullUrl, fullUrl) +
			"</li>"
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

	}

	/* 	h2 := func(w http.ResponseWriter, r *http.Request) {
	   		time.Sleep(1 * time.Second)
	   		log.Print("HTMX request recieved")
	   		log.Print(r.Header.Get("HX-Request"))
	   		title := r.PostFormValue("title")
	   		director := r.PostFormValue("director")
	   		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
	   		tmpl, _ := template.New("t").Parse(htmlStr)
	   		tmpl.Execute(w, nil)

	   	}
	*/
	//http.HandleFunc("/", h1)
	http.HandleFunc("/", handleTests)
	http.HandleFunc("/add-test/", handleAddTest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
