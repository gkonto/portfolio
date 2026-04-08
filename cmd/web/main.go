package main

import (
	"crypto/tls"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", "", "HTTP network address")
	tlsCertFile := flag.String("tls-cert", os.Getenv("TLS_CERT_FILE"), "TLS certificate file")
	tlsKeyFile := flag.String("tls-key", os.Getenv("TLS_KEY_FILE"), "TLS key file")
	flag.Parse()

	listenAddr := *addr
	if listenAddr == "" {
		port := os.Getenv("PORT")
		if port == "" {
			port = "4000"
		}
		listenAddr = ":" + port
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:         listenAddr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if *tlsCertFile != "" || *tlsKeyFile != "" {
		if *tlsCertFile == "" || *tlsKeyFile == "" {
			errorLog.Fatal("both -tls-cert and -tls-key must be provided")
		}

		srv.TLSConfig = &tls.Config{
			CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		}

		infoLog.Printf("Starting HTTPS server on %s", listenAddr)
		err = srv.ListenAndServeTLS(*tlsCertFile, *tlsKeyFile)
		errorLog.Fatal(err)
	}

	infoLog.Printf("Starting HTTP server on %s", listenAddr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
