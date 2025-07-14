<div align="center">
  <a href="https://github.com/gkits/gorecipe.md">
<img src="assets/logo.png" alt="Logo" width="240" height="240">
  </a>
  <h1 align="center">gorecipe.md</h1>
  <p align="center">
    Scrape your favorie recipes and convert them to markdown
  </p>
</div>

## Getting started

### Installation

#### Go (recommended)

To install `gorecipe.md` with go run:

```sh
go install github.com/gkits/gorecipe.md@latest
```

#### From source

To install `gorecipe.md` from the source follow these steps.

1. Clone this repository

```sh
git clone --depth 1 https://github.com/gkits/gorecipe.md
```

2. Move into the newly cloned repository

```sh
cd gorecipe.md
```

3. Build and install the binary

```sh
go build && go install
```

4. Add your go install path to your `$PATH`

```sh
export PATH=$PATH:$HOME/go/bin
```

By default go installs the binary of your local user to `$HOME/go/bin` if you have changed this
default behavior adjust accordingly.

### Usage

#### Examples

##### Print recipe to stdout

```sh
gorecipe.md https://www.noracooks.com/vegan-pancakes
```

##### Write recipe to file

```sh
gorecipe.md -o pancakes.md https://www.noracooks.com/vegan-pancakes
```

#### Options

```
-f, --force         force markdown by ignoring missing recipe parts
-h, --help          help for gorecipe.md
    --hugo          add hugo headers
-o, --out string    path to output file
    --tmpl string   custom markdown template
-v, --version       version for gorecipe.md
```
