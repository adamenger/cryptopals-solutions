package main

import (
  "encoding/hex"
  "fmt"
)

func main() {
  original_string := "1c0111001f010100061a024b53535009181c"
  decoded_string, _ := hex.DecodeString(original_string)
  xor_string,_ := hex.DecodeString("686974207468652062756c6c277320657965")

  n := len(decoded_string)
  b := make([]byte, n)

  if len(xor_string) > n {
     n = len(b)
  }

  for i := 0; i < n; i++ {
    b[i] = decoded_string[i] ^ xor_string[i]
  }

  fmt.Println(string(xor_string))
  fmt.Println(string(decoded_string))
  fmt.Printf("%x\n", string(b))
}
