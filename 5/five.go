package main

import (
  "fmt"
  "encoding/hex"
)

func main() {
  key := "ICE"
  expected_output := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
  string_to_encrypt := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

  // create empty byte array for us to fill with our xor'ed string
  encrypted_string := make([]byte, len(string_to_encrypt))
 
  // loop through our string we're going to encrypt
  for s := 0; s < len(string_to_encrypt); s++ {
    x := 0

    // for every character in the string, increment the loop once and reset the key index once it's past 3
    for i := 0; i < len(encrypted_string); i++ {
      if x == 3 {
        x = 0
      }

      // where we actually xor the string character vs the key
      encrypted_string[i] = string_to_encrypt[i] ^ key[x]
      x++
    }
  }

  // encode our finished string to hex
  finished_string := hex.EncodeToString(encrypted_string)

  // compare our output to what we expect
  if expected_output == finished_string {
    fmt.Println(finished_string)
    fmt.Println("Success!")
  }
}
