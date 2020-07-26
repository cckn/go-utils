package vehicle

import "log"

type Car string

func (c Car) Accelerate() {
	log.Println("Speeding up")
}
func (c Car) Break() {
	log.Println("Stopping")
}
func (c Car) Steer(diriection string) {
	log.Println("Turning", diriection)
}

type Truck string

func (t Truck) Accelerate() {
	log.Println("Speeding up")
}
func (t Truck) Break() {
	log.Println("Stopping")
}
func (t Truck) Steer(diriection string) {
	log.Println("Turning", diriection)
}
