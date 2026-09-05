package main

func main() {

}

type shape interface {
	area() float64
}
type circle struct {
	radius    float64
	hasRadius bool
}

// c, ok :=s.(circle)
