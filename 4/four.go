package main 

import (
  "fmt"
  "os"
  "bufio"
  "encoding/hex"
)

// character map to score a string
func generateCharacterMap() map[string]float64 {
  charmap := map[string]float64{
    "a": 0.0651738,
    "b": 0.0124248,
    "c": 0.0217339,
    "d": 0.0349835,
    "e": 0.1041442,
    "f": 0.0197881,
    "g": 0.0158610,
    "h": 0.0492888,
    "i": 0.0558094,
    "j": 0.0009033,
    "k": 0.0050529,
    "l": 0.0331490,
    "m": 0.0202124,
    "n": 0.0564513,
    "o": 0.0596302,
    "p": 0.0137645,
    "q": 0.0008606,
    "r": 0.0497563,
    "s": 0.0515760,
    "t": 0.0729357,
    "u": 0.0225134,
    "v": 0.0082903,
    "w": 0.0171272,
    "x": 0.0013692,
    "y": 0.0145984,
    "z": 0.0007836,
    " ": 0.1918182,
  }

  return charmap
}

// function to score incoming string
func score(s string) float64 {
  charmap := generateCharacterMap()
  score := 0.0
  for i := 0; i < len(s); i++ {
    char := s[i]
    if _, ok := charmap[string(char)]; ok {
      score = score + charmap[string(char)]
    } 
  }
  return score
}

func main() {
  file, _ := os.Open("4.txt")
  scanner := bufio.NewScanner(file)

   // loop through entire file
   for scanner.Scan() {
     line := scanner.Text()
     decoded_string, _ := hex.DecodeString(line)
    
     // loop through every known ascii code 
     for s := 0; s < 256; s++ {
       b := make([]byte, len(decoded_string))
       
       // loop through string after we've decoded it and xor it against current iteration of ascii code
       for i := 0; i < len(decoded_string); i++ {
         b[i] = decoded_string[i] ^ string(s)[0]
       }
       printscore := score(string(b))

       // arbitrary score to print, just sifting out bad results
       if printscore > 2.0 {
         // print out the score, the text and the key
         fmt.Println(score(string(b)), string(b), string(s))
       }
     }
   }

}
