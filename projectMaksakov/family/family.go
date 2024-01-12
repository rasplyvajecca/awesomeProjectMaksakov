package family

type Family struct {
	Member string
	Name   string
}

func FamilyInfo() []Family {
	return []Family{
		{Member: "Mom", Name: "Oksana"},
		{Member: "Dad", Name: "Ivan"},
		{Member: "Mom's son", Name: "Vlad"},
		{Member: "Dad's daughter", Name: "Masha"},
		{Member: "Cat", Name: "Simus"},
	}
}
