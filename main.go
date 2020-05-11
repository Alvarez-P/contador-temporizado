package main 

import (
    "fmt" 
	"os" 
	"bufio"
	"strings"
	"./clear"
	"time" 
	"strconv"
) 

func main() { 

	/*
		Contador de 40 digitos con limite de ejecucion dado por el usuario.
	*/

	menu := `Ingresa la cantidad en minutos a ejecutar el programa: `
	cont := 0

	fmt.Print(menu)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') 			// Leer hasta el separador de salto de línea
	selection := strings.TrimRight(input, "\r\n") 	// Remover el salto de línea de la entrada del usuario
	
	limit, _ := strconv.Atoi(selection) 
	t0 := time.Now()

	for {
		t1 := time.Now()
		tt := t1.Sub(t0)
		elapsed := tt.Minutes()						// Calculamos minutos transcurridos
		if (float64(limit) <= elapsed) {
			os.Exit(3)
		}
		clear.CallClear() 							// Limpia pantalla
		fmt.Printf("%040d\n", cont) 
		cont++ 
	}
    
} 