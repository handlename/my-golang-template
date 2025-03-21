package command

type Root struct {
	LogLevel string  `help:"Set log level (trace|debug|info|warn|error|panic)" default:"info"`
	Version  Version `cmd:"" help:"Show version"`
}
