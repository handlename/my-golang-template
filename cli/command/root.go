package command

type Root struct {
	LogLevel string  `cmd:"" help:"Set log level (trace|debug|info|warn|error|panic). Default is info."`
	Version  Version `cmd:"" help:"Show version"`
}
