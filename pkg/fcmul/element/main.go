package element

type Element interface {
	elementTag()
}

type Bool bool
func (Bool)elementTag() {}

type String string
func (String)elementTag() {}

type Int int64
func (Int)elementTag() {}

type Map map[Element] Element
func (Map)elementTag() {}

type List []Element
func (List)elementTag() {}
