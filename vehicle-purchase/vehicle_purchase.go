package purchase

import s "strings"
import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var needLisence bool
	if kind == "car" {
		needLisence = true
	} else if kind == "truck" {
		needLisence = true
	}
	return needLisence
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var car string
	comparedInteger := s.Compare(option1, option2)
	if comparedInteger == 1 {
		car = option2
	} else if comparedInteger == -1 {
		car = option1
	} else if comparedInteger == 0 {
		car = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", car)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellPrice float64
	if age < 3 {
		resellPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		resellPrice = originalPrice * 0.7

	} else if age >= 10 {
		resellPrice = originalPrice * 0.5
	}
	return resellPrice
}
