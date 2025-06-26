package server

import (
	"net/http"
	"crypto/tls"
	"log"
	"fmt"

	"github.com/th3corinthian/org-ssg/internal/models"
)

func Start(cfg *models.Config) error {
	// points at generated output dir
	fs 		:= http.FileServer(http.Dir("./myproj/_site"))
	handler := fs

	addr	:= fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	srv 	:= &http.Server{
		Addr:		addr,
		Handler:	handler,
	}
	log.Printf("Serving %q on http://%s\n", "./myproj/_site", addr)

	var err error
	if cfg.Server.UseTLS {
		srv.startTLS()
		err = srv.ListenAndServeTLS("cert.pem", "key.pem")
	} else {
		err = startTLS()
	}
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func startTLS() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Failed to load X509 key pair: %v", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	server := &http.Server{
		TLSConfig: config,
	}

	err = server.ListenAndServeTLS("","")
}

//func HelloServer(w http.ResponseWriter, req *http.Request) {
//	w.Header().Set("Content-Type", "text/plain")
//	w.Write([]byte("This is an example server.\n"))
//}

//func ListenAndPut() {
//	http.HandleFunc("/hello", HelloServer)
//	err := http.ListenAndServeTLS(":433", "server.crt", "server.key", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe:", err)
//	}
//}
