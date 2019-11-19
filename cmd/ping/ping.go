package ping

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type pingEntity struct {
	//等待时间
	timeOut time.Duration
	//回显次数
	count int
	//主机名
	host string
	//数据包大小
	size int
}

//icmp proto must run as su user
// https://blog.csdn.net/zxy_666/article/details/79958948
func (entity pingEntity) ping() {
	sendLen := entity.size + reqHeadLen
	recvLen := sendLen + replyHeadLen
	send := make([]byte, sendLen)
	// conn, err := net.DialTimeout("ip4:icmp", host, time.Duration(4*1000*1000))
	send[0] = 8 // echo
	send[1] = 0 // code 0
	send[2] = 0 // checksum
	send[3] = 0 // checksum
	send[4], send[5] = entity.host[0], entity.host[1]
	send[6], send[7] = 1>>8, 1&255
	check := checkSum(send)
	send[2] = byte(check >> 8)
	send[3] = byte(check & 255)
	for {
		conn, err := net.DialTimeout("ip4:icmp", entity.host, entity.timeOut)
		if err != nil {
			fmt.Println("err=", err)
			break
		}
		fmt.Println(send)
		conn.Write(send)

		recv := make([]byte, recvLen)
		conn.Read(recv)
		fmt.Println(recv)
		ttl := int(recv[8])
		fmt.Println("来自 ", conn.RemoteAddr(), "ms TTL=", strconv.Itoa(ttl))
		time.Sleep(1000 * 1000 * 1000 * 3)
	}
}
func checkSum(send []byte) uint16 {
	sum := 0
	length := len(send)
	for i := 0; i < length-1; i += 2 {
		sum += int(send[i])*256 + int(send[i+1])
	}
	if length%2 == 1 {
		sum += int(send[length-1]) * 256 // notice here, why *256?
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

const (
	reqHeadLen   = 8
	replyHeadLen = 20
)
