package main

import (
	"context"
	"fmt"
	"go_mysql/db"
	"log"
	"strconv"
)

func main() {
	db := db.GetCon()
	defer db.Close()

	ctx := context.Background()

	// script := `
	// 	INSERT INTO customer(id,name)
	// 	VALUES('budi','Budi Kurniawan')
	// `
	// _, err := db.ExecContext(ctx, script)

	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// fmt.Println("Success Insert")

	// script_all := `
	// 	SELECT *
	// 	FROM customer
	// `
	// rows, err := db.QueryContext(ctx, script_all)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var id, name string
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Panic(err.Error())
	// 	}
	// 	fmt.Println("ID :", id)
	// 	fmt.Println("NAME :", name)
	// }

	// DB_USER := "user"
	// username := "admin'; #"
	// password := "admin"

	// script_all := `
	// 	SELECT username
	// 	FROM ` + DB_USER + `
	// 	WHERE username=? AND password=?
	// 	LIMIT 1
	// `
	// rows, err := db.QueryContext(ctx, script_all, username, password)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// defer rows.Close()
	// if rows.Next() {
	// 	var username_db string
	// 	err := rows.Scan(&username_db)
	// 	if err != nil {
	// 		log.Panic(err.Error())
	// 	}
	// 	fmt.Println("USERNAME :", username_db)
	// } else {
	// 	fmt.Println("Gagal Login")
	// }

	// DB_USER := "user"
	// username := "admin"
	// password := "admin"

	// script_all := `
	// 	SELECT username
	// 	FROM ` + DB_USER + `
	// 	WHERE username=? AND password=?
	// 	LIMIT 1
	// `
	// stmt, err := db.PrepareContext(ctx, script_all)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// defer stmt.Close()

	// rows, err := stmt.QueryContext(ctx, username, password)

	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// if rows.Next() {
	// 	var username_db string
	// 	err := rows.Scan(&username_db)
	// 	if err != nil {
	// 		log.Panic(err.Error())
	// 	}
	// 	fmt.Println("USERNAME :", username_db)
	// } else {
	// 	fmt.Println("Gagal Login")
	// }
	// defer rows.Close()

	// DB_COMMENT := "comment"

	// script_all := `
	// 	INSERT
	// 	INTO ` + DB_COMMENT + `
	// 	(email,comment) VALUES (?,?)
	// `
	// stmt, err := db.PrepareContext(ctx, script_all)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// defer stmt.Close()

	// for i := 0; i < 100; i++ {
	// 	email := "anton" + strconv.Itoa(i) + "@gmail.com"
	// 	comment := "komentar ke " + strconv.Itoa(i)

	// 	result, err := stmt.ExecContext(ctx, email, comment)
	// 	if err != nil {
	// 		log.Panic(err.Error())
	// 	}

	// 	id, err := result.RowsAffected()
	// 	if err != nil {
	// 		log.Panic(err.Error())
	// 	}
	// 	fmt.Printf("Comment Status %d - %d\n", i, id)
	// }

	tx, err := db.BeginTx(ctx, nil)
	DB_COMMENT := "comment"
	if err != nil {
		log.Panic(err.Error())
	}
	script_all := `
		INSERT    
		INTO ` + DB_COMMENT + ` 
		(email,comment) VALUES (?,?) 
	`
	stmt, err := tx.PrepareContext(ctx, script_all)
	if err != nil {
		log.Panic(err.Error())
	}
	defer stmt.Close()

	for i := 101; i < 200; i++ {
		email := "anton" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			tx.Rollback()
			log.Panic(err.Error())
			return
		}

		id, err := result.RowsAffected()
		if err != nil {
			log.Panic(err.Error())
		}
		fmt.Printf("Comment Status %d - %d\n", i, id)
	}

	err = tx.Commit()
	if err != nil {
		log.Panic(err.Error())
	}
}
