package templating

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"html/template"
	"math/rand"
	"net/url"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
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
	"now": func() time.Time {
		return time.Now()
	},
	"timestamp": func() int64 {
		return time.Now().Unix()
	},
	"unix": func(time time.Time) int64 {
		return time.Unix()
	},
	"rand": func(min, max int64) int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(int(max-min)) + int(min)
	},
	"openssl_rand": OpenSslRand,
	"str_len":      func(str string) int { return len(str) },
	"len": func(obj interface{}) int {
		r := reflect.ValueOf(obj)
		return r.Len()
	},
	"trim": func(str string) string { return strings.TrimSpace(str) },
	"hash_md5": func(val interface{}) string {
		data := []byte(fmt.Sprintf("%s", val))
		return fmt.Sprintf("%x", md5.Sum(data))
	},
	"hash_sha256": func(val interface{}) string {
		data := []byte(fmt.Sprintf("%s", val))
		return fmt.Sprintf("%x", sha256.Sum256(data))
	},
	"url_encode": func(val interface{}) string {
		return url.QueryEscape(fmt.Sprintf("%s", val))
	},
}

func OpenSslRand(len int64, enc string) string {
	bytes := make([]byte, len)
	bytes, err := exec.Command("openssl", "rand", "-"+enc, strconv.Itoa(int(len))).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(bytes))
}
