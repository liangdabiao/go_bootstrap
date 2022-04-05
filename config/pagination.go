package config

import "gohub/pkg/config"

func init() {
	config.Add("pagination", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认每页条数
			"perpage": 10,

			// URL 中用以分辨多少页的参数
			"url_query": "page",
		}
	})
}
