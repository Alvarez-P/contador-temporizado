package main 

/* --------------------------------------------------------------------------------------------
					Contador de 40 digitos con o sin limite de tiempo   
---------------------------------------------------------------------------------------------*/

import (
    "fmt" 
	"os" 
	"bufio"
	"strings"
	"./clear"
	"time" 
	"strconv"
) 

/*
	withLimit realiza el ciclo mostrando un contador por un tiempo limite dado por el usuario
*/
func withLimit() {

	cont := 0
	menu_limit := "Ingrese el limite de tiempo: "
	fmt.Print(menu_limit)
	reader_limit := bufio.NewReader(os.Stdin)
	input_limit, _ := reader_limit.ReadString('\n') 			// Leer hasta el separador de salto de línea
	selection_limit := strings.TrimRight(input_limit, "\r\n") 	// Remover el salto de línea de la entrada del usuario
	
	limit, _ := strconv.ParseFloat(selection_limit, 64) 
	t0 := time.Now()

	for {
		t1 := time.Now()
		tt := t1.Sub(t0)
		elapsed := tt.Minutes()					// Calculamos minutos transcurridos
		if (limit <= elapsed) {					// Valida fin del ciclo
			break
		}
		clear.CallClear() 							
		fmt.Printf("%040d\n", cont) 
		cont++ 
	}
}

/*
	format da formato a la duracion y al contador.
*/
func format(duration time.Duration, cont int) string {
	duration = duration.Round(time.Millisecond)
	m := duration / time.Minute
	duration -= m * time.Minute
	s := duration / time.Second
	duration -= s * time.Second
	ms := duration / time.Millisecond	
	return fmt.Sprintf("[%02d:%02d:%02d] | %040d\n", m, s, ms, cont)
}

/*
	withoutLimit realiza un ciclo infinito imprimiendo un cronometro y un contador. 
*/
func withoutLimit() {

	cont := 0
	t0 := time.Now()
	for {
		t1 := time.Now()
		tt := t1.Sub(t0)
		clear.CallClear() 						// Limpia pantalla
		durStr := format(tt, cont)
		fmt.Println(durStr)
		cont++ 
	}
}

func main() { 

	menu := `
Ingrese una opción: 
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
		clear.CallClear()

		switch selection {
		case "1":
			withLimit()
		case "2":
			withoutLimit()
		case "3":
			fmt.Println("Saliendo ...")
			os.Exit(3)
		default:
			clear.CallClear() 
		}
	} 
} 