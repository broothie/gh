package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/broothie/gh/model"
	"github.com/samber/lo"
)

func main() {
	port := flag.Int("p", 8080, "port to serve on")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if filepath.Ext(r.URL.Path) == ".wasm" {
			prefix := strings.SplitN(r.URL.Path, ".", 2)[0]
			cmd := exec.CommandContext(r.Context(), "go", "build",
				"-o", filepath.Join("public", r.URL.Path),
				fmt.Sprintf("./www/%s/...", prefix),
			)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")

			fmt.Println(cmd.String())
			if err := cmd.Run(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		fileServer.ServeHTTP(w, r)
	})

	users := []model.User{
		{ID: "0", Username: "broothie", Age: 29},
		{ID: "1", Username: "chenstiffany", Age: 29},
	}

	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		users := lo.Filter(users, func(user model.User, i int) bool { return strings.Contains(user.Username, query) })

		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		fmt.Println("failed to start server", err)
	}
}
