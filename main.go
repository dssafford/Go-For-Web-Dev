package main


import (

	"fmt"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Page struct {
	Name string
	Job string
	DBStatus bool
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:mypassword@tcp(0.0.0.0:9010)/douggo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	p := Page{Name: "Doug"}

	//err = db.Ping()


	err = db.Ping()
	if err != nil {
		p.DBStatus = false
		panic(err.Error())
	} else {
		p.DBStatus = true
	}

	templates := template.Must(template.ParseFiles("templates/index.html"))


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {



		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}


		var myjob string

		err := db.QueryRow("SELECT job FROM test WHERE name=?", p.Name).Scan(&myjob)

		if err != nil {
			//http.Redirect(res, req, "/login", 301)
			fmt.Println("error ")

			return
		}

		p.Job = myjob;

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//w.Write([]byte("hello " + myjob))

		//fmt.Fprintf(w, "Hello, Go Web Development")
	})

	fmt.Println(http.ListenAndServe(":3000", nil))
}
