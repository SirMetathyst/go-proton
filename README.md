# Go Entitas

Go-Entitas is a package for modeling contexts / aliases / components / entity-indexes for [entitas](https://github.com/sschmid/Entitas-CSharp) code generation.

Immediate TODO:

- Cleanup proton package and write tests

TODO:

- Rewrite documentation after project restructuring...
- Currently all generators are off by default, create system to find newest version
and enable the latest generators/post-processors










## Installation

```
go get -u github.com/SirMetathyst/go-entitas
```

## Example

```go
import (
    entitas "github.com/SirMetathyst/go-entitas"
)

func main() {

    // Create model builder
    mb := entitas.NewModelBuilder();

    // Set target
    mb.SetTarget("entitas_csharp")

    // Set namespace
    mb.SetNamespace("my.game")

    // Create context
    err := mb.NewContext().SetID("game").Build()
    if err != nil {
        panic(err)
    }

    // Set default context
    mb.SetDefaultContext("game")

    // Create alias
    err = mb.NewAlias().SetID("go").SetValue("UnityEngine.GameObject").Build()
    if err != nil {
        panic(err)
    }

    // TODO: Continue documentation ...
}
```