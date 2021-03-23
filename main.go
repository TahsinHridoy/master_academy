package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	resp := []byte("\x00" + "thh34111@gmail.com" + "\x00" + "5858heart.")
	//fmt.Println(string(resp), resp)
	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)
}
