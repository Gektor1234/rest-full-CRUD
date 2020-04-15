package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_"github.com/lib/pq"
	"log"
	"net/http"
)

func main(){
	db, err := sql.Open("postgres", "user=postgres password=xyvafvyf123 dbname=testbd host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	e := echo.New()

	type Card struct {
		Id      int `json:"id"`
		Man_id  int `json:"man_id"`
		Balance int `json:"balance"`
		Number  int `json:"number"`
	}
	type Cards struct {
	Cards   []Card `json:"cards"`
}
	e.POST("/card", func(c echo.Context) error {
		u := new(Card)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "INSERT INTO card (man_id , balance, number)VALUES ($1, $2, $3)"
		res, err := db.Query(sqlStatement, u.Man_id, u.Balance, u.Number)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	})
	e.PUT("/card", func(c echo.Context) error {
		u := new(Card)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "UPDATE card SET man_id=$2,balance=$3,number=$4 WHERE id=$1"
		res, err := db.Query(sqlStatement, u.Id, u.Man_id, u.Balance, u.Number)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, string(u.Id))
	})
	e.GET("/card", func(c echo.Context) error {
		sqlStatment := "SELECT id, man_id, balance, number FROM card order by id"
		rows, err := db.Query(sqlStatment)
		if err != nil{
			fmt.Println(err)
		}
		defer rows.Close()
		result := Cards{}
		for rows.Next(){
			card := Card{}
			err2 := rows.Scan(&card.Id ,&card.Man_id ,&card.Balance ,&card.Number)
			if err2 != nil{
				return err2
			}
			result.Cards = append(result.Cards,card)
		}
		return c.JSON(http.StatusCreated, result)
	})
	e.DELETE("/card/:id", func(c echo.Context) error {
		id := c.Param("id")
		sqlStatement := "DELETE FROM card WHERE id = $1"
		res, err := db.Query(sqlStatement, id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusOK, "Deleted")
		}
		return c.String(http.StatusOK, id+"Deleted")
	})
	e.GET("/card/:id", func(c echo.Context) error {
		id := c.Param("id")
		sqlStatment := "SELECT * FROM card WHERE id = $1"
		res, err := db.Query(sqlStatment,id)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Close()
		result := Cards{}
		card := Card{}
		for res.Next(){
			err2 := res.Scan(&card.Id,&card.Man_id,&card.Balance,&card.Number)
			if err2 != nil{
				return err2
			}
		}
		result.Cards = append(result.Cards,card)
		return c.JSON(http.StatusCreated, result)
	})
	e.Start(":8080")
}