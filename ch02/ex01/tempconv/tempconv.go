// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius 摂氏
type Celsius float64

// Fahrenheit 華氏
type Fahrenheit float64

// Kelvin ケルビン
type Kelvin float64

const (
	// AbsoluteZeroC 最低零度(摂氏)
	AbsoluteZeroC Celsius = -273.15

	// FreezingC 凍る温度(摂氏)
	FreezingC Celsius = 0

	// BoilingC 沸点(摂氏)
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Kelvin) String() string     { return fmt.Sprintf("%g°K", f) }

//!-
