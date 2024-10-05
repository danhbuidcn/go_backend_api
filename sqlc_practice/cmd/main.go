package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/danhbuidcn/go_backend_api/sqlc_practice/internal/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mdb, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3308)/shopdevgo")
	if err != nil {
		panic(err)
	}
	defer mdb.Close() // Ensure that the connection will be closed when the main function ends

	// execution
	dao := database.New(mdb)

	ctx := context.Background()
	// insert shop
	err = dao.CreateShop(ctx, "shop one")
	if err != nil {
		log.Fatal(err)
	}

	// get list
	shopList, err := dao.GetShops(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(shopList)
}
