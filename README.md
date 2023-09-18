# [Stop](https://github.com/dbtedman/stop)

Go experiments monorepo.

## Experiments

An experiment for [every colour of the web](https://en.wikipedia.org/wiki/Web_colors).

### [Aquamarine](./aquamarine/)

Provide security by proxying requests to legacy applications.

```shell
go run ./aquamarine
```

### [Crimson](./crimson/)

An exploration into security headers with a gohtml site.

```shell
go run ./crimson
go run ./crimson -dev=true
```

## Workflow

```shell
# On updating go.work file.
go work sync
```

## Resources

Resources referenced during the development of these experiments.

- [Cache-Control for Civilians (csswizardry.com)](https://csswizardry.com/2019/03/cache-control-for-civilians/)
- [Content Security Policy Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/Content_Security_Policy_Cheat_Sheet.html)
- [Go by Example: Command-Line Flags (gobyexample.com)](https://gobyexample.com/command-line-flags)
- [HTTP Security Response Headers Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html)
- [HTTP headers (developer.mozilla.org)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
- [HTTP headers for the responsible developer (files.gotocon.com)](https://files.gotocon.com/uploads/slides/conference_16/1117/original/Stefan-Judis-http-headers-for-the-responsible-developer.pdf)
- [Password Storage Cheat Sheet (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- [Tutorial: Getting started with multi-module workspaces (go.dev)](https://go.dev/doc/tutorial/workspaces)
- [What is .crt and .key files and how to generate them? (serverfault.com)](https://serverfault.com/questions/224122#answer-224127)
