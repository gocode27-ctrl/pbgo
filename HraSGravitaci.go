package main

import (
	gke "github.com/Fanteria/go-krouzek-engine"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	gke.NastavSouradnicivouMrizku(50)
	gke.NastavUrovenLogovani(gke.LogError)
	gke.NastavPozadi("./Pozadi.png")
	gke.NastavRezimPozadi(gke.RezimPozadiVyplnit)
	gke.NastavGravitaci(0.1)

	kamen := gke.PridejBlok("./Rock Pile.png")
	gke.NastavZvetseni(kamen, 0.9)
	gke.NastavPozici(kamen, 175, 1850)
	gke.NastavBlokovani(kamen, true)

	kamen2 := gke.PridejBlok("./Rock Pile.png")
	gke.NastavZvetseni(kamen2, 0.3)
	gke.NastavPozici(kamen2, 100, 1746)
	gke.NastavBlokovani(kamen2, true)

	kamen3 := gke.PridejBlok("./Rock Pile.png")
	gke.NastavZvetseni(kamen3, 0.3)
	gke.NastavPozici(kamen3, 250, 1650)
	gke.NastavBlokovani(kamen3, true)

	kamen4 := gke.PridejBlok("./Rock Pile.png")
	gke.NastavZvetseni(kamen4, 0.3)
	gke.NastavPozici(kamen4, 550, 1450)
	gke.NastavBlokovani(kamen4, true)

	kamen5 := gke.PridejBlokSVyrezem("./platform2.png", gke.Vyrez{X1: 232, Y1: 78, X2: 661, Y2: 332})
	gke.NastavZvetseni(kamen5, 0.3)
	gke.NastavPozici(kamen5, 360, 1550)
	gke.NastavBlokovani(kamen5, true)

	//blok2 := gke.PridejBlokSVyrezem("./Pozadi.png", gke.Vyrez{X1: 0, Y1: 223, X2: 478, Y2: 221})
	//gke.NastavPozici(blok2, 0, 223)
	//gke.NastavBlokovani(blok2, true)
	//gke.NastavZvetseni(blok2, 0.9)

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
	gke.NastavRychlostPohybu(hratelna_postava, 1.25)
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
			if gke.ZjistiKontaktSHratelnouPostavou(enemy) {
				x := gke.ZjistitPoziciX(&hratelna_postava.Blok)
				y := gke.ZjistitPoziciY(&hratelna_postava.Blok)
				gke.NastavPozici(&hratelna_postava.Blok, x-100, y)
			}
			return []gke.Akce{}
		}
	}(),
	)

	nepritel := gke.PridejNepritele(
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
			gke.NastavPozici(&enemy.Blok, x_enemy, y_enemy)
			if gke.ZjistiKontaktSHratelnouPostavou(enemy) {
				gke.NastavPozici(&hratelna_postava.Blok, x-100, y)
			}
			return []gke.Akce{akce}
		},
	)

	gke.NastavPozici(&nepritel.Blok, 300.0, 1826.0)
	gke.NastavZvetseni(&nepritel.Blok, 0.5)
	gke.NastavBlokovani(&nepritel.Blok, false)
	gke.NastavAnimaci(nepritel, gke.AkceJdeVLevo, true,
		[]gke.Vyrez{
			{X1: 19, Y1: 19, X2: 194, Y2: 202},
		})
	gke.NastavAnimaci(nepritel, gke.AkceJdeVPravo, false,
		[]gke.Vyrez{
			{X1: 19, Y1: 19, X2: 194, Y2: 202},
		})

	gke.NastavPozici(&spike.Blok, 410, 1510)
	gke.NastavZvetseni(&spike.Blok, 0.4)

	gke.NastavBlokovani(&spike.Blok, false)
	gke.NastavAnimaci(spike, gke.AkceJdeVPravo, false,
		[]gke.Vyrez{{X1: 39, X2: 128, Y1: 67, Y2: 160}})

	gke.SpustHru()
}

//
