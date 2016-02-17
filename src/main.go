package main

import (
	"fmt"
	. "github.com/Dataman-Cloud/zookeeper-helper/src/config"
	"github.com/Dataman-Cloud/zookeeper-helper/src/logger"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/samuel/go-zookeeper/zk"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

func init() {
	conf := Pairs()
	logger.LoadLogConfig()

	numCPU := conf.NumCPU
	runtime.GOMAXPROCS(numCPU)
	log.Info("Runing with ", numCPU, " CPUs")
}

func main() {
	startServer()
}

func startServer() {
	log.Info("zookeeper helper server running...")
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.String(200, "pass")
	})
	statGroup := router.Group("/api/v1/zookeeper")
	{
		statGroup.GET("/stats", ServerStat)
	}

	host := os.Getenv("ZOOKEEPER_HELPER_HOST")
	port := os.Getenv("ZOOKEEPER_HELPER_PORT")
	zkPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		log.Warn(err)
	}
	if host == "" {
		host = "localhost"
	}
	if zkPort == 0 {
		zkPort = 5096
	}
	fmt.Println(host)
	fmt.Println(zkPort)
	addr := fmt.Sprintf("%s:%d", host, zkPort)
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Error("can't start server: ", err)
		panic(-1)
	}
}

func ServerStat(ctx *gin.Context) {
	Server := os.Getenv("ZOOKEEPER_HELPER_HOST")
	err := getServerStat([]string{Server}, Timeout)
	if err != nil {
		log.Error("getServerStat error: ", err)
		ReturnError(ctx, err)
		return
	}
	ReturnOK(ctx, "zookeeper running")
}

func ReturnError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{"code": CodeError, "data": "", "errors": err.Error()})
}

func ReturnOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": CodeOK, "data": data, "errors": ""})
}

func getServerStat(servers []string, timeout time.Duration) error {
	serverStats, ok := zk.FLWSrvr(servers, timeout)
	if !ok {
		err := serverStats[0].Error
		return err
	}
	return nil
}
