# cli

The `simurg` command line tool. It wraps any command, reports its start, finish, exit code and duration to Simurg, and never blocks or fails the wrapped command.

## Local setup

Prerequisites:

- [Docker](https://docs.docker.com/get-docker/) (OrbStack or Docker Desktop)
- [Lefthook](https://lefthook.dev/) for Git hooks

All language toolchains and quality gates run in pinned Docker images, so nothing else needs to be installed on the host.

Build the binaries for Linux and macOS on amd64 and arm64:

```sh
make build                          # writes dist/simurg-<os>-<arch>
./dist/simurg-darwin-arm64 version  # on an Apple silicon Mac
```

The version comes from `git describe`; set it with `make build VERSION=v0.1.0`.
`make check` runs all quality gates; `make help` lists every target.

## License

MIT. See [LICENSE](LICENSE).
