// Code generated from Pkl module `appconfig`. DO NOT EDIT.
package appconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Appconfig struct {
	Server *Server `pkl:"server"`

	Db *Database `pkl:"db"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Appconfig
func LoadFromPath(ctx context.Context, path string) (ret *Appconfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Appconfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Appconfig, error) {
	var ret Appconfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
