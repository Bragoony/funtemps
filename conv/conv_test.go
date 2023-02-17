package conv

import (
	"reflect"
	"testing"
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 0},
	}

	for _, tfc := range tests {
		got := FarhenheitToCelsius(tfc.input)
		if !reflect.DeepEqual(tfc.want, got) {
			t.Errorf("expected: %v, got: %v", tfc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	
	tests := []test{
		{input: 134.01, want: 329.82},
	}

	for _, tfk := range tests {
		got := FarhenheitToKelvin(tfk.input)
		if !reflect.DeepEqual(tfk.want, got) {
			t.Errorf("expected: %v, got: %v", tfk.want, got)
		}	
	}
}

func TestCelsiusToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	
	tests := []test{
		{input: -89.4, want: -128.92},
	}

	for _, tcf := range tests {
		got := CelsiusToFarhenheit(tcf.input)
		if !reflect.DeepEqual(tcf.want, got) {
			t.Errorf("expected: %v, got: %v", tcf.want, got)
		}	
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	
	tests := []test{
		{input: -89.4, want: 183.75}, //skal egentlig være + på want ikke - som Janis sier
	}

	for _, tck := range tests {
		got := CelsiusToKelvin(tck.input)
		if !reflect.DeepEqual(tck.want, got) {
			t.Errorf("expected: %v, got: %v", tck.want, got)
		}	
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	
	tests := []test{
		{input: 273.15, want: 0},
	}

	for _, tkc := range tests {
		got := KelvinToCelsius(tkc.input)
		if !reflect.DeepEqual(tkc.want, got) {
			t.Errorf("expected: %v, got: %v", tkc.want, got)
		}	
	}
}

func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	
	tests := []test{
		{input: 329.82, want: 134.01},
	}

	for _, tkf := range tests {
		got := KelvinToFarhenheit(tkf.input)
		if !reflect.DeepEqual(tkf.want, got) {
			t.Errorf("expected: %v, got: %v", tkf.want, got)
		}	
	}
}