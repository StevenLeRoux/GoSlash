package main

import (
    "io"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    fs, _ := ioutil.ReadDir("../ui/dist")
    out, _ := os.Create("ui.go")
    out.Write([]byte("package main \n\nconst (\n"))
    for _, f := range fs {
        if strings.HasSuffix(f.Name(), ".html") {
            out.Write([]byte(strings.TrimSuffix(f.Name(), ".html") + " = `"))
            f, _ := os.Open("../ui/dist/" + f.Name())
            io.Copy(out, f)
            out.Write([]byte("`\n"))
        }
    }
    out.Write([]byte(")\n"))
}
