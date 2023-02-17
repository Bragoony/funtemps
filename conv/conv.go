package conv

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
	return ((farhenheit - 32)*(5/9))
}

func FarhenheitToKelvin(farhenheit float64) float64 {
  return ((farhenheit - 32) * (5/9) + 273.15)
}

func CelsiusToFahrenheit(celsius float64) float64 {
  return (celsius*(9/5) + 32)
}

func CelsiusToKelvin(celsius float64) float64 {
  return (celsius + 273.15)
}

func KelvinToCelsius(kelvin float64) float64 {
  return (kelvin - 273.15)
}

func KelvinToFarhenheit(kelvin float64) float64 {
  return ((kelvin - 273.15)*(9/5) + 32)
}