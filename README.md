# K8s Resource Monitor

![CI](https://github.com/Qyroxen/K8s-Resource-Monitor/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/K8s-Resource-Monitor/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/K8s-Resource-Monitor?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/K8s-Resource-Monitor)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/K8s-Resource-Monitor)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/K8s-Resource-Monitor?style=social)](https://github.com/Qyroxen/K8s-Resource-Monitor/stargazers)

## What is it?

K8s Resource Monitor is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/K8s-Resource-Monitor.git
cd K8s-Resource-Monitor
go build -o k8sresourcemonitor .

# Run
./k8sresourcemonitor --help
```

## CLI Usage

```bash
# Basic usage
./k8sresourcemonitor

# With flags
./k8sresourcemonitor --verbose --output json

# Get help
./k8sresourcemonitor --help
```

## Examples

```bash
# Example 1
./k8sresourcemonitor example1

# Example 2
./k8sresourcemonitor example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o k8sresourcemonitor .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/K8s-Resource-Monitor/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/K8s-Resource-Monitor?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/K8s-Resource-Monitor/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/K8s-Resource-Monitor?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/K8s-Resource-Monitor/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/K8s-Resource-Monitor" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/K8s-Resource-Monitor/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/K8s-Resource-Monitor" alt="Pull Requests">
  </a>
</p>
