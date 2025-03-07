# BetterHook

BetterHook is a Git hook loader written in Go—similar to Husky for JavaScript. It allows users to keep their Git hook scripts (such as `pre-commit.sh` and `pre-push.sh`) in a dedicated `.betterhook` directory in their repository. The loader function then installs these scripts into the Git hooks folder so they run at the appropriate Git lifecycle events.

## Features

- **Automatic hook installation:** When invoked with a hook type (e.g. `pre-commit`), the loader checks for the corresponding script inside the `.betterhook` folder.
- **Error reporting:** If the required script does not exist, an error is returned.
- **Executable scripts:** The installed hook file is marked executable.

## Installation

Use `go get` to install the package:

```bash
go get github.com/yashGoyal40/BetterHook
