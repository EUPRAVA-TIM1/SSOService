package startup

import (
	"EuprvaSsoService/config"
	"EuprvaSsoService/handlers"
	"EuprvaSsoService/repo"
	"EuprvaSsoService/service"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
func (server Server) setup() handlers.SSOHandler {

	redisClient := server.intiRedisClient()
	repository := repo.NewSecretRepoRedis(redisClient)
	gradjaninRepo := repo.NewGrdjaninRepoSql(server.config.MysqlPort, server.config.MySqlRootPass, server.config.MySqlHost)
	ssoService := service.NewSSOService(repository)
	gradjaninService := service.NewGradjaniService(gradjaninRepo)
	return handlers.NewFilesHandler(ssoService, gradjaninService)
}

func (server Server) intiRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", server.config.RedisHost, server.config.RedisPort),
	})
}

func (server Server) Start() {

	r := mux.NewRouter()

	h := server.setup()
	h.Init(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.config.Port),
		Handler: r,
	}

	wait := time.Second * 15
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	log.Printf("Listening on port = %s\n", server.config.Port)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error shutting down server %s", err)
	}
	log.Println("server gracefully stopped")

}
