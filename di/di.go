package di

import (
	"fmt"
	"go.uber.org/fx"
)

func GetAllOptions() []fx.Option {
	fxOptions := AppProviders()

	fmt.Println("fxOptions_ASD", fxOptions)

	return fxOptions
}
