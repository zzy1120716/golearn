// 进行摄氏度和华氏度的转换
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZero   Celsius = -273.15 // 绝对零度
	FreezingC      Celsius = 0       // 结冰点温度
	BoilingC       Celsius = 100     // 沸水温度
	AbosoluteZeroK Kelvin  = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
