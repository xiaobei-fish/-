package main
import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strings"
)
func check(err error) {
	if err != nil{
	fmt.Println(err)
	panic(err)
	}
}
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	n:=4
	num:=0
	r.ParseForm()
	for k, v := range r.Form {
		num++
		fmt.Println("", k)
		fmt.Println("", strings.Join(v,""))
			if num==1 {
				db, err := sql.Open("mysql", "root:qwe123@tcp(127.0.0.1:3306)/bottle?charset=utf8")
				check(err)
				stem, err := db.Exec("INSERT INTO bottle(id,message,time)VALUE (?,?,?)", n, strings.Join(v, ""), strings.Join(v, ""))
				check(err)
				stem.LastInsertId()
				n++
			}
			if num==2{
				db, err := sql.Open("mysql", "root:qwe123@tcp(127.0.0.1:3306)/bottle?charset=utf8")
				check(err)
				stmt, err := db.Prepare(`UPDATE bottle SET time=? WHERE id=?`)
				check(err)
				res, err := stmt.Exec(strings.Join(v, ""),4)
				check(err)
				res.LastInsertId()
			}
	}
	mes :=r.Form["message"]
	tm :=r.Form["time"]
	for i,v :=range mes{
		fmt.Println(i)
		fmt.Fprintf(w,"内容:%v\n",v)
	}
	for k,n :=range tm{
		fmt.Println(k)
		fmt.Fprintf(w,"时间:%v\n",n)
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("message:", r.Form["message"])
		fmt.Println("time:", r.Form["time"])
	}
}
func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}