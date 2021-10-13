package main

import "fmt"

//func main() {
//	c := cli.NewCli(os.Args)
//	c.Run(app.Init())
//
//}

func main() {
    cetakHasil,angka := cetakAngka(343984983489348)
    fmt.Println(cetakHasil, angka)

}

func cetakAngka(input int) (string,int) {
    if input % 2 == 0 {
        return "Ini bilangan Genap",input
    } else {
        return "Bilangan Ganjl", input
    }
}
