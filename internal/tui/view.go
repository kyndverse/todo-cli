package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// renderHelpItem adalah fungsi pembantu (helper) untuk memformat satu pasang
// tombol pintasan dan teks penjelasannya pada menu bantuan di bagian bawah UI.
// Fungsi ini menggabungkan gaya (style) dari helpKeyStyle dan helpDescStyle
// agar tampilan menu konsisten.
func renderHelpItem(key, desc string) string {
	return helpKeyStyle.Render(key) + " " + helpDescStyle.Render(desc)
}

// View adalah fungsi inti ketiga dari The Elm Architecture pada Bubble Tea.
// Fungsi ini dipanggil setiap kali fungsi Update selesai memodifikasi Model,
// dan bertugas "menggambar ulang" seluruh antarmuka (UI) berdasarkan kondisi 
// (state) terbaru dari aplikasi.
//
// View pada implementasi ini merangkai 3 bagian utama:
// 1. Header (Judul aplikasi)
// 2. Content (Daftar todo atau form input, beserta pesan error jika ada)
// 3. Footer (Menu bantuan pintasan keyboard)
// 
// Semuanya digabungkan secara vertikal lalu dibungkus di dalam kotak aplikasi utama.
func (a App) View() tea.View {
	var content strings.Builder

	// 1. Render pesan error di bagian atas konten jika terdapat kesalahan.
	if a.err != "" {
		errMsg := fmt.Sprintf("Error: %s", a.err)
		fmt.Fprintf(&content, "%s\n\n", errorStyle.Render(errMsg))
	}

	// 2. Render bagian konten utama (Content) berdasarkan mode layar aktif.
	switch a.state {
	case stateListView:
		// Tampilan saat tidak ada data todo.
		if len(a.todos) == 0 {
			content.WriteString(lipgloss.NewStyle().Foreground(gray).Italic(true).Render("Belum ada tugas. Tekan 'a' untuk mulai mencatat."))
			content.WriteString("\n")
		} else {
			// Looping untuk merender setiap baris todo.
			for i, t := range a.todos {
				var row strings.Builder
				isCursor := a.cursor == i

				// Render kursor sorotan
				if isCursor {
					row.WriteString(cursorStyle.Render("❯ "))
				} else {
					row.WriteString("  ")
				}

				// Render ikon status selesai/belum
				if t.IsCompleted {
					row.WriteString(checkMarkStyle.Render("✓ "))
				} else {
					row.WriteString(lipgloss.NewStyle().Foreground(gray).Render("○ "))
				}

				// Terapkan coretan pada teks jika todo sudah selesai,
				// atau tebalkan teks jika sedang disorot oleh kursor.
				desc := t.Description
				if t.IsCompleted {
					desc = completedStyle.Render(desc)
				} else if isCursor {
					desc = cursorStyle.Render(desc)
				}

				row.WriteString(desc)
				fmt.Fprintln(&content, row.String())
			}
		}

	case stateInputView:
		// Tampilan mode form penulisan todo baru.
		content.WriteString(lipgloss.NewStyle().Foreground(primary).Bold(true).Render("Detail Tugas Baru:"))
		content.WriteString("\n")
		
		// Menambahkan kursor blok (█) di akhir teks yang sedang diketik.
		inputField := a.input + cursorStyle.Render("█")
		fmt.Fprintln(&content, inputBoxStyle.Render(inputField))
	}

	// 3. Render bagian Footer (Menu Bantuan)
	var footer string
	if a.state == stateListView {
		// Menggunakan lipgloss.JoinHorizontal untuk menata menu secara sejajar ke kanan.
		// Dibagi menjadi dua baris agar tidak menabrak batas lebar kotak aplikasi.
		
		// Baris 1: Aksi utama
		row1 := lipgloss.JoinHorizontal(
			lipgloss.Left,
			renderHelpItem("↑/↓", "Navigasi"),
			renderHelpItem("enter", "Tandai Selesai"),
			renderHelpItem("a", "Tambah"),
		)
		
		// Baris 2: Aksi destruktif/keluar
		row2 := lipgloss.JoinHorizontal(
			lipgloss.Left,
			renderHelpItem("d", "Hapus"),
			renderHelpItem("q", "Keluar"),
		)
		
		// Gabungkan baris atas dan bawah secara vertikal.
		footer = lipgloss.JoinVertical(lipgloss.Left, row1, row2)
		
	} else {
		// Footer untuk mode input
		footer = lipgloss.JoinHorizontal(
			lipgloss.Left,
			renderHelpItem("enter", "Simpan Todo"),
			renderHelpItem("esc", "Batal"),
		)
	}

	// 4. Susun Header
	header := titleStyle.Render(" ✓ TODO LIST ")

	// 5. Gabungkan Header, Content, dan Footer menjadi satu UI vertikal yang utuh.
	ui := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"", // Spacer atas
		content.String(),
		"", // Spacer bawah
		lipgloss.NewStyle().MarginTop(1).Render(footer), // Memberi sedikit jarak antara konten dan footer
	)

	// 6. Bungkus UI dengan styling kotak utama (appBoxStyle)
	finalRender := appBoxStyle.Render(ui)

	// Kembalikan ke Bubble Tea v2 dalam bentuk struct tea.View
	return tea.NewView(finalRender)
}