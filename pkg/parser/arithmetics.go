package parser

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// add returns the sum of a and b.
func add(b, a interface{}) interface{} {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() + bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() + int64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) + bv.Float()
		case reflect.String:
			c, _ := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() + c
		default:
			log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) + bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() + bv.Uint()
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) + bv.Float()
		case reflect.String:
			c, _ := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() + c
		default:
			log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() + float64(bv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() + float64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return av.Float() + bv.Float()
		case reflect.String:
			c, _ := strconv.ParseFloat(bv.String(), 64)
			return av.Float() + float64(c)
		default:
			log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.String:
		ca, _ := strconv.ParseInt(av.String(), 10, 64)
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ca + bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ca + int64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return float64(ca) + bv.Float()
		case reflect.String:
			cb, _ := strconv.ParseInt(bv.String(), 10, 64)
			return ca + cb
		default:
			log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", bv, b))
			return nil
		}
	default:
		log.Fatal(fmt.Errorf("add: unknown type for %q (%T)", av, a))
		return nil
	}
}

// subtract returns the difference of b from a.
func subtract(b, a interface{}) interface{} {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() - bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() - int64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) - bv.Float()
		default:
			log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) - bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() - bv.Uint()
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) - bv.Float()
		default:
			log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() - float64(bv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() - float64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return av.Float() - bv.Float()
		default:
			log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.String:
		ca, _ := strconv.ParseInt(av.String(), 10, 64)
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ca - bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ca - int64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return float64(ca) - bv.Float()
		case reflect.String:
			cb, _ := strconv.ParseInt(bv.String(), 10, 64)
			return ca - cb
		default:
			log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", bv, b))
			return nil
		}
	default:
		log.Fatal(fmt.Errorf("subtract: unknown type for %q (%T)", av, a))
		return nil
	}
}

// multiply returns the product of a and b.
func multiply(b, a interface{}) interface{} {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	//log.Fatalln("multiply kinds", av.Kind(), bv.Kind())

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() * bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() * int64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) * bv.Float()
		case reflect.String:
			c, _ := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() * c
		default:
			log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) * bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() * bv.Uint()
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) * bv.Float()
		case reflect.String:
			c, _ := strconv.ParseInt(bv.String(), 10, 64)
			return int64(av.Uint()) * c
		default:
			log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() * float64(bv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() * float64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return av.Float() * bv.Float()
		case reflect.String:
			c, _ := strconv.ParseFloat(bv.String(), 64)
			return av.Float() * c
		default:
			log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.String:
		ca, _e := strconv.ParseInt(av.String(), 10, 64)
		if _e != nil {
			return nil
		}
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ca * bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ca * int64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return av.Float() * bv.Float()
		case reflect.String:
			cb, _ := strconv.ParseInt(bv.String(), 10, 64)
			return ca * cb
		default:
			log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", bv, b))
			return nil
		}
	default:
		log.Fatal(fmt.Errorf("multiply: unknown type for %q (%T)", av, a))
		return nil
	}

}

// divide returns the division of b from a.
func divide(b, a interface{}) interface{} {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() / bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() / int64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) / bv.Float()
		case reflect.String:
			c, _ := strconv.ParseInt(bv.String(), 10, 64)
			return av.Int() / c
		default:
			log.Fatal(fmt.Errorf("divide: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) / bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() / bv.Uint()
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) / bv.Float()
		case reflect.String:
			c, _ := strconv.ParseInt(bv.String(), 10, 64)
			return int64(av.Uint()) / c
		default:
			log.Fatal(fmt.Errorf("divide: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() / float64(bv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() / float64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return av.Float() / bv.Float()
		case reflect.String:
			c, _ := strconv.ParseInt(bv.String(), 10, 64)
			return av.Float() / float64(c)
		default:
			log.Fatal(fmt.Errorf("divide: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.String:
		ca, _e := strconv.ParseInt(av.String(), 10, 64)
		if _e != nil {
			log.Fatal(fmt.Errorf("divide: unknown type for %q (%T)", av, a))
			return nil
		}
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ca / bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ca / int64(bv.Uint())
		case reflect.Float32, reflect.Float64:
			return av.Float() / bv.Float()
		case reflect.String:
			cb, _ := strconv.ParseInt(bv.String(), 10, 64)
			return ca / cb
		default:
			log.Fatal(fmt.Errorf("divide: unknown type for %q (%T)", bv, b))
			return nil
		}
	default:
		log.Fatal(fmt.Errorf("divide: unknown type for %q (%T)", av, a))
		return nil
	}
}

func umin(a interface{}) interface{} {
	av := reflect.ValueOf(a)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return -av.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return -av.Uint()
	case reflect.Float32, reflect.Float64:
		return -av.Float()
	case reflect.String:
		ca, _ := strconv.ParseInt(av.String(), 10, 64)
		return -ca
	default:
		log.Fatal(fmt.Errorf("unknown type for %q (%T)", av, a))
		return nil
	}
}

// mod returns a % b
func mod(b, a interface{}) interface{} {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() % bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() % int64(bv.Uint())
		default:
			log.Fatal(fmt.Errorf("mod: unknown type for %q (%T)", bv, b))
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) % bv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() % bv.Uint()
		default:
			log.Fatal(fmt.Errorf("mod: unknown type for %q (%T)", bv, b))
			return nil
		}
	default:
		log.Fatal(fmt.Errorf("mod: unknown type for %q (%T)", av, a))
		return nil
	}
}

func percent(a interface{}) float64 {
	av := reflect.ValueOf(a)
	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(av.Int()) / 100
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(av.Uint()) / 100
	case reflect.Float32, reflect.Float64:
		return av.Float() / 100
	case reflect.String:
		ca, _ := strconv.ParseFloat(av.String(), 10)
		return ca / 100
	default:
		return 0
	}
}
