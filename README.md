# Ekaki

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/ekaki.svg)](https://pkg.go.dev/github.com/aethiopicuschan/ekaki)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/ekaki)](https://goreportcard.com/report/github.com/aethiopicuschan/ekaki)
[![CI](https://github.com/aethiopicuschan/ekaki/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/ekaki/actions/workflows/ci.yaml)

Ekaki is a simple CLI tool and library to convert images.

## Installation

As a CLI tool:

```sh
go install github.com/aethiopicuschan/ekaki/cmd/ekaki@latest
ekaki input.png output.jpg
```

As a library:

```sh
go get -u github.com/aethiopicuschan/ekaki/pkg/ekaki
```

## Supported formats

- bmp
- jpeg
- png
- gif
- webp
- tiff

### TODO

- [ ] Conversion while maintaining GIF and WebP animations
