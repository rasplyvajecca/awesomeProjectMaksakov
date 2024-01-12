package relatives

type Relatives struct {
	Member string
	Name   string
}

func RelativesInfo() []Relatives {
	return []Relatives{
		{Member: "Grandmother", Name: "Ira"},
		{Member: "Grandfather", Name: "Oleg"},
		{Member: "Mother's sister", Name: "Liza"},
		{Member: "Father's brother", Name: "Dima"},
		{Member: "Mother's cousin", Name: "Dima"},
	}
}
