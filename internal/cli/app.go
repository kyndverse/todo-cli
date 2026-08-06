package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"todo-cli/internal/todo"
)

type TodoService interface {
	List() []todo.Todo
	Add(string) error
	Done(int) error
	Delete(int) error
}

type App struct {
	service TodoService
	reader *bufio.Reader
  writer io.Writer
}

func New(service TodoService, reader io.Reader, writer io.Writer) *App {
	return &App{
		service: service,
		reader: bufio.NewReader(reader),
		writer: writer,
	}
}

func (a *App) Run() {
	for {
		a.showMenu()

		choice, err := a.readChoice()
		if err != nil {
			fmt.Fprintln(a.writer, "Input tidak valid")
			continue
		}

		if !a.handleChoice(choice) {
			fmt.Fprintln(a.writer, "Sampai jumpa!")
			return
		}
	}
}

func (a *App) showMenu() {
	fmt.Fprintln(a.writer)
	fmt.Fprintln(a.writer, "\n=== TODO CLI ===")
	fmt.Fprintln(a.writer, "1. Lihat Todo")
	fmt.Fprintln(a.writer, "2. Tambah Todo")
	fmt.Fprintln(a.writer, "3. Tandai Selesai")
	fmt.Fprintln(a.writer, "4. Hapus Todo")
	fmt.Fprintln(a.writer, "5. Keluar")
	fmt.Fprint(a.writer, "Pilih: ")
}

func (a *App) readChoice() (int, error) {
	input, err := a.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return 0, err
	}

	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

func (a *App) handleChoice(choice int) bool {
	switch choice {

	case 1:
	todos := a.service.List()

	if len(todos) == 0 {
		fmt.Fprintln(a.writer, "Belum ada todo.")
		return true
	}

	for _, t := range todos {
		status := " "
		if t.IsCompleted {
			status = "x"
		}

		fmt.Fprintf(
			a.writer,
			"%d. [%s] %s\n",
			t.ID,
			status,
			t.Description,
		)
	}

	case 2:
		fmt.Fprint(a.writer, "Masukkan todo: ")

		text, err := a.reader.ReadString('\n')
		if err != nil && err != io.EOF {
    fmt.Fprintln(a.writer, "Error:", err)
		
    return true
}
		text = strings.TrimSpace(text)

		if err := a.service.Add(text); err != nil {
			fmt.Fprintln(a.writer, "Error:", err)
		} else {
			fmt.Fprintln(a.writer, "Todo berhasil ditambahkan.")
		}

	case 3:
		fmt.Fprint(a.writer, "Nomor todo: ")

		idx, err := a.readChoice()
		if err != nil {
			fmt.Fprintln(a.writer, "Nomor tidak valid.")
			return true
		}

		if err := a.service.Done(idx - 1); err != nil {
			fmt.Fprintln(a.writer, "Error:", err)
		} else {
			fmt.Fprintln(a.writer, "Todo ditandai selesai.")
		}

	case 4:
		fmt.Fprint(a.writer, "Nomor todo: ")

		id, err := a.readChoice()
		if err != nil {
			fmt.Fprintln(a.writer, "Nomor tidak valid.")
			return true
		}

		if err := a.service.Delete(id); err != nil {
			fmt.Fprintln(a.writer, "Error:", err)
		} else {
			fmt.Fprintln(a.writer, "Todo berhasil dihapus.")
		}

	case 5:
		return false

	default:
		fmt.Fprintln(a.writer, "Pilihan tidak tersedia.")
	}

	return true
}