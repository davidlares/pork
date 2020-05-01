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

### Clone functionality

The `clone` downloads a copy of a selecting repository to your local machine. The location is given by the `location` key set in the `pork.yaml`.

Internally expects 2 flags or sub-commands

1. The `--create` flag: will create a reference if not found
2. The `--ref` flag: will specify the `branch name` for the downloaded repository

Both of them are required

It's done like this: `pork clone davidlares/arp-spoofing --create --ref master`

### Fork functionality

The `fork` command actually generates and reference a copy of a target repository into your Github account. This one uses the `/repos/<OWNER>/<REPOSITORY>/forks` API Endpoint.

This is done like this: `pork fork yeyintminthuhtut/Awesome-Red-Teaming`

### Docs functionality

The `docs` command asks for `README` files using the `/repos/<OWNER>/<REPOSITORY>/readme` GitHub API endpoint. You will need to call this option like this

`pork docs davidlares/arp-spoofing`

### PRs functionalities

We are also able to generate `pull requests` to any public repository. Internally creates a `POST` request to the `/repos/<OWNER>/<PROJECT>/pulls` GitHub API endpoint with certain values that points to the `destination repository`, the `source repository` a `title` and a `message`

Here's an example

`pork pullrequest -d davidlares/arp-spoofing -t "This is my title" -m "This is my message" -s "davidlares:changes"`

The `source` repository it's a combination of the `owner:branch`

### Search functionality

The `search` command will actually send a `query` parameter to the `/search/repositories?q=<YOUR-SEARCH-CRITERIA>` GitHub API endpoint. This command will return a bunch of public repositories with the `owner/repository` format (actually it's a Go Slice element).

Here's an example of how to use it.

`pork search infosec` or `pork search topic:infosec`. Basically, anything that can be set up as a `query` criteria

## Running tests

For doing this, check the `go test -v` command.

This will check and evaluate all the test files for the project, referenced by `_test.go`

Must of them will crash if you don't change the default values set in there. The current `token` value for many of the test files are previous API Keys that were used during development phases.

## Credits

 - [David E Lares](https://twitter.com/davidlares3)

## License

 - [MIT](https://opensource.org/licenses/MIT)
