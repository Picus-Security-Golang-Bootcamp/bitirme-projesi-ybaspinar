package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/basket"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/categories"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/orders"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/products"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/database"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/graceful"
	logger "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/logging"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting the application...")
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatal("Error loading config file")
	}
	logger.NewLogger(cfg)
	defer logger.Close()

	DB := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s", cfg.ServerConfig.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	rootRouter := router.Group(cfg.ServerConfig.RoutePrefix)
	basketRouter := rootRouter.Group("/basket")
	categoryRouter := rootRouter.Group("/category")
	productRouter := rootRouter.Group("/product")
	userRouter := rootRouter.Group("/user")
	orderRouter := rootRouter.Group("/order")

	basketRepo := basket.NewBasketRepo(DB)
	basketRepo.Migrate()
	basket.NewBasketHandler(basketRouter, basketRepo)

	categoryRepo := categories.NewCategoryRepo(DB)
	categoryRepo.Migrate()
	categories.NewCategoriesHandler(categoryRouter, categoryRepo)

	productRepo := products.NewProductRepo(DB)
	productRepo.Migrate()
	products.NewProductHandler(productRouter, productRepo)

	userRepo := user.NewUserRepo(DB)
	userRepo.Migrate()
	user.NewUserHandler(userRouter, userRepo)

	orderRepo := orders.NewOrdersRepo(DB)
	orderRepo.Migrate()
	orders.NewOrdersHandler(orderRouter, orderRepo)
	srv.ListenAndServe()

	log.Println("Application started successfully")

	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))
}
