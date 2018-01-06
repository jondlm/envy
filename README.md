# envy

`envy` is a small CLI tool that can compile golang template files using
environment variables.

## Install

    go get github.com/jondlm/envy

## Usage

Given a `template.tpl` file containing:

    Hello {{ .NAME }}

You can run the following command to inject values:

    NAME=jon envy template.tpl output.txt

## Tests

    go test

