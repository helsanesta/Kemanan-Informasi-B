package main

// Import package fmt untuk formatting I/O (Input dan Output)
// Import package os untuk mengakses fitur sistem operasi
// Import package io untuk operasi input dan output
// Import package bufio untuk membaca inputan secara buffer
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Fungsi untuk mengecek error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Fungsi main untuk menjalankan program
func main() {
	fmt.Println("vim-go hoooi")
	readFiles()
	writeFiles()
}

// Fungsi untuk membaca file
func readFiles() {
	// Membaca file dengan nama dat
	dat, err := os.ReadFile("./tmp/dat")
	check(err)
	fmt.Println(string(dat))

	// Membuka file dengan nama dat
	f, err := os.Open("./tmp/dat")
	check(err)

	// Membuat buffer baru untuk membaca file
	b1 := make([]byte, 5) // Membuat buffer dengan panjang 5 byte
	n1, err := f.Read(b1) // Membaca file dengan buffer b1
	check(err)
	// Mencetak hasil pembacaan file dengan buffer b1 dengan panjang n1
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Mencari posisi offset dari file
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2) // Membuat buffer dengan panjang 2 byte
	n2, err := f.Read(b2) // Membaca file dengan buffer b2
	check(err)
	// Mencetak hasil pembacaan file dengan offset o2 dengan buffer b2
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	// Mencetak hasil pembacaan file dengan buffer b2 dengan panjang n2
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	// Membaca file minimal 2 byte dengan buffer b3
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	// Mencetak hasil pembacaan file dengan offset o3 dengan buffer b3
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Mengembalikan posisi offset ke awal file
	_, err = f.Seek(0, 0)
	check(err)

	// Membuat buffer baru untuk membaca file
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5) // Membaca 5 byte pertama dari file dengan buffer b4
	check(err)
	// Mencetak hasil pembacaan file dengan buffer b4 dengan panjang 5
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}

// Fungsi untuk menulis file
func writeFiles() {
	// Membuat slice byte baru
	d1 := []byte("hello\ngo\n")
	d1[0] = 'H'                                 // Mengubah elemen pertama slice byte d1 menjadi H
	fmt.Println(d1)                             // Mencetak slice byte d1
	fmt.Println(string(d1))                     // Mencetak slice byte d1 dengan konversi ke string
	err := os.WriteFile("./tmp/dat1", d1, 0644) // Menulis file dengan nama dat1
	check(err)

	// Membuka file dengan nama dat2 untuk ditulis
	f, err := os.Create("./tmp/dat2")
	check(err)

	defer f.Close()

	// Menulis slice byte d2 ke file
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2) // Menulis slice byte d2 ke file
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//  Menulis string ke file
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Memastikan data ditulis ke file
	f.Sync()

	// Membuat buffer baru untuk menulis file
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Memastikan data dalam buffer ditulis ke file
	w.Flush()
}
