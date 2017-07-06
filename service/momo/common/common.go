package common

import (
	"bytes"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func RandHex(n int) string {
	u1 := uuid.NewV4()
	d, _ := u1.MarshalBinary()
	buf := new(bytes.Buffer)
	for i := 0; i < n/2; i++ {
		buf.WriteString(fmt.Sprintf("%02x", d[i]))
	}

	return buf.String()
}

func UUID() string {
	u := uuid.NewV4()
	return u.String()
}
