package main

import "fmt"

type Ban interface {
	JenisBan() string
}

type BanKaret struct{}
func (kar BanKaret) JenisBan() string {
	return "Karet"
}

type BanKayu struct{}
func (kay BanKayu) JenisBan() string {
	return "Kayu"
}

type BanBesi struct{}
func (b BanBesi) JenisBan() string {
	return "Besi"
}

type Pintu struct {
	Posisi string
	BunyiKetuk string
	BunyiBuka string
}

func (p Pintu) Ketuk() {
	fmt.Println(p.BunyiKetuk)
}

func (p Pintu) Buka () {
	fmt.Println(p.BunyiBuka)
}

type Mobil struct {
	Roda [4]Ban
	Kanan Pintu
	Kiri Pintu
}

func (m *Mobil) GantiBan(i int, ban Ban){
	if i < 0 || i >= len(m.Roda){
		fmt.Println("Roda mobil invalid")
		return
	}
	m.Roda[i] = ban
}

func main() {
	mobil := Mobil{
		Roda: [4]Ban{BanKaret{}, BanKaret{}, BanKaret{}, BanKaret{}},
		Kanan: Pintu{Posisi: "kanan", BunyiKetuk: "Knock", BunyiBuka: "Beep"},
		Kiri:  Pintu{Posisi: "kiri", BunyiKetuk: "Beep", BunyiBuka: "Knock"},
	}


	mobil.GantiBan(0, BanKayu{})
	mobil.GantiBan(1, BanBesi{})
	mobil.GantiBan(2, BanKaret{})
	mobil.GantiBan(3, BanKayu{})

	fmt.Println("Bunyi Ketuk")
	fmt.Println("Kanan : ")
	mobil.Kanan.Ketuk()
	fmt.Println("Kiri : ")
	mobil.Kiri.Ketuk()

	// Mengetuk dan membuka pintu kiri
	fmt.Println("\n\nBunyi Buka:")
	fmt.Println("Kanan : ")
	mobil.Kanan.Buka()
	fmt.Println("Kiri : ")
	mobil.Kiri.Buka()
}