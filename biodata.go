package main

import (
	"fmt"
	"os"
)

// Membuat struct untuk mengsetor data
type Teman struct {
    Nama     	string
    Alamat  	string
    Pekerjaan	string
    Alasan   	string
}

// Membuat map untuk beberapa data yang sesuai dengan nomor absen
var sahabat = map[int]Teman{
    1: {"Dodi", "Merdeka 3 No. 12", "Student", "Ilmu coding Baru"},
    2: {"Andrew", "Neli Murni Raya 8", "Ex-Marketing", "Peluang kerja baru"},
	3: {"Ridho", "Menteng Asri", "Copywriter", "Referensi dari teman"},
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Pengunaan: go run biodata.go <nomor absen>")
        return
    }

    nomorAbsen := os.Args[1]
    IDTeman := 1 // Default absen teman pertama

    // Mengubah argument menjadi integer
    _, err := fmt.Sscanf(nomorAbsen, "%d", &IDTeman)
    if err != nil {
        fmt.Println("Input tidak Valid. Berikan nomor absen yang valid.")
        return
    }

    // Menerima data sesuai absen yang di dapat dan apabila tidak ada akan mengprint error
    teman, ok := sahabat[IDTeman]
    if !ok {
        fmt.Println("Teman tidak ditemukan.")
        return
    }

    // Mendisplay data
    fmt.Println("Nama:", teman.Nama)
    fmt.Println("Alamat:", teman.Alamat)
    fmt.Println("Pekerjaan:", teman.Pekerjaan)
    fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
