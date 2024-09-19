package decoder

type Format int

const (
	Json Format = iota
	Yaml
)

type Decorder interface {
	Decode(format Format) []byte
}
