package fileparser

type VaneJSON struct {
	Directory     string            `json:"directory"`
	Timeout       int               `json:"timeout"`
	Dependencies  map[string]string `json:"dependencies"`
	DevDependcies map[string]string `json:"devDependencies"`
}

func ReadJSONFile() {

}

type VaneJsoner interface {
	JsonReader
	JsonWriter
}

type JsonReader interface {
	ReadJson()
}

type JsonWriter interface {
	WriteJson()
}
