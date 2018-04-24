package main

import "testing"

func TestIsMutant(t *testing.T) {
  adn := [][]string{{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},{
            "ABCDABCDA",
            "EAGHEFGHB",
            "IJUMIJKMC",
            "LNIALNYOD",
            "ABCDXBCDE",
            "EFGHEAGHF",
            "IJKMIJAMG",
            "LNSOLNTAH",
            "LNUOLNXOA"},
            {"ABCD","AAGH","AJAM","ANNA"}
          }
  for i := 0; i < leng(adn); i++ {
    isMutant := isMutant(adn[i])
    if !isMutant {
      t.Errorf("isMutant was incorrect, got: %d, want: %d.", isMutant, true)
    }
  }
}
