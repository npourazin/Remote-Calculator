package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Println("Enter:")
		argis := readData()

		for i := 0; i < len(argis); i++ {
			//reader := bufio.NewReader(os.Stdin)
			//use your own data reader to read data
			//nums, _ := reader.ReadString('\n')
			//fmt.Println(nums)
			nums := argis[i]

			if nums == "quit\n" {
				fmt.Println("suicide")
				break
			}

			//prepare the string to be sent
			write := nums + "."

			_, err = connection.Write([]byte(write))
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println("request sent")

			res, _ := bufio.NewReader(connection).ReadString('\n')

			fmt.Println("Result:", res)
		}
	}
	_, err = connection.Write([]byte("quit"))
	//unhandled error
	connection.Close()
	return

}

func readData() []string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	argis := []string{}
	j := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			s := text[j:i]
			j = i + 1
			if j > len(text) {
				break
			}
			argis = append(argis, s)
			//fmt.Println("--",s)
		}
	}
	if j < len(text) {
		s := text[j : len(text)-1]
		argis = append(argis, s)
	}

	return argis
}

/**
Issues:
1- wrong input not handled
*/
