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

func (m *Must) Run(state State) (err error) {
	// eval the line,
	// fatal if isMust()
	// just log if isHould()

	if err == nil {
		if m.IsMust() || m.IsMustNot() {
			state.SuccessMust ++
		} else if m.IsShould() || m.IsShouldNot() {
			state.SuccessShould++
		}
	} else {
		if m.IsShould() || m.IsShouldNot() {
			state.FailedShould ++
		} else {
			return err
		}
	}

	return err
}
