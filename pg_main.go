package main
import(
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)

const(
	USER="postgres"
	PASSWORD="gechman"
	DB_NAME="first_time"
)

//Db pointer to a database connnection
var Db *sql.DB 

func init()  {
	var err error //This is because we are not using colon
	info:=fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",USER,DB_NAME,PASSWORD)
	Db,err=sql.Open("postgres",info)
	checkErr(err)
	fmt.Println("Successfully Connected")
}


type Post struct{
	ID int
	Author string
	Content string
}

func main()  {
	post1:=Post{Author:"Getachew Tebikew",Content:"political Case"}
	//err:=post1.insert()
	// if err!=nil {
	// 	panic(err)
	// }
p,er:=getpost(3)
checkErr(err)


fmt.Println(p.ID,p.Content,p.Author)
	post1.update()

	update(1)
	update(4)
	delete(5)

	rows, err := Db.Query("SELECT * FROM posts")
	checkErr(err)


	fmt.Println("postId |      Author     |      Content     |")
	for rows.Next() {
		var id int
		var author string
		var content string

		err = rows.Scan(&id, &author, &content)
		checkErr(err)

		s:=fmt.Sprintf("%6v | %15v | %16v |", id, author, content)
		fmt.Println(s)

	}

}

func delete(id int) {
	stmt,err:=Db.Prepare("delete from posts where id=$1")
	checkErr(err)
	_,err=stmt.Exec(id)
	checkErr(err)

}
func update(id int) {
	stmt,err:=Db.Prepare("update posts set author=$1 where id=$2")
	checkErr(err)
	_,err=stmt.Exec("Abebe Hlala",id)
	checkErr(err)

}
func (post *Post) insert()(err error)  {
	statement:="insert into posts(content,author) values($1,$2) returning id"	
	stmt,err:=Db.Prepare(statement)
	checkErr(err)
	defer stmt.Close()
	err=stmt.QueryRow(post.Content,post.Author).Scan(&post.ID)//here we are assigning the returnerd auto incremented id to our post
	checkErr(err)
	
	return
}

func getpost(id int)(post Post,err error)   {
	post=Post{}
	err=Db.QueryRow("select id,content,author from posts where id=$1",id).Scan(&post.ID,&post.Content,&post.Author)
	checkErr(err)
	return 
}
func checkErr(err error)  {
	if err != nil {
		pancic(err)
	}
}