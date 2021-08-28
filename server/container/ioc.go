package container

import "github.com/sarulabs/di/v2"

var container di.Container

func Init(instances ...di.Def) error {
	builder, err := di.NewBuilder()
	if err != nil {
		return err
	}
	builder.Add(instances...)

	// build container
	container = builder.Build()
	return nil
}

func GetInstance() di.Container {
	return container
}
