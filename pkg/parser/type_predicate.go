package parser

import "fmt"

type PredType int

const (
	PredTypeAnd PredType = iota
	PredTypeOr
	PredTypeXor
	PredTypeNot
	PredTypeMatch
	PredTypeStartsWith
	PredTypeEndsWith
	PredTypeStringEqual
	PredTypeNumberEqual
	PredTypeGreaterThen
	PredTypeLessThen
	PredTypeEqualGreaterThen
	PredTypeEqualLessThen
	PredTypeNilCheck
	PredTypeTrueCheck
	PredTypeFalseCheck
)

func (p PredType) String() string {
	switch p {
	case PredTypeAnd:
		return "And"
	case PredTypeOr:
		return "Or"
	case PredTypeXor:
		return "Exclusive Or"
	case PredTypeNot:
		return "Not"
	case PredTypeMatch:
		return "Match"
	case PredTypeStartsWith:
		return "Starts With"
	case PredTypeEndsWith:
		return "Ends With"
	case PredTypeStringEqual:
		return "String Equal"
	case PredTypeNumberEqual:
		return "Number Equal"
	case PredTypeGreaterThen:
		return "Greater Then"
	case PredTypeLessThen:
		return "Less Then"
	case PredTypeEqualGreaterThen:
		return "Greater Then Equal"
	case PredTypeEqualLessThen:
		return "Less Then Equal"
	case PredTypeNilCheck:
		return "Null Check"
	case PredTypeTrueCheck:
		return "True Check"
	case PredTypeFalseCheck:
		return "False Check"
	default:
		panic("Impossible predicate type")
	}
}

// A basic interface for all predicate types, with typing and evaluation methods
type Evaluable interface {
	Type() PredType
	Eval() bool
	// Produces a string which completes the sentence "Expected ..."
	ErrorString() string
}

// A simple binary logic op, which contains the type (And, Or, Xor) and the left and right sides
type PredLogicOp struct {
	t           PredType
	left, right Evaluable
}

func (p PredLogicOp) Type() PredType {
	p.enforceTypes()
	return p.t
}
func (p PredLogicOp) Eval() bool {
	p.enforceTypes()
	switch p.t {
	case PredTypeAnd:
		return p.left.Eval() && p.right.Eval()
	case PredTypeOr:
		return p.left.Eval() || p.right.Eval()
	case PredTypeXor:
		// This is a slightly tricky simplification of XOR, which should be (left || right) && !(left && right)
		// Because we're dealing with boolean types, this allows us to only eval each side once and requires no
		// temporary variables
		return p.left.Eval() != p.right.Eval()
	default:
		panic(fmt.Sprintf("Invalid contained predicate type: %s", p.t.String()))
	}
}
func (p PredLogicOp) ErrorString() string {
	p.enforceTypes()
	f := "the %s of `%s` %s `%s`"
	switch p.t {
	case PredTypeAnd:
		return fmt.Sprintf(f, "conjunction", p.left.ErrorString(), "AND", p.right.ErrorString())
	case PredTypeOr:
		return fmt.Sprintf(f, "disjunction", p.left.ErrorString(), "OR", p.right.ErrorString())
	case PredTypeXor:
		return fmt.Sprintf(f,"exclusive disjunction", p.left.ErrorString(), "XOR", p.right.ErrorString())
	default:
		panic(fmt.Sprintf("Invalid contained predicate type: %s", p.t.String()))
	}
}
func (p PredLogicOp) enforceTypes() {
	if p.t != PredTypeAnd && p.t != PredTypeOr && p.t != PredTypeXor {
		panic(fmt.Sprintf("Predicate type %s is not a valid type for a logic operation", p.t.String()))
	}
}

// A simple unary NOT operation. Although we don't explicitly support an '!' style not operator (that is to say,
// we cannot just drop a '!' anywhere in the grammar), it simplifies greatly reasoning about the parsed syntax
// to emit the normal type wrapped in a NOT for these cases
type PredNotOp struct {
	v Evaluable
}
func (p PredNotOp) Type() PredType {
	return PredTypeNot
}
func (p PredNotOp) Eval() bool {
	return !p.v.Eval()
}
func (p PredNotOp) ErrorString() string {
	return fmt.Sprintf("the opposite of %s", p.v.ErrorString())
}

// A simple binary operation which handles all string type operations (match, starts with, ends with, equal)
// Precondition: v and r are ValTypeString or ValTypeTemplate
type PredStringOp struct {
	t PredType
	// The value upon which to perform the string op (left hand side)
	v Valueable
	// The reference to use for the string op (right hand side)
	r Valueable
}

