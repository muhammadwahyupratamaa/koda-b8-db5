package main

import (
	"bufio"
	"context"
	"fmt"
	"koda-b8-db5/config"
	"koda-b8-db5/models"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func showMenu(){
	fmt.Println("==== CONTACT ====")
	fmt.Println("1.List Contact")
	fmt.Println("2.Tambah list contact")
	fmt.Println("3.Edit Contact")
	fmt.Println("4.Hapus Contact")
	fmt.Println("5.Exit")
}

func listContact(conn *pgx.Conn){
	rows, err := conn.Query(context.Background(), `
		SELECT	id,nama,email,phone,created_at,updated_at FROM contacts
	`)
	if err != nil {
		panic(err)
	}

	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Contact])
	if err != nil {
		panic(err)
	}

	fmt.Println("=== LIST CONTACT ===")
	for _, contact := range contacts{
		fmt.Printf("ID : %d \n Nama: %s \n Email : %s \n Phone: %s \n\n", contact.ID, contact.Nama,contact.Email,contact.Phone)
	}
}

func addContact(conn *pgx.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name :")
	nama , _ := reader.ReadString('\n')
	fmt.Print("Email :")
	email, _ := reader.ReadString('\n')
	fmt.Print("Phone :")
	phone, _ := reader.ReadString('\n')

	nama = strings.TrimSpace(nama)
	email = strings.TrimSpace(email)
	phone = strings.TrimSpace(phone)

	_, err := conn.Exec(context.Background(),`
	INSERT INTO contacts (nama,email,phone) VALUES ($1,$2,$3)`, nama,email,phone)

	if err != nil {
	fmt.Println("Gagal menambah contact")
	return
	}
	fmt.Println("Contact berhasil ditambahkan")
}

func main(){
	
	conn := config.Conn()
	defer conn.Close(context.Background())

	showMenu()
	var choice int
	fmt.Print("Choose :")
	fmt.Scan(&choice)

	switch choice {
	case 1: listContact(conn)
	case 2 : addContact(conn)
	case 3 : fmt.Println("Masih tahap pembuatan :)")
	case 4 : fmt.Println("Masih tahap pembuatan :')")
	case 5 : fmt.Println("Thanks ya")
			 os.Exit(0)
	default : fmt.Println("menu tidak tersedia")
	}

	
}