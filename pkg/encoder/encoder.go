package encoder

import (
	"encoding/json"
	"errors"

	"github.com/naonao2323/ecs-task-def/pkg/plovider/ecs"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Format int

const (
	Json Format = iota
	Yaml
	Unknow
)

func GetFormat(ext string) Format {
	switch ext {
	case ".json":
		return Json
	case ".yaml":
		return Yaml
	case ".yml":
		return Yaml
	default:
		return Unknow
	}
}

type (
	Encoder[P ecs.EcsTarget] interface {
		Encode(in []byte, format Format) (*P, error)
	}
	EncoderImpl[P ecs.EcsTarget] struct {
		logger *zap.Logger
	}
)

func NewEncoder[P ecs.EcsTarget](logger *zap.Logger) Encoder[P] {
	return EncoderImpl[P]{logger}
}

func (e EncoderImpl[P]) Encode(in []byte, format Format) (*P, error) {
	switch format {
	case Json:
		def, err := e.EncodeJson(in)
		if err != nil {
			e.logger.Error("fail to encode json file", zap.Error(err))
			return nil, errors.New("fail to encode json file")
		}
		return def, nil
	case Yaml:
		def, err := e.EncodeYaml(in)
		if err != nil {
			e.logger.Error("fail to encode yaml file", zap.Error(err))
			return nil, errors.New("fail to encode yaml file")
		}
		return def, nil
	default:
		return nil, errors.New("unknown extension")
	}
}

func (e EncoderImpl[P]) EncodeJson(in []byte) (*P, error) {
	var def P
	err := json.Unmarshal(in, &def)
	if err != nil {
		return nil, err
	}
	return &def, nil
}

func (e EncoderImpl[P]) EncodeYaml(in []byte) (*P, error) {
	var def P
	err := yaml.Unmarshal(in, &def)
	if err != nil {
		return nil, err
	}
	return &def, nil
}
