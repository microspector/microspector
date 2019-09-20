package parser

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// add returns the sum of a and b.
func add(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() + bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() + int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) + bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() + c, _e
		default:
			log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) + bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() + bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) + bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() + c, _e
		default:
			log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() + float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() + float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() + bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() + c, _e
		default:
			log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	case reflect.String:
		ca, _ := strconv.ParseInt(av.String(), 10, 64)
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ca + bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ca + int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(ca) + bv.Float(), nil
		case reflect.String:
			cb, _e := strconv.ParseInt(bv.String(), 10, 64)
			return ca + cb, _e
		default:
			log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	default:
		log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", av, a))
		return nil, fmt.Errorf("add: unknown type for %q (%T)", av, a)
	}
}

// subtract returns the difference of b from a.
func subtract(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() - bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() - int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) - bv.Float(), nil
		default:
			log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("subtract: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) - bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() - bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) - bv.Float(), nil
		default:
			log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("subtract: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() - float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() - float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() - bv.Float(), nil
		default:
			return nil, fmt.Errorf("subtract: unknown type for %q (%T)", bv, b)
		}
	case reflect.String:
		ca, _ := strconv.ParseInt(av.String(), 10, 64)
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ca - bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ca - int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(ca) - bv.Float(), nil
		case reflect.String:
			cb, _e := strconv.ParseInt(bv.String(), 10, 64)
			return ca - cb, _e
		default:
			log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("subtract: unknown type for %q (%T)", bv, b)
		}
	default:
		log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", av, a))
		return nil, fmt.Errorf("subtract: unknown type for %q (%T)", av, a)
	}
}

// multiply returns the product of a and b.
func multiply(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	//log.Fatalln("multiply kinds", av.Kind(), bv.Kind())

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() * bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() * int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) * bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() * c, _e
		default:
			log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("multiply: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) * bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() * bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) * bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return int64(av.Uint()) * c, _e
		default:
			log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("multiply: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() * float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() * float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() * bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return av.Float() * float64(c), _e
		default:
			log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("multiply: unknown type for %q (%T)", bv, b)
		}
	case reflect.String:
		ca, _e := strconv.ParseInt(av.String(), 10, 64)
		if _e != nil {
			return nil, fmt.Errorf("multiply: unknown type for %q (%T)", av, a)
		}
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ca * bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ca * int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() * bv.Float(), nil
		case reflect.String:
			cb, _e := strconv.ParseInt(bv.String(), 10, 64)
			return ca * cb, _e
		default:
			log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("multiply: unknown type for %q (%T)", bv, b)
		}
	default:
		log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", av, a))
		return nil, fmt.Errorf("multiply: unknown type for %q (%T)", av, a)
	}

}

// divide returns the division of b from a.
func divide(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() / bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() / int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) / bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() / c, _e
		default:
			return nil, fmt.Errorf("divide: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) / bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() / bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) / bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return int64(av.Uint()) / c, _e
		default:
			return nil, fmt.Errorf("divide: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() / float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() / float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() / bv.Float(), nil
		case reflect.String:
			c, _e := strconv.ParseInt(bv.String(), 10, 64)
			return av.Float() / float64(c), _e
		default:
			return nil, fmt.Errorf("divide: unknown type for %q (%T)", bv, b)
		}
	case reflect.String:
		ca, _e := strconv.ParseInt(av.String(), 10, 64)
		if _e != nil {
			return nil, fmt.Errorf("divide: unknown type for %q (%T)", av, a)
		}
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ca / bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ca / int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() / bv.Float(), nil
		case reflect.String:
			cb, _e := strconv.ParseInt(bv.String(), 10, 64)
			return ca / cb, _e
		default:
			log.Fatal(fmt.Errorf("divide: unknown type for %q (%T)", bv, b))
			return nil, fmt.Errorf("divide: unknown type for %q (%T)", bv, b)
		}
	default:
		return nil, fmt.Errorf("divide: unknown type for %q (%T)", av, a)
	}
}

func umin(a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return -av.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return -av.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return -av.Float(), nil
	case reflect.String:
		ca, _ := strconv.ParseInt(av.String(), 10, 64)
		return -ca, nil
	default:
		return nil, fmt.Errorf("unknown type for %q (%T)", av, a)
	}
}

// mod returns a % b
func mod(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() % bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() % int64(bv.Uint()), nil
		default:
			return nil, fmt.Errorf("mod: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) % bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() % bv.Uint(), nil
		default:
			return nil, fmt.Errorf("mod: unknown type for %q (%T)", bv, b)
		}
	default:
		return nil, fmt.Errorf("mod: unknown type for %q (%T)", av, a)
	}
}
