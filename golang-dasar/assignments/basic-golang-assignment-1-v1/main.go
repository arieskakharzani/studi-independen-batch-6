package main

import (
	"a21hc3NpZ25tZW50/helper"
	"fmt"
	"strings"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	// Pemisahan data mahasiswa menjadi slice
	studentsData := strings.Split(Students, ", ")

	// Perulangan untuk mencari data mahasiswa
	for _, studentDetail := range studentsData {
		// Pemisahan data mahasiswa menjadi slice (detail)
		splitStudent := strings.Split(studentDetail, "_")

		// Percabangan apakah ID dan Nama cocok
		if splitStudent[0] == id && splitStudent[1] == name {
			return fmt.Sprintf("Login berhasil: %s (%s)", name, splitStudent[2])
		}
	}

	// Jika tidak ditemukan
	return "Login gagal: data mahasiswa tidak ditemukan"
}

func Register(id string, name string, major string) string {
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	// Pemisahan data mahasiswa menjadi slice
	studentsData := strings.Split(Students, ", ")
	// Perulangan untuk mencari data mahasiswa
	for _, studentDetail := range studentsData {
		// Pemisahan data mahasiswa menjadi slice (detail)
		splitStudent := strings.Split(studentDetail, "_")
		if id == splitStudent[0] {
			return "Registrasi gagal: id sudah digunakan"
		}
	}
	studentsData = append(studentsData, fmt.Sprintf("%s_%s_%s", id, name, major))
	Students = strings.Join(studentsData, ", ")

	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

func GetStudyProgram(code string) string {
	if code == "" {
		return "Code is undefined!"
	}
	// Pemisahan data program studi menjadi slice
	studyPrograms := strings.Split(StudentStudyPrograms, ", ")

	// Perulangan untuk mencari data program studi
	for _, studyProgram := range studyPrograms {
		// Pemisahan data program studi menjadi slice (detail)
		splitStudyProgram := strings.Split(studyProgram, "_")

		// Percabangan apakah kode cocok
		if code == splitStudyProgram[0] {
			// Jika kode cocok, maka tampilkan nama program studi
			return splitStudyProgram[1]
		}
	}
	return "Kode program studi tidak ditemukan"
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
