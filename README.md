# [Stop](https://github.com/dbtedman/stop)

[![CI GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/stop/ci.yml?branch=main&style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/stop/actions/workflows/ci.yml?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/dbtedman/stop?style=for-the-badge)](https://goreportcard.com/report/github.com/dbtedman/stop)
[![Latest Release](https://img.shields.io/github/v/release/dbtedman/stop?style=for-the-badge&logo=github&color=blue)](https://github.com/dbtedman/stop/releases)

Go experiments monorepo.

- [Experiments](#experiments)
- [Resources](#resources)
- [License](#license)

## Experiments

An experiment for [every colour of the web](https://en.wikipedia.org/wiki/Web_colors).

### [Aquamarine](cmd/aquamarine/)

Provide security by proxying requests to legacy applications.

```shell
brew install dbtedman/tap/aquamarine
```

```shell
aquamarine serve --from=:3000 --to=https://example.com
```

```shell
curl http://localhost:3000 --head --header "Host: example.com"
```

### [Crimson](cmd/crimson/)

An exploration into security headers with a gohtml site.

```shell
brew install dbtedman/tap/crimson
```

```shell
crimson -cert=host.cert -key=host.key
```

### Khaki

Tool for sanitising data files of sensitive information through substitution with fake information.

```shell
khaki --in=./dump.sql --out=./safe.sql --preset=wordpress
```

## Resources

Resources referenced during the development of these experiments.

- [Assigning permissions to jobs (docs.github.com)](https://docs.github.com/en/actions/using-jobs/assigning-permissions-to-jobs)
- [Cache-Control for Civilians (csswizardry.com)](https://csswizardry.com/2019/03/cache-control-for-civilians/)
- [Cobra: A Framework for Modern CLI Apps in Go (cobra.dev)](https://cobra.dev)
- [Configuration options for the dependabot.yml file (docs.github.com)](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file#package-ecosystem)
- [Content Security Policy Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/Content_Security_Policy_Cheat_Sheet.html)
- [Go by Example: Command-Line Flags (gobyexample.com)](https://gobyexample.com/command-line-flags)
- [Go by Example: Regular Expressions (gobyexample.com)](https://gobyexample.com/regular-expressions)
- [Golang UK Conference 2017 | Ian Kent - Production-ready Go (youtube.com)](https://www.youtube.com/watch?v=YF1qSfkDGAQ)
- [Goreleaser - homebrew (goreleaser.com)](https://goreleaser.com/customization/homebrew/)
- [HTTP Security Response Headers Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html)
- [HTTP headers (developer.mozilla.org)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
- [HTTP headers for the responsible developer (files.gotocon.com)](https://files.gotocon.com/uploads/slides/conference_16/1117/original/Stefan-Judis-http-headers-for-the-responsible-developer.pdf)
- [How to Create a Reverse Proxy using Golang (codedodle.com)](https://www.codedodle.com/go-reverse-proxy-example.html)
- [How to Hash and Verify Passwords With Argon2 in Go (alexedwards.net)](https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go)
- [How to Use go:embed in Go (blog.jetbrains.com)](https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/)
- [How to properly seed random number generator (stackoverflow.com)](https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator)
- [Password Storage Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- [Security Headers (securityheaders.com)](https://securityheaders.com)
- [Serve embedded filesystem from root path of URL (stackoverflow.com)](https://stackoverflow.com/questions/66248258)
- [Testing Your (HTTP) Handlers in Go (blog.questionable.services)](https://blog.questionable.services/article/testing-http-handlers-go/)
- [Using a nonce with CSP (content-security-policy.com)](https://content-security-policy.com/nonce/)
- [V. Single Host Reverse Proxy (fideloper.com)](https://fideloper.com/golang-single-host-reverse-proxy)
- [What is .crt and .key files and how to generate them? (serverfault.com)](https://serverfault.com/questions/224122#answer-224127)
- [What's the best way to bundle static resources in a Go program? (stackoverflow.com)](https://stackoverflow.com/questions/13904441)

## License

See [LICENSE.md](./LICENSE.md) for details.
