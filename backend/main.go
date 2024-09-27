package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"ppo/internal/app"
	"ppo/internal/config"
	loggerPackage "ppo/pkg/logger"
	"ppo/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var tokenAuth *jwtauth.JWTAuth

func newConn(ctx context.Context, cfg *config.Database) (pool *pgxpool.Pool, err error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s", cfg.Driver, cfg.User, cfg.Password,
		cfg.Host, cfg.Port, cfg.Name)

	pool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("подключение к БД: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("пинг БД: %w", err)
	}

	return pool, nil
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Println("создание директории для логов:", err)
	}

	logPath := filepath.Join(logDir, "logs.log")
	logFileW, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Failed to open log file:", err)
		return
	}
	defer logFileW.Close()

	logger := loggerPackage.NewLogger(cfg.Logger.Level, logFileW)
	if err != nil {
		log.Fatalln("cоздание логгера:", err)
	}

	tokenAuth = jwtauth.New("HS256", []byte(cfg.Server.JwtKey), nil)

	pool, err := newConn(context.Background(), &cfg.Database)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	a := app.NewApp(pool, cfg, logger)

	mux := chi.NewMux()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Use(middleware.Logger)

	mux.Route("/entrepreneurs", func(r chi.Router) {
		r.Get("/{id}", web.GetEntrepreneur(a))
		r.Get("/", web.ListEntrepreneurs(a))
		r.Get("/{id}/rating", web.CalculateRating(a))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateAdminRoleJWT)

			r.Patch("/{id}", web.UpdateEntrepreneur(a))
			r.Delete("/{id}", web.DeleteEntrepreneur(a))
		})
	})

	mux.Route("/contacts", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Get("/", web.ListEntrepreneurContacts(a))
			r.Post("/", web.CreateContact(a))
			r.Get("/{id}", web.GetContact(a))
			r.Patch("/{id}", web.UpdateContact(a))
			r.Delete("/{id}", web.DeleteContact(a))
		})
	})

	mux.Route("/activity_fields", func(r chi.Router) {
		r.Get("/{id}", web.GetActivityField(a))
		r.Get("/", web.ListActivityFields(a))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateAdminRoleJWT)

			r.Post("/", web.CreateActivityField(a))
			r.Patch("/{id}", web.UpdateActivityField(a))
			r.Delete("/{id}", web.DeleteActivityField(a))
		})
	})

	mux.Route("/companies", func(r chi.Router) {
		r.Get("/{id}", web.GetCompany(a))
		r.Get("/", web.ListEntrepreneurCompanies(a))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Post("/", web.CreateCompany(a))
			r.Patch("/{id}", web.UpdateCompany(a))
			r.Delete("/{id}", web.DeleteCompany(a))
		})

		r.Route("/{id}/financials", func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Post("/", web.CreateReport(a))
			r.Get("/", web.ListCompanyReports(a))
		})
	})

	mux.Route("/financials", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Get("/", web.GetEntrepreneurFinancials(a))
			r.Delete("/{id}", web.DeleteFinReport(a))
			r.Patch("/{id}", web.UpdateFinReport(a))
		})
	})

	mux.Post("/login", web.LoginHandler(a))
	mux.Post("/signup", web.RegisterHandler(a))

	go func() {
		metricsAddress := fmt.Sprintf("%s:%s", cfg.Server.MetricsHost, cfg.Server.MetricsPort)

		metricsMux := http.NewServeMux()
		metricsMux.Handle("/metrics", promhttp.Handler())

		fmt.Printf("сервер метрик прослушивает адрес: %s\n", metricsAddress)
		http.ListenAndServe(metricsAddress, metricsMux)
	}()

	serverAddress := fmt.Sprintf("%s:%s", cfg.Server.ServerHost, cfg.Server.ServerPort)
	fmt.Printf("сервер прослушивает адрес: %s\n", serverAddress)
	logger.Infof("сервер прослушивает адрес: %s\n", serverAddress)
	http.ListenAndServe(serverAddress, mux)
}
