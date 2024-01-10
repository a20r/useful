package units

import (
	"fmt"
	"time"

	"github.com/a20r/falta"
	"github.com/a20r/useful/generics"
)

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

var distanceUnitNames = map[string]Distance{
	"nm":         Nanometer,
	"nanometer":  Nanometer,
	"nanometers": Nanometer,
	"m":          Meter,
	"meter":      Meter,
	"meters":     Meter,
	"km":         Kilometer,
	"kilometer":  Kilometer,
	"kilometers": Kilometer,
	"ft":         Foot,
	"foot":       Foot,
	"feet":       Foot,
	"yd":         Yard,
	"yard":       Yard,
	"yards":      Yard,
	"mi":         Mile,
	"mile":       Mile,
	"miles":      Mile,
}

// ErrUnsupportedUnit is returned when the unit is not supported
var ErrUnsupportedUnit = falta.Newf(`units: unit "%s" is not supported`)

// MultDistance is a generic function that let's you multiply a distance by a number
func MultDistance[T generics.Number](d Distance, coeff T) Distance {
	return Distance(float64(d) * float64(coeff))
}

// Nanometers converts a number of nanometers into a Distance
func Nanometers[T generics.Number](d T) Distance {
	return MultDistance(Nanometer, d)
}

// Meters converts a number of meters into a Distance
func Meters[T generics.Number](d T) Distance {
	return MultDistance(Meter, d)
}

// Kilometers converts a number of kilometers into a Distance
func Kilometers[T generics.Number](d T) Distance {
	return MultDistance(Kilometer, d)
}

// Feet converts a number of feet into a Distance
func Feet[T generics.Number](d T) Distance {
	return MultDistance(Foot, d)
}

// Yards converts a number of yards into a Distance
func Yards[T generics.Number](d T) Distance {
	return MultDistance(Yard, d)
}

// Miles converts a number of miles into a Distance
func Miles[T generics.Number](d T) Distance {
	return MultDistance(Mile, d)
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

func (d Distance) Per(dur Duration) Speed {
	hours := float64(dur.Nanoseconds()) / float64(time.Hour)
	return Kph(d.Kilometers() / hours)
}

// String returns a nicely formatted distance string w/ standard units
func (d Distance) String() string {
	switch {
	case d.Kilometers() >= 1:
		return fmt.Sprintf("%fkm", d.Kilometers())
	default:
		return fmt.Sprintf("%fm", d.Meters())
	}
}

// MarshalJSON marshals the distance into a JSON string
func (d Distance) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalJSON unmarshals the distance from a JSON string
func (d *Distance) UnmarshalJSON(data []byte) error {
	var dist float64
	var unit string

	if _, err := fmt.Sscanf(string(data), "%f%s", &dist, &unit); err != nil {
		return err
	}

	coeff, ok := distanceUnitNames[unit]

	if !ok {
		return ErrUnsupportedUnit.New(unit)
	}

	*d = MultDistance(coeff, dist)
	return nil
}
