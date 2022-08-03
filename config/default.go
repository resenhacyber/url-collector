package config

//DefaultConf 默认配置
var DefaultConf AppConfig = AppConfig{
	RoutineCount: 5,
	SearchEngine: "google-image",
	Format:       "url",
	BaseURL: map[string]string{
		"google":       "https://www.google.com.hk/search?q=$keyword",
		"google-image": "https://g.luciaz.me/search?q=$keyword",
		"bing":         "https://cn.bing.com/search?q=$keyword",
		"baidu":        "https://www.baidu.com/s?wd=$keyword",
	},
	BlackList: []string{
		"gov",
		"baidu.com",
		"csdn.net",
		"voidcc.com",
		"blog.sina.com.cn",
		"segmentfault.com",
		"youku.com",
		"wenwen.sogou.com",
		"hao123.com",
		"cache.baiducontent.com",
		"g3.luciaz.me",
		"www.youtube.com",
		"gitee.com",
		"github.com",
		"stackoverflow.com",
		"developer.aliyun.com",
		"cloud.tencent.com",
		"www.zhihu.com/question",
		"blog.51cto.com",
		"zhidao.baidu.com",
		"www.cnblogs.com",
		"coding.m.imooc.com",
		"weibo.cn",
		"www.taobao.com",
		"www.google.com",
		"go.microsoft.com",
		"facebook.com",
		"blog.csdn.net",
		"books.google.com",
		"policies.google.com",
		"webcache.googleusercontent.com",
		"translate.google.com",
		"support.google.com",
	},
}
