package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
type into struct {
	id int `db:"id"`
	message string `db:"message"`
	time string `db:"time"`
}
func check(err error) {
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}
func main() {
	db, err := sql.Open("mysql", "root:qwe123@tcp(127.0.0.1:3306)/bottle?charset=utf8")
	check(err)
	rows,err:=db.Query("SELECT *FROM bottle.bottle ORDER BY  RAND() LIMIT 4")
	check(err)
	for rows.Next(){
		var s into
		err=rows.Scan(&s.id,&s.message,&s.time)
		check(err)
		fmt.Println(s)
		break
	}
	rows.Close()
	var n int
	fmt.Scanf("%d",&n)
	res,err:=db.Exec("DELETE FROM bottle.bottle where id=?",n)
	check(err)
	q,err:=res.RowsAffected()
	check(err)
	fmt.Println(q)
}