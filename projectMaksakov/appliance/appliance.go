package appliance

type Appliance struct {
	Name  string
	Price float64
}

func ApplianceInfo() []Appliance {
	return []Appliance{
		{Name: "Telephone", Price: 3},
		{Name: "Iphone 12", Price: 1},
		{Name: "Iphone 15 pro", Price: 10},
		{Name: "AppleWatch", Price: 1},
		{Name: "Mac", Price: 5},
	}
}
