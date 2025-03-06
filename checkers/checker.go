package checkers

type Checker interface {
	Check(host, port string) bool
}

var Checkers = make(map[string]Checker)

func RegisterChecker(protocol string, checker Checker) {
	Checkers[protocol] = checker
}
