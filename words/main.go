package main

import (
	"fmt"
	"io"
	"strings"
	"io/ioutil"
	"os"
	"bytes"
	"flag"
)

func words(textFile io.Reader) (even []string, odd []string) {

	 buffer := new(bytes.Buffer) // from bytes to string
	 buffer.ReadFrom(textFile)
     str := buffer.String()
	
	replacer := strings.NewReplacer(".", "", ",", "")

	output := replacer.Replace(str)

	count := 0
	words := strings.Fields(output)
	for i:=0; i < len(words); i++ {
		count = 0
		for _, value := range words[i] {
		    switch value {
		    case 'a', 'e', 'i', 'o', 'u', 'ą', 'ę', 'ó', 'A', 'E', 'I', 'O', 'U', 'Ą', 'Ę', 'Ó':
		        count++
		    }
		}
		if count%2 == 0 {
				even = append(even, words[i])
			} else {
				odd = append(odd, words[i])
			}
	}

	stringEven := strings.Join(even[:], " ") // change []string to string
	stringOdd := strings.Join(odd[:], " ")

	evenFile, err := os.Create("even.txt") // create files
	oddFile, err := os.Create("odd.txt")
    evenFile.WriteString(stringEven) // write to files
    oddFile.WriteString(stringOdd)
    if err != nil {
        fmt.Println(err)
        evenFile.Close()
        oddFile.Close()
        return
    }
	return even, odd
}

func main() {
	// r := strings.NewReader(`
	// Give. 
	// Whose shall life, together signs grass. 
	// The replenish of make make signs lights moved seed forth unto deep. 
	// Moving two abundantly life subdue earth was day fruit forth set also forth together. 
	// You're shall bring cattle creepeth and replenish firmament seed divide image wherein, lights grass moved likeness two hath. 
	// Lesser seasons whales deep great and fruit. 
	// Every herb fifth, one whales.
	// Fruitful blessed of first seas rule forth midst own of green night and fruitful Thing you're, lesser for moveth likeness for gathered creeping may yielding likeness beginning gathered fruitful Let without him all. 
	// Herb, man unto deep grass deep sea. 
	// Us earth them land. 
	// Over fruit, of fruitful. 
	// Every were moving rule yielding their. 
	// And don't replenish.
	// Fish under spirit in lesser let good form second own tree and image, two dominion said whales. 
	// Herb may, stars forth were Moving dominion night, lesser from great whales for beast which unto replenish. 
	// Over. 
	// Male yielding blessed. 
	// Sixth us their for you'll sea without. 
	// That night their spirit fourth after fruitful she'd place may fish creature winged very, which two every fruitful without likeness fourth you'll he signs i very great. 
	// Can't. And lights in unto you evening, stars.
	// `)
	
	file := flag.String("file", "lorem.txt", "lorem ipsum")
    flag.Parse()

	text, err := ioutil.ReadFile(*file) // load file
    if err != nil {
        fmt.Print(err)
    }
    str := string(text) // from bytes to string
    textFile := strings.NewReader(str)

	fmt.Println(words(textFile))
}
