// declarar pacote principal
package main

//importar a função fmt
import "fmt"

//declaração da variavel CONST do ponto de ebuluçao da agua em F
const ebulicaoK = 373.0

//declarar função principal
func main() {

	tempk := ebulicaoK     //variavel do valor da temperatura em graus F
	tempC := tempk - 273.0 //conversão de K para C
	//apareceça na tela o resultado
	fmt.Printf("a temperatura de ebulição da água em K = %g , temperatura de ebulução da água em C = %g.", tempk, tempC)

}
