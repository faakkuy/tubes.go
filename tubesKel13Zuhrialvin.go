// DESKRIPSI:
// Aplikasi ini digunakan oleh pasien dan dokter.
// Data yang diolah adalah data konsultasi antara
// pasien dan dokter

// SPESIFIKASI:
// 1. Pengguna bisa mendaftar ke aplikasi
// sebagai pasien
// 2.Pasien bisa berkonsultasi dengan
// posting ppertanyaan melalui aplikasi
// 3.Dokter dan Pasien bisa memberikan tanggapan
// terhadap pertanyaan dari pasien
// 4.Pengguna yang tidak mendaftar hanya dapat
// melihat forum konsultasi antara dokter dan pasien
// 5.Pertanyaan perlu diberi tag tertentu
// untuk memudahkanpencarian dan pengurutan
// 6.Pengguna bisa mencari pertanyaan tertentu
// berdasarkan tag yang dicari
// 7.Dokter bisa menampilkan topik atau
// tag terurut berdasarkan jumlah pertanyaannya

package main

import "fmt"

const NMAX int = 3000

type ConsultationData struct {
	UserDetails       User
	Question          string
	Responses         [NMAX]ResponseData
	NumberOfResponses int
	Tags              string
}

type User struct {
	Name, Email, Password, Role string
}

type ResponseData struct {
	UserDetails User
	Response    string
}

var NumberOfUsers int
var NumberOfConsultations int

var LoggedInUser User

var Users [NMAX]User
var Consultations [NMAX]ConsultationData

func main() {
	var choice int

	for {

		fmt.Println("               Main Menu               ")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Forum Konsultasi")
		fmt.Println("4. Exit")

		fmt.Print("Pilih: ")
		fmt.Scan(&choice)

		if choice == 1 {
			Register()
		} else if choice == 2 {
			Login()
		} else if choice == 3 {
			ViewForum()
		} else if choice == 4 {
			return
		} else {
			fmt.Println("Invalid choice")
		}
	}
}

func Register() {
	var newUser User

	fmt.Println("                Registration                  ")
	fmt.Print("Kamu dokter atau pasien? : ")
	fmt.Scan(&newUser.Role)
	fmt.Print("Nama: ")
	fmt.Scan(&newUser.Name)
	fmt.Print("Password: ")
	fmt.Scan(&newUser.Password)

	if isAlreadyRegistered(newUser.Password) {
		fmt.Println("Password Terdaftar")
	} else {
		Users[NumberOfUsers] = newUser
		NumberOfUsers++
		fmt.Println("Pendaftaran sukses")
	}
}

func isAlreadyRegistered(password string) bool {
	for i := 0; i < NumberOfUsers; i++ {
		if Users[i].Password == password {
			return true
		}
	}
	return false
}

func Login() {
	var status, password string

	fmt.Println("                 Login                  ")
	fmt.Print("Login dokter atau pasien? ")
	fmt.Scan(&status)
	fmt.Print("Password : ")
	fmt.Scan(&password)
	fmt.Println("------------------------------------------")

	if isAlreadyRegistered(password) {
		fmt.Println("Login successful")

		// Temukan pengguna yang login
		for i := 0; i < NumberOfUsers; i++ {
			if Users[i].Password == password {
				LoggedInUser = Users[i]
				break
			}
		}

		if status == "Pasien" || status == "pasien" {
			PatientConsultationMenu()
		} else if status == "Dokter" || status == "dokter" {
			DoctorConsultationMenu()
		}
	} else {
		fmt.Println("Login failed")
	}
}

