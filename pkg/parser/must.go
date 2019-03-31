package parser

type Must struct {
	Token Token
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

func (m *Must) performMatch(i int) error {

	return nil
}

func (m *Must) performIn(i int) error {

	return nil
}

func (m *Must) performEqual(i int) error {
	return nil
}

func (m *Must) performNotEqual(i int) error {
	return nil
}

func (m *Must) performEndWith(i int) error {
	return nil
}

func (m *Must) performStartWith(i int) error {
	return nil
}

func (m *Must) performOperator(i int) error {
	return nil
}

func (m *Must) performGT(i int) error {
	return nil
}

func (m *Must) performLT(i int) error {
	return nil
}

func (m *Must) performGTE(i int) error {
	return nil
}

func (m *Must) performLTE(i int) error {
	return nil
}
