package middleware

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/huchengbei/for-my-girl/pkg/logging"
	_ "github.com/huchengbei/for-my-girl/statik"
	"github.com/rakyll/statik/fs"
)

type ginFS struct {
	FS http.FileSystem
}

var staticFS static.ServeFileSystem

// open 打开文件
func (b *ginFS) Open(name string) (http.File, error) {
	return b.FS.Open(name)
}

// exists 文件是否存在
func (b *ginFS) Exists(prefix string, filepath string) bool {

	if _, err := b.FS.Open(filepath); err != nil {
		return false
	}
	return true
}

// FrontendFileHandler 前端静态文件处理
func FrontendFileHandler() gin.HandlerFunc {
	var err error
	staticFS = &ginFS{}
	staticFS.(*ginFS).FS, err = fs.New()
	if err != nil {
		logging.Error("无法初始化静态资源, %s", err)
	}
	//staticFS, err := fs.New()
	if err != nil {
		logging.Error("无法初始化静态资源, %s", err)
	}

	ignoreFunc := func(c *gin.Context) {
		c.Next()
	}

	// 读取index.html
	file, err := staticFS.Open("/index.html")
	if err != nil {
		logging.Error("静态文件[index.html]不存在，可能会影响首页展示")
		return ignoreFunc
	}

	fileContentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logging.Error("静态文件[index.html]读取失败，可能会影响首页展示")
		return ignoreFunc
	}
	fileContent := string(fileContentBytes)

	fileServer := http.FileServer(staticFS)
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// API 跳过
		if strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/auth") {
			c.Next()
			return
		}

		// 不存在的路径和index.html均返回index.html
		if (path == "/index.html") || (path == "/") || !staticFS.Exists("/", path) {
			// 读取、替换站点设置

			c.Header("Content-Type", "text/html")
			c.String(200, fileContent)
			c.Abort()
			return
		}

		// 存在的静态文件
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
