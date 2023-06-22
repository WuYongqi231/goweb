package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

var Id int
var Username string
var Password string
var Content string
var db *sql.DB

func initDB() {
	var err error
	// 设置一下dns charset:编码方式 parseTime:是否解析time类型 loc:时区
	dsn := "root:wyqsdyp123@tcp(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开mysql驱动
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
}

func AddUser(username, password string) {
	initDB()
	defer db.Close()
	sqlStr := "insert into user(username,password) values (?,?)"
	_, err := db.Exec(sqlStr, username, password)
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	initDB()
	defer db.Close()
	sqlStr := "SELECT id, username, password FROM user WHERE username = ?"
	err := db.QueryRow(sqlStr, username).Scan(&Id, &Username, &Password)
	log.Println(Id, Username, Password)
	if err != nil {
		log.Println("The query defeated")
		return false
	} else {
		log.Println("The query succeeded")
		return true
	}
}

func SelectPasswordFromUsername(username string) string {
	initDB()
	defer db.Close()
	sqlStr := "SELECT id, username, password FROM user WHERE username = ?"
	db.QueryRow(sqlStr, username).Scan(&Id, &Username, &Password)
	return Password
}
func AddQuestion(user_id, content string) bool {
	initDB()
	defer db.Close()
	uid, _ := strconv.Atoi(user_id)
	sqlStr := "insert into questions(userid,content) values (?,?)"
	_, err := db.Exec(sqlStr, uid, content)
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return false
	}
	log.Println("insert success")
	return true
}

func DeleteQuestion(question_id string) bool {
	initDB()
	defer db.Close()
	qid, _ := strconv.Atoi(question_id)
	sqlStr := "delete from questions where id=?"
	_, err := db.Exec(sqlStr, qid)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return false
	}
	log.Println("delete success")
	return true
}
func CreateAnswer(question_id, content, user_id string) bool {
	initDB()
	defer db.Close()
	qid, _ := strconv.Atoi(question_id)
	uid, _ := strconv.Atoi(user_id)
	sqlStr := "insert into answers(question_id,user_id,content) values (?,?,?)"
	_, err := db.Exec(sqlStr, qid, uid, content)
	if err != nil {
		log.Printf("answer failed, err:%v\n", err)
		return false
	}
	log.Printf("%s answer to %d success", user_id, qid)
	return true
}
func DeleteAnswer(answer_id string) bool {
	initDB()
	defer db.Close()
	aid, _ := strconv.Atoi(answer_id)
	sqlStr := "delete from answers where id=?"
	_, err := db.Exec(sqlStr, aid)
	if err != nil {
		log.Printf("delete answer failed, err:%v\n", err)
		return false
	}
	log.Println("delete answer success")
	return true
}
func ShowQuestion(userid string) (bool, string) {
	initDB()
	defer db.Close()
	var ans string
	uid, _ := strconv.Atoi(userid)
	sqlStr := "select content from questions where userid = ?"
	rows, err := db.Query(sqlStr, uid)
	if err != nil {
		log.Printf("Select question of %s failed, err:%v\n", userid, err)
		return false, ""
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Content)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			return false, ans
		}
		ans += Content
		log.Printf("Select question of %s success,", userid)
	}
	return true, ans
}
func ShowAnswer(asker, question_id string) (bool, string) {
	initDB()
	defer db.Close()
	var answer_id int
	var ans, temp string
	qid, _ := strconv.Atoi(question_id)
	sqlStr := "select user_id,content from answers where question_id=?"
	rows, err := db.Query(sqlStr, qid)
	if err != nil {
		log.Printf("Select answer of %s for %s failed, err:%v\n", question_id, asker, err)
		return false, ""
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&answer_id, &Content)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			return false, ans
		}
		a := strconv.Itoa(answer_id)
		temp = "Answer: " + Content + ".answered by " + a + " ,asked by " + asker
		log.Printf("Answer:%s ,answered by %d,asked by %s", Content, answer_id, asker)
		ans += temp + "\n"
	}
	return true, ans
}
