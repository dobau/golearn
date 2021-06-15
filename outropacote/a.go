package outropacote

import "fmt"

var a int = 1
var b = 1

// c := 1
var (
	d    int
	e, f = 1, 2
)

const (
	g, h, i, j = "E", 'f', 10000, 3.41
)

// FuncaoPublica os comentários de uma função começam pelo nome na entidade (pacote, funcao, struct, ...)
func FuncaoPublica() {
	fmt.Println("Chamando FuncaoPublica()")

	funcaoPrivada1()
	funcaoPrivada2()
}

func funcaoPrivada1() {
	fmt.Println("Chamando funcaoPrivada1()")
}
