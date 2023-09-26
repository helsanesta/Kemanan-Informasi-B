// Description: Enkripsi dan dekripsi file dengan algoritma AES
package main

import (
	"bufio"
	"crypto/aes" // Library AES untuk enkripsi dan dekripsi AES
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
	// encrypt()
	decrypt()
}

func encrypt() {

	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8} // Replace with your AES key

	plaintext, err := os.ReadFile("./hello.go.txt")
	check(err)

	fmt.Println(len(plaintext))
	var num int
	num = len(plaintext) / aes.BlockSize
	fmt.Println("jumlah blok : ", num)

	block, err := aes.NewCipher(key)
	check(err)

	ciphertext := make([]byte, len(plaintext))

	// idx adalah nomor urutan blok dimulai dari 1
	idx := 1
	end := 0
	for num > 0 {
		// Menghitung index awal untuk blok yang akan dienkripsi
		// karena index dimulai dari 0, maka idx dikurangi 1
		// aes.BlockSize adalah panjang blok AES (16 byte)
		start := aes.BlockSize * (idx - 1)

		// Menghitung index akhir untuk blok yang akan dienkripsi
		end = aes.BlockSize * idx

		// Jika index akhir melebihi panjang plaintext
		// Maka index akhir diubah menjadi panjang plaintext
		// Memastikan bahwa index akhir tidak melebihi panjang plaintext
		if end > len(plaintext) {
			end = len(plaintext)
		}

		// Jika panjang blok kurang dari panjang blok AES (16 byte)
		// Maka panjang blok ditambah padding dengan panjang yang dibutuhkan
		// Padding ditambahkan ke akhir plaintext
		if end-start < aes.BlockSize {
			padding := make([]byte, aes.BlockSize-(end-start))
			plaintext = append(plaintext, padding...)
		}

		fmt.Println("blok ke ", idx, " : ")
		fmt.Println(string(plaintext[start:end]))
		fmt.Println(plaintext[start:end])
		block.Encrypt(ciphertext[start:end], plaintext[start:end])
		fmt.Println(ciphertext[start:end])
		fmt.Println("--------------")

		idx++
		num--
	}

	err = os.WriteFile("./ciphertext.txt", ciphertext[0:end], 0644)
	check(err)
}

func decrypt() {
	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8} // Replace with your AES key

	ciphertext, err := os.ReadFile("./ciphertext.txt")
	check(err)

	fmt.Println(len(ciphertext))
	var num int
	num = len(ciphertext) / aes.BlockSize
	fmt.Println("jumlah blok : ", num)

	block, err := aes.NewCipher(key)
	check(err)

	plaintext := make([]byte, len(ciphertext))

	idx := 1
	end := 0
	for num > 0 {
		start := aes.BlockSize * (idx - 1)
		end = aes.BlockSize * idx

		if end > len(ciphertext) {
			end = len(ciphertext)
		}

		fmt.Println("blok ke ", idx, " : ")
		fmt.Println(ciphertext[start:end])
		block.Decrypt(plaintext[start:end], ciphertext[start:end])
		fmt.Println(ciphertext[start:end])
		fmt.Println(string(plaintext[start:end]))
		fmt.Println("--------------")

		idx++
		num--
	}

	err = os.WriteFile("./plaintext.txt", plaintext[0:end], 0644)
	check(err)
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
	err := os.WriteFile("./tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("./tmp/dat2")
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
