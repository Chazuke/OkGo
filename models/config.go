package models

type (
	OkGoConfig struct {
		OkGo struct {
			ProjectId    string
			OTLPExporter *OTLPExporter
			StdOut       *StdOut
		}
	}

	OTLPExporter struct {
		URI        string
		SampleRate float64
	}

	StdOut struct {
		SampleRate float64
	}
)
