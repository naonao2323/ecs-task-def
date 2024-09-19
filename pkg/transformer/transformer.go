package transformer

type Transformer interface {
	Transform(tag string, appName string)
}
