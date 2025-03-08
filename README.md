# BetterHook

BetterHook is a Git hook loader written in Go—similar to Husky for JavaScript. It allows users to manage their Git hook scripts (such as `pre-commit.sh` and `pre-push.sh`) in a dedicated `.betterhook` directory within their repository. The loader function automatically installs these scripts into the Git hooks folder, ensuring they execute at the appropriate Git lifecycle events.

## Features

- 🚀 **Automatic Hook Installation** – Detects and installs hook scripts from the `.betterhook` directory.
- ⚠️ **Error Reporting** – Notifies users if the required script is missing.
- 🔧 **Executable Scripts** – Ensures installed hook files have executable permissions.

## Installation

Install BetterHook using `go get`:

```bash
go get github.com/yashGoyal40/BetterHook
```

## Usage

1. Create a `.betterhook` directory in the root of your repository.
2. Add your Git hook scripts (e.g., `pre-commit.sh`, `pre-push.sh`) inside `.betterhook`.
3. Reference example scripts:
   - **Pre-commit hook:** [pre-commit.sh](https://github.com/yashGoyal40/BetterHook/blob/main/example/.betterhook/pre-commit.sh)
   - **Pre-push hook:** [pre-push.sh](https://github.com/yashGoyal40/BetterHook/blob/main/example/.betterhook/pre-push.sh)
4. Use BetterHook in your Go code:

```go
betterhook.SyncHook("pre-commit")
betterhook.SyncHook("pre-push")
```

## Contributing

Feel free to open issues or submit pull requests to improve BetterHook. Contributions are always welcome! 🚀

## License

BetterHook is open-source and available under the MIT License.

