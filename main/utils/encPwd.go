package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// 产生一个长度为6的 salt 值
func NewSalt() string {
	// 使用当前时间戳作为随机数种子，产生一个[0, 1000000)的6位伪随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	return fmt.Sprintf("%06v", r)
}

// 对密码字符串进行加盐 md5，返回 md5 的十六进制
func EncPasswd(pwd string, salt string) (pwdmd5 string) {
	pwd = pwd + salt
	return fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
}
