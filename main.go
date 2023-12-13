package main

import (
	"fmt"
	"math/rand"
	// "main/cmd"
)

type Contact struct {
	firstName   string
	lastName    string
	phoneNumber string
	address     string
}

type ContactBook struct {
	book []Contact
}

type SearchObject struct {
	name        string
	phoneNumber string
}

func searchForContact(contactBook []Contact, searchable SearchObject) {
	for _, v := range contactBook {
		if searchable.name != "" && v.firstName == searchable.name {
			fmt.Println("Found the name")
		} else if searchable.phoneNumber != "" && v.phoneNumber == searchable.phoneNumber {
			fmt.Print("found the number")
		}
	}
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}

func populateContactBook(contactBook *ContactBook, numUsers int) {
	// rand.Seed(time.Now().UnixNano())
	for i := 0; i < numUsers; i++ {
		userFirstName := generateRandomString(10)
		userSurName := generateRandomString(10)
		phoneNumber := generateRandomString(10)
		address := generateRandomString(20)
		contactBook.book = append(contactBook.book, Contact{userFirstName, userSurName, phoneNumber, address})
	}
	// fmt.Print(contactBook)
}

func benchmarkContactBookSearch(contactBook *ContactBook) {

}

func main() {
	demoContact := &Contact{"Aneesh", "Chawla", "test", "test"}
	demoContact2 := &Contact{"Aneesh", "Chawla", "test", "test"}
	demoContact3 := &Contact{"Aneesh", "Chawla", "test", "test"}
	demoContact4 := &Contact{"Aneesh", "Chawla", "test", "test"}
	demoContact5 := &Contact{"Aneesh", "Chawla", "test", "test"}

	var contactBook ContactBook
	contactBook.book = append(contactBook.book, *demoContact, *demoContact2, *demoContact3, *demoContact4, *demoContact5)
	populateContactBook(&contactBook, 500)
	searchFor := SearchObject{"", "test"}
	fmt.Println(contactBook)
	go searchForContact(contactBook.book[:100], searchFor)
	go searchForContact(contactBook.book[100:200], searchFor)
	go searchForContact(contactBook.book[200:300], searchFor)
	go searchForContact(contactBook.book[300:400], searchFor)
	// cmd.Execute()
}
