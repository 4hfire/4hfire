package jwt

import (
	"fmt"
	"testing"
)

func TestCreateJwt(t *testing.T) {
	jwt := NewJWT(
		"561528c8-2d58-46df-a075-516bef5b7f80",
		86400,
	)
	payload := make(Payload)
	payload["uid"] = "xx-aa-ww"
	payload["zz"] = 123456
	token, err := jwt.Token(payload)
	t.Log(token, err)
}

func TestParse(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50IjoieHVsZWkwMDEiLCJkZXZpY2UiOiI0QjRGRjQyMy02MzEyLTQwNjQtODk2Qy1FMTM5NEQ0QTFEM0QiLCJleHAiOjE2OTA2MjMwNjYsImlhdCI6MTY5MDM2Mzg2NiwibmFtZSI6Inh1bGVpMDAxIiwidWlkIjoiZjkzZGVkMzYtMGM0Ni00ZjhkLTllM2EtM2UzYjVjMTVhYWMyIn0.qgzwaqhq0Wqrqpygq_-3I_7e28EsL1Rkrgv-B1PJLeg"
	parser := NewJWT("561528c8-2d58-46df-a075-516bef5b7f88", 259200)
	p, err := parser.Parse(token)
	fmt.Println(p, err)
}
