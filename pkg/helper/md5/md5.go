package md5

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// 加密密码
func hashPassword(password string) (string, error) {
	// 使用bcrypt生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// 验证密码是否匹配
func checkPasswordHash(password, hash string) bool {
	// 使用bcrypt.CompareHashAndPassword比较密码和哈希值
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
