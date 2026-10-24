# AGENTS.md

Project knowledge for coding agents and humans working in this repository.

## Purpose

`cli` builds the `simurg` command line tool. It wraps any command, reports
start, finish, exit code, duration and output tail using the public ping
protocol, and syncs monitor definitions from a YAML file.

The one rule that overrides everything else: the CLI never blocks or fails the
wrapped command. If the service is unreachable, the command still runs and its
exit code is returned unchanged.

This repository is public and MIT licensed. It must stay self-contained: do
not reference private repositories or internal documents.

## Repository layout

| Path | Contents |
| --- | --- |
| `main.go` | Entry point; builds the root command. |
| `internal/` | Command implementations and helpers. |
| `internal/product/` | The single product name constant. |
| `scripts/ci/` | Project scripts (cross-platform build). |

## Naming

- Standard Go conventions, formatted with `gofmt`. Package names are short, lowercase, singular and without underscores.
- Exported identifiers have doc comments. Errors are wrapped with `%w` and start with a lowercase letter.
- Subcommands are single lowercase words (`run`, `sync`, `version`); flags are `kebab-case`.

## Rules

- Network calls have short timeouts and their failures are never returned as the wrapped command's failure.
- Standard library first. Any new dependency needs a reason in the pull request.
- Never hardcode the product name in user-facing text; use `product.Name` from `internal/product`.
- Never create or commit secrets, tokens or real credentials.
- Code, comments, commit messages and documents are in English.

## Implementation

- Meet every need the task states; check them one by one before calling the work done.
- Follow the best practices of the language and its tools. Keep the code DRY and the design SOLID, without layers that only restate configuration.
- Weigh performance and security in every change.
