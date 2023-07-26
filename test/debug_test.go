package test

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	nilaiAsli := 42

	// Mendeklarasikan sebuah pointer dan menetapkan alamat dari 'nilaiAsli' ke pointer tersebut
	pointer := &nilaiAsli

	// Mencetak nilai dari pointer (alamat memori)
	fmt.Printf("Nilai pointer: %p\n", pointer)

	// Mencetak nilai asli yang ditunjuk oleh pointer (dereferencing)
	fmt.Println("Nilai asli:", *pointer)

	// Anda juga bisa mengubah nilai asli melalui pointer
	*pointer = 100

	// Sekarang nilai dari nilaiAsli telah diubah
	fmt.Println("Nilai asli yang telah diubah:", nilaiAsli)
}
