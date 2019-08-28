package engine

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

func NilParserFunc(content []byte) ParserResult {
	return ParserResult{}
}
