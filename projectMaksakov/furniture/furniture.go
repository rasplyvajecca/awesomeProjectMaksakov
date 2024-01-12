package furniture

type Furniture struct {
	Name string
	Size float64
}

func FurnitureInfo() []Furniture {
	return []Furniture{
		{Name: "Mom's bed", Size: 3},
		{Name: "Dad's bed", Size: 2.5},
		{Name: "Grandma's bed", Size: 7},
		{Name: "Grandpa's bed", Size: 7},
		{Name: "Kids' bed", Size: 5},
	}
}
