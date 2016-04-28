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
	// 	Content     string    `json:"content,omitempty"`
	// 	ID          string    `json:"id,omitempty"`
	// 	PublishedAt time.Time `json:"publishedAt,omitempty"`
	// 	Title       string    `json:"title,omitempty"`
	// }
	// type Machine struct {
	// 	ID   string `json:"id,omitempty"`
	// 	Name string `json:"name,omitempty"`
	// }
	// type User struct {
	// 	AddressCity  string    `json:"addressCity,omitempty"`
	// 	AddressLine1 string    `json:"addressLine1,omitempty"`
	// 	AddressLine2 string    `json:"addressLine2,omitempty"`
	// 	AddressState string    `json:"addressState,omitempty"`
	// 	AddressZip   string    `json:"addressZip,omitempty"`
	// 	Birthday     string    `json:"birthday"`
	// 	FirstName    string    `json:"firstName"`
	// 	ID           string    `json:"id"`
	// 	Infos        []Info    `json:"infos,omitempty"`
	// 	LastName     string    `json:"lastName"`
	// 	Machine      *Machine  `json:"machine,omitempty"`
	// 	RegisteredAt time.Time `json:"registeredAt"`
	// 	Tel          string    `json:"tel,omitempty"`
	// }
	// type InfoInstancesRequest struct {
	// }
	// type InfoInstancesResponse []Info
	// type UserCreateRequest struct {
	// 	Birthday  string `json:"birthday"`
	// 	FirstName string `json:"firstName"`
	// 	LastName  string `json:"lastName"`
	// 	Password  string `json:"password"`
	// }
	// type UserCreateResponse User
	// type UserSelfRequest struct {
	// }
	// type UserSelfResponse User
}
