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
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter:")
		nums, _ := reader.ReadString('\n')
		//fmt.Println(nums)
		if nums == "quit\n" {
			fmt.Println("suicide")
			break
		}

		//prepare the string to be sent
		write := nums[0:len(nums)-1] + "."

		_, err = connection.Write([]byte(write))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("request sent")

		res, _ := bufio.NewReader(connection).ReadString('\n')

		fmt.Println("Result:", res)
	}
	_, err = connection.Write([]byte("quit"))
	//unhandled error
	connection.Close()
	return

}

/**
Issues:
1- if file is closed server is dooooomed
2- wrong input not handled
*/
