//package Server
package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func main() {
	client, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		connection, err := client.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			receive, err := bufio.NewReader(connection).ReadString('.')
			if err != nil {
				fmt.Println(err)
				break
			}

			ans := calculate(receive)
			fmt.Println(receive + " = " + strconv.Itoa(ans))

			if receive == "quit" {
				break
			}

			//errors are to be handled below
			//if 1 == 0 {
			//	_, err = connection.Write([]byte("Invalid Input!"))
			//	if err != nil {
			//		fmt.Println(err)
			//	}
			//	connection.Close()
			//	return
			//}

			_, err = connection.Write([]byte(strconv.Itoa(ans) + "\n"))
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		connection.Close()
		//unhandled error

	}

}

func calculate(n string) int {
	text := n
	if len(n) > 0 && n[len(n)-1] == '\n' {
		text = n[0 : len(n)-1]
	}
	//or -3
	argis := []int{}
	j := 0
	for i := 0; i < len(text); i++ {
		if text[i] == '+' {
			s := text[j:i]
			if j != 0 {
				d := text[j-1 : j]
				//fmt.Println(text[j-1:j])
				s = d + s
			}

			j = i + 1
			if j > len(text) {
				break
			}
			p, _ := strconv.Atoi(s)
			argis = append(argis, p)
			//fmt.Println("--",p)
		}
		if text[i] == '-' {
			s := text[j:i]
			if j != 0 {
				d := text[j-1 : j]
				//fmt.Println(text[j-1:j])
				s = d + s
			}
			d := text[j:j]
			s = d + s
			j = i + 1
			if j > len(text) {
				break
			}
			p, _ := strconv.Atoi(s)
			argis = append(argis, p)
			//fmt.Println("--",p)
		}
	}
	if j < len(text) {
		s := text[j : len(text)-1]
		if j != 0 {
			d := text[j-1 : j]
			//fmt.Println(text[j-1:j])
			s = d + s
		}
		p, _ := strconv.Atoi(s)
		argis = append(argis, p)
		//fmt.Println("--",p)
	}

	//fmt.Println(len(text))
	//fmt.Println("len ",len(argis), "\n", argis)
	sum := 0
	for i := 0; i < len(argis); i++ {
		sum += argis[i]
	}
	//fmt.Println(sum)
	return sum
}

/**
Issues:
1- if file is closed server is dooooomed
2- wrong input not handled
3- explain input way to the idiot client
4- handle multiple clients?

*/
