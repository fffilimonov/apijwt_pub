package main

import (
    "io"
    "io/ioutil"
    "os"
)

// Reads files in the settings/keys
// and encodes them as strings literals in textfiles.go
func main() {
    fs, _ := ioutil.ReadDir("./settings/keys/")
    out, _ := os.Create("core/authentication/keys.go")
    out.Write([]byte("package authentication \n\nconst (\n"))
    for _, f := range fs {
	out.Write([]byte(f.Name() + " = `"))
	f, _ := os.Open("./settings/keys/" + f.Name())
	io.Copy(out, f)
	out.Write([]byte("`\n"))
    }
    out.Write([]byte(")\n"))
}

