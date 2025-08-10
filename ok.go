package okgo

import (
	"okgo/service"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logo = `
   ____  __   ______
  / __ \/ /__/ ____/___
 / / / / //_/ / __/ __ \
/ /_/ / ,< / /_/ / /_/ /
\____/_/|_|\____/\____/
`

var log *zap.Logger

func init() {
	// Create a production logger with timestamp
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	
	var err error
	log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func New(version string) *cobra.Command {
	return &cobra.Command{
		Use:     "okgo",
		Long:    logo,
		Version: version,
	}
}

func config[T service.Service](service T, environment string, configFolderSuffix string) {

}
