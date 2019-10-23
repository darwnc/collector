package exercises

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//BinWatch 二进制手表问题
func BinWatch(c *gin.Context) {
	qNum := c.Query("num")
	if num, err := strconv.Atoi(qNum); err != nil {
		c.JSON(http.StatusOK, gin.H{"num": "num为数字"})
	} else {
		watch := newBinaryWatch()
		result := watch.compute(num)
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

//RemoveDigits 移除k位，求剩下最小的
func RemoveDigits(c *gin.Context) {
	qNum := c.Query("num")
	qK := c.Query("k")
	k, kErr := strconv.Atoi(qK)
	num, nErr := strconv.Atoi(qNum)
	if kErr != nil || nErr != nil {
		c.JSON(http.StatusOK, gin.H{"num&k": "num与k需为数字", "errnummsg": nErr, "errkmsg": kErr})
	} else {
		digits := newDigits()
		result := digits.removeKdigits(strconv.Itoa(num), k)
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}
