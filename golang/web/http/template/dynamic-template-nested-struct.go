package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type person struct { //Nested structure : Parent
	Name    string
	Age     int
	Address       // nested struct with no field name
	Phone   Phone //nested struct with field name same as struct type
	O       Occp  //nested struct with field name different from the struct type
}
type Occp struct { //Nested in Parent structure
	Name   string //ambiguous with person.name - should address with alias while referring
	Salary string //not ambiguous. but since the struct nesting is done with name, address this field with alias struct
}
type Address struct { //Nested in Parent structure
	Name string //ambiguous with parent
	Zip  string //not ambiguous
}
type Phone struct { //Nested in Parent structure
	Name string //ambiguous with parent
	Type string //not ambiguous
}

const dtmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>dynamic-template-nested</title>
</head>
<body>
    <b><u>Name :</u></b>{{.Name}}<br/>
    <b><u>Age :</u></b>{{.Age}}<br/>

    <b><u>Details:</u></b><br/>
    Department: {{.O.Name}} {{.O.Salary}}<br/>
	Address: {{.Address.Name}} {{.Zip}}<br/>
	Phone: {{.Phone.Name}} {{.Phone.Type}}<br/>
</body>
</html>
	`

func main() {
	o1 := Occp{"Software Engineer", "100K"}
	a1 := Address{"Hyderabad", "500001"}
	ph1 := Phone{"040", "home"}
	p1 := person{Name: "praveen", Age: 35, O: o1, Address: a1, Phone: ph1}

	tmpl, err := template.New("dyn-tmpl-nest-str").Parse(dtmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, p1)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