func (p PredStringOp) Type() PredType {
	p.enforceTypes()
	return p.t
}
func (p PredStringOp) Eval() bool {
	p.enforceTypes()
	// TODO: Complete the implementation of this for the regex cases (can starts with and ends with contain regex?)
	// rgx := regex.Compile(p.r)
	// etc.
	switch p.t {
	case PredTypeMatch:
		fallthrough
	case PredTypeStartsWith:
		fallthrough
	case PredTypeEndsWith:
		return true
	case PredTypeStringEqual:
		return p.vStr() == p.rStr()
	default:
		panic(fmt.Sprintf("Invalid contained predicate type: %s", p.t.String()))
	}
}
func (p PredStringOp) ErrorString() string {
	p.enforceTypes()
	switch p.t {
	case PredTypeMatch:
		return fmt.Sprintf("%s to match %s", p.vStr(), p.rStr())
	case PredTypeStartsWith:
		return fmt.Sprintf("%s to start with %s", p.vStr(), p.rStr())
	case PredTypeEndsWith:
		return fmt.Sprintf("%s to end with %s", p.vStr(), p.rStr())
	case PredTypeStringEqual:
		return fmt.Sprintf("%s to equal %s", p.vStr(), p.rStr())
	default:
		panic(fmt.Sprintf("Invalid contained predicate type: %s", p.t.String()))
	}
}
func (p PredStringOp) vStr() string {
	return p.str(p.v)
}
func (p PredStringOp) rStr() string {
	return p.str(p.r)
}
func (p PredStringOp) str(i Valueable) string {
	var s string
	if p.v.Type() == ValTypeString {
		s = i.GetVal().(string)
	} else {
		// TODO: Force evaluation
		s = i.GetVal().(string)
	}
	return s
}
func (p PredStringOp) enforceTypes() {
	if p.t != PredTypeMatch && p.t != PredTypeStartsWith && p.t != PredTypeEndsWith && p.t != PredTypeStringEqual {
		panic(fmt.Sprintf("Predicate type %s is not a valid type for a string operation", p.t.String()))
	}
	if p.v.Type() != ValTypeString && p.v.Type() != ValTypeTemplate {
		panic(fmt.Sprintf("Value type %s is not valid for a string %s op", p.v.Type().String(), p.t.String()))
	}
	if p.r.Type() != ValTypeString && p.r.Type() != ValTypeTemplate {
		panic(fmt.Sprintf("Match type %s is not valid for a string %s op", p.r.Type().String(), p.t.String()))
	}
}

// A binary operation on two numeric types. Handles all of the ==, >, <. >=, <= cases
// Precondition: left and right are ValTypeInt, ValTypeFloat or ValTypeTemplate(single evaluation only)
type PredNumericOp struct {
	t           PredType
	left, right Valueable
}
func (p PredNumericOp) Type() PredType {
	p.enforceTypes()
	return p.t
}
func (p PredNumericOp) Eval() bool {
	p.enforceTypes()
	// TODO: Implement this, probably need to add a helper like pureInt for comparisons that can be done in a pure
	// integer space, and otherwise cast down to float
	return true
}
func (p PredNumericOp) ErrorString() string {
	p.enforceTypes()
	// TODO: Implement this, see notes on Eval()
	// Very much a placeholder error string
	return "some number to have a relation to another number"
}
func (p PredNumericOp) enforceTypes() {
	if p.t != PredTypeNumberEqual && p.t != PredTypeGreaterThen && p.t != PredTypeLessThen {
		panic(fmt.Sprintf("Predicate type %s is not a valid type for a string operation", p.t.String()))
	}
	if p.left.Type() != ValTypeInt && p.left.Type() != ValTypeFloat && p.left.Type() != ValTypeTemplate {
		panic(fmt.Sprintf("Value type %s is not valid for a numeric %s op", p.left.Type().String(), p.t.String()))
	}
	if p.right.Type() != ValTypeInt && p.right.Type() != ValTypeFloat && p.right.Type() != ValTypeTemplate {
		panic(fmt.Sprintf("Match type %s is not valid for a numeric %s op", p.right.Type().String(), p.t.String()))
	}
}

// A simple unary op which determines if its value evaluates to null, true or false
// Precondition: v is a ValTypeTemplate(single evaluation only)
type PredBoolCheckOp struct {
	t PredType
	v	Valueable
}
func (p PredBoolCheckOp) Type() PredType {
	p.enforceTypes()
	return p.t
}
func (p PredBoolCheckOp) Eval() bool {
	p.enforceTypes()
	if p.t == PredTypeNilCheck {
		return p.v.GetVal() == nil
	}
	b, ok := p.v.GetVal().(bool)
	if !ok {
		panic("tried to corece a bool from a non-bool value")
	}
	if p.t == PredTypeTrueCheck {
		return b == true
	} else { // PredTypeFalseCheck
		return b == false
	}
}
func (p PredBoolCheckOp) ErrorString() string {
	p.enforceTypes()
	switch p.t {
	case PredTypeNilCheck:
		if p.v.GetVal() == nil {
			return "nil to be nil"
		}
		return fmt.Sprintf("the value %v to be nil", p.v.GetVal())
	case PredTypeTrueCheck:
		// TODO: Better message formatting (including type check) for these
		return fmt.Sprintf("the value of %v to be true", p.v.GetVal())
	case PredTypeFalseCheck:
		return fmt.Sprintf("the value of %v to be false", p.v.GetVal())
	default:
		panic(fmt.Sprintf("impossible predicate type %s", p.t.String()))
	}

}
func (p PredBoolCheckOp) enforceTypes() {
	if p.t != PredTypeNilCheck && p.t != PredTypeTrueCheck && p.t != PredTypeFalseCheck {
		panic(fmt.Sprintf("Predicate type %s is not a valid type for a bool operation", p.t.String()))
	}
}