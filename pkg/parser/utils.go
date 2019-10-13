package parser

import (
	"fmt"
	"github.com/microspector/microspector/pkg/lookup"
	"github.com/microspector/microspector/pkg/templating"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	Version = "unknown"
	Build   = "unknown"
	Verbose = false
)

//query basically ancodes objects and then tries to find values from objects by their path like variable.sub.value
func query(fieldPath string, thevars map[string]interface{}) interface{} {
	//b, err := json.Marshal(thevars)
	//if err != nil {
	//	fmt.Println("error finding variable value", err)
	//}
	//jq := gojsonq.New()
	//found := jq.JSONString(string(b)).Find(strings.TrimSpace(fieldPath))
	//return found

	v, err := lookup.LookupString(thevars, fieldPath)
	if err != nil {
		return nil
	}
	return v.Interface()
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
	if obj == nil {
		return false
	}

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

func ToVariableName(str string) string {
	segments := strings.Split(str, "-")
	for key, value := range segments {
		segments[key] = strings.Title(value)
	}

	return strings.Join(segments, "")
}

func runop(left, operator, right interface{}) (eq bool) {
	if operator == "!=" {
		operator = "NOTEQUALS"
	}
	op := strings.TrimPrefix(strings.ToUpper(operator.(string)), "NOT")
	not := op != operator

	eq = runOpPositive(left, op, right)

	if not {
		return !eq
	}

	return eq
}

func runOpPositive(left interface{}, operator string, right interface{}) (eq bool) {

	switch operator {
	case "CONTAINS", "CONTAIN":
		switch reflect.TypeOf(left).Kind() { //if left is an array, search right in left :thumbs-up:
		case reflect.Array, reflect.Slice:
			return runOpPositive(right, "IN", left)
		default:
			return strings.Contains(fmt.Sprintf("%s", left), fmt.Sprintf("%s", right))
		}
	case "STARTSWITH", "STARTWITH":
		return strings.HasPrefix(fmt.Sprintf("%s", left), fmt.Sprintf("%s", right))
	case "EQUALS", "==", "EQUAL":
		if oneIsInt(left, right) {
			l, r := convertToInt(left, right)
			return l == r
		} else if bothAreTime(left, right) {
			return left.(time.Time).Unix() == right.(time.Time).Unix()
		}
		return left == right
	case "LT", "GT", ">", "<":

		if bothAreTime(left, right) {
			return left.(time.Time).Unix() < right.(time.Time).Unix()
		} else if oneIsInt(left, right) {
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
		if bothAreTime(left, right) {
			if operator == "GE" || operator == ">=" {
				return left.(time.Time).Unix() >= right.(time.Time).Unix()
			} else {
				return left.(time.Time).Unix() <= right.(time.Time).Unix()
			}
		} else if oneIsInt(left, right) {
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
	case "IN":
		switch reflect.TypeOf(right).Kind() {
		case reflect.Array, reflect.Slice:
			r := reflect.ValueOf(right)
			for i := 0; i < r.Len(); i++ {
				if left == r.Index(i).Interface() {
					return true
				}
			}

		default:
			return false //TODO: or throw an error?
		}
	case "AND":
		return IsTrue(left) && IsTrue(right)
	case "OR":
		return IsTrue(left) || IsTrue(right)

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
	case time.Time:
		return float64(obj.(time.Time).Unix())
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
	case time.Time:
		return obj.(time.Time).Unix()
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

func bothAreTime(left, right interface{}) bool {

	leftTime := false
	rightTime := false

	switch left.(type) {
	case time.Time:
		leftTime = true
	}

	switch right.(type) {
	case time.Time:
		rightTime = true
	}

	return leftTime && rightTime
}

func funcCall(funcName string, args []interface{}) interface{} {
	if x, ok := templating.Functions[funcName]; ok {
		vals := make([]reflect.Value, len(args))

		for index, arg := range args {
			vals[index] = reflect.ValueOf(arg)
		}

		return reflect.ValueOf(x).Call(vals)[0].Interface()

	}
	return nil
}
