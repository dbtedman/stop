# [Stop](https://github.com/dbtedman/stop)

[![CI GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/stop/ci.yml?branch=main&style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/stop/actions/workflows/ci.yml?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/dbtedman/stop?style=for-the-badge)](https://goreportcard.com/report/github.com/dbtedman/stop)
[![Latest Release](https://img.shields.io/github/v/release/dbtedman/stop?style=for-the-badge&color=red)](https://github.com/dbtedman/stop/releases)

Go experiments monorepo.

## Experiments

An experiment for [every colour of the web](https://en.wikipedia.org/wiki/Web_colors).

### [Aquamarine](cmd/aquamarine/)

Provide security by proxying requests to legacy applications.

```shell
go run ./aquamarine --from=:3000 --to=https://example.com

curl http://localhost:3000 --head --header "Host: example.com"
```

### [Crimson](cmd/crimson/)

An exploration into security headers with a gohtml site.

```shell
go run ./crimson
go run ./crimson -dev=true
```

## Resources

Resources referenced during the development of these experiments.

- [Accretion (github.com)](https://github.com/dbtedman/accretion)
- [Cache-Control for Civilians (csswizardry.com)](https://csswizardry.com/2019/03/cache-control-for-civilians/)
- [Cobra: A Framework for Modern CLI Apps in Go (cobra.dev)](https://cobra.dev)
- [Content Security Policy Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/Content_Security_Policy_Cheat_Sheet.html)
- [Go by Example: Command-Line Flags (gobyexample.com)](https://gobyexample.com/command-line-flags)
- [HTTP Security Response Headers Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html)
- [HTTP headers (developer.mozilla.org)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
- [HTTP headers for the responsible developer (files.gotocon.com)](https://files.gotocon.com/uploads/slides/conference_16/1117/original/Stefan-Judis-http-headers-for-the-responsible-developer.pdf)
- [How to Create a Reverse Proxy using Golang (codedodle.com)](https://www.codedodle.com/go-reverse-proxy-example.html)
- [Password Storage Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- [Scrutinise (github.com)](https://github.com/dbtedman/scrutinise)
- [Security Headers (securityheaders.com)](https://securityheaders.com)
- [V. Single Host Reverse Proxy (fideloper.com)](https://fideloper.com/golang-single-host-reverse-proxy)
- [What is .crt and .key files and how to generate them? (serverfault.com)](https://serverfault.com/questions/224122#answer-224127)
