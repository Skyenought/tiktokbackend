package jwtx

import (
	"fmt"
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	token, err := GenToken("uOvKLmVfztaXGpNYd4Z0I1SiT7MweJhl", time.Now().Unix(), 4548470, 2)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	claim, err := ParseToken("eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTY1ODQ1NzEwOX0.A0CY6KohUQMyEzFeeDlyIHqzvzCLym4oGD9aeUVRiudTQLvzdX5r2R-SChTe6Lih", "uOvKLmVfztaXGpNYd4Z0I1SiT7MweJhl")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	println(claim.UserID)
}

func TestTimeNow(t *testing.T) {
	r := time.Now().Unix()
	println(r)
	println(time.Unix(1653983992759, 0).Format("2006-01-02 15:04:05"))
}
