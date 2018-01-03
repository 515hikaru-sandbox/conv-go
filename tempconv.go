package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahreheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string   { return fmt.Sprintf("%g℃", c) }
func (f Fahreheit) String() string { return fmt.Sprintf("%g^oF", f) }
