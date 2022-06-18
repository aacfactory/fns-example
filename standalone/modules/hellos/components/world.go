package components

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/logs"
)

const (
	worldName = "world"
)

func GetWorld(ctx context.Context) (w World) {
	v, has := service.GetComponent(ctx, worldName)
	if !has {
		panic(fmt.Sprintf("%+v", errors.Warning("there is no component in context").WithMeta("component", worldName)))
		return
	}
	w, has = v.(World)
	if !has {
		panic(fmt.Sprintf("%+v", errors.Warning("invalid component type").WithMeta("component", worldName)))
		return
	}
	return
}

// NewWorld
// component world
func NewWorld() service.Component {
	return &world{
		log:   nil,
		value: "",
	}
}

type World interface {
	Value() string
}

type world struct {
	log   logs.Logger
	value string
}

func (w *world) Name() (name string) {
	name = worldName
	return
}

func (w *world) Build(options service.ComponentOptions) (err error) {
	w.log = options.Log
	_, err = options.Config.Get("value", &w.value)
	if err != nil {
		if w.log.ErrorEnabled() {
			w.log.Error().Cause(err).Caller().Message("world component build failed")
		}
		err = errors.Warning("world component build failed").WithCause(err)
		return
	}
	return
}

func (w *world) Close() {

}

func (w *world) Value() string {
	return w.value
}
