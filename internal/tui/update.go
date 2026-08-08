package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

// Update adalah fungsi inti dari antarmuka Model pada Bubble Tea yang merespons
// event (pesan) yang masuk, seperti ketikan keyboard, ukuran layar, atau timer.
// 
// Pada implementasi ini, Update bertindak sebagai "Router" pusat. Fungsi ini 
// mencegat input global (seperti ctrl+c untuk keluar), lalu mendelegasikan 
// sisa event keyboard ke handler (fungsi pembantu) yang spesifik berdasarkan 
// mode layar (state) aplikasi saat ini.
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		// ctrl+c adalah aksi global untuk menghentikan aplikasi kapan saja.
		if msg.String() == "ctrl+c" {
			return a, tea.Quit
		}

		// Teruskan pesan ke handler spesifik berdasarkan state aktif
		switch a.state {
		case stateListView:
			return a.updateList(msg)
		case stateInputView:
			return a.updateInput(msg)
		}
	}

	return a, nil
}

// updateList menangani semua interaksi keyboard saat pengguna berada di 
// mode daftar todo (stateListView).
//
// Fitur yang ditangani meliputi navigasi kursor (up/down), menandai 
// status selesai (enter/space), menghapus item (d), dan berpindah 
// ke mode penambahan todo baru (a).
func (a App) updateList(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return a, tea.Quit
	case "up":
		if a.cursor > 0 {
			a.cursor--
		}
	case "down":
		if a.cursor < len(a.todos)-1 {
			a.cursor++
		}
	case "enter", "space":
		if len(a.todos) > 0 {
			t := a.todos[a.cursor]
			if err := a.service.Done(t.ID); err != nil {
				a.err = err.Error()
			}
			a.todos = a.service.List()
		}
	case "d":
		if len(a.todos) > 0 {
			t := a.todos[a.cursor]
			if err := a.service.Delete(t.ID); err != nil {
				a.err = err.Error()
			}
			a.todos = a.service.List()
			
			// Menjaga agar kursor tidak melebihi batas indeks ketika
			// item terakhir pada list dihapus.
			if a.cursor >= len(a.todos) && a.cursor > 0 {
				a.cursor--
			}
		}
	case "a":
		// Pindah ke mode input dan reset nilai input & error sebelumnya.
		a.state = stateInputView
		a.input = ""
		a.err = ""
	}
	return a, nil
}

// updateInput menangani interaksi keyboard saat pengguna berada di mode
// form penulisan todo baru (stateInputView).
//
// Alih-alih mengeksekusi perintah (command), mode ini menangkap karakter
// yang diketik untuk dirangkai menjadi string (a.input). Fungsi ini juga 
// menangani penghapusan karakter (backspace), penyimpanan data (enter), 
// dan pembatalan operasi (esc).
func (a App) updateInput(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		a.state = stateListView
	case "enter":
		if strings.TrimSpace(a.input) != "" {
			if err := a.service.Add(strings.TrimSpace(a.input)); err != nil {
				a.err = err.Error()
			}
			a.todos = a.service.List()
		}
		a.state = stateListView
	case "backspace":
		if len(a.input) > 0 {
			// Mengonversi input ke slice rune agar penghapusan karakter
			// multibyte (seperti emoji atau karakter non-Latin) aman.
			runes := []rune(a.input)
			a.input = string(runes[:len(runes)-1])
		}
	case "space":
		a.input += " "
	default:
		// Menangkap input karakter standar. Pemeriksaan panjang string == 1
		// memastikan bahwa input dari tombol kontrol (misal: "tab", "shift") 
		// tidak ikut tercetak ke layar.
		if len(msg.String()) == 1 {
			a.input += msg.String()
		}
	}
	return a, nil
}