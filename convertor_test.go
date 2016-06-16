package main

import "os"

func ExampleConvertor() {
	con, _ := ReadFile("./example.json")
	con.Extract()
	con.Write(os.Stdout)
	// Output:
	// package main
	//
	// import (
	// 	"time"
	// )
	//
	// type Info struct {
	// 	Content     string    `json:"content,omitempty" schema:"content"`
	// 	ID          string    `json:"id,omitempty" schema:"id"`
	// 	PublishedAt time.Time `json:"publishedAt,omitempty" schema:"publishedAt"`
	// 	Title       string    `json:"title,omitempty" schema:"title"`
	// }
	// type Machine struct {
	// 	ID   int64  `json:"id,omitempty" schema:"id"`
	// 	Name string `json:"name,omitempty" schema:"name"`
	// }
	// type User struct {
	// 	AddressCity  string    `json:"addressCity,omitempty" schema:"addressCity"`
	// 	AddressLine1 string    `json:"addressLine1,omitempty" schema:"addressLine1"`
	// 	AddressLine2 string    `json:"addressLine2,omitempty" schema:"addressLine2"`
	// 	AddressState string    `json:"addressState,omitempty" schema:"addressState"`
	// 	AddressZip   string    `json:"addressZip,omitempty" schema:"addressZip"`
	// 	Birthday     string    `json:"birthday" schema:"birthday"`
	// 	FirstName    string    `json:"firstName" schema:"firstName"`
	// 	ID           string    `json:"id" schema:"id"`
	// 	Infos        []Info    `json:"infos,omitempty" schema:"infos"`
	// 	LastName     string    `json:"lastName" schema:"lastName"`
	// 	Machine      *Machine  `json:"machine,omitempty" schema:"machine"`
	// 	RegisteredAt time.Time `json:"registeredAt" schema:"registeredAt"`
	// 	Tel          string    `json:"tel,omitempty" schema:"tel"`
	// }
	// type InfoInstancesRequest struct {
	// }
	// type InfoInstancesResponse []Info
	// type UserCreateRequest struct {
	// 	Birthday  string `json:"birthday" schema:"birthday"`
	// 	FirstName string `json:"firstName" schema:"firstName"`
	// 	LastName  string `json:"lastName" schema:"lastName"`
	// 	Password  string `json:"password" schema:"password"`
	// }
	// type UserCreateResponse User
	// type UserSelfRequest struct {
	// }
	// type UserSelfResponse User
}
