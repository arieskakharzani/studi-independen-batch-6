package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	sync.Mutex
	students             []model.Student
	studentStudyPrograms map[string]string
	//add map for tracking login attempts here
	failedLoginAttempts map[string]int
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
		//inisialisasi failedLoginAttempts di sini:
		failedLoginAttempts: make(map[string]int),
	}
}

func ReadStudentsFromCSV(filename string) ([]model.Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // ID, Name and StudyProgram

	var students []model.Student
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		student := model.Student{
			ID:           record[0],
			Name:         record[1],
			StudyProgram: record[2],
		}
		students = append(students, student)
	}
	return students, nil
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	return sm.students
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	sm.Lock()
	defer sm.Unlock()

	var foundStudent model.Student
	for _, student := range sm.students {
		if student.ID == id && student.Name == name {
			foundStudent = student
			break
		}
	}

	if foundStudent.ID == "" {
		if _, ok := sm.failedLoginAttempts[id]; ok {
			sm.failedLoginAttempts[id]++
		} else {
			sm.failedLoginAttempts[id] = 1
		}
		if sm.failedLoginAttempts[id] > 3 {
			return "", fmt.Errorf("Login gagal: Batas maksimum login terlampaui")
		}
		return "", fmt.Errorf("Login gagal: data mahasiswa tidak ditemukan")
	}

	// Reset failed attempts setelah berhasil login
	delete(sm.failedLoginAttempts, id)

	studyProgram, ok := sm.studentStudyPrograms[foundStudent.StudyProgram]
	if !ok {
		return "", fmt.Errorf("Login gagal: Program studi tidak ditemukan untuk mahasiswa dengan ID %s", id)
	}
	welcomeMsg := fmt.Sprintf("Login berhasil: Selamat datang %s! Kamu terdaftar di program studi: %s", name, studyProgram)
	return welcomeMsg, nil
}

func (sm *InMemoryStudentManager) RegisterLongProcess() {
	// 30ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	sm.RegisterLongProcess() // 30ms delay to simulate slow processing. DO NOT REMOVE THIS LINE
	sm.Lock()
	defer sm.Unlock()

	if id == "" || name == "" || studyProgram == "" {
		return "", fmt.Errorf("ID, Name or StudyProgram is undefined!")
	}
	// Periksa apakah StudyProgram valid
	if _, ok := sm.studentStudyPrograms[studyProgram]; !ok {
		return "", fmt.Errorf("Study program %s is not found", studyProgram)
	}
	// Periksa apakah ID sudah digunakan
	for _, student := range sm.students {
		if student.ID == id {
			return "", fmt.Errorf("Registrasi gagal: id sudah digunakan")
		}
	}

	sm.students = append(sm.students, model.Student{
		ID:           id,
		Name:         name,
		StudyProgram: studyProgram,
	})
	registrationSuccess := fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram)
	return registrationSuccess, nil
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	program, ok := sm.studentStudyPrograms[code]
	if !ok {
		return "", fmt.Errorf("Program studi dengan kode %s tidak ditemukan", code)
	}
	return program, nil
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	sm.Lock()
	defer sm.Unlock()

	for i, student := range sm.students {
		if student.Name == name {
			err := fn(&sm.students[i])
			if err != nil {
				return "", err
			}
			return "Program studi mahasiswa berhasil diubah.", nil
		}
	}
	return "", fmt.Errorf("Mahasiswa dengan nama %s tidak ditemukan", name)
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		if programStudi == "" {
			return fmt.Errorf("program studi tidak boleh kosong")
		}
		s.StudyProgram = programStudi
		return nil
	}
}

func (sm *InMemoryStudentManager) ImportStudents(filenames []string) error {
	start := time.Now()

	// Create buffered channel untuk mengirim hasil impor dari goroutine
	resultChan := make(chan error, len(filenames))

	// Create WaitGroup untuk menyinkronkan goroutine
	var wg sync.WaitGroup

	// Launch goroutine untuk setiap file csv
	for _, filename := range filenames {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()

			// Read student data from CSV file
			students, err := ReadStudentsFromCSV(filename)
			if err != nil {
				resultChan <- err
				return
			}

			// Register student secara concurrent
			errChan := make(chan error)
			for _, student := range students {
				go func(student model.Student) {
					// Register student
					_, err := sm.Register(student.ID, student.Name, student.StudyProgram)
					errChan <- err
				}(student)
			}

			// Cek registrasi error
			for range students {
				if err := <-errChan; err != nil {
					resultChan <- err
					return
				}
			}

			resultChan <- nil
		}(filename)
	}

	go func() {
		wg.Wait()         // Menunggu hingga semua goroutine selesai
		close(resultChan) // Tutup channel
	}()

	// Tunggu hingga semua goroutine selesai dan cek hasilnya
	for result := range resultChan {
		if result != nil {
			return result
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Importing students took %s\n", elapsed)

	return nil
}

func (sm *InMemoryStudentManager) SubmitAssignmentLongProcess() {
	// 3000ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) SubmitAssignments(numAssignments int) {
	start := time.Now()

	assignmentQueue := make(chan int, numAssignments)
	done := make(chan bool) // Channel untuk sinyal bahwa semua tugas selesai

	// Start worker goroutines
	for i := 1; i <= 3; i++ {
		go func(workerID int) {
			for {
				assignment, more := <-assignmentQueue // Menerima tugas dari channel
				if !more {
					done <- true // Kirim sinyal bahwa semua tugas selesai
					return       // Keluar dari goroutine jika channel ditutup
				}
				fmt.Printf("Worker %d: Processing assignment %d\n", workerID, assignment)
				sm.SubmitAssignmentLongProcess()
				fmt.Printf("Worker %d: Finished assignment %d\n", workerID, assignment)
			}
		}(i)
	}

	// Queue assignments
	for i := 1; i <= numAssignments; i++ {
		assignmentQueue <- i
	}

	// Close channel
	close(assignmentQueue)

	// Menunggu semua worker selesai
	for i := 1; i <= 3; i++ {
		<-done
	}

	elapsed := time.Since(start)
	fmt.Printf("Submitting %d assignments took %s\n", numAssignments, elapsed)
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Bulk Import Student")
		fmt.Println("6. Submit assignment")
		fmt.Println("7. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "5":
			helper.ClearScreen()
			fmt.Println("=== Bulk Import Student ===")

			// Define the list of CSV file names
			csvFiles := []string{"students1.csv", "students2.csv", "students3.csv"}

			err := manager.ImportStudents(csvFiles)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println("Import successful!")
			}

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')

		case "6":
			helper.ClearScreen()
			fmt.Println("=== Submit Assignment ===")

			// Enter how many assignments you want to submit
			fmt.Print("Enter the number of assignments you want to submit: ")
			numAssignments, _ := reader.ReadString('\n')

			// Convert the input to an integer
			numAssignments = strings.TrimSpace(numAssignments)
			numAssignmentsInt, err := strconv.Atoi(numAssignments)

			if err != nil {
				fmt.Println("Error: Please enter a valid number")
			}

			manager.SubmitAssignments(numAssignmentsInt)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "7":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
