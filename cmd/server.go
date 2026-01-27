package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
	"github.com/spf13/cobra"

	"github.com/ljgago/sift/internal/config"
	"github.com/ljgago/sift/internal/log"
	"github.com/ljgago/sift/internal/middleware"
	"github.com/ljgago/sift/internal/repository"

	"github.com/ljgago/sift/api"
	"github.com/ljgago/sift/ui"
)

const DefaultServerPort = "8090"
const DefaultNPMRegistryURL = "https://registry.npmjs.com"

type ServerFlags struct {
	Port string
	URL  string
	JSON bool
}

// rootCmd represents the root
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  `Start the server`,
	Run: func(cmd *cobra.Command, args []string) {
		serverRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("port", "p", DefaultServerPort, "server port")
	serverCmd.Flags().StringP("url", "u", DefaultNPMRegistryURL, "npm registry url")
	serverCmd.Flags().BoolP("json", "j", false, "show log format in json")
}

func parseServerFlags(cmd *cobra.Command) (*ServerFlags, error) {
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		return nil, fmt.Errorf("Error getting server port flag -> %w", err)
	}

	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return nil, fmt.Errorf("Error getting npm registry url flag -> %w", err)
	}

	json, err := cmd.Flags().GetBool("json")
	if err != nil {
		return nil, fmt.Errorf("Error getting log flag -> %w", err)
	}

	return &ServerFlags{
		Port: port,
		URL:  url,
		JSON: json,
	}, nil
}

func serverRun(cmd *cobra.Command, _ []string) {
	flags, err := parseServerFlags(cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	log.Setup(flags.JSON, "debug", "sift")

	cfg := config.Config{}
	cfg.Server.URL = flags.URL

	r := chi.NewRouter()
	r.Use(chimiddleware.Recoverer)
	r.Use(httplog.RequestLogger(log.Get()))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://demo.com"},
		AllowedMethods: []string{"GET"},
		// AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRF-Token"},
		// // ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		// MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.CSRF)
	r.Use(middleware.HTTPCache)

	repo := repository.New(cfg)
	api.RegistryRoutes(r, repo)

	r.Handle("/_app/*", http.FileServerFS(ui.AssetsDirFS))
	r.HandleFunc("/*", handleFallback)

	log.Info("Server start at :" + flags.Port)
	server := &http.Server{
		Addr:    ":" + flags.Port,
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Errorf("Server failed: %s\n", err)
	}
}

func handleFallback(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, ui.AssetsDirFS, "index.html")
}
