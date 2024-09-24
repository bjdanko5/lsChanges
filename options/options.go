package options

import (
	"html/template"
	"net/http"
	"reflect"
	"strconv"
)

var tmplOptions *template.Template

func init() {
	tmplOptions = template.Must(template.ParseFiles("options/Options.html"))
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

type SelectDataForTemplate struct {
	Options        []interface{}
	SelectedOption interface{}
}
type Option interface {
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

func GetIDNameOptions(w http.ResponseWriter, r *http.Request) {
	selectedValue, err := GetParamAsInt(r, "idName")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var Data SelectDataForTemplate

	Data.Options = []interface{}{
		IDNameOption{Value: 1, Id: "201000003125", IdName: "Экоград Азов"},
		IDNameOption{Value: 2, Id: "201000003592", IdName: "Экоград Новочеркасск"},
	}

	selectedOption := findOptionByValue(Data.Options, selectedValue)
	if selectedOption != nil {
		if option, ok := selectedOption.(Option); ok {
			Data.SelectedOption = option
		} else {
			http.Error(w, "Invalid option type", http.StatusInternalServerError)
			return
		}
	}

	err = tmplOptions.ExecuteTemplate(w, "idNameOptions", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetMODENameOptions(w http.ResponseWriter, r *http.Request) {
	selectedValue, err := GetParamAsInt(r, "modeName")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var Data SelectDataForTemplate

	Data.Options = []interface{}{
		MODENameOption{Value: 1, Mode: "status", ModeName: "Статус изменений ЛС"},
		MODENameOption{Value: 2, Mode: "changes", ModeName: "Изменения ЛС"},
		MODENameOption{Value: 3, Mode: "status_pay", ModeName: "Статус оплат ЛС"},
		MODENameOption{Value: 4, Mode: "changes_pay", ModeName: "Оплаты ЛС"},
	}

	selectedOption := findOptionByValue(Data.Options, selectedValue)
	if selectedOption != nil {
		if option, ok := selectedOption.(Option); ok {
			Data.SelectedOption = option
		} else {
			http.Error(w, "Invalid option type", http.StatusInternalServerError)
			return
		}
	}

	err = tmplOptions.ExecuteTemplate(w, "modeNameOptions", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func GetBASENameOptions(w http.ResponseWriter, r *http.Request) {
	selectedValue, err := GetParamAsInt(r, "baseName")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var Data SelectDataForTemplate

	Data.Options = []interface{}{
		BASENameOption{Value: 1, Base: "04", BaseName: "г.Азов"},
	}

	selectedOption := findOptionByValue(Data.Options, selectedValue)
	if selectedOption != nil {
		//	if option, ok := selectedOption.(BASENameOption); ok {
		if option, ok := selectedOption.(Option); ok {
			Data.SelectedOption = option
		} else {
			http.Error(w, "Invalid option type", http.StatusInternalServerError)
			return
		}
	}

	err = tmplOptions.ExecuteTemplate(w, "baseNameOptions", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
