package main

import "github.com/theartefak/artefak/bootstrap"

func main() {
	artefak := bootstrap.Run()
    artefak.Listen("127.0.1.1:8000")
}
