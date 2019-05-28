package server

import (
	sve "{{ .pkgName}}/server"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:  "server",
	RunE: run,
}

var (
	mysqlAddr  string
	enableMock bool
	addr       string
)

func init() {
	ServerCmd.Flags().StringVar(&addr, "addr", ":6000", "web server running address")
	ServerCmd.Flags().BoolVar(&enableMock, "mock", false, "enable mock database engine")
	ServerCmd.Flags().StringVar(&mysqlAddr, "dsn", "", "mysql address")
}

func run(cmd *cobra.Command, args []string) error {
	var opts = make([]sve.Option, 0)

	if mysqlAddr != "" {
		opts = append(opts, sve.WithMysqlDSN(mysqlAddr))
	}

	if enableMock {
		opts = append(opts, sve.EnableMockStorage())
	}

	if addr == "" {
		panic(`addr can not be empty`)
	}

	opts = append(opts, sve.WithAddr(addr))

	server := sve.NewServer(opts...)

	return server.Run()
}
