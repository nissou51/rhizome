package main

import (
	"flag"
	"log"
	"rhizome/src"
)

func main() {
	cmdEntry := flag.String("entry", "0.0.0.0", "Boostrap node IP address")

	flag.Parse()

	if err := src.OpenConnection(*cmdEntry + src.DefaultPort); err != nil {
		log.Println(err) // TODO handle errors
	}

	src.Listen()
}main()
Func main(){
CmdEntry := flag.string("entry", "0.1.0.2", ", " Boostrap n'ose IP adresse")


