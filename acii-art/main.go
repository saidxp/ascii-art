package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

func main(){
		 
		if len(os.Args[:1]) != 1 {
			return
		}
		word := os.Args[1]
		file := "standard.txt"
		alphabetascciart , err := os.Open(file)
		
		if err != nil {
			fmt.Println("not open")
		}
		contianalphabet  := bufio.NewScanner(alphabetascciart)
		m := make(map[string]string)

		curent := ""
		count := 32
		for contianalphabet.Scan() {
			line := contianalphabet.Text() 
			if line == "" {
				if curent != "" && count <= 126 { 
				 
				m[string(rune(count))] = curent
				curent = ""
				count++
				}
			}else {
			curent += line + "\n"
			}
		}
	defer alphabetascciart.Close()
	
	if curent != "" && count <= 126 {
		m[string(rune(count))] = curent
	} 

	output := make([]string, 8)
	i := 0
	for i < len(word) {
		charStr := string(word[i])
		if word[i] == '\\' && i + 1 < len(word) && word[i + 1] == 'n' {
			for _, char := range output {
			// her i print the content !
				fmt.Println(char)
			}
			// Her I print New line 
			fmt.Println()
			output = make([]string, 8)
			i += 2
			continue
		}

		lines := strings.Split(m[charStr], "\n")
	 
		for j, line := range lines {
			if j < 8 {
				output[j] += line + "  "
			}
		}
		i++
	}

	for _, line := range output {
		  
		fmt.Printf("%s\n",line)
		}
	}
	
}