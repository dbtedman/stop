package main

import (
	"crypto/rand"
	"embed"
	"encoding/hex"
	"flag"
	"fmt"
	"html/template"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/argon2"
)

//go:embed static/*
var embeddedFS embed.FS

type IndexData struct {
	Nonce string
	Title string
}

func main() {
	dev := flag.Bool("dev", false, "enable development mode")
	listen := flag.String("listen", ":3000", "address to listen on")
	cert := flag.String("cert", "./crimson/host.cert", "")
	key := flag.String("key", "./crimson/host.key", "")
	root := flag.String("root", "./crimson", "")

	flag.Parse()

	server := &http.Server{
		Addr:         *listen,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonce := generateNonce()
		w.WriteHeader(http.StatusOK)
		setHeaders(w, nonce)
		theTemplate, err := templateForFile(*dev, "static/index.gohtml", *root)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = theTemplate.Execute(w, IndexData{Nonce: nonce, Title: "Crimson"})
		if err != nil {
			fmt.Println(err)
		}
	}))

	fmt.Printf("Listening %s; dev=%t\n", *listen, *dev)
	err := server.ListenAndServeTLS(*cert, *key)

	if err != nil {
		fmt.Println(err)
	}
}

func templateForFile(dev bool, filePath string, rootPath string) (*template.Template, error) {
	// When in development, load html directly from file system, and when in production
	// load from embedded file system.
	if dev {
		return templateForFileFromFileSystem(filePath, rootPath)
	} else {
		return templateForFileFromEmbeddedFileSystem(filePath)
	}
}

func templateForFileFromFileSystem(filePath string, rootPath string) (*template.Template, error) {
	fmt.Println("Reading file from file system.")
	indexHTMLBytes, _ := os.ReadFile(rootPath + "/" + filePath)
	return template.New(filePath).Parse(string(indexHTMLBytes))
}

func templateForFileFromEmbeddedFileSystem(filePath string) (*template.Template, error) {
	fmt.Println("Reading file from embedded file system.")
	indexHTMLBytes, _ := embeddedFS.ReadFile(filePath)
	return template.New(filePath).Parse(string(indexHTMLBytes))
}

func generateNonce() string {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	return hex.EncodeToString(argon2.IDKey(
		[]byte(time.Now().Format("2006-01-02 15:04:05.000000000")),
		[]byte(strconv.FormatUint(nBig.Uint64(), 10)),
		uint32(2),
		uint32(15*1024),
		uint8(1),
		uint32(16),
	))
}

func setHeaders(w http.ResponseWriter, nonce string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "max-age=0, private, must-revalidate")
	w.Header().Set("Content-Security-Policy", contentSecurityPolicy(nonce))
	w.Header().Set("Permissions-Policy", "geolocation=(), camera=(), microphone=(), display-capture=()")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=300; includeSubDomains")

	// Demonstrate secure cookie assignment.
	w.Header().Set("Set-Cookie", "example=example; SameSite=Strict; HttpOnly; Secure")
}

func contentSecurityPolicy(nonce string) string {
	return "default-src 'none'; style-src 'nonce-" + nonce + "'; upgrade-insecure-requests; frame-ancestors 'none'; form-action 'none'; sandbox"
}
