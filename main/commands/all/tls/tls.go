package tls

import (
	"github.com/GFW-knocker/Xray-core/main/commands/base"
)

// CmdTLS holds all tls sub commands
var CmdTLS = &base.Command{
	UsageLine: "{{.Exec}} tls",
	Short:     "TLS tools",
	Long: `{{.Exec}} {{.LongName}} provides tools for TLS.
`,
	Commands: []*base.Command{
		cmdCert,
		cmdPing,
		cmdCertChainHash,
	},
}
