package utils

import (
	"fmt"
	"testing"
)

/**
* @author wangpengcheng@ccssoft.com.cn
* @date 2018/1/20 14:29
 */

func TestGetToken(t *testing.T) {
	openshiftUrl := "https://openshift-master:8443"
	username := "wpc"
	password := "wpc"

	var err error
	var token *string
	token, err = GetToken(openshiftUrl, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Token:", *token)
}

func TestTokenToBase64Encode(t *testing.T) {
	tokenStr := "w8ZxAJRFI7eZiJu04x1uvt5dqaZcJH4GMsrrxpved6w"
	fmt.Println(TokenToBase64Encode(tokenStr))
}
