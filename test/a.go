package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	_ "io"
	"math/rand"
	"time"
)

const (
	APPID     string = "1252228829"
	SECRETID  string = "AKIDWL4fDbyrncGvSoH4vdBXhgV8bu33qX0J"
	SECRETKEY string = "IQNUTIwPhniYFlATcaNzQhms0DTcaRBl"
	BUCKET    string = "airdb"
	FOLDER    string = "wechat"
)

func main() {
	// conn, err := gorm.Open("mysql", beego.AppConfig.String("gdbc"))
	// if err != nil {
	// 	panic("connect to databases failed!")
	// }
	// defer conn.Close()
	fmt.Println(time.Now().Unix())
	curtime := time.Now().Unix()
	// Original = "a=[appid]&b=[bucket]&k=[SecretID]&e=[expiredTime]&t=[currentTime]&r=[rand]&f="
	// fileid := "/200001/newbucket/tencent_test.jpg"
	randnum := rand.Intn(1000000000)
	multi_effect_signature := fmt.Sprintf("a=%s&b=%s&k=%s&e=%d&t=%d&r=%d&f=", APPID, BUCKET, SECRETID, curtime+60, curtime, randnum)
	// multi_effect_signature = "a=200001&b=newbucket&k=AKIDUfLUEUigQiXqm7CVSspKJnuaiIKtxqAv&e=1437995704&t=1437995644&r=2081660421&f="
	fmt.Println(multi_effect_signature)

	mac := hmac.New(sha1.New, []byte(SECRETKEY))
	mac.Write([]byte(multi_effect_signature))
	macSum := mac.Sum(nil)
	fmt.Printf("mac: %s\n", macSum)

	// $multi_effect_signature = base64_encode(hash_hmac('SHA1', $multi_effect_signature, $secret_key, true).$multi_effect_signature);
	data64 := base64.StdEncoding.EncodeToString([]byte(string(macSum) + multi_effect_signature))

	fmt.Println("data64: ", data64)

}
