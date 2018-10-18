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
