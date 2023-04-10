package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	// 开启跨域
	r.Use(Cors())

	r.GET(`/index`, controller.Index)
	r.GET(`/navCategory`, controller.CategoriesInfo)
	r.GET(`/filmDetail`, controller.FilmDetail)
	r.GET(`/filmPlayInfo`, controller.FilmPlayInfo)
	r.GET(`/searchFilm`, controller.SearchFilm)
	r.GET(`/filmCategory`, controller.FilmCategory)

	// 触发spider
	spiderRoute := r.Group(`/spider`)
	{
		// 清空全部数据并从零开始获取数据
		spiderRoute.GET("/SpiderRe", controller.SpiderRe)
		// 获取影片详情, 用于网路不稳定导致的影片数据缺失
		spiderRoute.GET(`/FixFilmDetail`, controller.FixFilmDetail)
		spiderRoute.GET(`/RefreshSitePlay`, controller.RefreshSitePlay)
	}

	return r
}

// Cors 开启跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session, Content-Type")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
