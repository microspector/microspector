package templatefuncs

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func OpenSslRand(len int, enc string) string {
	bytes := make([]byte, len)
	bytes, err := exec.Command("openssl", "rand", "-"+enc, strconv.Itoa(len)).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(bytes))
}