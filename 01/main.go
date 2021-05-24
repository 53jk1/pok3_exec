package main

import (
	"fmt"
	"log"

	"github.com/53jk1/pok2/interpolate"
	"github.com/53jk1/pok2/interpolate/linear"
)

// W przypadku interpolacji pojedynczej wartości musimy zdefiniować wycinki dla x i y,
// które zostaną przekonwertowane na pary współrzędnych i posortowane.
// Dopasowujemy plasterki do naszego typu interpolacji i wywołujemy funkcję interpolate,
// która odbiera dwa parametry - pierwszym parametrem jest dowolny typ, który jest zgodny
// z interfejsem validateInterpolator (w ten sposób implementujemy metody Interpolate i Validate),
// a drugim parametrem jest wartość zmiennoprzecinkowa,
// której potrzebujemy interpolować. Zwraca wartość interpolowaną i błąd (jeśli istnieje).

func main() {
	x := []float64{3.0, 750.0}
	y := []float64{4.0, 1000.0}
	valToInterp := 3.2

	li := linear.New()
	li.Fit(x, y)

	estimate, err := interpolate.WithSingle(li, valToInterp)
	fmt.Printf("Oszacowana wartosc w czasie %.1f jest %f\n", valToInterp, estimate)
	fmt.Println(li)

	if err != nil {
		log.Fatal(err)
	}
}
