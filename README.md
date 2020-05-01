# Pork

The `Pork` program, it's a `CLI project` that allows us to `fork`, `PRs`, `search`, `read` docs, `clone`, and many other functionalities from GitHub public accounts. This tool makes things easier for developers in order to find repositories and kind of automate development and integration processes

## Dependencies

There's a lot of them

Check the `glide.yaml` file for it. For this project, we use four main dependencies.

1. The `NAP` library (explained below) as the HTTP wrapper encapsulation for Rest API's interactions.
2. `Viper` for handling configurations options for files. Check the library [here](https://github.com/spf13/viper)
3. `Cobra` is used for CLI apps generation. Check the library [here])(https://github.com/spf13/cobra)
4. `gopkg.in/src-d/go-git.v4` this one is a Git implementation for Go. Check it [here](https://gopkg.in/src-d/go-git.v4)

### NAP Library

[NAP](https://github.com/davidlares/nap) is a library that wraps up and interacts with the `HTTP`'s Golang built-in package.

### Creating GitHub API KEY

This process is quite easy, and it's very important because is the main window for GitHub interaction.

You should go to your GitHub `Settings` profile page first, then, check the `Developer Settings`, then move to the `Personal Access tokens` section and `Generate new Token` with `public_repo` and `read:user`. Grab the generated `token` and paste it the `Pork.yaml`

### Filling pork.yaml file

The `Pork.yaml` (please rename it), contains the defaults for GitHub interaction, actually it holds two variables, the `location` (for repo interaction, things like: downloads and clones), and the main one which is the `token`.

This particular file serves as a `secret` file for holding our access credentials and environment setups.

You should rename the `pork_example.yaml` to `pork.yaml` and change the values for your own. It will be called many times by the `os` Golang package to

## Usage

I assume you already have Go installed, with the `GOPATH` set.

Clone this repository inside your `src` directory and install it using the following command

`go install ./cmd/...`

Then, inside `cmd/pork` directory just, `pork []`

For help: `pork -help`

## Commands

Soon.

## Running tests

For doing this, check the `go test -v` command.

This will check and evaluate all the test files for the project, referenced by `_test.go`

Must of them will crash if you don't change the default values set in there. The current `token` value for many of the test files are previous API Keys that were used during development phases.

## Credits

 - [David E Lares](https://twitter.com/davidlares3)

## License

 - [MIT](https://opensource.org/licenses/MIT)
