package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type Must struct {
	Token Token
	State *State
}

func (m *Must) IsMust() bool {
	return m.Token.Tree[0].Text == "MUST"
}

func (m *Must) IsMustNot() bool {
	return m.Token.Tree[0].Text == "MUSTNOT"
}

func (m *Must) IsShould() bool {
	return m.Token.Tree[0].Text == "SHOULD"
}

func (m *Must) IsShouldNot() bool {
	return m.Token.Tree[0].Text == "SHOULDNOT"
}

func (m *Must) Run(state *State) (err error) {
	m.State = state

	for i := 0; i < len(m.Token.Tree); i++ {
		token := m.Token.Tree[i]

		if token.Type == KEYWORD {

			// MATCH
			switch token.Text {
			case "MATCH":
				err = m.performMatch(i)
				break

			case "IN":
				err = m.performIn(i)
				break

			case "EQUAL", "EQUALS":
				err = m.performEqual(i)
				break

			case "ENDWITH":
				err = m.performEndWith(i)
				break

			case "STARTWITH":
				err = m.performStartWith(i)
				break
			}

		} else if token.Type == OPERATOR {

			switch token.Text {
			case "<":
				err = m.performLT(i)
				break
			case ">":
				err = m.performGT(i)
				break
			case ">=":
				err = m.performGTE(i)
				break
			case "<=":
				err = m.performLTE(i)
				break
			case "=", "==":
				err = m.performEqual(i)
				break
			case "!=", "!==":
				err = m.performNotEqual(i)
				break
			}
		}
	}

	// eval the line,
	// fatal if isMust()
	// just log if isHould()

	if err == nil {
		if m.IsMust() || m.IsMustNot() {
			state.SuccessMust++
		} else if m.IsShould() || m.IsShouldNot() {
			state.SuccessShould++
		}
	} else {
		if m.IsShould() || m.IsShouldNot() {
			state.FailedShould++
		} else {
			return err
		}
	}

	return err
}

func (m *Must) performMatch(i int) (err error) {
	leftToken := m.Token.Tree[i-1]
	rightToken := m.Token.Tree[i+1]
	_, err = regexp.MatchString(rightToken.Text, leftToken.Text)
	return m.actuallyFailed(err)
}

func (m *Must) actuallyFailed(err error) error {
	if err != nil {
		if m.IsShouldNot() || m.IsMustNot() {
			return nil
		} else {
			return err
		}
	}

	return nil
}

func (m *Must) performIn(i int) (err error) {

	return nil
}

func (m *Must) performEqual(i int) (err error) {

	leftToken := m.valueOf(m.Token.Tree[i-1])
	rightToken := m.valueOf(m.Token.Tree[i+1])

	if leftToken != rightToken {
		err = fmt.Errorf("they are not equal")
	}

	return m.actuallyFailed(err)
}

func (m *Must) performNotEqual(i int) (err error) {

	leftToken := m.valueOf(m.Token.Tree[i-1])
	rightToken := m.valueOf(m.Token.Tree[i+1])

	if leftToken == rightToken {
		err = fmt.Errorf("they are equal")
	}

	return m.actuallyFailed(err)
}

func (m *Must) performEndWith(i int) (err error) {
	leftToken := m.valueOf(m.Token.Tree[i-1])
	rightToken := m.valueOf(m.Token.Tree[i+1])

	if !strings.HasSuffix(leftToken, rightToken) {
		err = fmt.Errorf("%s doesn't end with %s", leftToken, rightToken)
	}

	return m.actuallyFailed(err)
}

func (m *Must) performStartWith(i int) (err error) {
	leftToken := m.valueOf(m.Token.Tree[i-1])
	rightToken := m.valueOf(m.Token.Tree[i+1])

	if !strings.HasPrefix(leftToken, rightToken) {
		err = fmt.Errorf("%s doesn't start with %s", leftToken, rightToken)
	}

	return m.actuallyFailed(err)
}

func (m *Must) performGT(i int) (err error) {
	leftToken := m.valueOf(m.Token.Tree[i-1])
	rightToken := m.valueOf(m.Token.Tree[i+1])

	if !(leftToken > rightToken) {
		err = fmt.Errorf("it is not gt")
	}

	return m.actuallyFailed(err)
}

func (m *Must) performLT(i int) (err error) {
	leftToken := m.valueOf(m.Token.Tree[i-1])
	rightToken := m.valueOf(m.Token.Tree[i+1])

	if !(leftToken < rightToken) {
		err = fmt.Errorf("it is not lt")
	}

	return m.actuallyFailed(err)
}

func (m *Must) performGTE(i int) (err error) {
	leftToken := m.Token.Tree[i-1]
	rightToken := m.Token.Tree[i+1]

	if !(leftToken.Text >= rightToken.Text) {
		err = fmt.Errorf("it is not gte")
	}

	return m.actuallyFailed(err)
}

func (m *Must) performLTE(i int) (err error) {
	leftToken := m.Token.Tree[i-1]
	rightToken := m.Token.Tree[i+1]

	if !(leftToken.Text <= rightToken.Text) {
		err = fmt.Errorf("it is not lte")
	}

	return m.actuallyFailed(err)
}

func (m *Must) valueOf(token Token) string {
	return valueOf(token, m.State)
}
