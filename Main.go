package main

import (

	"fmt"
	"os"
	"github.com/meritos/go-study/strings"
)


func main () {

	strings.CountWord("Hello World")

	if len(os.Args) < 2 {

		fmt.Println("\n For Usage: ") 
		fmt.Println("           --help -h\n") 
	} else {

		if os.Args[1] == "--help" || os.Args[1] == "-h" {

			printHelp()
		}

		if os.Args[1] == "--run" || os.Args[1] == "-r" {

			if len(os.Args) < 3 {

				fmt.Println("\n Specify one of the samples to run\n")
			} else {

				sample := os.Args[2]

				if sample == "CountWord" {

					if (len(os.Args) < 4) {

						fmt.Println("\n Specify string to run by CountWord sample \n  e.g: \"Hello World\" ")
					} else {

						fmt.Println("\n Result: ", strings.CountWord(os.Args[3]), "\n")
					}
				}

			}
		}

	}
}

/**
 * printHelp()
 *    nothing fancy command line options 
 */
func printHelp(){

	fmt.Println("Run Samples ")
	fmt.Println("")
	fmt.Println("Usage: ")
	fmt.Println("")
	fmt.Println(" --help -h: ....... print help message     ")
	fmt.Println(" --run  -r: ....... run single sample     ")
	fmt.Println("")
	fmt.Println("Samples:")
	fmt.Println("")
	fmt.Println("  CountWord: gets string e.g. \"Hello World\" and returns words sizes")
	fmt.Println("             inside the string result:  \"Hello5 World5\" ")
	fmt.Println();
}
