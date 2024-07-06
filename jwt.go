package controllerx

import (
	"sync"

	"github.com/shanluzhineng/configurationx"
	optCasdoor "github.com/shanluzhineng/configurationx/options/casdoor"
	"github.com/shanluzhineng/irisx/casdoor"
)

var (
	_casdoorOptions casdoor.CasdoorOptions
	_cm             *casdoor.Middleware
	_sync           sync.Once
)

func GetCasdoorMiddleware() *casdoor.Middleware {
	_sync.Do(func() {
		casdoorOpt := &optCasdoor.CasdoorOptions{}
		configurationx.GetInstance().UnmarshalPropertiesTo(optCasdoor.ConfigurationKey, casdoorOpt)
		_casdoorOptions = casdoor.CasdoorOptions{
			CasdoorOptions: *casdoorOpt,
			Extractor: casdoor.FromFirst(casdoor.FromHeader("Authorization"),
				casdoor.FromAuthHeader),
		}
		_cm = casdoor.New(_casdoorOptions)
	})
	return _cm
}
