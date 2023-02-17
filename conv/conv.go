package conv

import "math"
/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(farhenheit float64) float64 {
	celsiusF := (farhenheit - 32)*(5.0/9.0)
  return (math.Round(celsiusF*100)/100)
}

func FarhenheitToKelvin(farhenheit float64) float64 {
  kelvinF := (farhenheit - 32) * (5.0/9.0) + 273.15
  return (math.Round(kelvinF*100)/100)
}

func CelsiusToFarhenheit(celsius float64) float64 {
  farhenheitC := (celsius*(9.0/5.0)) + 32
  return (math.Round(farhenheitC*100)/100)
}

func CelsiusToKelvin(celsius float64) float64 {
  kelvinC := (celsius + 273.15)
  return (math.Round(kelvinC*100)/100)
}

func KelvinToCelsius(kelvin float64) float64 {
  celsiusK := (kelvin - 273.15)
  return (math.Round(celsiusK*100)/100)
}

func KelvinToFarhenheit(kelvin float64) float64 {
  farhenheitK := (kelvin - 273.15)*(9.0/5.0) + 32
  return (math.Round(farhenheitK*100)/100)
}