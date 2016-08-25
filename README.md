# BitBucket Archiver

_This project is still in the very early stage, bugs and problems are expected. Use with caution._

## Usage

### Pre-requirements

* Linux or macOS operating system
* Latest Golang installed and configured properly
* BitBucket account and an Oauth client
* SSH public key should be added to the account
* Git to be installed

When creating the Oauth client, you can specify any Return URI, for example http://localhost, which will not open but instead give you an "unavailable" page. This is fine because all you need is the URI parameter.

### Compile

```bash
// clone the repository
git clone ...

// compile
go build .

// or alternatively install
go install .

// create a directory "archive" to store repositories
mkdir archive
```

### Config

1. Run `./bitbucket-archiver config`
1. Copy the _Client Key_ and paste to the prompt, hit Enter
1. Copy the _Client Secret_ and paste to the prompt, hit Enter
1. Visit the URI showing on your terminal, and grant access
1. Look at the URI of the page you should see a URI parameter `code=[code]`, copy that _code_ and paste to the prompt, hit Enter.
1. Enter the team name / username you want to archive, seperate with Enter, and finally ends with an empty line.

That's it, now you should see a configuration file `client-config.json` stored on your current directory

### Archive

1. Run `./bitbucket-archive archive`

This will save repositories to `./archive/{username/team name}/{repository}`, and the program only works with git projects, and maybe in the future I'll add support for other version control manager.

## Roadmap

* Better output, currently it's quite messy
* More configuration, less hard-coding
* Unit Testing
