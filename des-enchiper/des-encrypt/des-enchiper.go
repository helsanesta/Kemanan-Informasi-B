// Description: Enchiper file dengan algoritma DES
package main

import (
	"bufio"
	"crypto/des" // library des untuk enkripsi dan dekripsi DES
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// key untuk enkripsi dan dekripsi DES dengan panjang 8 byte (64 bit)
	key := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("vim-go hoooi")

	// Membaca file hello.go.txt
	plaintext, err := os.ReadFile("./hello.go.txt")
	check(err)

	fmt.Println(len(plaintext))

	// Menghitung jumlah blok yang akan dienkripsi
	// Dengan cara membagi panjang plaintext dengan panjang blok (8 byte)
	var num int
	num = len(plaintext) / 8
	fmt.Println("jumlah blok : ", num)

	// Membuat chiper block dengan key yang telah ditentukan
	block, err := des.NewCipher(key)
	check(err)

	// Inisisasi ciphertext dengan panjang yang sama dengan plaintext
	ciphertext := make([]byte, len(plaintext))

	// Mencetak 8 byte pertama dari plaintext
	fmt.Println(plaintext[0:8])
	fmt.Println(string(plaintext[0:8]))

	// Melakukan enkripsi pada blok pertama plaintext saja
	block.Encrypt(ciphertext, plaintext)

	// Mencetak 8 byte pertama dari ciphertext
	fmt.Println(ciphertext[0:8])
	fmt.Println(string(ciphertext[0:8]))
}

func readFiles() {
	dat, err := os.ReadFile("./tmp/dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("./tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}

func writeFiles() {

	d1 := []byte("hello\ngo\n")
	d1[0] = 'H'
	fmt.Println(d1)
	fmt.Println(string(d1))
	err := os.WriteFile("./tmp/dat", d1, 0644)
	check(err)

	f, err := os.Create("./tmp/dat")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
