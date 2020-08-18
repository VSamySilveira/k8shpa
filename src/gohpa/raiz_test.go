package main

import "testing"

func TestRaiz(t *testing.T) {
	resultado := raiz()
	if resultado != "<b>Code.education Rocks!</b>" {
	  t.Errorf("Greeting esperada: <b>Code.education Rocks!</b>, obtida: %s", resultado)
	}
}