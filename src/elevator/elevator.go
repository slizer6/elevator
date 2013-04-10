
package elevator

import "elevdriver"

type Elevatorinfo struct {
	state State
	event Event
	internal_orders [][]int
	external_orders [][]int
	last_floor int
	last_direction elevdriver.Direction
}

var N_FLOORS, N_BUTTONS int = 4, 3
	
func (elevinf *Elevatorinfo) HandleElevator() {
	
	elevinf.state = IDLE
	elevinf.event = NO_EVENT
	
	// Order Arrays
	elevinf.internal_orders = make([][]int, N_FLOORS)
	for i := range(elevinf.internal_orders){
		elevinf.internal_orders[i] = make([]int, N_BUTTONS)
	}
	elevinf.external_orders = make([][]int, N_FLOORS)
	for i := range(elevinf.internal_orders){
		elevinf.external_orders[i] = make([]int, N_BUTTONS-1)
	}
	
	elevinf.BootStatemachine()
	
	for {
		elevinf.UpdateLastDirection()
		elevinf.RunStatemachine()	
	}
	
}
