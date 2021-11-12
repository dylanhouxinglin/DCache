package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type inParam struct {
	Key	string	`json:"key" binding:"required"`
	Val	[]byte	`json:"val"`
}

func cacheGet(ctx *gin.Context)  {
	var in inParam
	if err := ctx.ShouldBindJSON(&in); err != nil {
		fmt.Println(err)
	}
	fmt.Println(in.Key)
	ret, _ := Cache.Get(in.Key)
	fmt.Println(string(ret))
}
