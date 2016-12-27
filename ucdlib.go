package runefinder

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// AnalisarLinha devolve a runa e o nome de uma linha do UnicodeData.txt
func AnalisarLinha(linha string) (rune, string) {
	campos := strings.Split(linha, ";")
	código, _ := strconv.ParseInt(campos[0], 16, 32)
	return rune(código), campos[1]
}

// Listar exibe na saída padrão o código, a runa e o nome dos caracteres Unicode
// cujo nome contem o texto da consulta // <1>
func Listar(texto io.Reader, consulta string) {
	varredor := bufio.NewScanner(texto) // <2>
	for varredor.Scan() {               // <3>
		linha := varredor.Text()            // <4>
		if strings.TrimSpace(linha) == "" { // <5>
			continue
		}
		runa, nome := AnalisarLinha(linha)    // <6>
		if strings.Contains(nome, consulta) { // <7>
			fmt.Printf("U+%04X\t%[1]c\t%s\n", runa, nome) // <8>
		}
	}
}
