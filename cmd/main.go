package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/swastiijain24/psp/internals/handlers"
	"github.com/swastiijain24/psp/internals/repositories"
	"github.com/swastiijain24/psp/internals/kafka"
	"github.com/swastiijain24/psp/internals/redis"
	"github.com/swastiijain24/psp/internals/routes"
	"github.com/swastiijain24/psp/internals/services"
	"github.com/swastiijain24/psp/internals/workers"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	err := godotenv.Load()
	if err != nil {
		log.Print("no .env file found")
	}

	dsn := os.Getenv("GOOSE_DBSTRING")
	log.Print(dsn)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	redisAddr := os.Getenv("REDIS_ADDR")
	kafkaAddr := os.Getenv("KAFKA_ADDR")
	port := os.Getenv("PORT")

	redisStore := repository.NewRedisStore(redisAddr)
	kafkaProducer := kafka.NewProducer(kafkaAddr)

	repo := repo.New(pool)
	vpaSvc := services.NewVpaService(repo)

	paymentSvc := services.NewPaymentService(vpaSvc, kafkaProducer, redisStore)
	paymentHandler := handlers.NewPaymentHandler(paymentSvc)

	consumer := kafka.NewConsumer([]string{kafkaAddr}, "payment.response.v1", "psp-grp")
	defer consumer.Reader.Close()

	worker := workers.NewResponseWorker(consumer, paymentSvc)

	go worker.StartConsumingResponse(ctx)
	r := gin.Default()
	routes.RegisterRoutes(r, paymentHandler)
	log.Print("initialized all")


	//since the .Run runs infinitley and if we press ctlr c the consumer goes into infinite loop and the main never ends giving infinite errors 
	// log.Println("PSP API starting on :" + port)
	// if err := r.Run(":" + port); err != nil {
	// 	log.Fatal(err)
	// } 

	srv := &http.Server{
		Addr: ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s\n", err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutdown signal received...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}

}
