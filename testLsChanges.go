package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
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
type IDNameOption struct {
	Value         int
	Id            string
	IdName        string
	SelectedValue int
}
type MODENameOption struct {
	Value         int
	Mode          string
	ModeName      string
	SelectedValue int
}
type BASENameOption struct {
	Value         int
	Base          string
	BaseName      string
	SelectedValue int
}

func findOptionByValue(options []interface{}, value int) interface{} {
	for _, option := range options {
		field := reflect.ValueOf(option).FieldByName("Value")
		if field.IsValid() {
			if field.Kind() == reflect.Int && field.Int() == int64(value) {
				return option
			}
		}
	}
	return nil
}
func findIdByValue(options []IDNameOption, value int) string {
	for _, option := range options {
		if option.Value == value {
			return option.Id
		}
	}
	return ""
}
func findModeByValue(options []MODENameOption, value int) string {
	for _, option := range options {
		if option.Value == value {
			return option.Mode
		}
	}
	return ""
}

/*
	 func findBaseByValue(options []BASENameOption, value int) string {
		for _, option := range options {
			if option.Value == value {
				return option.Base
			}
		}
		return ""
	}
*/
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
func GetIDNameOptions(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	idNameValue := q.Get("idName")
	if idNameValue == "" {
		idNameValue = "1"
	}
	idNameOptions := []IDNameOption{
		{Value: 1, Id: "201000003125", IdName: "Экоград Азов", SelectedValue: convertToInt(idNameValue)},
		{Value: 2, Id: "201000003592", IdName: "Экоград Новочеркасск", SelectedValue: convertToInt(idNameValue)},
	}
	selectedId := findIdByValue(idNameOptions, convertToInt(idNameValue))
	fmt.Println(selectedId)

	tmpl, err := template.ParseFiles("idNameOptions.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		IDNameOptions []IDNameOption
	}{
		IDNameOptions: idNameOptions,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func GetMODENameOptions(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	modeNameValue := q.Get("modeName")
	if modeNameValue == "" {
		modeNameValue = "1"
	}
	modeNameOptions := []MODENameOption{
		{Value: 1, Mode: "status", ModeName: "Статус изменений ЛС", SelectedValue: convertToInt(modeNameValue)},
		{Value: 2, Mode: "changes", ModeName: "Изменения ЛС", SelectedValue: convertToInt(modeNameValue)},
		{Value: 3, Mode: "status_pay", ModeName: "Статус оплат ЛС", SelectedValue: convertToInt(modeNameValue)},
		{Value: 4, Mode: "changes_pay", ModeName: "Оплаты ЛС", SelectedValue: convertToInt(modeNameValue)},
	}
	selectedMode := findModeByValue(modeNameOptions, convertToInt(modeNameValue))
	fmt.Println(selectedMode)

	tmpl, err := template.ParseFiles("modeNameOptions.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		MODENameOptions []MODENameOption
	}{
		MODENameOptions: modeNameOptions,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func GetParamAsInt(r *http.Request, name string) (int, error) {
	q := r.URL.Query()
	value := q.Get(name)
	if value == "" {
		return 1, nil // or return an error, depending on your needs
	}
	paramAsInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return paramAsInt, nil
}
func GetBASENameOptions(w http.ResponseWriter, r *http.Request) {
	selectedValue, _ := GetParamAsInt(r, "baseName")
	options := []interface{}{
		BASENameOption{
			Value:         1,
			Base:          "04",
			BaseName:      "г.Азов",
			SelectedValue: selectedValue,
		},
		// добавьте другие элементы аналогичным образом
	}
	var Data struct {
		BASENameOptions       []interface{}
		SelectedBaseNameValue int
		SelectedBaseName      string
		SelectedBase          string
	}

	selectedOption := findOptionByValue(options, selectedValue)
	if option, ok := selectedOption.(BASENameOption); ok {
		fmt.Println(option.BaseName)
		Data.BASENameOptions = options
		Data.SelectedBaseNameValue = option.Value
		Data.SelectedBaseName = option.BaseName
		Data.SelectedBase = option.Base
	}
	//fmt.Println(selectedOption.BaseName)

	tmpl, err := template.ParseFiles("baseNameOptions.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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
		idNameText := r.PostFormValue("idNameText")
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
				"<a class='text-warning' href='%s'>%s </a>", idNameText, modeText, fullUrl, fullUrl) +
			"</li>"
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

	}

	http.HandleFunc("/", handleTests)
	http.HandleFunc("/add-test/", handleAddTest)
	http.HandleFunc("/get-id-name-options", GetIDNameOptions)     // GetIDNameOptions
	http.HandleFunc("/get-mode-name-options", GetMODENameOptions) // GetMODENameOptions
	http.HandleFunc("/get-base-name-options", GetBASENameOptions) // GetBASENameOptions

	log.Fatal(http.ListenAndServe(":8000", nil))
}