func PatientConsultationMenu() {
	var choice int

	for {
		fmt.Println("               Consultation Menu               ")

		fmt.Println("1. Post")
		fmt.Println("2. Search Pertanyaan")
		fmt.Println("3. Komen Pertanyaan")
		fmt.Println("4. Lihat pertanyaan menurun")
		fmt.Println("5. Lihat pertanyaan menaik")
		fmt.Println("6. Back")

		fmt.Print("Pilih : ")
		fmt.Scan(&choice)

		if choice == 1 {
			PostQuestion()
		} else if choice == 2 {
			SearchQuestion()
		} else if choice == 3 {
			RespondToQuestion()
		} else if choice == 4 {
			ViewQuestionsDescending()
		} else if choice == 5 {
			ViewQuestionsAscending()
		} else if choice == 6 {
			return
		} else {
			fmt.Println("Invalid choice")
		}
	}
}

func DoctorConsultationMenu() {
	var choice int

	for {
		fmt.Println("               Consultation Menu               ")
		fmt.Println("1. Komentar")
		fmt.Println("2. Search Pertanyaan")
		fmt.Println("3. Lihat topik")
		fmt.Println("4. Lihat pertanyaan menurun")
		fmt.Println("5. Lihat pertanyaan menaik")
		fmt.Println("6. Back")

		fmt.Print("pilih : ")
		fmt.Scan(&choice)

		if choice == 1 {
			Comment()
		} else if choice == 2 {
			SearchQuestion()
		} else if choice == 3 {
			ViewTopics()
		} else if choice == 4 {
			ViewQuestionsDescending()
		} else if choice == 5 {
			ViewQuestionsAscending()
		} else if choice == 6 {
			return
		} else {
			fmt.Println("Invalid choice")
		}
	}
}

func SearchQuestion() {
	var choice int
	var tag string

	fmt.Println("             Search Question Menu            ")
	fmt.Println("1. Search pertanyaan menurun")
	fmt.Println("2. Search pertanyaan menaik")
	fmt.Println("3. Back to Consultation Menu")
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("tag pertanyaan yang ingin dicari ")
		fmt.Scan(&tag)
		SearchQuestionsDescending(tag)
	} else if choice == 2 {
		fmt.Print("tag pertanyaan yang ingin dicari ")
		fmt.Scan(&tag)
		SearchQuestionsAscending(tag)
	} else if choice == 3 {
		return
	} else {
		fmt.Println("Invalid choice")
	}
}

func SearchQuestionsDescending(tag string) {
	var tags [NMAX]string
	var tagIndex int

	for i := 0; i < NumberOfConsultations; i++ {
		tags[tagIndex] = Consultations[i].Tags
		tagIndex++
	}

	// Binary Search
	index := binarySearch(tags, tagIndex, tag)
	if index != -1 {
		for i := NumberOfConsultations - 1; i >= 0; i-- {
			if Consultations[i].Tags == tag {
				DisplayQuestionAndResponses(i)
			}
		}
	} else {
		fmt.Println("Tag tidak ditemukan")
	}
}

func SearchQuestionsAscending(tag string) {
	var tags [NMAX]string
	var tagIndex int

	for i := 0; i < NumberOfConsultations; i++ {
		tags[tagIndex] = Consultations[i].Tags
		tagIndex++
	}

	// Binary Search
	index := binarySearch(tags, tagIndex, tag)
	if index != -1 {
		for i := 0; i < NumberOfConsultations; i++ {
			if Consultations[i].Tags == tag {
				DisplayQuestionAndResponses(i)
			}
		}
	} else {
		fmt.Println("Tag tidak ditemukan")
	}
}

func SequentialSearch(tags [NMAX]string, size int, tag string) int {
	for i := 0; i < size; i++ {
		if tags[i] == tag {
			return i
		}
	}
	return -1
}

