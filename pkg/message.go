package pkg

type Message struct {
	Value   string
	Topic   []string
	Headers []Header
}

type Header struct {
	Key   string
	Value string
}
