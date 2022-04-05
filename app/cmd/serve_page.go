package cmd

import (
	"embed"
	"gohub/app/http/middlewares"
	"gohub/bootstrap"
	c "gohub/pkg/config"
	"gohub/pkg/logger"
	"net/http"

	"github.com/spf13/cobra"
)

//go:embed public/*
var staticFS embed.FS

// CmdServe represents the available web sub-command.
var CmdPageServe = &cobra.Command{
	Use:   "serve_page",
	Short: "Start web page server",
	Run:   runPageWeb,
	Args:  cobra.NoArgs,
}

func runPageWeb(cmd *cobra.Command, args []string) {

	// 初始化路由绑定
	router := bootstrap.SetupWebRoute(staticFS)

	err := http.ListenAndServe(":"+c.GetString("app.port2"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
