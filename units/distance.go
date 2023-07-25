package units

import "golang.org/x/exp/constraints"

// Distance is a nanometers representation of distance
type Distance int64

const (
	Nanometer Distance = 1
	Meter     Distance = 1e+9
	Kilometer Distance = 1e+12
	Foot      Distance = 3.048e+8
	Yard      Distance = 9.144e+8
	Mile      Distance = 1.609e+12
)

// MultDistance is a generic function that let's you multiply a distance by a number
func MultDistance[T constraints.Float | constraints.Integer](d Distance, coeff T) Distance {
	return Distance(float64(d) * float64(coeff))
}

// Nanometers converts a number of nanometers into a Distance
func Nanometers[T constraints.Float | constraints.Integer](d T) Distance {
	return Distance(d) * Nanometer
}

// Meters converts a number of meters into a Distance
func Meters[T constraints.Float | constraints.Integer](d T) Distance {
	return Distance(d) * Meter
}

// Kilometers converts a number of kilometers into a Distance
func Kilometers[T constraints.Float | constraints.Integer](d T) Distance {
	return Distance(d) * Kilometer
}

// Feet converts a number of feet into a Distance
func Feet[T constraints.Float | constraints.Integer](d T) Distance {
	return Distance(d) * Foot
}

// Yards converts a number of yards into a Distance
func Yards[T constraints.Float | constraints.Integer](d T) Distance {
	return Distance(d) * Yard
}

// Miles converts a number of miles into a Distance
func Miles[T constraints.Float | constraints.Integer](d T) Distance {
	return Distance(d) * Mile
}

// Nanometers converts the distance to nanometers. It just returns the distance
// value as a float because it is already represented as nanometers
func (d Distance) Nanometers() float64 {
	return float64(d) / float64(Nanometer)
}

// Meters converts the distance from nanometers to meters
func (d Distance) Meters() float64 {
	return float64(d) / float64(Meter)
}

// Kilometers converts the distance from nanometers to kilometers
func (d Distance) Kilometers() float64 {
	return float64(d) / float64(Kilometer)
}

// Feet converts the distance from nanometers to feet
func (d Distance) Feet() float64 {
	return float64(d) * float64(Foot)
}

// Yards converts the distance from nanometers to yards
func (d Distance) Yards() float64 {
	return float64(d) / float64(Yard)
}

// Miles converts the distance from nanometers to miles
func (d Distance) Miles() float64 {
	return float64(d) / float64(Mile)
}
