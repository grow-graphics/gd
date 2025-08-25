---
title: The Development Workflow
slug: guide/workflow
sidebar:
  order: 10
---

Start with a `main.go` file, model your project in Go using structs to represent the
world, space or state of your project. With its simple type-safe syntax, Go is an
excellent language for theory building and conceptual representation.

```go
package idea

import "graphics.gd/classdb/Node"

type Project struct {
	Node.Extension[Project]
}
````

Keep creating more data structures to represent the concepts of what you are aiming to
build. Explore the `classdb` to discover all the functionality the engine provides.
Think about how you can use these features to build your project.

When you want to organise the visual appearance of the project, use the `gd` command to
launch the engine's editor. The editor is excellent for importing media, managing assets
and designing the visual + spatial elements of your project.
