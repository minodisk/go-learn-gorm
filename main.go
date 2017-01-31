package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	if err := _main(); err != nil {
		panic(err)
	}
}

func _main() error {
	db, err := gorm.Open("mysql", "root:@tcp(mysql:3306)/test")
	defer db.Close()
	if err != nil {
		return err
	}

	// for {
	// 	if err := db.Ping(); err != nil {
	// 		fmt.Println(err)
	// 		time.Sleep(time.Second)
	// 		continue
	// 	}
	// 	break
	// }

	if err := readyTables(db); err != nil {
		return err
	}
	if err := create(db); err != nil {
		return err
	}
	if err := readList(db); err != nil {
		return err
	}
	// if err := union(db); err != nil {
	// 	return err
	// }

	return nil
}

type User struct {
	ID     int
	Name   sql.NullString
	Gender sql.NullString
	Age    sql.NullInt64
}

type Post struct {
	ID     int
	Body   sql.NullString
	UserID sql.NullInt64
	User   User `gorm:"ForeignKey:UserID"`
}

func readyTables(db *gorm.DB) error {
	db.CreateTable(&User{}, &Post{})
	return nil
}

func NewNullString(s string) sql.NullString {
	return sql.NullString{Valid: true, String: s}
}

func NewNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Valid: true, Int64: i}
}

func NewNullBool(b bool) sql.NullBool {
	return sql.NullBool{Valid: true, Bool: b}
}

func create(db *gorm.DB) error {
	for _, u := range []User{
		{
			Name:   NewNullString("Foo"),
			Gender: NewNullString("male"),
			Age:    NewNullInt64(29),
		},
		{
			Name:   NewNullString("Bar"),
			Gender: NewNullString("female"),
			Age:    NewNullInt64(17),
		},
		{
			Name:   NewNullString("Baz"),
			Gender: NewNullString("male"),
			Age:    NewNullInt64(41),
		},
		{
			Name:   NewNullString("Qux"),
			Gender: NewNullString("female"),
			Age:    NewNullInt64(32),
		},
		{
			Name:   NewNullString("Hoge"),
			Gender: NewNullString("male"),
			Age:    NewNullInt64(11),
		},
		{
			Name:   NewNullString("Fuga"),
			Gender: NewNullString("female"),
			Age:    NewNullInt64(51),
		},
	} {
		db.Create(&u)
		// _, err := sess.InsertInto("users").Columns("name", "gender", "age").Record(&u).Exec()
		// if err != nil {
		// 	return err
		// }
	}

	for _, p := range []Post{
		{
			UserID: NewNullInt64(1),
			Body:   NewNullString("AAAAAAAAAAA"),
		},
		{
			UserID: NewNullInt64(1),
			Body:   NewNullString("BBBBBBBBBBBBB"),
		},
		{
			UserID: NewNullInt64(2),
			Body:   NewNullString("CCCCCC"),
		},
	} {
		db.Create(&p)
		// _, err := sess.InsertInto("posts").Columns("user_id", "body").Record(&u).Exec()
		// if err != nil {
		// 	return err
		// }
	}

	return nil
}

func readList(db *gorm.DB) error {
	var ps []Post
	db.Find(&ps)
	for i, _ := range ps {
		db.Model(&ps[i]).Related(&ps[i].User)
	}
	posts, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("posts = %v\n", string(posts))

	return nil
}

//
// func uniq(ns []sql.NullInt64) []int64 {
// 	m := map[int64]bool{}
// 	for _, n := range ns {
// 		if n.Valid {
// 			m[n.Int64] = true
// 		}
// 	}
// 	var s []int64
// 	for i, _ := range m {
// 		s = append(s, i)
// 	}
// 	return s
// }
//
// func union(conn *sql.Connection) error {
// 	var us []User
// 	sess := conn.NewSession(nil)
// 	if _, err := sess.Select("*").From(
// 		sql.Union(
// 			sql.Select("*").From("users").Where(
// 				sql.And(
// 					sql.Gt("age", 20),
// 					sql.Eq("gender", "male"),
// 				),
// 			),
// 			sql.Select("*").From("users").Where(
// 				sql.And(
// 					sql.Lt("age", 50),
// 					sql.Eq("gender", "female"),
// 				),
// 			),
// 		).As("uni"),
// 	).Load(&us); err != nil {
// 		return err
// 	}
//
// 	users, err := json.MarshalIndent(us, "", "  ")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("union users = %v\n", string(users))
//
// 	return nil
// }
