package main

import (
	"bjdanko5/lsChanges/options"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"rsc.io/quote"
)

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

func convertToInt(idNameValue interface{}) int {
	if idNameValue == nil || idNameValue == "" {
		return 0
	}
	id, err := strconv.Atoi(idNameValue.(string))
	if err != nil {
		return 0
	}
	return id
}
func convertDate(dt string) (string, error) {
	//if layout == "DD.MM.YYYY" {
	//	dt = strings.Split(dt, "-")[2] + "-" + strings.Split(dt, "-")[1] + "-" + strings.Split(dt, "-")[0] */
	//}
	layout := "2006-01-02"
	t, err := time.Parse(layout, dt)
	if err != nil {
		return "", err
	}
	return t.Format("02.01.2006"), nil
}

func constructUrl(r *http.Request, params string) string {
	//lsChangesScriptName := "lsChanges"
	fullUrl := "http://" + strings.Split(r.Host, ":")[0] + "/lsChanges/lsChanges.php" + "?" + params
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
		//tmpl.Execute(w, tests)

		currentDate := time.Now().Format("2006-01-02")
		type Data struct {
			CurrentDate string
			Tests       []Tests
		}
		data := Data{
			CurrentDate: currentDate,
			Tests:       tests["Tests"],
		}

		tmpl.Execute(w, data)

	}
	handleAddTest := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		log.Print("HTMX request recieved in handleAddTest")
		log.Print(r.Header.Get("HX-Request"))
		//idName := r.PostFormValue("idName")
		idText := r.PostFormValue("idText")
		//modeName := r.PostFormValue("modeName")
		modeText := r.PostFormValue("modeText")
		id := r.PostFormValue("id")
		base := r.PostFormValue("base")
		dt := r.PostFormValue("dt")
		newDt, err := convertDate(dt)
		if err != nil {
			fmt.Println(err)
		} else {
			dt = newDt
		}

		mode := r.PostFormValue("mode")
		start := r.PostFormValue("start")
		if start == "" {
			start = "1"
		}
		end := r.PostFormValue("end")
		if end == "" {
			end = "1"
		}

		params := fmt.Sprintf("id=%s&base=%s&dt=%s&mode=%s&start=%s&end=%s",
			id, base, dt, mode, start, end)
		fullUrl := constructUrl(r, params)

		htmlStr := "<li class='list-group-item bg-primary text-white'>" +
			fmt.Sprintf("<h3>%s - %s</h3>"+
				"<a class='text-warning' href='%s'>%s </a>", idText, modeText, fullUrl, fullUrl) +
			"</li>"
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

	}

	http.HandleFunc("/", handleTests)
	http.HandleFunc("/add-test/", handleAddTest)
	http.HandleFunc("/get-id-name-options", options.GetIDNameOptions)     // GetIDNameOptions
	http.HandleFunc("/get-mode-name-options", options.GetMODENameOptions) // GetMODENameOptions
	http.HandleFunc("/get-base-name-options", options.GetBASENameOptions) // GetBASENameOptions

	log.Fatal(http.ListenAndServe(":8000", nil))
}
