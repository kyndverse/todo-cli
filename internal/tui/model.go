package tui

import "todo-cli/internal/todo"

// TodoService mendefinisikan kontrak (interface) untuk operasi data Todo.
// Dengan menggunakan interface, lapisan TUI tidak perlu tahu bagaimana data disimpan
// (apakah di memori, file, atau database), sehingga kode menjadi lebih modular dan mudah diuji.
type TodoService interface {
	// List mengembalikan semua daftar todo yang tersimpan.
	List() []todo.Todo
	// Add menambahkan todo baru berdasarkan teks deskripsi yang diberikan.
	Add(string) error
	// Done menandai sebuah todo sebagai selesai berdasarkan ID todo tersebut.
	Done(int) error
	// Delete menghapus todo secara permanen berdasarkan ID todo tersebut.
	Delete(int) error
}

// sessionState merepresentasikan mode layar (state) aplikasi saat ini.
// Tipe ini digunakan sebagai penentu (router) logika Update dan View mana
// yang harus dieksekusi oleh Bubble Tea.
type sessionState int

const (
	// stateListView adalah mode utama di mana pengguna melihat daftar todo,
	// melakukan navigasi kursor, menghapus, atau menandai tugas sebagai selesai.
	stateListView sessionState = iota

	// stateInputView adalah mode form di mana pengguna sedang mengetik todo baru.
	// Pada mode ini, sebagian besar input keyboard akan ditangkap sebagai teks.
	stateInputView
)

// App adalah struktur data utama (Model) dalam The Elm Architecture yang digunakan Bubble Tea.
// Struct ini bertugas menyimpan seluruh kondisi aktual (state) aplikasi pada momen tertentu.
type App struct {
	// service adalah instance yang menangani operasi logika bisnis (CRUD) pada data todo.
	service TodoService

	// todos menyimpan salinan sementara dari daftar todo yang diambil dari service.
	// Data ini digunakan secara langsung oleh View untuk merender daftar ke terminal.
	todos []todo.Todo

	// cursor menyimpan posisi indeks yang sedang disorot (highlight) oleh pengguna
	// pada daftar todo saat aplikasi berada di mode stateListView.
	cursor int

	// state menyimpan mode layar aktif saat ini (contoh: sedang melihat list atau mengisi input).
	state sessionState

	// input menampung karakter teks (string) yang sedang diketik oleh pengguna
	// secara real-time saat aplikasi berada di mode stateInputView.
	input string

	// err menampung pesan kesalahan operasional (jika ada) untuk ditampilkan ke UI,
	// misalnya gagal menyimpan, menghapus, atau masalah koneksi data.
	err string
}