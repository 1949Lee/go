package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const pathLabel = "path"

func main() {
	// 初始化gin框架
	r := gin.Default()

	// 初始化一个request counter计数器，用来记录每个http请求的数量
	reqCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		// 这些参数都会以文本的形式显示在/metrics页面中
		Namespace: "imooc_com_ccmouse_learngo",
		Subsystem: "ginserver",
		Name:      "request_count",
		Help:      "request counter",
	}, []string{pathLabel}) // 加一个label以便根据请求的path进行汇总
	prometheus.MustRegister(reqCounter)

	// 通过Use注册一个叫做middleware的函数。许多go的框架使用这种方法
	// 来实现类似Java中AOP的能力。每一个http request都会运行一遍
	// 这个middleware函数。我们将在这里完成请求计数。
	r.Use(func(c *gin.Context) {
		// 为请求计数，并且标记Path以便分类汇总
		reqCounter.WithLabelValues(c.Request.URL.Path).Inc()
		c.Next()
	})

	// Handler必须在r.Use之后注册
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"token": rand.Int(),
		})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 服务器会开在localhost:8080
	// 请尝试localhost:8080/hello以及localhost:8080/metrics
	r.Run()
}
