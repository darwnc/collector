package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	listen()
}

func dial() {
	// addr, _ := net.ResolveTCPAddr("tcp", ":8080")
	conn, DialErr := net.Dial("tcp", ":8080")
	// conn, DialErr := net.Dial("tcp", ":8080")
	if DialErr != nil {
		fmt.Println(DialErr)
		return
	}
	// conn.SetDeadline(time.Now().Add(time.Duration(1000 * 1000 * 1000)))
	// scan := bufio.NewScanner(os.Stdin)
	scan := bufio.NewScanner(os.Stdin)
	i := 0
	for {
		var input string
		if scan.Scan() {
			input = scan.Text()
		}
		i++
		input = "input" + strconv.Itoa(i) + "\n" + input
		wn, err := conn.Write([]byte(input))
		// n, err := conn.Write([]byte(input))
		fmt.Println("input=", "n", wn, "err", err)
		if i == 100 {
			break
		}
		buff := make([]byte, 100)
		rn, _ := conn.Read(buff)
		fmt.Println(string(buff[:rn]))

	}

	// go func() {
	// 	scan := bufio.NewScanner(os.Stdin)
	// 	for {
	// 		if scan.Scan() {
	// 			input := scan.Text()
	// 			w := bufio.NewWriter(conn)
	// 			n, err := w.WriteString(input)
	// 			w.Flush()
	// 			// n, err := conn.Write([]byte(input))
	// 			fmt.Println("scan=", input, "n", n, "err", err)
	// 			if input == "exit" {
	// 				break
	// 			}
	// 		}
	// 	}
	// }()
	// go func() {

	// }()
	// fmt.Println("addr=", conn.LocalAddr())
	// go func() {
	// 	for {
	// 		fmt.Println("dial read in")
	// 		buff, _ := ioutil.ReadAll(conn)
	// 		fmt.Println("return", string(buff))
	// 	}
	// }()
	fmt.Println("dial read in")
	buff := make([]byte, 100)
	rn, _ := conn.Read(buff)
	fmt.Println("return", string(buff[:rn]))
	conn.Close()
}
func listen() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8080")
	// addr := net.TCPAddr{}
	// addr.Port = 8080
	lis, err := net.ListenTCP("tcp", addr)
	// net.ListenTCP(network, laddr)
	if err != nil {
		fmt.Println(err)
	}
	// conn, lisError := lis.Accept()
	over := make(chan int)
	conn, lisError := lis.AcceptTCP()
	if lisError != nil {
		fmt.Println(lisError)
		return
	}
	// fmt.Println("read")
	// buff, _ := ioutil.ReadAll(conn)
	// in := string(buff)
	// fmt.Println("read in", in)
	// sign <- 1
	reader := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	go func() {
		for {
			fmt.Println("server read")
			in, readErr := reader.ReadString('\n')
			// buff, readErr := ioutil.ReadAll(conn)
			fmt.Println("readErr", readErr)
			// in := string(buff)
			fmt.Println("read in", in)
			// bufio.NewReadWriter(r, w)
			// bufio.NewWriter(conn)
			w.WriteString("echo->" + string(in))

			w.Flush()
			// conn.Write([]byte())
			if readErr == io.EOF {
				over <- 1
				break
			}
		}
	}()
	<-over
	// <-over
}
