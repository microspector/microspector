package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"github.com/tufanbarisyildirim/microspector/pkg/templating"
	"html/template"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var Version = "unknown"
var Build = "unknown"

//compiles strings using golang template engine and returns the result as string
func executeTemplate(text string, state map[string]interface{}) (string, error) {
	t := template.New("microspector").Funcs(templating.Functions)
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

//query basically ancodes objects and then tries to find values from objects by their path like variable.sub.value
func query(fieldPath string, thevars map[string]interface{}) interface{} {
	b, err := json.Marshal(thevars)
	if err != nil {
		fmt.Println("error finding variable value", err)
	}
	jq := gojsonq.New()
	found := jq.JSONString(string(b)).Find(strings.TrimSpace(fieldPath))
	return found
}

func isTemplate(text string) bool {
	return strings.Contains(text, "{{") && strings.Contains(text, "}}")
}

//current unix timestamp
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// IsTrue reports whether the value is 'true', in the sense of not the zero of its type,
// and whether the value has a meaningful truth value. This is the definition of
// truth used by if and other such actions.
func IsTrue(val interface{}) bool {
	return isTrue(reflect.ValueOf(val))
}

func isTrue(val reflect.Value) (truth bool) {
	if !val.IsValid() {
		// Something like var x interface{}, never set. It's a form of nil.
		return false
	}
	switch val.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		truth = val.Len() > 0
	case reflect.Bool:
		truth = val.Bool()
	case reflect.Complex64, reflect.Complex128:
		truth = val.Complex() != 0
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Interface:
		truth = !val.IsNil()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		truth = val.Int() != 0
	case reflect.Float32, reflect.Float64:
		truth = val.Float() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		truth = val.Uint() != 0
	case reflect.Struct:
		truth = true // Struct values are always true.
	default:
		return false
	}
	return truth
}

var typeMap = map[reflect.Kind]string{
	reflect.Array:   "array",
	reflect.Map:     "array",
	reflect.Slice:   "array",
	reflect.String:  "string",
	reflect.Int:     "integer",
	reflect.Int8:    "integer",
	reflect.Int16:   "integer",
	reflect.Int32:   "integer",
	reflect.Int64:   "integer",
	reflect.Uint:    "integer",
	reflect.Uint8:   "integer",
	reflect.Uint16:  "integer",
	reflect.Uint32:  "integer",
	reflect.Uint64:  "integer",
	reflect.Uintptr: "integer",
	reflect.Struct:  "object",
	reflect.Float32: "float",
	reflect.Float64: "float",
	reflect.Bool:    "boolean",
}

func IsTypeOf(obj interface{}, typeName string) bool {
	kind := reflect.TypeOf(obj).Kind()

	typeName = strings.ToLower(typeName)
	if typeName == "int" {
		typeName = "integer"
	} else if typeName == "bool" {
		typeName = "boolean"
	}

	if name, ok := typeMap[kind]; ok {
		return typeName == name
	}

	return false
}

var zero reflect.Value

func toVariableName(str string) string {
	segments := strings.Split(str, "-")
	for key, value := range segments {
		segments[key] = strings.Title(value)
	}

	return strings.Join(segments, "")
}

func runop(left, operator, right interface{}) bool {
	switch strings.ToUpper(operator.(string)) {
	case "EQUALS", "==", "EQUAL":
		if oneIsInt(left, right) {
			l, r := convertToInt(left, right)
			return l == r
		}
		return left == right
	case "NOTEQUALS", "!=", "NOTEQUAL":
		if oneIsInt(left, right) {
			l, r := convertToInt(left, right)
			return l != r
		}
		return left != right
	case "CONTAINS":
		return strings.Contains(fmt.Sprintf("%s", left), fmt.Sprintf("%s", right))
	case "STARTSWITH":
		return strings.HasPrefix(fmt.Sprintf("%s", left), fmt.Sprintf("%s", right))
	case "LT", "GT", ">", "<":

		if oneIsInt(left, right) {
			l, r := convertToInt(left, right)
			if operator == "GT" || operator == ">" {
				return l > r
			} else {
				return l < r
			}
		} else if oneIsFloat(left, right) {
			l, r := convertToFloat(left, right)
			if operator == "GT" || operator == ">" {
				return l > r
			} else {
				return l < r
			}
		} else {
			if operator == "GT" || operator == ">" {
				return floatVal(left) > floatVal(right)
			} else {
				return floatVal(left) < floatVal(right)
			}

		}
	case "LE", "GE", "<=", ">=":
		if oneIsInt(left, right) {
			l, r := convertToInt(left, right)
			if operator == "LE" || operator == "<=" {
				return l <= r
			} else {
				return l >= r
			}
		} else if oneIsFloat(left, right) {
			l, r := convertToFloat(left, right)
			if operator == "GE" || operator == ">=" {
				return l >= r
			} else {
				return l <= r
			}
		} else {
			if operator == "GE" || operator == ">=" {
				return floatVal(left) >= floatVal(right)
			} else {
				return floatVal(left) <= floatVal(right)
			}

		}
	case "MATCHES", "MATCH":
		match, _ := regexp.MatchString(fmt.Sprintf("%s", left), fmt.Sprintf("%s", right))
		return match
	case "IS":
		return IsTypeOf(left, right.(string))
	case "ISNOT":
		return !IsTypeOf(left, right.(string))

	}

	return false
}

func convertToInt(left, right interface{}) (l, r int64) {
	return intVal(left), intVal(right)
}

func convertToFloat(left, right interface{}) (l, r float64) {
	return floatVal(left), floatVal(right)
}

func floatVal(obj interface{}) float64 {

	switch obj.(type) {
	case int:
		return float64(obj.(int))
	case int64:
		return float64(obj.(int64))
	case int32:
		return float64(obj.(int32))
	case float32:
		return float64(obj.(float32))
	case float64:
		return obj.(float64)
	default:
		f, _err := strconv.ParseFloat(fmt.Sprintf("%s", obj), 64)
		if _err == nil {
			return f
		}
		return 0
	}
}

func intVal(obj interface{}) int64 {
	switch obj.(type) {
	case int:
		return int64(obj.(int))
	case int64:
		return obj.(int64)
	case int32:
		return int64(obj.(int32))
	case float32:
		return int64(obj.(float32))
	case float64:
		return int64(obj.(float64))
	default:
		f, _err := strconv.Atoi(fmt.Sprintf("%s", obj))
		if _err == nil {
			return int64(f)
		}
		return 0
	}
}

func oneIsInt(left, right interface{}) bool {
	switch left.(type) {
	case int, int64, int32, uint, uint32, uint64:
		return true
	}

	switch right.(type) {
	case int, int64, int32, uint, uint32, uint64:
		return true
	}

	return false
}

func oneIsFloat(left, right interface{}) bool {
	switch left.(type) {
	case float32, float64:
		return true
	}

	switch right.(type) {
	case float32, float64:
		return true
	}

	return false
}
