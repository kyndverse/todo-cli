package tui

import "charm.land/lipgloss/v2"

var (
	// ==========================================
	// PALET WARNA
	// ==========================================
	
	primary = lipgloss.Color("#7D56F4") // Ungu utama khas ekosistem Charm
	accent  = lipgloss.Color("#EE6FF8") // Magenta untuk elemen yang disorot (highlight)
	text    = lipgloss.Color("#FAFAFA") // Putih terang untuk teks utama
	gray    = lipgloss.Color("#626262") // Abu-abu redup untuk elemen pasif
	green   = lipgloss.Color("#04B575") // Hijau untuk indikator sukses/selesai
	red     = lipgloss.Color("#E88388") // Merah untuk peringatan/kesalahan

	// ==========================================
	// GAYA TATA LETAK UTAMA (LAYOUT)
	// ==========================================

	// appBoxStyle adalah gaya untuk kontainer (jendela) utama aplikasi.
	// Gaya ini memberikan border melengkung, jarak dalam/luar (padding/margin), 
	// dan mengunci lebar aplikasi (Width: 65) agar konsisten seperti "kartu".
	appBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(primary).
		Padding(1, 2).
		Margin(1, 2).
		Width(65)

	// titleStyle digunakan untuk merender header/judul aplikasi di bagian atas.
	// Ditampilkan dengan huruf tebal dan latar belakang warna blok (primary).
	titleStyle = lipgloss.NewStyle().
		Background(primary).
		Foreground(text).
		Bold(true).
		Padding(0, 1)

	// ==========================================
	// GAYA DAFTAR TODO (LIST VIEW)
	// ==========================================

	// cursorStyle menyorot teks todo dan ikon kursor (❯) pada baris 
	// yang saat ini sedang dipilih/difokuskan oleh pengguna.
	cursorStyle = lipgloss.NewStyle().Foreground(accent).Bold(true)

	// completedStyle digunakan untuk merender todo yang sudah selesai.
	// Teks akan dicoret (strikethrough) dan warnanya diredupkan (gray).
	completedStyle = lipgloss.NewStyle().Strikethrough(true).Foreground(gray)

	// checkMarkStyle memberikan warna hijau khusus untuk ikon centang (✓).
	checkMarkStyle = lipgloss.NewStyle().Foreground(green)

	// ==========================================
	// GAYA FORM INPUT (INPUT VIEW)
	// ==========================================

	// inputBoxStyle membingkai area teks untuk menyerupai text-area saat 
	// pengguna mengetik todo baru. Lebarnya disesuaikan (61) agar pas 
	// di dalam appBoxStyle tanpa menabrak border utama.
	inputBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(accent).
		Padding(0, 1).
		Width(61)

	// ==========================================
	// GAYA STATUS & FOOTER (BANTUAN)
	// ==========================================

	// errorStyle digunakan untuk merender pesan kesalahan dengan warna merah tebal.
	errorStyle = lipgloss.NewStyle().Foreground(red).Bold(true)

	// helpKeyStyle menyorot tombol pintasan (shortcut) keyboard pada menu bantuan 
	// di bagian bawah layar (misal: teks "enter", "q", "a").
	helpKeyStyle = lipgloss.NewStyle().
		Foreground(accent).
		Bold(true)

	// helpDescStyle digunakan untuk teks penjelasan aksi dari tombol pintasan 
	// (misal: "Selesai", "Keluar"). Diberikan MarginRight(3) untuk menciptakan 
	// jarak antar-menu secara otomatis saat disejajarkan secara horizontal.
	helpDescStyle = lipgloss.NewStyle().
		Foreground(gray).
		MarginRight(3)
)