package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vishnusomank/sbom-poc/src/controller"
	"github.com/vishnusomank/sbom-poc/utils/constants"
)

// Create once per go file and re use log
var configFilePath *string

func main() {

	configFilePath = flag.String("config-path", "conf/", "conf/")
	flag.Parse()

	loadConfig()
	gin.SetMode(gin.ReleaseMode)
	models.ConnectDatabase()
	r := gin.New()
	//Allow CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	setupRoutes(r)
	startServer(r)
}
func loadConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("No config file found at %s\n", *configFilePath)
		} else {
			log.Fatal("Error reading config file: %s\n", readErr)
		}
	}
}

// startServer - Start server
func startServer(r *gin.Engine) {
	srv := &http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...\n")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: %s\n", err)
	}

	log.Println("Server exiting\n")

}

//setupRoutes -
func setupRoutes(r *gin.Engine) {

	application := r.Group(viper.GetString("server.basepath"))
	{
		v1 := application.Group("/api/v1")
		{
			scanimage := v1.Group("/scan-image")
			{
				scanimage.POST(constants.ADD_IMAGE, controller.AddImage)
			}
			outputscan := v1.Group("/show-output")
			{
				outputscan.POST(constants.GET_OP, controller.ShowData)
			}
		}
	}
}
