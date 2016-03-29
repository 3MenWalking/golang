package main

import (
  "bufio"
  "fmt"
  //"log"
  "os"
  "strings"
  "net/http"
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

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}

var lines, err = readLines("stdmsg.txt")

func main() {
 http.HandleFunc("/query", FindMsg)
 http.ListenAndServe(":8080",nil)
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


func FindMsg(w http.ResponseWriter, r *http.Request){
   stdMsgs := lookUp(r.URL.Query().Get("msgNo"))
   for _, stdMsg := range stdMsgs {
     fmt.Fprintln(w,stdMsg)
   }
}
