package rooms

type Rooms struct {
	Name string
	Size float64
}

func RoomsInfo() []Rooms {
	return []Rooms{
		{Name: "Kitchen", Size: 4.5},
		{Name: "Bathroom", Size: 3.3},
		{Name: "Mom's room", Size: 5},
		{Name: "Dad's room", Size: 4.4},
		{Name: "Kids' room", Size: 6},
	}
}
