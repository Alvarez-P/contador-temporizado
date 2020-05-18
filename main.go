package main 

/* --------------------------------------------------------------------------------------------
					Contador de 40 digitos con o sin limite de tiempo   
---------------------------------------------------------------------------------------------*/

import (
    "fmt" 
	"os" 
	"bufio"
	"strings"
	"./cicle"
	"github.com/inancgumus/screen"
) 

func main() { 
	menu := `Ingrese una opción: 
[1] Con limite de tiempo
[2] Ciclo infinito
[3] Salir
`
	selection := ""
	for ok := true; ok; ok = (selection != "3") {
    	fmt.Print(menu)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n') 			// Leer hasta el separador de salto de línea
		selection = strings.TrimRight(input, "\r\n")  	// Remover el salto de línea de la entrada del usuario
		screen.Clear()

		switch selection {
		case "1":
			cicle.WithLimit()
		case "2":
			cicle.WithoutLimit()
		case "3":
			fmt.Println("Saliendo ...")
			os.Exit(3)
		default:
			screen.Clear() 
		}
	} 
} 