package config

import "gohub/pkg/config"

func init() {
	config.Add("session", func() map[string]interface{} {
		return map[string]interface{}{
			// 目前只支持 Cookie
			"default": config.Env("SESSION_DRIVER", "cookie"),

			// 会话的 Cookie 名称
			"session_name": config.Env("SESSION_NAME", "goblog-session"),
		}
	})
}
