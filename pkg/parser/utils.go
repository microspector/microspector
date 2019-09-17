package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func executeTemplate(text string, state map[string]interface{}) (string, error) {
	t := template.New("microspector")
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

var zero reflect.Value

func toVariableName(str string) string {
	segments := strings.Split(str, "-")
	for key, value := range segments {
		segments[key] = strings.Title(value)
	}

	return strings.Join(segments, "")
}

func runop(left, operator, right interface{}) bool {
	switch operator {
	case "EQUALS", "==":
		if oneIsInt(left, right) {
			l, r := convertToInt(left, right)
			return l == r
		}
		return left == right
	case "NOTEQUALS", "!=":
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

	}

	return false
}

func convertToInt(left, right interface{}) (l, r int) {
	return intVal(left), intVal(right)
}

func convertToFloat(left, right interface{}) (l, r float64) {
	return floatVal(left), floatVal(right)
}

func floatVal(obj interface{}) float64 {
	switch obj.(type) {
	case int, int64, int32:
		return obj.(float64)
	case float32, float64:
		return obj.(float64)
	default:
		f, _err := strconv.ParseFloat(fmt.Sprintf("%s", obj), 64)
		if _err == nil {
			return f
		}
		return 0
	}
}

func intVal(obj interface{}) int {
	switch obj.(type) {
	case int, int64, int32:
		return obj.(int)
	case float32, float64:
		return obj.(int)
	default:
		f, _err := strconv.Atoi(fmt.Sprintf("%s", obj))
		if _err == nil {
			return f
		}
		return 0
	}
}

func oneIsInt(left, right interface{}) bool {
	switch left.(type) {
	case int, int64, int32:
		return true
	}

	switch right.(type) {
	case int, int64, int32:
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
