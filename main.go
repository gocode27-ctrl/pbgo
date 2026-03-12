package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"math/rand/v2"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"
)

type Misto struct {
	x int
	y int
}

type AnimovanaPostava struct {
	obrazek           *ebiten.Image // Tady si uložíme obrázek naší postavy
	obdelnikySAnimaci []image.Rectangle
	index             int
	rychlostAnimace   int
	misto             Misto
}

type Blok struct {
	obrazek          *ebiten.Image
	vyrezany_obrazek image.Rectangle
	misto            Misto
	zvetseniZmenseni float32
}

// Hra je náš hlavní herní objekt - drží v sobě všechny věci, které hra potřebuje
type Hra struct {
	pozadi               *ebiten.Image
	obrazekPostavy       AnimovanaPostava // Tady si uložíme obrázek naší postavy
	poleNpc              []AnimovanaPostava
	poleBloku            []Blok
	sirkaHry             int // Šířka okna ve kterém máme hru
	vyskaHry             int // Výška okna ve kterém máme hru
	konechHry            bool
	casKdySkoncilaHra    time.Time
	hudbaVPozadiContext  *audio.Context
	hudbaVPozadiPlayer   *audio.Player
	hudbaGameOverContext *audio.Context
	hudbaGameOverPlayer  *audio.Player
}

func NactuHudbuKonecHry(nazev_souboru string) (*audio.Context, *audio.Player) {
	context := audio.NewContext(160000)
	soubor, err := os.ReadFile(nazev_souboru)
	if err != nil {
		fmt.Println("Soubor s hudbou se nepodařilo načíst", err)
		os.Exit(1)

	}
	hudba, err := mp3.DecodeWithSampleRate(160000, bytes.NewReader(soubor))
	if err != nil {
		fmt.Println("Soubor s hudbou se nepodařilo dekódovat", err)
		os.Exit(1)
	}

	smycka := audio.NewInfiniteLoop(hudba, hudba.Length())
	player, err := context.NewPlayer(smycka)
	if err != nil {
		fmt.Println("Nepovedlo se vytvořit přehrávač", err)
		os.Exit(1)
	}
	return context, player

}

func NaraziDoBloku(postava AnimovanaPostava, blok Blok) bool {
	obdelnikPostavy := resolv.NewRectangleFromTopLeft(
		float64(postava.misto.x),
		float64(postava.misto.y),
		float64(postava.obdelnikySAnimaci[0].Dx())*3,
		float64(postava.obdelnikySAnimaci[0].Dy())*3)
	obdelnikBloku := resolv.NewRectangleFromTopLeft(
		float64(blok.misto.x),
		float64(blok.misto.y),
		float64(blok.vyrezany_obrazek.Dx())*float64(blok.zvetseniZmenseni),
		float64(blok.vyrezany_obrazek.Dy())*float64(blok.zvetseniZmenseni))
	return obdelnikPostavy.IsIntersecting(obdelnikBloku)
}

func NaraziDoNejakehoBloku(postava AnimovanaPostava, bloky []Blok) bool {
	for index := range bloky {
		blok := bloky[index]
		fmt.Println(postava.misto, postava.obdelnikySAnimaci[0].Dx(), postava.obdelnikySAnimaci[0].Dy(), "------", blok.misto, blok.vyrezany_obrazek.Dx(), blok.vyrezany_obrazek.Dy())
		if NaraziDoBloku(postava, blok) {
			return true
		}
	}
	return false
}

