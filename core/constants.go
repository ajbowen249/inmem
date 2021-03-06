package inmem

const errUnknownCommand = "(error) unknown Command '%s'"
const errArgumentCount = "(error) ERR wrong number of arguments for '%s' command"
const errWrongType = "(error) WRONGTYPE Operation against a key holding the wrong kind of value"
const errNotIntOrOutOfRange = "(error) ERR value is not an integer or out of range"

const resOK = "OK"
const resNil = "(nil)"
const resInt = "(integer) %v"

const cmdSet = "set"
const cmdGet = "get"
const cmdAppend = "append"
const cmdIncr = "incr"

var argCounts = map[string]int{
	cmdSet: 2,
	cmdGet: 1,
	cmdAppend: 2,
	cmdIncr: 1,
}

type keyType int
const (
	tString keyType = iota
	tHash
	tList
	tSet
)
