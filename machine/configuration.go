package machine

type Configuration struct {
	Operations        map[string][]Operation
	NextConfiguration byte
}