// Tady budeme pohybovat postavami
func (h *Hra) Update() error {
	if h.konechHry == true {
		aktualni_cas := time.Now()
		pred_jak_dlouhou_dobou_skoncila_hra := aktualni_cas.Sub(h.casKdySkoncilaHra)
		if pred_jak_dlouhou_dobou_skoncila_hra > time.Duration(3*time.Second) {
			os.Exit(0)
		}
		return nil
	}
	h.obrazekPostavy.index += 1
	var zmacknute_klavesy []ebiten.Key
	zmacknute_klavesy = inpututil.AppendPressedKeys(zmacknute_klavesy)
	for _, zmacknuta_klavesa := range zmacknute_klavesy {
		if zmacknuta_klavesa == ebiten.KeyD {
			h.obrazekPostavy.misto.x += 1
			fmt.Println("Jdi do prava")
			if NaraziDoNejakehoBloku(h.obrazekPostavy, h.poleBloku) {
				fmt.Println("BLOK")
				h.obrazekPostavy.misto.x -= 1
			}
		}
		if zmacknuta_klavesa == ebiten.KeyA {
			h.obrazekPostavy.misto.x -= 1
			fmt.Println("Jdi do leva")
			if NaraziDoNejakehoBloku(h.obrazekPostavy, h.poleBloku) {
				fmt.Println("BLOK")
				h.obrazekPostavy.misto.x += 1
			}
		}
		if zmacknuta_klavesa == ebiten.KeyS {
			h.obrazekPostavy.misto.y += 1
			fmt.Println("Jdi do dolu")
			if NaraziDoNejakehoBloku(h.obrazekPostavy, h.poleBloku) {
				fmt.Println("BLOK")
				h.obrazekPostavy.misto.y -= 1
			}
		}
		if zmacknuta_klavesa == ebiten.KeyW {
			h.obrazekPostavy.misto.y -= 1
			fmt.Println("Jdi do nahoru")
			if NaraziDoNejakehoBloku(h.obrazekPostavy, h.poleBloku) {
				fmt.Println("BLOK")
				h.obrazekPostavy.misto.y += 1
			}
		}
	}
	for index := range h.poleNpc {
		if Protinajise(h.obrazekPostavy, h.poleNpc[index]) {
			h.konechHry = true
			h.casKdySkoncilaHra = time.Now()
			h.hudbaGameOverPlayer.Play()
		}
		h.poleNpc[index].index += 1
		nahodny_smer_horizontalne := rand.Int() % 10
		//nahodny_smer_vertikalne := rand.Int() % 50
		if nahodny_smer_horizontalne == 0 {
			if h.obrazekPostavy.misto.x > h.poleNpc[index].misto.x {
				h.poleNpc[index].misto.x += index + 1
			}
			if h.obrazekPostavy.misto.x < h.poleNpc[index].misto.x {
				h.poleNpc[index].misto.x -= index + 1
			}
		}
		if nahodny_smer_horizontalne == 0 {
			if h.obrazekPostavy.misto.y > h.poleNpc[index].misto.y {
				h.poleNpc[index].misto.y += index + 1
			}
			if h.obrazekPostavy.misto.y < h.poleNpc[index].misto.y {
				h.poleNpc[index].misto.y -= index + 1
			}
		}
	}
	return nil
}

func NakresliAnimovanouPostavu(postava AnimovanaPostava, obrazovka *ebiten.Image) {
	vlastnosti_obrazku := &ebiten.DrawImageOptions{}
	vlastnosti_obrazku.GeoM.Scale(3, 3)
	vlastnosti_obrazku.GeoM.Translate(float64(postava.misto.x), float64(postava.misto.y))
	pocet_obrazku := len(postava.obdelnikySAnimaci)
	index_obrazku := (postava.index / postava.rychlostAnimace) % pocet_obrazku
	ctverec_obrazku := postava.obdelnikySAnimaci[index_obrazku]
	oriznuty_obrazek := postava.obrazek.SubImage(ctverec_obrazku)
	obrazovka.DrawImage(ebiten.NewImageFromImage(oriznuty_obrazek), vlastnosti_obrazku)
}

// Draw se také volá každý snímek - tady kreslíme vše na obrazovku
func (h *Hra) Draw(obrazovka *ebiten.Image) {
	vlastnosti_pozadi := &ebiten.DrawImageOptions{}
	obrazovka.DrawImage(h.pozadi, vlastnosti_pozadi)

	NakresliAnimovanouPostavu(h.obrazekPostavy, obrazovka)
	for index := range h.poleNpc {
		NakresliAnimovanouPostavu(h.poleNpc[index], obrazovka)
	}

	for index := range h.poleBloku {
		blok := h.poleBloku[index]
		vlastnosti_bloku := &ebiten.DrawImageOptions{}
		vlastnosti_bloku.GeoM.Scale(float64(blok.zvetseniZmenseni), float64(blok.zvetseniZmenseni))
		vlastnosti_bloku.GeoM.Translate(float64(blok.misto.x), float64(blok.misto.y))
		oriznuty_obrazek := blok.obrazek.SubImage(blok.vyrezany_obrazek)
		obrazovka.DrawImage(ebiten.NewImageFromImage(oriznuty_obrazek), vlastnosti_bloku)
	}

}

// Layout říká, jak velké má být herní okno
func (h *Hra) Layout(vnejsiSirka int, vnejsiVyska int) (int, int) {
	// Vrátíme stejnou velikost, jakou má okno
	h.sirkaHry = vnejsiSirka
	h.vyskaHry = vnejsiVyska
	return vnejsiSirka, vnejsiVyska
}

func NactiObrazek(cesta_k_obrazku string) *ebiten.Image {
	// Načteme obrázek postavy z knihovny
	obrazek_ze_souboru, err := os.ReadFile(cesta_k_obrazku)
	if err != nil {
		fmt.Println("Nepodařilo se otevřít obrázek", cesta_k_obrazku, "chyba:", err)
		os.Exit(1)
	}
	obrazek, _, err := image.Decode(bytes.NewReader(obrazek_ze_souboru))
	if err != nil {
		// Pokud se obrázek nepodařilo načíst, vypíšeme chybu a ukončíme program
		fmt.Println("Obrázek nešel načíst chyba:", err)
		os.Exit(1)
	}
	return ebiten.NewImageFromImage(obrazek)
}

