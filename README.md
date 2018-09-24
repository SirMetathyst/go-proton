# Overview
Proton is an [Entitas](https://github.com/sschmid/Entitas-CSharp) Code-Generator written in golang. Distributed as a single executable you are able compose contexts, components and entity index's in a custom DSL ([Proton-Lang](https://github.com/SirMetathyst/go-proton-lang)
) based on [Entitas-Lang](https://github.com/sschmid/Entitas-CSharp/wiki/Tutorial-%E2%80%90-Entitas%E2%80%90lang).
```
// CSharp.
#include "~CSharp/System";
#include "~CSharp/System.Collections.Generic";

// Unity.
#include "~Unity/UnityEngine";

context Input, Command, Core(default), State, Meta;

alias IConfiguration = "IConfiguration";
alias IServiceConfiguration = "IServiceConfiguration";


// Command Components
component EndGameCommand in Command;
component NewGameCommand in Command;
component StartGameCommand in Command;
component StartGameFromSnapshotCommand in Command as int;

// Core Components
component Position(eventTarget:self) as Vector3; 
component Surface(eventTarget:self) as "Surface";
component Tile(eventTarget:self) as int; 
component Move(unique);
component Blank;
component Content as string;
component View as "IView";
component Destroyed;

component ID in Core, State {
    value(entityIndex:single) as int;
}

// State Components
component Snapshot(unique) in State as "Snapshot";

// Meta Components
component Configuration(unique) in Meta as IConfiguration;
component ServiceConfiguration(unique) in Meta as IServiceConfiguration;

// Entity Index's
index BlankEntityIndex(primary) in Core {
    method GetBlankWithCoordinate {
        c as "CrossCoordinate";
    }
}
```

# Installation & Execution

1) Download latest release from the releases page and extract the zip file. 
2) Copy ```proton.exe``` to your project folder or add ```proton.exe``` to your ```PATH``` environment variable.
3) Create a file called ```my_project.proton```
4) add a context and component ```context core(default);``` ```component test as "int";```
5) Execute ```proton``` in watch mode: ```proton.exe --WatchFileEnable=true --WatchFile=./my_project.proton --File=./my_project.proton``` or without: ```proton.exe --File=./my_project.proton```

# Proton-Lang

## Include Statement

```
#include "your_file.proton";
```
When Proton-Lang parses an include statement that file will then be immediately parsed before moving onto the rest of the file.

### Reserved Includes

Proton-Lang uses [Packr](https://github.com/gobuffalo/packr) and converts .proton files into go code so you can use common C# and UnityEngine types without having to write your own. Reserved proton includes start with ```~```
```
#include "~CSharp/System";
#include "~CSharp/System.Collections.Generic";
```

## Target Statement

```
target entitas_csharp;
``` 
The target statement defines what target language/runtime you want to produce entitas code for; this statement only works in the first file that gets parsed. While the target statement is implemented for future compatibily between languages/runtimes only csharp is supported. Because of this you don't need to specify a target.

## Namespace Statement

```
namespace my.game;
```
The namespace statement defines a namespace to wrap your entitas code around. While the namespace statement and model supports namespaces there are currently no generators which support it.

## Context Statement

```
context core(default);
context meta, state;
```
You can define a single context or multiple on the same line however without specifying a default context you will need to explicitly specify a context for each component.

## Alias Statement

```
alias Int32 = "System.Int32";
```
```
alias String = "System.String", Int64 = "System.Int64";
```
You can define a single alias or multiple on the same line within an include file or main file but cannot be overidden if already defined.

## Component Statement

```
component name1(key:value), name2 in context1, context2 as "int";
```
```
component name1(key:value), name2 in context1, context2 as int;
```
```
component name1(key:value), name2 in context1, context2
{
    // single variable with alias
    value1 as int;

    // multiple variable with alias.
    value2, value3 as int;

    // multiple variable with value
    value4, value5 as "int";

    // single variable with alias and primary entity index
    value6(entityIndex:single) as int;

    // single variable with alias and entity index
    value7(entityIndex:multiple) as int;
}
```

Component Options.
```
unique
flagPrefix:my_prefix
eventTarget:removed
eventTarget:added
eventType:self
eventType:any
```

*eventPriority is not yet implemented in the language.

You can define a single component or multiple on the same line that share the same variables and contexts. If you define a component with a single variable and want to use the first style of defining it ```component name in context as "int";``` you wont be able to assign an entity index to it. 

You would have to write it like so, for example:
```
component name in context {
    value(entiyIndex:single) as int;
}
```

## Custom Entity Index's

```
index EntityIndexType(primary) in Core {
    method TheMethodToCall {
        p1 as "param1";
        p2 as "param2";
    }
}
```

```EntityIndexType``` needs to match the class name of the entity index. You can assign your entity as ``primary``` which will return a single entity. By default, not adding ```primary``` will return the hashset of entities.

```TheMethodToCall``` is what you would normally add your ```EntityIndexMethod``` attribute to (Entitas C# Runtime) but in Proton-Lang you need to specifiy the method and the parameters of that method. You are not limited to a single a method; you can have as many as your class supports.

```
index EntityIndexType(primary) in Core {
    method TheMethodToCall1 {
        p1 as "param1";
        p2 as "param2";
    }
    method TheMethodToCall2 {
        p1 as "param1";
    }
}
```