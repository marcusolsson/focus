focus
=====

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/marcusolsson/focus)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

Subscribe to window focus events. Currently only supports Linux.

## Background

I wrote this mainly to solve the problem of application-specific keyboard layouts in Linux. Simply put, I want to use Swedish layout when browsing but US layout when I am coding and I could not find any simple way of doing this.

The `xkbd` command in this project is a uses the package to solve this by executing `setxkbmap` whenever I switch between browser and terminal. 

If you know of any better way of doing this, please let me know.

Note: Honestly, this is yak shaving at it's finest. Nowadays I'm using: `setxkbmap -layout us -variant altgr-intl -option ctrl:nocaps`
