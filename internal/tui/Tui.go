package tui

import tea "charm.land/bubbletea/v2"

// New membuat dan mengembalikan instansiasi baru dari App (Model Bubble Tea).
// Fungsi ini menginisialisasi state awal aplikasi: memasukkan service,
// mengambil daftar todo pertama (List()), dan mengatur mode tampilan
// awal ke daftar todo (stateListView).
func New(service TodoService) App {
	return App{
		service: service,
		todos:   service.List(),
		state:   stateListView,
	}
}

// Init adalah fungsi bawaan (lifecycle) dari The Elm Architecture di Bubble Tea
// yang dijalankan persis satu kali saat aplikasi pertama kali dimulai.
// Fungsi ini mengembalikan tea.Cmd (perintah I/O asynchronous).
// Karena kita tidak memerlukan inisialisasi I/O tambahan (seperti tick timer
// atau fetching API), kita cukup mengembalikan nil.
func (a App) Init() tea.Cmd {
	return nil
}

// RunApp adalah titik masuk (entry point) utama untuk menjalankan antarmuka
// terminal (TUI) dari luar package (misalnya dari file main.go).
//
// Fungsi ini menerima TodoService untuk injeksi dependensi, membungkus
// instansiasi App ke dalam tea.NewProgram(), lalu mengeksekusinya (Run()).
// Akan mengembalikan error jika terminal gagal dirender atau terjadi masalah internal.
func RunApp(service TodoService) error {
	p := tea.NewProgram(New(service))
	_, err := p.Run()
	return err
}