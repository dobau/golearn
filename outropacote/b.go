package outropacote

import "fmt"

func funcaoPrivada2() {
	fmt.Println("Chamando funcaoPrivada2() (inicio)")

	str := fmt.Sprintf("a(%v) b(%v) d(%v) e(%v) f(%v) g(%v) h(%v) j(%v)",
		a, b, d, e, f, g, h, j)
	fmt.Println(str)

	fmt.Println("Chamando funcaoPrivada2() (fim)")
}
