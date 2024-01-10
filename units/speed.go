package units

import (
	"fmt"
	"time"

	"github.com/a20r/useful/generics"
)

type Speed int64

const (
	NanometersPerSecond Speed = 1
	MetersPerSecond     Speed = 1e+9
	FeetPerSecond       Speed = 3.281e+9
	KilometersPerHour   Speed = 3.6e+9
	MilesPerHour        Speed = 2.237e+9
)

type Duration interface {
	Nanoseconds() int64
}

var speedUnitNames = map[string]Speed{
	"nm/s":         NanometersPerSecond,
	"nanometer/s":  NanometersPerSecond,
	"nanometers/s": NanometersPerSecond,
	"m/s":          MetersPerSecond,
	"meter/s":      MetersPerSecond,
	"meters/s":     MetersPerSecond,
	"ft/s":         FeetPerSecond,
	"foot/s":       FeetPerSecond,
	"feet/s":       FeetPerSecond,
	"kph":          KilometersPerHour,
	"km/h":         KilometersPerHour,
	"kilometer/h":  KilometersPerHour,
	"kilometers/h": KilometersPerHour,
	"mph":          MilesPerHour,
	"mi/h":         MilesPerHour,
	"mile/h":       MilesPerHour,
	"miles/h":      MilesPerHour,
}

func NmPerSec[T generics.Number](s T) Speed {
	return Speed(s) * NanometersPerSecond
}

func MsPerSec[T generics.Number](s T) Speed {
	return Speed(s) * MetersPerSecond
}

func FtPerSec[T generics.Number](s T) Speed {
	return Speed(s) * FeetPerSecond
}

func Kph[T generics.Number](s T) Speed {
	return Speed(s) * KilometersPerHour
}

func Mph[T generics.Number](s T) Speed {
	return Speed(s) * MilesPerHour
}

func (s Speed) Distance(dur Duration) Distance {
	secs := float64(dur.Nanoseconds()) / float64(time.Second)
	return Distance(s.NmPerSec() * secs)
}

func (s Speed) Duration(dist Distance) time.Duration {
	return time.Duration(float64(time.Hour) * dist.Kilometers() / s.Kph())
}

func (s Speed) NmPerSec() float64 {
	return float64(s) / float64(NanometersPerSecond)
}

func (s Speed) MsPerSec() float64 {
	return float64(s) / float64(MetersPerSecond)
}

func (s Speed) FtPerSec() float64 {
	return float64(s) / float64(FeetPerSecond)
}

func (s Speed) Kph() float64 {
	return float64(s) / float64(KilometersPerHour)
}

func (s Speed) Mph() float64 {
	return float64(s) / float64(MilesPerHour)
}

func (s Speed) String() string {
	return fmt.Sprintf("%fkph", s.Kph())
}

func (s Speed) MarshalJSON() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *Speed) UnmarshalJSON(data []byte) error {
	var unit string
	var val float64

	if _, err := fmt.Sscanf(string(data), "%f%s", &val, &unit); err != nil {
		return err
	}

	coeff, ok := speedUnitNames[unit]

	if !ok {
		return ErrUnsupportedUnit.New(unit)
	}

	*s = Speed(val * float64(coeff))
	return nil
}
