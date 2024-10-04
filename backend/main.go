package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"ppo/internal/app"
	"ppo/internal/config"
	loggerPackage "ppo/pkg/logger"
	"ppo/web"
	v1 "ppo/web/v1"
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

// @title Cервис поиска партнеров-предпринимателей
// @version 1.0
// @description Сервис призван помочь с поиском партнеров по бизнесу.

// @contact.name Dmitry Lankin
// @contact.url @lankind
// @contact.email lankindl@student.bmstu.ru

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8081
// @BasePath /api/v2
// @query.collection.format multi

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
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

	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"), //The url pointing to API definition
	))

	mux.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/entrepreneurs", func(r chi.Router) {
				r.Get("/{id}", v1.GetEntrepreneur(a))
				r.Get("/", v1.ListEntrepreneurs(a))
				r.Get("/{id}/rating", v1.CalculateRating(a))
				r.Get("/companies", v1.ListEntrepreneurCompanies(a))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateAdminRoleJWT)

					r.Patch("/{id}", v1.UpdateEntrepreneur(a))
					r.Delete("/{id}", v1.DeleteEntrepreneur(a))
				})
			})

			r.Route("/contacts", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Get("/", v1.ListEntrepreneurContacts(a))
					r.Post("/", v1.CreateContact(a))
					r.Get("/{id}", v1.GetContact(a))
					r.Patch("/{id}", v1.UpdateContact(a))
					r.Delete("/{id}", v1.DeleteContact(a))
				})
			})

			r.Route("/activity_fields", func(r chi.Router) {
				r.Get("/{id}", v1.GetActivityField(a))
				r.Get("/", v1.ListActivityFields(a))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateAdminRoleJWT)

					r.Post("/", v1.CreateActivityField(a))
					r.Patch("/{id}", v1.UpdateActivityField(a))
					r.Delete("/{id}", v1.DeleteActivityField(a))
				})
			})

			r.Route("/companies", func(r chi.Router) {
				r.Get("/{id}", v1.GetCompany(a))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/", v1.CreateCompany(a))
					r.Patch("/{id}", v1.UpdateCompany(a))
					r.Delete("/{id}", v1.DeleteCompany(a))
				})

				r.Route("/{id}/financials", func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/", v1.CreateReport(a))
					r.Get("/", v1.ListCompanyReports(a))
				})
			})

			r.Route("/financials", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Get("/", v1.GetEntrepreneurFinancials(a))
					r.Delete("/{id}", v1.DeleteFinReport(a))
					r.Patch("/{id}", v1.UpdateFinReport(a))
				})
			})

			r.Post("/login", v1.LoginHandler(a))
			r.Post("/signup", v1.RegisterHandler(a))
		})

		r.Route("/v2", func(r chi.Router) {
			r.Route("/entrepreneurs", func(r chi.Router) {
				r.Get("/{id}", v1.GetEntrepreneur(a))
				r.Get("/", v1.ListEntrepreneurs(a))
				r.Get("/{id}/rating", v1.CalculateRating(a))
				r.Get("/companies", v1.ListEntrepreneurCompanies(a))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateAdminRoleJWT)

					r.Patch("/{id}", v1.UpdateEntrepreneur(a))
					r.Delete("/{id}", v1.DeleteEntrepreneur(a))
				})
			})

			r.Route("/contacts", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Get("/", v1.ListEntrepreneurContacts(a))
					r.Post("/", v1.CreateContact(a))
					r.Get("/{id}", v1.GetContact(a))
					r.Patch("/{id}", v1.UpdateContact(a))
					r.Delete("/{id}", v1.DeleteContact(a))
				})
			})

			r.Route("/activity_fields", func(r chi.Router) {
				r.Get("/{id}", v1.GetActivityField(a))
				r.Get("/", v1.ListActivityFields(a))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateAdminRoleJWT)

					r.Post("/", v1.CreateActivityField(a))
					r.Patch("/{id}", v1.UpdateActivityField(a))
					r.Delete("/{id}", v1.DeleteActivityField(a))
				})
			})

			r.Route("/companies", func(r chi.Router) {
				r.Get("/{id}", v1.GetCompany(a))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/", v1.CreateCompany(a))
					r.Patch("/{id}", v1.UpdateCompany(a))
					r.Delete("/{id}", v1.DeleteCompany(a))
				})

				r.Route("/{id}/financials", func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/", v1.CreateReport(a))
					r.Get("/", v1.ListCompanyReports(a))
				})
			})

			r.Route("/financials", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Get("/", v1.GetEntrepreneurFinancials(a))
					r.Delete("/{id}", v1.DeleteFinReport(a))
					r.Patch("/{id}", v1.UpdateFinReport(a))
				})
			})

			r.Post("/login", v1.LoginHandler(a))
			r.Post("/signup", v1.RegisterHandler(a))
		})
	})

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
