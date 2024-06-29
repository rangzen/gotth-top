# top tool with the GoTTH Stack

## General

Simple top implementation providing a dynamic real-time view of a running system build with the GoTTH Stack for fun.

## GoTTH Stack

- [Go](https://go.dev), the backend language used for server-side logic and handling requests
- [Templ](https://templ.guide), a templating language for generating HTML templates in Go format
- [Tailwind CSS](https://tailwindcss.com), a utility-first CSS framework for styling the application
- [HTMX](https://htmx.org), a library for adding interactivity and dynamic updates to web pages without relying heavily on JavaScript

Add-ons:

- [gopsutil](https://github.com/shirou/gopsutil), psutil (process and system utilities) port in Go
- [Echo](https://echo.labstack.com), a high-performance, extensible web framework for Go
- [daisyUI](https://daisyui.com), a plugin for Tailwind CSS that adds a set of beautiful components and utilities

## Requirements

```shell
go install templ
npm install onchange
```

## References

- https://github.com/danawoodman/go-echo-htmx-templ
- https://github.com/emarifer/go-echo-templ-htmx
