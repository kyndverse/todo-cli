package tui

import (
	"fmt"
	"strings"
	"todo-cli/internal/todo"

	tea "charm.land/bubbletea/v2"
)

const (
	reset         = "\033[0m"
	bold          = "\033[1m"
	dim           = "\033[2m"
	strikethrough = "\033[9m"
	blink         = "\033[5m"

	// Warna Foreground (Teks)
	fgCyan  = "\033[36m"
	fgGreen = "\033[32m"
	fgRed   = "\033[31m"
)

// Service interface tetap sama seperti sebelumnya
type TodoService interface {
	List() []todo.Todo
	Add(string) error
	Done(int) error
	Delete(int) error
}

// Kita definisikan state untuk membedakan tampilan
type sessionState int

const (
	stateListView sessionState = iota
	stateInputView
)

// App sekarang bertindak sebagai "Model" di Bubble Tea
type App struct {
	service TodoService
	todos   []todo.Todo
	cursor  int          // Menandai posisi kursor di list
	state   sessionState // Mode saat ini (List atau Input)
	input   string       // Menyimpan teks saat mode penambahan todo
	err     string       // Menyimpan pesan error jika ada
}

func New(service TodoService) App {
	return App{
		service: service,
		todos:   service.List(),
		state:   stateListView,
	}
}

// 1. Init: Dijalankan saat aplikasi dimulai
func (a App) Init() tea.Cmd {
	return nil
}

// 2. Update: Menangani semua interaksi keyboard
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		
		// Jika pengguna menekan Ctrl+C kapan saja, keluar dari aplikasi
		if msg.String() == "ctrl+c" {
			return a, tea.Quit
		}

		// Logika berbeda berdasarkan layar yang sedang aktif
		switch a.state {

		// --- TAMPILAN LIST TODO ---
		case stateListView:
			switch msg.String() {
			case "q": // Keluar
				return a, tea.Quit
			case "up": // Kursor naik
				if a.cursor > 0 {
					a.cursor--
				}
			case "down": // Kursor turun
				if a.cursor < len(a.todos)-1 {
					a.cursor++
				}
			case "enter", "space": // Tandai selesai
				if len(a.todos) > 0 {
					// Ambil objek todo berdasarkan posisi kursor saat ini
					t := a.todos[a.cursor]
					
					// Gunakan t.ID sebagai parameter untuk Done()
					if err := a.service.Done(t.ID); err != nil {
						a.err = err.Error()
					}
					
					// Perbarui daftar todo dari service
					a.todos = a.service.List()
				}
			case "d": // Hapus todo
				if len(a.todos) > 0 {
					// Delete() menggunakan ID todo
					t := a.todos[a.cursor]
					if err := a.service.Delete(t.ID); err != nil {
						a.err = err.Error()
					}
					a.todos = a.service.List()

					// Cegah kursor keluar batas setelah penghapusan
					if a.cursor >= len(a.todos) && a.cursor > 0 {
						a.cursor--
					}
				}
			case "a": // Pindah ke layar tambah todo
				a.state = stateInputView
				a.input = ""
				a.err = ""
			}

		// --- TAMPILAN TAMBAH TODO (INPUT) ---
		case stateInputView:
			switch msg.String() {
			case "esc": // Batal, kembali ke list
				a.state = stateListView
			case "enter": // Simpan todo
				if strings.TrimSpace(a.input) != "" {
					if err := a.service.Add(strings.TrimSpace(a.input)); err != nil {
						a.err = err.Error()
					}
					a.todos = a.service.List()
				}
				a.state = stateListView
			case "backspace": // Hapus karakter terakhir
				if len(a.input) > 0 {
					// Konversi ke rune agar aman untuk karakter Unicode/Emoji
					runes := []rune(a.input)
					a.input = string(runes[:len(runes)-1])
				}
			case "space": // Tangani spasi secara khusus di v2
				a.input += " "
			default:
				// Tambahkan karakter biasa ke input (abaikan tombol kontrol)
				if len(msg.String()) == 1 {
					a.input += msg.String()
				}
			}
		}
	}

	return a, nil
}

func (a App) View() tea.View {
	var b strings.Builder

	// Beri jarak atas
	b.WriteString("\n")

	// Judul (Konstanta digabungkan dengan Fprintf agar efisien)
	fmt.Fprintf(&b, "  %s%s=== TODO TUI ===%s\n\n", bold, fgCyan, reset)

	// Error
	if a.err != "" {
		fmt.Fprintf(&b, "  %s%sError: %s%s\n\n", bold, fgRed, a.err, reset)
	}

	// --- TAMPILAN BERDASARKAN STATE ---
	switch a.state {
	case stateListView:
		if len(a.todos) == 0 {
			fmt.Fprintf(&b, "  %sBelum ada todo. Tekan 'a' untuk menambahkan.%s\n", dim, reset)
		} else {
			for i, t := range a.todos {
				isCursor := a.cursor == i

				// 1. Indikator Kursor
				if isCursor {
					fmt.Fprintf(&b, "  %s%s❯ %s", bold, fgCyan, reset)
				} else {
					b.WriteString("    ") // Spasi kosong agar sejajar
				}

				// 2. Ikon Checkbox
				if t.IsCompleted {
					fmt.Fprintf(&b, "%s✓ %s", fgGreen, reset)
				} else {
					fmt.Fprintf(&b, "%s○ %s", dim, reset)
				}

				// 3. Teks Todo
				if t.IsCompleted {
					// Jika selesai: diredupkan dan dicoret
					fmt.Fprintf(&b, "%s%s%s%s\n", dim, strikethrough, t.Description, reset)
				} else {
					if isCursor {
						// Jika sedang dipilih dan belum selesai: ditebalkan
						fmt.Fprintf(&b, "%s%s%s\n", bold, t.Description, reset)
					} else {
						// Normal
						fmt.Fprintf(&b, "%s\n", t.Description)
					}
				}
			}
		}

		// Menu Bantuan
		fmt.Fprintf(&b, "\n  %s[↑/↓]: Navigasi • [enter/spasi]: Selesai%s\n", dim, reset)
		fmt.Fprintf(&b, "  %s[d]: Hapus • [a]: Tambah Todo • [q]: Keluar%s\n", dim, reset)

	case stateInputView:
		fmt.Fprintf(&b, "  %s%sMasukkan todo baru: %s", bold, fgCyan, reset)
		
		// Input teks dengan kursor berkedip (blink underscore)
		fmt.Fprintf(&b, "%s%s_%s\n\n", a.input, blink, reset)

		fmt.Fprintf(&b, "  %s[enter]: Simpan • [esc]: Batal%s\n", dim, reset)
	}

	// Beri jarak bawah
	b.WriteString("\n")

	return tea.NewView(b.String())
}

// Helper untuk menjalankan aplikasi dari main.go
func RunApp(service TodoService) error {
	p := tea.NewProgram(New(service))
	_, err := p.Run()
	return err
}