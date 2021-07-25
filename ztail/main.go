package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		if os.Args[1] == "-c" {
			ls := "0123456798"
			summa := 0
			for _, val := range os.Args[2] {
				sov := 1
				idx1 := 0
				for idx, l := range ls {
					if l == val {
						sov = 0
						idx1 = idx
					}
				}
				summa = (summa + idx1) * 10
				if sov == 1 {
					os.Exit(1)
				}
			}
			summa = summa / 10
			if summa > 0 {
				error1 := 0
				for i := 3; i < len(os.Args); i++ {
					content, err := os.Open(os.Args[i])
					if err != nil {
						fmt.Printf("%s\n", err)
						error1 = 1
						content.Close()
					} else {
						stat, _ := content.Stat()
						b1 := make([]byte, stat.Size())
						n1, erro1 := content.Read(b1)
						if erro1 != nil {
							os.Exit(1)
						}
						content.Close()
						content1 := (string(b1[0:n1]))
						if i != len(os.Args) && i != 3 {
							fmt.Printf("%s", "\n")
						}
						if summa > len(string(content1))+1 {
							fmt.Printf("==> %s <==\n", os.Args[i])
							fmt.Printf("%s", content1)
						} else {
							fmt.Printf("==> %s <==\n", os.Args[i])
							fmt.Printf("%s", content1[len(content1)-summa:])
						}
					}
				}
				if error1 == 1 {
					os.Exit(1)
				}
			} else {
				os.Exit(1)
			}
		} else {
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}
}
