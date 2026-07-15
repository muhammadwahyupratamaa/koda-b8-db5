package main

import (
	"bufio"
	"context"
	"fmt"
	"koda-b8-db5/config"
	"koda-b8-db5/models"
	"os"
	"strings"
	"os/exec"
	"runtime"

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

	fmt.Print()
	fmt.Println("Tekan enter bila ingin ke menu")
	fmt.Scanln()

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

func editContact(conn *pgx.Conn) {
	var id int
	var nama string
	var email string
	var phone string

	fmt.Print("Masukan ID :")
	fmt.Scan(&id)
	fmt.Print("Masukan Nama :")
	fmt.Scan(&nama)
	fmt.Print("Masukan Email : ")
	fmt.Scan(&email)
	fmt.Print("Masukan Nomornya :")
	fmt.Scan(&phone)

	commandTag, err := conn.Exec(context.Background(),`
	UPDATE contacts SET nama=$1, email=$2, phone=$3, updated_at= NOW() WHERE ID=$4`, nama,email,phone,id)

	if err != nil {
	fmt.Println("Gagal mengedit contact")
	return
	}
	if commandTag.RowsAffected() == 0 {
    fmt.Println("")
	return
	}
	fmt.Println("Berhasil mengedit kontak")
}

func deleteContact(conn *pgx.Conn){
		var id int

		fmt.Print("Masukan id:")
		fmt.Scan(&id)
	
	commandTag,err := conn.Exec(context.Background(), `
	DELETE from contacts WHERE ID=$1`, id)

	if err != nil {
		fmt.Println("Gagal menghapus kontak")
		return
	}
	if commandTag.RowsAffected() == 0 {
		fmt.Println("Kontak tidak di temukan")
		return
	}
	fmt.Println("Berhasil menghapus kontak")
}
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main(){
	
	conn := config.Conn()
	defer conn.Close(context.Background())
for{
	ClearScreen()
	showMenu()
	var choice int
	fmt.Print("Choose :")
	fmt.Scan(&choice)

	switch choice {
	case 1: ClearScreen()
			listContact(conn)
	case 2 :ClearScreen()
			 addContact(conn)
	case 3 :ClearScreen()
			 editContact(conn)
	case 4 :ClearScreen()
			 deleteContact(conn)
	case 5 : fmt.Println("Thanks ya")
			 os.Exit(0)
	default : fmt.Println("menu tidak tersedia")
	}

}
	
}