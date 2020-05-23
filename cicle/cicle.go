package cicle

import (
	"fmt"
	"time" 
	"bufio"
	"os"
	"strconv"
	"github.com/inancgumus/screen"
	"strings"
)

/*
	WithLimit realiza el ciclo mostrando un contador por un tiempo limite dado por el usuario
*/
func WithLimit () {
	cont := 0
	menu_limit := "\rIngrese el limite de tiempo(minutos) con un máximo de 5: "
	fmt.Print(menu_limit)
	reader_limit := bufio.NewReader(os.Stdin)
	input_limit, _ := reader_limit.ReadString('\n') 			// Leer hasta el separador de salto de línea
	selection_limit := strings.TrimRight(input_limit, "\r\n") 	// Remover el salto de línea de la entrada del usuario
	limit, _ := strconv.Atoi(selection_limit)
	
	if limit > 5 {
		fmt.Println("\rDebe ser un valor igual o menor a 5")
		return 
	} 
	selection_limit = strconv.Itoa(limit) + "m"
	limitms, _ := time.ParseDuration(selection_limit)
	screen.Clear() 
	screen.MoveTopLeft()
	tf := time.Now().Add(limitms)

	for {
		tn := time.Now()
		if (tn.After(tf)) {					// Valida fin del ciclo
			break
		}							
		fmt.Printf("\r%040d", cont) 
		cont++ 
	}
}

/*
	format da formato a la duracion y al contador.
*/
func format(duration time.Duration, cont int) string {
	duration = duration.Round(time.Minute)
	h := duration / time.Hour
	duration -= h * time.Hour
	m := duration / time.Minute
	return fmt.Sprintf("\r[%02d:%02d] | %040d", h, m, cont)
} 

/*
	WithoutLimit realiza un ciclo infinito imprimiendo un cronometro y un contador. 
*/

func WithoutLimit () {
	cont := 0
	screen.MoveTopLeft()
	t0 := time.Now()
	for {
		t1 := time.Now()
		tt := t1.Sub(t0)
		durStr := format(tt, cont)
		fmt.Print(durStr)
		cont++ 
	}
}
