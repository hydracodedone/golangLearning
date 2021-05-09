package somepkg

type private struct {
}

func (p *private) PrivateHasPublicMethod() {

}

type Public struct {
	private
}

func (p Public) publicHasPrivateMethod() {

}
