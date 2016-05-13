package main

import (
  "io/ioutil"
  "encoding/hex"
  "fmt"
  "regexp"
  "strings"
  "sort"
)

type PairList []Pair

type Pair struct {
  Key string
  Value int
}

func rankByLetterCount(letterFrequencies map[string]int) PairList{
  pl := make(PairList, len(letterFrequencies))
  i := 0
  for k, v := range letterFrequencies {
    pl[i] = Pair{k,v}
    i++
  }
  sort.Sort(sort.Reverse(pl))
  return pl
}

func (p PairList) Len() int { return len(p)}
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func main() {
  alpha := regexp.MustCompile(`[a-zA-Z]`)
  char_map := make(map[string]int)
  file, _ := (ioutil.ReadFile("text.txt"))

  for i := 0; i < len(file); i++ {
    char := string(file[i])
    if alpha.MatchString(char) {
      upper_char := strings.ToUpper(char)
      if _, ok := char_map[upper_char]; ok {
        char_map[upper_char] = char_map[upper_char] + 1
      } else {
        char_map[upper_char] = 1
      }
    }
  }

  hexstr, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

  for _, v := range rankByLetterCount(char_map) {
    b   := make([]byte, len(hexstr))
    for i := 0; i < len(hexstr); i++ {
      b[i] = hexstr[i] ^ v.Key[0]
    }

    fmt.Println(v.Key, string(b))
  }
}
