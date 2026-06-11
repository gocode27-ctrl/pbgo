package main

import (
	"fmt"
	"os"
	"time"

	gke "github.com/Fanteria/go-krouzek-engine"
	"github.com/hajimehoshi/ebiten/v2"
)

func ProhralJsi(zivoty *int, textSmrti string) {
	obrazovkaSmrti := gke.NastavKonecovouObrazovku(textSmrti, "Zkus to znovu")
	gke.PridejTlacitko(obrazovkaSmrti, "Zkusit znovu", func() {
		*zivoty = 5
		gke.ResetujHru()
	})

	gke.PridejTlacitko(obrazovkaSmrti, "Ukoncit hru", func() { os.Exit(0) })

	gke.ZobrazKonecovouObrazovku(obrazovkaSmrti)
}

func pevnyBlokMraku(x float64, y float64, velikost float64) {
	blok := gke.PridejBlokSVyrezem("./mrak.png", gke.Vyrez{X1: 524, Y1: 425, X2: 124, Y2: 196})
	gke.NastavPozici(blok, x, y)
	gke.NastavZvetseni(blok, velikost)
	gke.NastavBlokovani(blok, true)
}

func pevnyBlok(x float64, y float64, velikost float64) {
	blok := gke.PridejBlok("./Rock Pile.png")
	gke.NastavPozici(blok, x, y)
	gke.NastavZvetseni(blok, velikost)
	gke.NastavBlokovani(blok, true)
}

func pevnyBlokplatforma(x float64, y float64, velikost float64) {
	blok := gke.PridejBlokSVyrezem("./platform2.png", gke.Vyrez{X1: 232, Y1: 78, X2: 661, Y2: 332})
	gke.NastavPozici(blok, x, y)
	gke.NastavZvetseni(blok, velikost)
	gke.NastavBlokovani(blok, true)
}

func pridejNepritele_pacman(hratelna_postava *gke.Postava, vyska float64, levaHranice float64, pravaHranice float64, HracovyZivoty *int,
	CasovaMezeraMeziRanou time.Duration) {
	PosledniRana := time.Now()
	jdeDoPrava := true
	nepritel := gke.PridejNepritele(
		"./pacman.png",
		func(enemy *gke.Postava) []gke.Akce {
			x := gke.ZjistitPoziciX(&hratelna_postava.Blok)
			y := gke.ZjistitPoziciY(&hratelna_postava.Blok)
			x_enemy := gke.ZjistitPoziciX(&enemy.Blok)
			var akce gke.Akce
			if jdeDoPrava {
				x_enemy += 0.5
				akce = gke.AkceJdeVPravo
			} else {
				x_enemy -= 0.5
				akce = gke.AkceJdeVLevo
			}
			if x_enemy > pravaHranice {
				jdeDoPrava = false
			}
			if x_enemy < levaHranice {
				jdeDoPrava = true
			}
			gke.NastavPozici(&enemy.Blok, x_enemy, vyska)
			if gke.ZjistiKontaktSHratelnouPostavou(enemy) {
				gke.NastavPozici(&hratelna_postava.Blok, x-50, y)
				now := time.Now()
				if now.Sub(PosledniRana) >= CasovaMezeraMeziRanou {
					*HracovyZivoty -= 1
					fmt.Println("Moje životy:", *HracovyZivoty)
					PosledniRana = now
				}
				if *HracovyZivoty == 0 {
					ProhralJsi(HracovyZivoty, "Snedl te pacman :(")
				}
			}
			return []gke.Akce{akce}
		},
	)

	gke.NastavPozici(&nepritel.Blok, 300.0, 1826.0)
	gke.NastavZvetseni(&nepritel.Blok, 0.3)
	gke.NastavRychlostPohybu(nepritel, 0.5)
	gke.NastavBlokovani(&nepritel.Blok, false)
	gke.NastavAnimaci(nepritel, gke.AkceJdeVLevo, true,
		[]gke.Vyrez{
			{X1: 19, Y1: 19, X2: 194, Y2: 202},
		})
	gke.NastavAnimaci(nepritel, gke.AkceJdeVPravo, false,
		[]gke.Vyrez{
			{X1: 19, Y1: 19, X2: 194, Y2: 202},
		})
}

