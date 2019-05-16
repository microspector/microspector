package parser

type ValType int
const (
	ValTypeString ValType = iota
	ValTypeFloat
	ValTypeInt
	ValTypeBool
	ValTypeNil
)
func(v ValType) String() string {
	switch v {
	case ValTypeString:
		return "String"
	case ValTypeFloat:
		return "Float"
	case ValTypeInt:
		return "Int"
	case ValTypeBool:
		return "Bool"
	case ValTypeNil:
		return "Nil"
	default:
		panic("impossible value type")
	}
}

type Valueable interface {
	Type() ValType
	Equals(v Valueable) bool
	GetVal() interface{}
}

type Val struct {
	typ ValType
	str string
	flt float64
	itg int64
	bl bool
}
func (v Val) Type() ValType {
	return v.typ
}
func (v Val) Equals(o Valueable) bool {
	if v.typ != o.Type() {
		return false
	}
	switch v.typ {
	case ValTypeString:
		return v.str == o.GetVal().(string)
	case ValTypeFloat:
		return v.flt == o.GetVal().(float64)
	case ValTypeInt:
		return v.itg == o.GetVal().(int64)
	default:
		panic("impossible value type")
	}
}
func (v Val) GetVal() interface{} {
	switch v.typ {
	case ValTypeString:
		return v.str
	case ValTypeFloat:
		return v.flt
	case ValTypeInt:
		return v.itg
	case ValTypeBool:
		return v.bl
	case ValTypeNil:
		return nil
	default:
		panic("impossible value type")
	}
}
