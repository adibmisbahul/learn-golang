package main

import (
	"fmt"
	"strings"
)

func manipulasi(){
	type AtributMahasiswa struct{
		Nim int `json:"nim"`
		Nama string `json:"nama"`
		Prodi string `json:"prodi"`
		Fakultas string `json:"fakultas"`
		Angkatan int `json:"angkatan"`
	}

	Mahasiswa := []AtributMahasiswa{
	{Nim: 251242, Nama: "Adib", Prodi: "Teknik Elektro", Fakultas: "Teknik", Angkatan: 25},
    {Nim: 252352, Nama: "Imam", Prodi: "Teknik Elektro", Fakultas: "Teknik", Angkatan: 24},
    {Nim: 253461, Nama: "Budi", Prodi: "Teknik Informatika", Fakultas: "Teknik", Angkatan: 23},
    {Nim: 254572, Nama: "Andi", Prodi: "Pendidikan Teknik Informatika", Fakultas: "Teknik", Angkatan: 25},
    {Nim: 255683, Nama: "Rizky", Prodi: "Teknik Elektro", Fakultas: "Teknik", Angkatan: 24},
    {Nim: 256794, Nama: "Dimas", Prodi: "Teknik Sipil", Fakultas: "Teknik", Angkatan: 23},
    {Nim: 257815, Nama: "Fajar", Prodi: "Pendidikan Teknik Mesin", Fakultas: "Teknik", Angkatan: 25},
    {Nim: 258926, Nama: "Rafi", Prodi: "Teknik Informatika", Fakultas: "Teknik", Angkatan: 24},
    {Nim: 259137, Nama: "Yoga", Prodi: "Teknik Elektro", Fakultas: "Teknik", Angkatan: 23},
    {Nim: 251348, Nama: "Ilham", Prodi: "Pendidikan Teknik Elektronika", Fakultas: "Teknik", Angkatan: 25},
    {Nim: 252459, Nama: "Farhan", Prodi: "Teknik Sipil", Fakultas: "Teknik", Angkatan: 24},
    {Nim: 253560, Nama: "Aldi", Prodi: "Teknik Informatika", Fakultas: "Teknik", Angkatan: 23},
    {Nim: 254671, Nama: "Rendra", Prodi: "Teknik Elektro", Fakultas: "Teknik", Angkatan: 25},
    {Nim: 255782, Nama: "Bagas", Prodi: "Pendidikan Teknik Mesin", Fakultas: "Teknik", Angkatan: 24},
    {Nim: 256893, Nama: "Daffa", Prodi: "Teknik Informatika", Fakultas: "Teknik", Angkatan: 23},
    {Nim: 257914, Nama: "Arif", Prodi: "Teknik Elektro", Fakultas: "Teknik", Angkatan: 25},
    {Nim: 258125, Nama: "Hendra", Prodi: "Teknik Sipil", Fakultas: "Teknik", Angkatan: 24},
    {Nim: 259236, Nama: "Wahyu", Prodi: "Pendidikan Teknik Informatika", Fakultas: "Teknik", Angkatan: 23},
    {Nim: 251457, Nama: "Kevin", Prodi: "Teknik Elektro", Fakultas: "Teknik", Angkatan: 25},
    {Nim: 252568, Nama: "Reza", Prodi: "Teknik Informatika", Fakultas: "Teknik", Angkatan: 24},
	}

	// create
	Mahasiswa = append(Mahasiswa, AtributMahasiswa{
		Nim: 254356,
		Nama: "fazri",
		Prodi: "Psikologi",
		Fakultas: "Psikologi",
		Angkatan: 23,
	}) 

	for _, v := range Mahasiswa{
		if strings.Contains(strings.ToLower(v.Prodi), "pendidikan")  {
			fmt.Println("->", v )
		}
	}

}

func main() {
	var angka int = 10
	fmt.Println("angka lama ", angka)

	angka = 20 	
	fmt.Println("angka baru", angka)

	for i := 0 ; i < 5; i++{
		fmt.Println(i +1 )
	}

	var x int = 10 
	x += 5

	fmt.Println("x->", x)

	var nama = [3]string{"adib","fajri","imam"}
	fmt.Println(nama)

	for i:= 0 ; i < len(nama); i++ {
		fmt.Println("halo" , nama[i])
	} 

	var randomNumber = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	var number int = randomNumber[0]
	for i := 0 ; i < len(randomNumber); i++{
		if number < randomNumber[i]{
			number = randomNumber[i]
		}
	}
	fmt.Println("big number -> " , number)

	var totalNumber int = 0 
	for i := 0 ; i < len(randomNumber); i++{
		totalNumber += randomNumber[i]
	} 
	fmt.Println("total number ->", totalNumber)
	manipulasi()
}