// main je hlavní funkce - tady naše hra začína
func main() {
	// Vytvoříme si novou hru
	var hra Hra

	// Převedeme obrázek do formátu, který umí Ebiten používat

	hra.obrazekPostavy.obrazek = NactiObrazek("./Ninja.png")
	hra.obrazekPostavy.rychlostAnimace = 30
	hra.obrazekPostavy.obdelnikySAnimaci = append(hra.obrazekPostavy.obdelnikySAnimaci, image.Rect(11, 11, 25, 25))
	hra.obrazekPostavy.obdelnikySAnimaci = append(hra.obrazekPostavy.obdelnikySAnimaci, image.Rect(52, 11, 70, 25))
	hra.obrazekPostavy.obdelnikySAnimaci = append(hra.obrazekPostavy.obdelnikySAnimaci, image.Rect(92, 11, 108, 25))
	hra.pozadi = NactiObrazek("./Background_1.png")
	hra.obrazekPostavy.misto.x = 100
	hra.obrazekPostavy.misto.y = 100

	var npc1 AnimovanaPostava
	npc1.obrazek = NactiObrazek("./ghost.png")
	npc1.rychlostAnimace = 30
	npc1.obdelnikySAnimaci = append(npc1.obdelnikySAnimaci, image.Rect(9, 82, 37, 125))
	npc1.obdelnikySAnimaci = append(npc1.obdelnikySAnimaci, image.Rect(57, 82, 85, 125))
	npc1.obdelnikySAnimaci = append(npc1.obdelnikySAnimaci, image.Rect(104, 82, 133, 125))
	npc1.misto.x = -1000000000
	npc1.misto.y = -1000000000
	hra.poleNpc = append(hra.poleNpc, npc1)

	var npc2 AnimovanaPostava
	npc2.obrazek = NactiObrazek("./ghost.png")
	npc2.rychlostAnimace = 30
	npc2.obdelnikySAnimaci = append(npc2.obdelnikySAnimaci, image.Rect(9, 82, 37, 125))
	npc2.obdelnikySAnimaci = append(npc2.obdelnikySAnimaci, image.Rect(57, 82, 85, 125))
	npc2.obdelnikySAnimaci = append(npc2.obdelnikySAnimaci, image.Rect(104, 82, 133, 125))
	npc2.misto.x = -1000000000
	npc2.misto.y = -1000000000
	hra.poleNpc = append(hra.poleNpc, npc2)

	hra.hudbaVPozadiContext, hra.hudbaVPozadiPlayer = NactuHudbu("./with_me.mp3")
	hra.hudbaGameOverContext, hra.hudbaGameOverPlayer = NactuHudbu("./Game Over.mp3")
	hra.hudbaVPozadiPlayer.Play()

	var blok1 Blok
	blok1.obrazek = NactiObrazek("./Rock Pile.png")
	blok1.vyrezany_obrazek = image.Rect(0, 0, 160, 160)
	blok1.misto.x = 400
	blok1.misto.y = 200
	blok1.zvetseniZmenseni = 1
	hra.poleBloku = append(hra.poleBloku, blok1)

	// Spustíme hru! Tady už se program dostane do smyčky a začne volat Update a Draw
	ebiten.RunGame(&hra)
}

func NactuHudbu(nazev_souboru string) (*audio.Context, *audio.Player) {
	context := audio.CurrentContext()
	if context == nil {
		context = audio.NewContext(160000)
	}

	soubor, err := os.ReadFile(nazev_souboru)
	if err != nil {
		fmt.Println("Soubor s hudbou se nepodařilo načíst", err)
		os.Exit(1)
	}
	hudba, err := mp3.DecodeWithSampleRate(160000, bytes.NewReader(soubor))
	if err != nil {
		fmt.Println("Soubor s hudbou se nepodařilo dekódovat", err)
		os.Exit(1)
	}

	smycka := audio.NewInfiniteLoop(hudba, hudba.Length())
	player, err := context.NewPlayer(smycka)
	if err != nil {
		fmt.Println("Nepovedlo se vytvořit přehrávač", err)
		os.Exit(1)
	}
	return context, player
}
func Protinajise(a AnimovanaPostava, b AnimovanaPostava) bool {
	obdelnikA := resolv.NewRectangleFromTopLeft(float64(a.misto.x), float64(a.misto.y), float64(a.obdelnikySAnimaci[0].Dx())*3, float64(a.obdelnikySAnimaci[0].Dy())*3)
	obdelnikB := resolv.NewRectangleFromTopLeft(float64(b.misto.x), float64(b.misto.y), float64(b.obdelnikySAnimaci[0].Dx())*3, float64(b.obdelnikySAnimaci[0].Dy())*3)

	return obdelnikA.IsIntersecting(obdelnikB)
}
