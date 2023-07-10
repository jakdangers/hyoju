package errors

type UserName string
type Op string
type Kind uint8

type Error struct {
	User UserName
	Op   Op
	Kind Kind
	Err  error
	// Stack information; used only when the 'debug' build tag is set.
	stack
}
