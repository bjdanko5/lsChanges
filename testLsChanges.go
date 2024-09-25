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

var tmplTests *template.Template
var tests TestsType

func init() {
	tmplTests = template.Must(template.ParseFiles("Tests.html"))
	tests = []Test{
		//id=201000003125&base=04&dt=19.09.2024&mode=changes&start=1&end=100
		{FullUrl: "", IdName: "Экоград Азов", ModeName: "Статус изменений ЛС", Id: "201000003125", Base: "04", Dt: "19.09.2024", Mode: "status", Start: "1", End: "100"},
		{FullUrl: "", IdName: "Экоград Азов", ModeName: "Изменения ЛС", Id: "201000003125", Base: "04", Dt: "19.09.2024", Mode: "changes", Start: "1", End: "100"},
		{FullUrl: "", IdName: "Экоград Азов", ModeName: "Статус Оплат ЛС", Id: "201000003125", Base: "04", Dt: "19.09.2024", Mode: "status_pay", Start: "1", End: "100"},
		{FullUrl: "", IdName: "Экоград Азов", ModeName: "Оплаты ЛС", Id: "201000003125", Base: "04", Dt: "19.09.2024", Mode: "changes_pay", Start: "1", End: "100"},
		{FullUrl: "", IdName: "Экоград Новочеркасск", ModeName: "Статус изменений ЛС", Id: "201000003592", Base: "04", Dt: "19.09.2024", Mode: "status", Start: "1", End: "100"},
		{FullUrl: "", IdName: "Экоград Новочеркасск", ModeName: "Изменения ЛС", Id: "201000003592", Base: "04", Dt: "19.09.2024", Mode: "changes", Start: "1", End: "100"},
		{FullUrl: "", IdName: "Экоград Новочеркасск", ModeName: "Статус Оплат ЛС", Id: "201000003592", Base: "04", Dt: "19.09.2024", Mode: "status_pay", Start: "1", End: "100"},
		{FullUrl: "", IdName: "Экоград Новочеркасск", ModeName: "Оплаты ЛС", Id: "201000003592", Base: "04", Dt: "19.09.2024", Mode: "changes_pay", Start: "1", End: "100"},
	}

}

type Test struct {
	IdName   string
	ModeName string
	Id       string
	Base     string
	Dt       string
	Mode     string
	Start    string
	End      string
	FullUrl  string
	Added    bool
}
type TestsType []Test
type TestsDataForTemplate struct {
	Tests TestsType
}

func (t *TestsType) Template(w http.ResponseWriter, r *http.Request) {
	var Data TestsDataForTemplate
	Data.Tests = *t
	err := tmplTests.ExecuteTemplate(w, "tests", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (t *TestsType) GetTests(r *http.Request) {
	for i, test := range *t {
		params := fmt.Sprintf("id=%s&base=%s&dt=%s&mode=%s&start=%s&end=%s",
			test.Id, test.Base, test.Dt, test.Mode, test.Start, test.End)
		(*t)[i].FullUrl = constructUrl(r, params)
	}
}
func (t *TestsType) AddTest(w http.ResponseWriter, r *http.Request) {

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
	tests = append([]Test{Test{
		IdName:   idText,
		ModeName: modeText,
		Id:       id,
		Base:     base,
		Dt:       dt,
		Mode:     mode,
		Start:    start,
		End:      end,
		FullUrl:  "",
		Added:    true,
	}}, tests...)

}

func constructUrl(r *http.Request, params string) string {
	//lsChangesScriptName := "lsChanges"
	fullUrl := "http://" + strings.Split(r.Host, ":")[0] + "/lsChanges/lsChanges.php" + "?" + params
	return fullUrl
}

func convertDate(dt string) (string, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dt)
	if err != nil {
		return "", err
	}
	return t.Format("02.01.2006"), nil
}

func handleTestLsChanges(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("testLsChanges.html"))
	currentDate := time.Now().Format("2006-01-02")
	type Data struct {
		CurrentDate string
	}
	data := Data{
		CurrentDate: currentDate,
	}
	tmpl.Execute(w, data)
}
func handleGetTests(w http.ResponseWriter, r *http.Request) {
	tests.GetTests(r)
	tests.Template(w, r)
}
func handleAddTest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(500 * time.Millisecond)
	tests.AddTest(w, r)
	tests.GetTests(r)
	tests.Template(w, r)
}

func main() {

	fmt.Println(quote.Go())

	http.HandleFunc("/", handleTestLsChanges)
	http.HandleFunc("GET /tests", handleGetTests)
	http.HandleFunc("POST /tests", handleAddTest)
	http.HandleFunc("GET /get-id-name-options", options.GetIDNameOptions)
	http.HandleFunc("GET /get-mode-name-options", options.GetMODENameOptions)
	http.HandleFunc("GET /get-base-name-options", options.GetBASENameOptions)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

// пока не использую-----------------------------------------------
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
