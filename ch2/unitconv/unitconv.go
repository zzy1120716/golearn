package unitconv

import "fmt"

type Meter float64
type Foot float64
type Kilogram float64
type Pound float64

func (m Meter) String() string     { return fmt.Sprintf("%gm", m) }
func (ft Foot) String() string     { return fmt.Sprintf("%gft", ft) }
func (lb Pound) String() string    { return fmt.Sprintf("%glb", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
