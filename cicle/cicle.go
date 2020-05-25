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
	format da formato a la duracion y al contador.
*/
func format(duration time.Duration, cont int) string {
	h := duration / time.Hour
	duration -= h * time.Hour
	m := duration / time.Minute
	return fmt.Sprintf("\r[%02d:%02d] | %040d", h, m, cont)
} 

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
	limit = limit * 60000
	selection_limit = strconv.Itoa(limit) + "ms"
	limitms, _ := time.ParseDuration(selection_limit)
	screen.Clear() 
	screen.MoveTopLeft()
	t0 := time.Now()
	for {
		t1 := time.Now()
		tt := t1.Sub(t0)
		durStr := format(tt, cont)
		fmt.Print(durStr) 
		if (limitms < tt) {				// Valida fin del ciclo
			break
		}							
		cont++ 
	}
}

/*
	WithoutLimit realiza un ciclo infinito imprimiendo un cronometro y un contador. 
*/

func WithoutLimit () {
	cont := 0
	screen.MoveTopLeft()
	for {
		fmt.Printf("\r%040d", cont)
		cont++ 
	}
}