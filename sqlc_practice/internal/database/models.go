// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

type Product struct {
	Productid   int32
	Shopid      int32
	Productname string
}

type Shop struct {
	Shopid   int32
	Shopname string
}
