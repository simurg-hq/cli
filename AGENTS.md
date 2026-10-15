# AGENTS.md

Instructions for coding agents and humans working in this repository.

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
| `scripts/ci/` | CI helper scripts and their tests. |

Some of these paths are created by later tasks; keep to this layout when you
add them.

## Naming

- Standard Go conventions, formatted with `gofmt`. Package names are short, lowercase, singular and without underscores.
- Exported identifiers have doc comments. Errors are wrapped with `%w` and start with a lowercase letter.
- Subcommands are single lowercase words (`run`, `sync`, `version`); flags are `kebab-case`.

## Rules

- One task, one branch, one pull request against `main`. Keep the diff outside tests and generated files under 400 lines; if it will not fit, stop and ask for the task to be split.
- Write the failing test first. Every acceptance criterion maps to at least one test.
- Implement only what the task asks for.
- Network calls have short timeouts and their failures are never returned as the wrapped command's failure.
- Standard library first. Any new dependency needs a reason in the pull request.
- Never lower a threshold, disable or skip a check, delete or weaken a test, or add an ignore comment, suppression or baseline to make a gate pass. Such a change needs its own pull request approved by a human.
- Do not change `.github/`, `AGENTS.md`, `CLAUDE.md`, `.golangci.yml` or test thresholds unless the task is explicitly about them.
- Never hardcode the product name in user-facing text; use `product.Name` from `internal/product`.
- Never create or commit secrets, tokens or real credentials.
- Code, comments, commit messages and documents are in English.
- Commit messages use a type prefix: `feat`, `fix`, `test`, `refactor`, `docs`, `ci`, `chore`.

## Quality gates

Every gate runs in Docker through `make`; nothing but Docker and Lefthook is
needed on the host. Run the gates before pushing and show their output as
evidence. "It should work" is not evidence.

```sh
make help               # list all targets
make lint-ci-scripts    # shellcheck for scripts/ci
make test-ci-scripts    # bats tests for scripts/ci
make pr-size            # measure this branch against origin/main
make check              # all gates below
make fmt-check          # gofmt
make lint               # golangci-lint (.golangci.yml)
make test               # go test -race
make vulncheck          # govulncheck
make secrets            # gitleaks over the git history
make build              # build Linux/macOS × amd64/arm64 into dist/; check the injected version
```

A task that adds a gate also adds its `make` target and lists it here.

## When you are stuck

1. If the same gate fails three times, do not weaken the test or the gate.
2. Hand the task to the next implementer tier.
3. If the top tier also fails three times on the same gate, stop and write a blocker note: which gate failed, the exact error, what was tried and why it did not work.

## Definition of done

The checklist in `.github/pull_request_template.md` must be complete before a
pull request is merged.
