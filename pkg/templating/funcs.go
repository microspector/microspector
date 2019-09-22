package templating

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"html/template"
	"os/exec"
	"strconv"
	"strings"
)

//compiles strings using golang template engine and returns the result as string
func ExecuteTemplate(text string, state map[string]interface{}) (string, error) {
	t := template.New("microspector").Funcs(Functions)
	_, err := t.Parse(text)

	if err != nil {
		return "", nil
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, state); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

var Functions = template.FuncMap{
	"openssl_rand": OpenSslRand,
	"str_len":      func(str string) int { return len(str) },
	"hash_md5": func(val interface{}) string {
		data := []byte(fmt.Sprintf("%s", val))
		return fmt.Sprintf("%x", md5.Sum(data))
	},
	"hash_sha256": func(val interface{}) string {
		data := []byte(fmt.Sprintf("%s", val))
		return fmt.Sprintf("%x", sha256.Sum256(data))
	},
}

func OpenSslRand(len int, enc string) string {
	bytes := make([]byte, len)
	bytes, err := exec.Command("openssl", "rand", "-"+enc, strconv.Itoa(len)).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(bytes))
}
