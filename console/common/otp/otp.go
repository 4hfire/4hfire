/*
 * Author: lihy lihy@zhiannet.com
 * Date: 2023-03-11 16:10:43
 * LastEditors: lihy lihy@zhiannet.com
 * Note: Need note condition
 */
package otp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// 生成32位随机序列
var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	codeLen = len(codes)
)

func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func RandNewStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}

	return string(data)
}

/*
Create
协议格式：otpauth://totp/demo.com?secret=sds&period=60&digits=8&iiuser=Esd
参数说明：

	demo.com 账号名称 显示在左下角
	secret 种子密钥
	period  生成值的有效期 30s
	digits 	口令长度 6
	issuer 备注信息 左上角
*/
func Create(name, secretId string) (string, error) {
	return "otpauth://totp/4hfire-" + name + "?issuer=" + name + "&secret=" + secretId, nil
}
func Check(secretId, code string) bool {
	//判断google验证码与code是否一致
	result := ReturnCode(secretId)
	//result需要与code类型一致
	return code == fmt.Sprint(result)
}
