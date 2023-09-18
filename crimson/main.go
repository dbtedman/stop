package main

import (
	"crypto/rand"
	"embed"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/argon2"
	"html/template"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
)

//- https://cheatsheetseries.owasp.org/cheatsheets/Content_Security_Policy_Cheat_Sheet.html
//- https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html
//- https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#content-security-policy-csp
//- https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#permissions-policy-formerly-feature-policy
//- https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#referrer-policy
//- https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#strict-transport-security-hsts
//- https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#x-frame-options
//- https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
//- https://csp-evaluator.withgoogle.com
//- https://csswizardry.com/2019/03/cache-control-for-civilians/
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/sandbox
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/upgrade-insecure-requests
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
//- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
//- https://files.gotocon.com/uploads/slides/conference_16/1117/original/Stefan-Judis-http-headers-for-the-responsible-developer.pdf
//- https://github.com/CAFxX/httpcompression
//- https://github.com/google/brotli
//- https://hstspreload.org
//- https://serverfault.com/questions/224122#answer-224127
//- https://whatdoesmysitecost.com
//- https://webhint.io/
//- https://securityheaders.com

//go:embed static/*
var embeddedFS embed.FS

type IndexData struct {
	Nonce string
	Title string
}

func main() {
	// TODO: Read from cli/env
	dev := true
	server := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonce := generateNonce()
		w.WriteHeader(http.StatusOK)
		setHeaders(w, nonce)
		err := templateForFile(dev, "static/index.gohtml").Execute(w, IndexData{Nonce: nonce, Title: "Crimson"})
		if err != nil {
			fmt.Println(err)
		}
	}))

	err := server.ListenAndServeTLS("./crimson/host.cert", "./crimson/host.key")

	if err != nil {
		fmt.Println(err)
	}
}

func templateForFile(dev bool, filePath string) *template.Template {
	var indexHTMLTemplate *template.Template

	// When in development, load html directly from file system, and when in production
	// load from embedded file system.
	if dev {
		fmt.Println("Reading file from file system.")
		indexHTMLBytes, _ := os.ReadFile("crimson/" + filePath)
		indexHTMLTemplate, _ = template.New(filePath).Parse(string(indexHTMLBytes))
	} else {
		fmt.Println("Reading file from embedded file system.")
		indexHTMLBytes, _ := embeddedFS.ReadFile(filePath)
		indexHTMLTemplate, _ = template.New(filePath).Parse(string(indexHTMLBytes))
	}

	return indexHTMLTemplate
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
	// Cache-Control: max-age=31536000, public, immutable
	// TODO: no-transform?
	w.Header().Set("Content-Security-Policy", "default-src 'none'; style-src 'nonce-"+nonce+"'; upgrade-insecure-requests; frame-ancestors 'none'; form-action 'none'; sandbox")
	w.Header().Set("Permissions-Policy", "geolocation=(), camera=(), microphone=(), display-capture=()")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=300; includeSubDomains")

	// Example
	w.Header().Set("Set-Cookie", "example=example; SameSite=Strict; HttpOnly; Secure")

	// TODO: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-CH
	// TODO: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary
	// TODO: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Save-Data
	// TODO: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-Prefers-Reduced-Motion
	// TODO: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-Prefers-Reduced-Transparency
	// TODO: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-Prefers-Color-Scheme

	// TODO: webp and other image formats
	// TODO: Client hints

	// TODO: rel=preload; as=image; no-push
}