func main() {
	HracovyZivoty := 5
	PosledniRana := time.Now()
	CasovaMezeraMeziRanou := 1 * time.Second
	gke.NastavSouradnicivouMrizku(50)
	gke.NastavUrovenLogovani(gke.LogError)
	gke.NastavPozadi("./Pozadi.png")
	gke.NastavRezimPozadi(gke.RezimPozadiVyplnit)
	gke.NastavGravitaci(0.1)

	kamen := gke.PridejBlok("./Rock Pile.png")
	gke.NastavZvetseni(kamen, 0.9)
	gke.NastavPozici(kamen, 175, 1850)
	gke.NastavBlokovani(kamen, true)

	pevnyBlok(200, 1860, 0.3)

	pevnyBlok(160, 1860, 0.3)

	pevnyBlok(100, 1746, 0.3)

	pevnyBlok(250, 1650, 0.3)

	pevnyBlok(550, 1450, 0.3)

	pevnyBlokplatforma(360, 1550, 0.3)

	pevnyBlokplatforma(700, 1360, 0.3)

	pevnyBlokplatforma(800, 1360, 0.3)

	pevnyBlokMraku(1040, 1160, 0.25)

	//pevnyBlokMraku(300, 1850, 0.25)

	vyherniCoin := gke.PridejNepritele("./coin 2.png", func(enemy *gke.Postava) []gke.Akce {
		if gke.ZjistiKontaktSHratelnouPostavou(enemy) {
			ProhralJsi(&HracovyZivoty, "Vyhrál jsi! :)")
		}
		return []gke.Akce{}
	})
	gke.NastavPozici(&vyherniCoin.Blok, 750, 1244)
	gke.NastavZvetseni(&vyherniCoin.Blok, 1)
	gke.NastavBlokovani(&vyherniCoin.Blok, false)
	gke.NastavAnimaci(vyherniCoin, gke.AkceJdeVPravo, false,
		[]gke.Vyrez{
			{X1: 0, Y1: 0, X2: 24, Y2: 24}})

	blok_cislo_n := 0.0
	for blok_cislo_n <= 100 {
		blok1 := gke.PridejBlokSVyrezem("./Gras.png", gke.Vyrez{X1: 62, Y1: 63, X2: 16, Y2: 17})
		gke.NastavPozici(blok1, blok_cislo_n*32, 1900)
		gke.NastavBlokovani(blok1, true)
		blok_cislo_n += 1
	}

	strom_cislo_n := 0
	animace_stromu := []gke.Vyrez{}
	for strom_cislo_n <= 38 {
		animace_stromu = append(animace_stromu, gke.Vyrez{X1: strom_cislo_n * 64, Y1: 0, X2: (strom_cislo_n + 1) * 64, Y2: 64})
		strom_cislo_n += 1
	}

	animovany_blok := gke.PridejAnimovanyBlok("./animated tree.png", 0.1, animace_stromu...)
	gke.NastavZvetseni(animovany_blok, 2.0)
	gke.NastavPozici(animovany_blok, 460, 1770)

	animovany_blok2 := gke.PridejAnimovanyBlok("./animated tree.png", 0.1, animace_stromu...)
	gke.NastavZvetseni(animovany_blok2, 2.0)
	gke.NastavPozici(animovany_blok2, -100, 1770)
	gke.NastavBlokovani(animovany_blok2, true)

	animovany_blok3 := gke.PridejAnimovanyBlok("./animated tree.png", 0.1, animace_stromu...)
	gke.NastavZvetseni(animovany_blok3, 2.0)
	gke.NastavPozici(animovany_blok3, 750, 1244)

	//70 25 51 11 / 65 55 51 40/
	hratelna_postava := gke.PrijdejHratelnouPostavu("./Ninja.png",
		0.1, map[ebiten.Key]gke.Akce{
			ebiten.KeyA:     gke.AkceJdeVLevo,
			ebiten.KeyD:     gke.AkceJdeVPravo,
			ebiten.KeySpace: gke.AkceSkace,
		},
	)
	gke.NastavZvetseni(&hratelna_postava.Blok, 3.0)
	gke.NastavPozici(&hratelna_postava.Blok, 480.0, 1826.0)
	gke.NastavRychlostPohybu(hratelna_postava, 1.6)
	gke.NastavAnimaci(hratelna_postava, gke.AkceStoji, false,
		[]gke.Vyrez{
			{X1: 11, Y1: 11, X2: 25, Y2: 25},
			{X1: 26, Y1: 56, X2: 10, Y2: 40},
			{X1: 26, Y1: 86, X2: 10, Y2: 70},
			{X1: 26, Y1: 116, X2: 10, Y2: 100},
		},
	)

	gke.NastavAnimaci(hratelna_postava, gke.AkceJdeVPravo, false,
		[]gke.Vyrez{
			{X1: 71, Y1: 26, X2: 51, Y2: 10},
			{X1: 65, Y1: 54, X2: 51, Y2: 40},
			{X1: 67, Y1: 86, X2: 50, Y2: 70},
			{X1: 73, Y1: 116, X2: 51, Y2: 100},
			{X1: 69, Y1: 145, X2: 51, Y2: 130},
			{X1: 51, Y1: 160, X2: 73, Y2: 176},
		},
	)

	gke.NastavAnimaci(hratelna_postava, gke.AkceJdeVLevo, true,
		[]gke.Vyrez{
			{X1: 71, Y1: 26, X2: 51, Y2: 10},
			{X1: 65, Y1: 54, X2: 51, Y2: 40},
			{X1: 67, Y1: 86, X2: 50, Y2: 70},
			{X1: 73, Y1: 116, X2: 51, Y2: 100},
			{X1: 69, Y1: 145, X2: 51, Y2: 130},
			{X1: 51, Y1: 160, X2: 73, Y2: 176},
		},
	)

	gke.NastavAnimaci(hratelna_postava, gke.AkceSkace, true,
		[]gke.Vyrez{
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
		},
	)

	gke.NastavAnimaci(hratelna_postava, gke.AkceSkaceVLevo, true,
		[]gke.Vyrez{
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
		},
	)

	gke.NastavAnimaci(hratelna_postava, gke.AkceSkaceVPravo, false,
		[]gke.Vyrez{
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
			{X1: 91, Y1: 6, X2: 109, Y2: 24},
		},
	)

	gke.NastavKameru(hratelna_postava)
	gke.NastavOkrajeKamery(100, 100, 100, 20)

	spike := gke.PridejNepritele("./pixel_art (1).png", func() func(*gke.Postava) []gke.Akce {

		return func(enemy *gke.Postava) []gke.Akce {
			//fmt.Println("Moje životy:", HracovyZivoty)
			if gke.ZjistiKontaktSHratelnouPostavou(enemy) {
				//x := gke.ZjistitPoziciX(&hratelna_postava.Blok)
				//y := gke.ZjistitPoziciY(&hratelna_postava.Blok)
				//gke.NastavPozici(&hratelna_postava.Blok, x-50, y)
				now := time.Now()
				if now.Sub(PosledniRana) >= CasovaMezeraMeziRanou {
					HracovyZivoty -= 1
					fmt.Println("Moje životy:", HracovyZivoty)
					PosledniRana = now
				}
				if HracovyZivoty == 0 {
					ProhralJsi(&HracovyZivoty, "Napichl ses na spike :(")
				}
			}
			return []gke.Akce{}
		}
	}(),
	)

	pridejNepritele_pacman(hratelna_postava, 1620, 100, 360, &HracovyZivoty, CasovaMezeraMeziRanou)
	pridejNepritele_pacman(hratelna_postava, 1400, 100, 360, &HracovyZivoty, CasovaMezeraMeziRanou)

	/*nepritel := gke.PridejNepritele(
		"./pacman.png",
		func(enemy *gke.Postava) []gke.Akce {
			x := gke.ZjistitPoziciX(&hratelna_postava.Blok)
			y := gke.ZjistitPoziciY(&hratelna_postava.Blok)
			x_enemy := gke.ZjistitPoziciX(&enemy.Blok)
			y_enemy := gke.ZjistitPoziciY(&enemy.Blok)
			var akce gke.Akce
			if x > x_enemy {
				x_enemy += 0.5
				akce = gke.AkceJdeVPravo
			}
			if x < x_enemy {
				x_enemy -= 0.5
				akce = gke.AkceJdeVLevo
			}
			if y > y_enemy {
				y_enemy += 1
			}
			if y < y_enemy {
				y_enemy -= 1
			}
			gke.NastavPozici(&enemy.Blok, x_enemy, 1828)
			if gke.ZjistiKontaktSHratelnouPostavou(enemy) {
				gke.NastavPozici(&hratelna_postava.Blok, x-100, y)
			}
			return []gke.Akce{akce}
		},
	)

	gke.NastavPozici(&nepritel.Blok, 300.0, 1828.0)
	gke.NastavZvetseni(&nepritel.Blok, 0.26)
	gke.NastavRychlostPohybu(nepritel, 0.5)
	gke.NastavBlokovani(&nepritel.Blok, false)
	gke.NastavAnimaci(nepritel, gke.AkceJdeVLevo, true,
		[]gke.Vyrez{
			{X1: 19, Y1: 19, X2: 194, Y2: 202},
		})
	gke.NastavAnimaci(nepritel, gke.AkceJdeVPravo, false,
		[]gke.Vyrez{
			{X1: 19, Y1: 19, X2: 194, Y2: 202},
		})
	*/
	gke.NastavPozici(&spike.Blok, 410, 1510)
	gke.NastavZvetseni(&spike.Blok, 0.4)

	gke.NastavBlokovani(&spike.Blok, false)
	gke.NastavAnimaci(spike, gke.AkceJdeVPravo, false,
		[]gke.Vyrez{{X1: 39, X2: 128, Y1: 67, Y2: 160}})

	gke.SpustHru()
}

//
