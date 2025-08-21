---
title: Dependencies
slug: guide/dependencies
sidebar:
  order: 2
---

## Go

If you haven't already done so, you will need to [Download and Install Go](https://go.dev/dl/) so you will be able to compile any `.go` source files into native executables.

After Go is installed, make sure `$GOPATH/bin` is part of your `PATH`, so that any tools that you decide to `go install` will be available for use in your terminal.

##### Windows
```powershell
setx PATH "%PATH%;%GOPATH%\bin"
```

##### macOS
```sh
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.zshrc && source ~/.zshrc
```

##### Linux
```sh
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.bashrc && source ~/.bashrc
```

## Zig
In order for Go code to work together with the C++ code in the Godot Engine, you'll need a C compiler. We officially
support [The Zig Compiler](https://ziglang.org/download/) as it supports full cross-compilation.

> **BTW** graphics.gd can build projects for supported platforms on any operating system (this includes building for macOS on Windows and Linux without access to Apple hardware).

##### Windows
Please ensure that Zig is added to your `PATH` after you've extracted it to a convieniant location on your computer.

```powershell
setx PATH "%PATH%;C:\place\you\extracted\zig-windows-x86_64-your-version"
```

##### macOS
You can use [Homebrew](https://brew.sh/) to install Zig:
```sh
brew install zig
```

##### Linux
You should be able to install Zig using your package manager.
```sh
snap install zig
```
