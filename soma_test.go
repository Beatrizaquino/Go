package main

import "testing"

func TestSoma(t *testing.T) {
  
  numero := [5]int{1, 2, 3, 4, 5}
  
  resultado := Soma(numero)

  esperado := 15
  
  if resultado != esperado {
    t.Errof("Resultado %d, esperado %d, dado %v", resultado, esperado, numero)
  }
}