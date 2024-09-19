package encoder

type Format int

const (
	Json Format = iota
	Yaml
)

type TaskEncoder interface {
	Encode(in []byte, format Format)
}
