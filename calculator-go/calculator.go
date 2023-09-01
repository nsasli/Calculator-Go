package main

import (
	"fmt"
)

func main() {
    fmt.Println("Program Penghitung Rata-rata")
    
    var n int
    var total float64
    
    fmt.Print("Masukkan jumlah angka yang akan dihitung: ")
    fmt.Scan(&n)
    
    if n <= 0 {
        fmt.Println("Tidak ada angka yang dimasukkan.")
        return
    }
    
    for i := 1; i <= n; i++ {
        var angka float64
        fmt.Printf("Masukkan angka ke-%d: ", i)
        fmt.Scan(&angka)
        total += angka
    }
    
    rataRata := total / float64(n)
    
    fmt.Printf("Rata-rata dari %d angka yang dimasukkan adalah %.2f\n", n, rataRata)
}
