package okgo

import (
	"okgo/service"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const logo = `
   ____  __   ______    
  / __ \/ /__/ ____/___ 
 / / / / //_/ / __/ __ \
/ /_/ / ,< / /_/ / /_/ /
\____/_/|_|\____/\____/ 
`

var log zerolog.Logger

func init() {
	log = zerolog.New(os.Stdout).With().Timestamp().Logger()
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
