/*
 * Author: lihy lihy@zhiannet.com
 * Date: 2022-11-15 11:23:31
 * LastEditors: lihy lihy@zhiannet.com
 * Note: Need note condition
 */
package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	rndVerifyCode = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func SHA1(src ...string) string {
	var data []byte
	buf := bytes.NewBuffer(data)
	for _, v := range src {
		buf.WriteString(v)
	}
	h := sha1.New()
	h.Write(buf.Bytes())
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 求值MD5码
func MD5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func GenPassword(pwd string, salt string) string {
	return SHA1(pwd, salt)
}

func GenStoreSn() uint64 {
	source := rand.NewSource(time.Now().UnixNano())
	var max int64 = 999999999
	var min int64 = 100000000
	return uint64(rand.New(source).Int63n(max-min) + min)
}

func CheckPassword(pwd string, src string, salt string) bool {
	check := SHA1(pwd, salt)
	return strings.Compare(check, src) == 0
}

func GenVerifyCode() string {
	return fmt.Sprintf("%06v", rndVerifyCode.Int31n(1000000))
}

func VerifyPhone(phone string) bool {
	rule := `^1[3-9]\d{9}$`
	rgx := regexp.MustCompile(rule)
	return rgx.MatchString(phone)
}

func IsSameDay(t1 time.Time, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

func TrimJsonArray(s string) string {
	if len(s) >= 2 {
		if s[0] == '[' || s[len(s)-1] == ']' {
			s = s[1 : len(s)-1]
		}
	}
	return s
}

// type T []string|int64

func IsStrEmpty(strs ...string) bool {
	for _, str := range strs {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

func IsInSlice(value int64, slice []int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func IsInSliceString(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func SliceRemoveDuplicates(nums []int64) []int64 {
	encountered := map[int64]bool{}
	result := []int64{}

	for v := range nums {
		if encountered[nums[v]] == false {
			encountered[nums[v]] = true
			result = append(result, nums[v])
		}
	}

	return result
}

func GetMaxLenStr(strs ...string) string {
	m := 0
	i := 0
	for j, str := range strs {
		l := len(str)
		if l > m {
			m = l
			i = j
		}
	}
	return strs[i]
}

func GetMinLenStr(strs ...string) string {
	m := int(^uint(0) >> 1)
	i := 0
	for j, str := range strs {
		l := len(str)
		if l < m {
			m = l
			i = j
		}
	}
	return strs[i]
}

// GetIpsFromDomain 从域名获取解析ip列表
func GetIpsFromDomain(v string) []string {
	var res []string
	ips, err := net.LookupIP(v)
	if err != nil {
		logx.Error(err)
		return res
	}
	for _, ip := range ips {
		res = append(res, ip.String())
	}
	return res
}

// UnicodeToString unicoode 转义
func UnicodeToString(str string) string {
	// fmt.Println(str)
	var needFix bool
	if strings.Contains(str, `\\\\u`) {
		needFix = true
		str = `"` + strings.ReplaceAll(str, `\\\\u`, `\u`) + `"`
	}
	if strings.Contains(str, `\\\u`) {
		needFix = true
		str = `"` + strings.ReplaceAll(str, `\\\u`, `\u`) + `"`
	}
	if strings.Contains(str, `\\u`) {
		needFix = true
		str = `"` + strings.ReplaceAll(str, `\\u`, `\u`) + `"`
	}
	if strings.Contains(str, `\u`) {
		needFix = true
		str = `"` + str + `"`
	}
	if !needFix {
		return str
	}
	if s, err := strconv.Unquote(str); err != nil {
		logx.Error(err)
		return ""
	} else {
		return s
	}
}

func MatchAccount(account string) bool {
	reg := regexp.MustCompile(`^[A-Za-z0-9_]{1,20}$`)
	return reg.MatchString(account)
}

func ToJSONString(obj interface{}) string {

	bs, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bs)
}
func JSONStringToObject(data []byte, obj any) error {
	tp := reflect.TypeOf(obj)
	if tp.Kind() != reflect.Ptr {
		return errors.New("obj not ptr")
	}
	return json.Unmarshal(data, obj)
}