func binarySearch(array [NMAX]string, size int, key string) int {
	low := 0
	high := size - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == key {
			return mid
		} else if array[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func ViewForum() {
	var choice int

	fmt.Println("          Consultation Forum          ")
	fmt.Println("1. Lihat pertanyaan secara terurut menurun")
	fmt.Println("2. Lihat pertanyaan secara terurut menaik")
	fmt.Println("3. Exit")
	fmt.Print("Pilih : ")
	fmt.Scan(&choice)

	if choice == 1 {
		ViewQuestionsDescending()
	} else if choice == 2 {
		ViewQuestionsAscending()
	} else if choice == 3 {
		return
	} else {
		fmt.Println("Invalid choice")
	}
}

func DisplayQuestionAndResponses(index int) {
	fmt.Printf("%d. Pertanyaan: %s ; Penanya: %s (%s)\n", (index + 1), Consultations[index].Question, Consultations[index].UserDetails.Name, Consultations[index].UserDetails.Role)
	for j := 0; j < Consultations[index].NumberOfResponses; j++ {
		fmt.Printf("\tResponse %d: %s ; Responded by: %s (%s)\n", (j + 1), Consultations[index].Responses[j].Response, Consultations[index].Responses[j].UserDetails.Name, Consultations[index].Responses[j].UserDetails.Role)
	}
}

func ViewQuestionsDescending() {
	for i := 1; i < NumberOfConsultations; i++ {
		key := Consultations[i]
		j := i - 1

		for j >= 0 && Consultations[j].Tags < key.Tags {
			Consultations[j+1] = Consultations[j]
			j = j - 1
		}
		Consultations[j+1] = key
	}

	for i := 0; i < NumberOfConsultations; i++ {
		DisplayQuestionAndResponses(i)
	}
}

func ViewQuestionsAscending() {
	for i := 0; i < NumberOfConsultations-1; i++ {
		minIndex := i
		for j := i + 1; j < NumberOfConsultations; j++ {
			if Consultations[j].Tags < Consultations[minIndex].Tags {
				minIndex = j
			}
		}
		Consultations[i], Consultations[minIndex] = Consultations[minIndex], Consultations[i]
	}

	for i := 0; i < NumberOfConsultations; i++ {
		DisplayQuestionAndResponses(i)
	}
}

func PostQuestion() {
	var question, tag string

	fmt.Println("                Post Question              ")
	fmt.Print("Kasih Pertanyaan: ")
	fmt.Scan(&question)
	fmt.Println("Berikan tag pertanyaan dengan suatu kata agar mudah untuk mencari pertanyaan mu")
	fmt.Print("tag pertanyaan : ")
	fmt.Scan(&tag)

	Consultations[NumberOfConsultations].UserDetails = LoggedInUser
	Consultations[NumberOfConsultations].Question = question
	Consultations[NumberOfConsultations].Tags = tag
	Consultations[NumberOfConsultations].NumberOfResponses = 0
	NumberOfConsultations++

	fmt.Println("Question posted successfully")
}

func RespondToQuestion() {
	var questionNumber int
	var response string

	fmt.Print("Enter the number of the question you want to respond to: ")
	fmt.Scan(&questionNumber)

	if questionNumber > 0 && questionNumber <= NumberOfConsultations {
		fmt.Print("Enter your response: ")
		fmt.Scan(&response)

		index := questionNumber - 1
		Consultations[index].Responses[Consultations[index].NumberOfResponses].UserDetails = LoggedInUser
		Consultations[index].Responses[Consultations[index].NumberOfResponses].Response = response
		Consultations[index].NumberOfResponses++
		fmt.Println("Response added successfully")
	} else {
		fmt.Println("Invalid question number")
	}
}

func Comment() {
	var questionNumber int
	var response string

	fmt.Print("Pilih nomor pertanyaan yang ingin ditanggapi: ")
	fmt.Scan(&questionNumber)

	if questionNumber > 0 && questionNumber <= NumberOfConsultations {
		fmt.Print("tanggapan anda: ")
		fmt.Scan(&response)

		index := questionNumber - 1
		Consultations[index].Responses[Consultations[index].NumberOfResponses].UserDetails = LoggedInUser
		Consultations[index].Responses[Consultations[index].NumberOfResponses].Response = response
		Consultations[index].NumberOfResponses++
		fmt.Println("Comment added successfully")
	} else {
		fmt.Println("Invalid question number")
	}
}

func ViewTopics() {
	for i := 0; i < NumberOfConsultations; i++ {
		DisplayQuestionAndResponses(i)
	}
}
