package house

import (
	"awesomeProjectMaksakov/projectMaksakov/appliance"
	"awesomeProjectMaksakov/projectMaksakov/family"
	"awesomeProjectMaksakov/projectMaksakov/furniture"
	"awesomeProjectMaksakov/projectMaksakov/relatives"
	"awesomeProjectMaksakov/projectMaksakov/rooms"
	"fmt"
)

type House struct {
	HouseWidth    float64
	HouseLength   float64
	HouseHeight   float64
	FamilyInfo    []family.Family
	RelativesInfo []relatives.Relatives
	FurnitureInfo []furniture.Furniture
	ApplianceInfo []appliance.Appliance
	RoomsInfo     []rooms.Rooms
}

func createMaksakovHouse() House {
	return House{
		HouseWidth:    4.5,
		HouseLength:   3.55,
		HouseHeight:   3,
		FamilyInfo:    family.FamilyInfo(),
		RelativesInfo: relatives.RelativesInfo(),
		FurnitureInfo: furniture.FurnitureInfo(),
		RoomsInfo:     rooms.RoomsInfo(),
		ApplianceInfo: appliance.ApplianceInfo(),
	}
}

func MaksakovHouse() {
	house := createMaksakovHouse()
	fmt.Printf("House size: %.3f x %.3f x %.3f \n", house.HouseWidth, house.HouseLength, house.HouseHeight)
	showFamilyInfo(house.FamilyInfo)
	showRelativesInfo(house.RelativesInfo)
	showFurnitureInfo(house.FurnitureInfo)
	showRoomsInfo(house.RoomsInfo)
	showApplianceInfo(house.ApplianceInfo)
}

func showFamilyInfo(objects []family.Family) {
	for _, object := range objects {
		fmt.Printf("Family:\n")
		fmt.Printf("- %s %s\n", object.Member, object.Name)
	}
}

func showRelativesInfo(objects []relatives.Relatives) {
	for _, object := range objects {
		fmt.Printf("Relatives:\n")
		fmt.Printf("- %s %s\n", object.Member, object.Name)
	}
}

func showFurnitureInfo(objects []furniture.Furniture) {
	for _, object := range objects {
		fmt.Printf("Furniture:\n")
		fmt.Printf("- %s %.2f m\n", object.Name, object.Size)
	}
}

func showRoomsInfo(objects []rooms.Rooms) {
	for _, object := range objects {
		fmt.Printf("Rooms:\n")
		fmt.Printf("- %s %.2f m\n", object.Name, object.Size)
	}
}

func showApplianceInfo(objects []appliance.Appliance) {
	for _, object := range objects {
		fmt.Printf("Appliance:\n")
		fmt.Printf("- %s %.3f $\n", object.Name, object.Price)
	}
}
