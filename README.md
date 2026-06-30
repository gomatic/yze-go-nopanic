# yze-go-nopanic

A [`yze`](https://github.com/gomatic/yze) analyzer (category `patterns`) that forbids calls to the built-in `panic` outside test files, per the gomatic Go standard that errors are returned rather than raised by panicking.

- **Rule:** `yze/nopanic`
- **Library:** exports `Analyzer` (a standard `go/analysis` analyzer) and `Registration` for the [`yze`](https://github.com/gomatic/yze) aggregator and [`stickler`](https://github.com/gomatic/stickler) runner.
- **Binary:** `cmd/yze-go-nopanic` runs it standalone (`text`/`-json`, and as a `go vet -vettool`).

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.
