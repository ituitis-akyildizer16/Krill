# Changelog

All notable changes to krill are documented here. The format is based on
[Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project
adheres to [Semantic Versioning](https://semver.org/).

## [0.9.2] - 2026-03-09

### Fixed
- Cache keys no longer collide when a file path contains `|`
- `suggest --inline` respects the configured shell

### Changed
- Lowered default model to `qwen2.5-coder:7b` (better latency on M-series)

## [0.9.0] - 2026-01-19

### Added
- `krill models` command
- PowerShell completion script
- `--context-only` flag on `ask` for debugging prompts

## [0.8.0] - 2025-11-02

### Added
- Python plugin SDK (`plugins/`) with context providers
- `krill init` writes a default config

### Changed
- Context collection runs fully in parallel

## [0.7.0] - 2025-06-12

### Added
- VS Code extension (`extension/`) with ask/review/suggest commands
- Safety deny/warn patterns for `suggest`

## [0.6.0] - 2025-01-27

### Added
- `krill status` repo digest
- TTL cache for model responses

## [0.5.0] - 2024-10-08

### Added
- `krill review` for staged diff summaries
- markdown-stripping terminal renderer

## [0.4.0] - 2024-05-21

### Added
- `krill suggest` with shell-specific prompts
- zsh and fish completions

## [0.3.0] - 2023-12-14

### Added
- Blame-aware answers with commit citations
- Bounded context assembly (max_blame_lines / max_diff_lines)

## [0.2.0] - 2023-04-03
