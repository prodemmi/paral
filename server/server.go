package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	d "paral/core/dag"
	"paral/core/runtime"
	"path"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

func watchAndRestartIfChanged(filename string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("failed to create file watcher: %v", err)
	}
	defer watcher.Close()

	if err := watcher.Add(filename); err != nil {
		log.Fatalf("failed to watch file: %v", err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Printf("file changed: %s. Restarting server...", event.Name)
				time.Sleep(500 * time.Millisecond)

				binary, err := exec.LookPath(os.Args[0])
				if err != nil {
					log.Fatalf("failed to find binary: %v", err)
				}

				if err := syscall.Exec(binary, os.Args, os.Environ()); err != nil {
					log.Fatalf("failed to exec: %v", err)
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("watcher error: %v", err)
		}
	}
}

func NewGraphServer(runtime *runtime.Runtime) {
	dag := d.NewDAG(runtime)
	dag = dag.GenerateDAG()
	graph, err := dag.MarshalJSON()
	if err != nil {
		runtime.Reporter.ThrowRuntimeError(fmt.Sprintf("failed to generate DAG: %s", err), nil)
		return
	}

	go watchAndRestartIfChanged(runtime.Metadata.Filename)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(cors.Default())

	// Static files (serve any files in /public/, fallback to directory)
	r.Static("/assets", "./public")

	// Template files
	r.LoadHTMLGlob("public/*.tmpl")

	r.GET("/graph", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "graph.tmpl", gin.H{
			"title":    "Paral Graph Viewer",
			"filename": path.Base(runtime.Metadata.Filename),
			"graph":    graph,
			"content":  runtime.Metadata.Content,
		})
	})

	r.GET("/api/dag", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "application/json", graph)
	})

	log.Println("🚀 Paral Graph Server is running at: http://localhost:8091/graph")
	if err := r.Run(":8091"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
