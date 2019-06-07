# L i t t l e  f l a g s

Small and simple feature flags based on environment variables for your project.

## Status

[![Build Status](https://secure.travis-ci.org/roughy/little-flags.svg?branch=master)](https://travis-ci.org/roughy/little-flags)

## What are feature flags

Feature flags (also feature toggles) are a pattern to en-/disable certain features in a software.
Once you are using feature flags you are able to commit and deloy unfinished features to your production environment very fast and release it (aka enable the feature) when it is finished. This gives you a simple and powerful feedback mechanism that your changes are always running in production.

## Overview
You just want to en- or disable your features quick and easy. `L i t t l e  f l a g s` are meant to be very light weight and simple to use. With a few lines of code and a environment variable you are already ready to use feature flags.

## Installation

 * Install via `go get`
   * `go get -a github.com/roughy/little-flags`
 * Install via release
   * [https://github.com/roughy/little-flags/releases](https://github.com/roughy/little-flags/releases)
 * Build from source
```
git clone git@github.com:roughy/little-flags.git
cd little-flags
go build
```

## Example

### Implementation

```Go
import (
    "fmt"
    "github.com/roughy/little-flags"
)

func main() {
    littleflags.Initialize("DOC_EXAMPLE")

    fmt.Print("You")

    feature, _ := littleflags.GetFeature("feature")
    if feature.IsEnabled() {
        fmt.Print(" are")
    }

    feature, _ = littleflags.GetFeature("wonder")
    if feature.IsEnabled() {
        fmt.Print(" wonderful")
    }

    fmt.Print("\n")
}
```

### Usage

```bash
go run main.go " Prints: You"
DOC_EXAMPLE_FEATURE=1 go run main.go " Prints: You are"
DOC_EXAMPLE_ALPHA=1 go run main.go " Prints: You are wonderful"
```

`Alpha` is a special key word to enable all features. You can read it like "Enable all features in alpha status".

#### Options of value

To enable a feature you can set the value `1` or `on` to your environment variable. You don't have to disable anything explicitly because all `L i t t l e  f l a g s` are disabled by default. If you still want to do it explicit you can assign the value `0` or `""` to the environment variable.
