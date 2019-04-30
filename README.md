[![Build Status](https://travis-ci.org/SirMetathyst/go-proton.svg?branch=develop&dummy=unused)](https://travis-ci.org/SirMetathyst/go-proton)  [![Coverage Status](https://coveralls.io/repos/github/SirMetathyst/go-proton/badge.svg?branch=develop&kill_cache=1&dummy=unused)](https://coveralls.io/github/SirMetathyst/go-proton?branch=develop)  [![Go Report Card](https://goreportcard.com/badge/github.com/SirMetathyst/go-proton?dummy=unused)](https://goreportcard.com/report/github.com/SirMetathyst/go-proton)

# Overview
Proton is an [Entitas](https://github.com/sschmid/Entitas-CSharp) Code-Generator written in golang. Distributed as a single executable you are able to compose contexts, components and entity index's in a custom DSL([Proton-Lang](https://github.com/SirMetathyst/go-proton/language/)) inspired by [Entitas-Lang](https://github.com/sschmid/Entitas-CSharp/wiki/Tutorial-%E2%80%90-Entitas%E2%80%90lang).

For anyone who wants to use this or has questions/requests/feedback on the project you can find me on entitas's [gitter page](https://gitter.im/sschmid/Entitas-CSharp) under the handle @T2RKUS or @SirMetathyst. You can also create an issue here on github.


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

```
xxx@xxx-PC /x/x/x/x/x/proton> ./proton.exe help
proton is a tool to generate entitas source code

Usage:
  proton [command]

Available Commands:
  generate    generate entitas source code
  help        Help about any command

Flags:
  -h, --help   help for proton

Use "proton [command] --help" for more information about a command.
```

```
xxx@xxx-PC /x/x/x/x/x/proton> ./proton.exe help generate
you can enable/disable generators and post-processors with flags,
change the project path (which is based on excuting directory by default),
output folder where code is written to and keep the generator alive.
This will find any *.proton files in the project path, track them for changes
and re-generate the code.

Usage:
  proton generate [flags]

Flags:
      --CSharpComponentGenerator_E_1_4_2
      ...




      --CleanTargetDirectoryPostProcessor_C_1_4_2    (default true)
      --FileHeaderPostProcessor_C_1_4_2              (default true)
      --MergeContentPostProcessor_C_1_4_2            (default true)
      --PrintFileContentPostProcessor_C_1_4_2
      --PrintFilePostProcessor_C_1_4_2
      --WriteToDiskPostProcessor_C_1_4_2             (default true)
  -d, --daemonize                                   daemonize application
  -h, --help                                        help for generate
  -o, --output-folder string                        output folder (default "src-gen")
  -p, --project-path string                         project path (default "./")
```

# TODO

- Cleanup packages and write tests, coverage: at least 80%
- Rewrite documentation
- Currently all generators are off by default, create system to find newest version
and enable the latest generators/post-processors. Default would be set to latest but allow change of generator/post-processor through command line flag. Only one be allowed to be set per "set" under generator name. Also, move more work out of generators into code-generation package
- Add system syntax (cleanup, reactive, initialize etc) to DSL
- Add context divert syntax back e.g. "context PhysicsKit : Game" for project modules.
- Add "init" command for proton file generation. Tutorial of dsl syntax and project quickstart

## Compiling

```
git clone github.com/SirMetathyst/go-proton
cd go-proton/proton
make
```
