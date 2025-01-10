package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/alielmi98/Unit-Converter-Web-APP-Go-Implementation/services"
)

// LengthHandler handles length conversion
func LengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		value, _ := strconv.ParseFloat(r.FormValue("lengthValue"), 64)
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")
		result := services.ConvertLength(value, fromUnit, toUnit)

		tmpl := template.Must(template.ParseFiles("templates/length.html"))
		tmpl.Execute(w, map[string]interface{}{
			"Result": result,
			"ToUnit": toUnit,
		})
	} else {
		tmpl := template.Must(template.ParseFiles("templates/length.html"))
		tmpl.Execute(w, nil)
	}
}
