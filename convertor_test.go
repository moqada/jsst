package main

import "os"

func ExampleConvertor() {
	con, _ := ReadFile("./example.json")
	con.Extract()
	con.Write(os.Stdout)
	// Output:
	// package main
	//
	// type Info struct {
	// 	Content     string `json:"content"`
	// 	ID          string `json:"id"`
	// 	PublishedAt string `json:"publishedAt"`
	// 	Title       string `json:"title"`
	// }
	// type User struct {
	// 	AddressCity  string `json:"addressCity"`
	// 	AddressLine1 string `json:"addressLine1"`
	// 	AddressLine2 string `json:"addressLine2"`
	// 	AddressState string `json:"addressState"`
	// 	AddressZip   string `json:"addressZip"`
	// 	Birthday     string `json:"birthday"`
	// 	FirstName    string `json:"firstName"`
	// 	ID           string `json:"id"`
	// 	Infos        []Info `json:"infos"`
	// 	LastName     string `json:"lastName"`
	// 	RegisteredAt string `json:"registeredAt"`
	// 	Tel          string `json:"tel"`
	// }

}
