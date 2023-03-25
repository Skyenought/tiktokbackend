package bcryptx

import (
	"fmt"
	"log"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestCheckPassword(t *testing.T) {
	password := "123456"

	// 生成哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hashed password:", string(hashedPassword))

	// 模拟用户登录
	loginPassword := "123456"

	// 比较哈希值和密码
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(loginPassword))
	if err != nil {
		fmt.Println("Wrong password")
	} else {
		fmt.Println("Correct password")
	}
}
