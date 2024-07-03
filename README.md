# GoTTH top

This is a simple implementation of the `top` command, which provides a dynamic real-time view of a running system. It is built using the GoTTH Stack for fun.

- [GoTTH top](#gotth-top)
  - [The GoTTH Stack](#the-gotth-stack)
  - [Suplementary tools](#suplementary-tools)
  - [Installation](#installation)
    - [Binary Installation with Go installed](#binary-installation-with-go-installed)
  - [Requirements for dev](#requirements-for-dev)
  - [References](#references)

## The GoTTH Stack

- [Go](https://go.dev): The backend language used for server-side logic and handling requests.
- [Templ](https://templ.guide): A templating language for generating HTML templates in Go format.
- [Tailwind CSS](https://tailwindcss.com): A utility-first CSS framework for styling the application.
- [HTMX](https://htmx.org): A library for adding interactivity and dynamic updates to web pages without relying heavily on JavaScript.

## Suplementary tools

- [gopsutil](https://github.com/shirou/gopsutil): A Go port of psutil (process and system utilities). GoTTH top relies heavily on this.
- [Echo](https://echo.labstack.com): A high-performance, extensible web framework for Go.
- [daisyUI](https://daisyui.com): A plugin for Tailwind CSS that adds a set of beautiful components and utilities.
- [Alpine.js](https://alpinejs.dev/): A minimal JavaScript framework for building interactive web interfaces.

## Installation

### Binary Installation with Go installed

```shell
go install github.com/rangzen/gotth-top/cmd/gotth-top@latest
```

## Requirements for dev

```shell
go install github.com/a-h/templ/cmd/templ@latest # CLI tool to watch templates
npm install onchange                             # Files Watcher
```

## References

- [A GitHub Repository with an example of the stack](https://github.com/danawoodman/go-echo-htmx-templ).
- [Another GitHub Repository with an example of the stack](https://github.com/emarifer/go-echo-templ-htmx).
