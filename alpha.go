package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

var lines, err = readLines("stdmsg.txt")

func main() {
   stdMsgs := lookUp(os.Args[1])
   for i, stdMsg := range stdMsgs {
     fmt.Println(i,stdMsg)
   } 
}

func lookUp(stdMsgNo string) ([]string){
  var results []string
  found := false
  for _, line := range lines {
    if strings.Contains(line,stdMsgNo){
       results = append(results, line)
       found =true
    }
  }
  if found==false{
       results = append(results, "Not found")
  }
  return results
}


