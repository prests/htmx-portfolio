# Portfolio

## Installation

### Setup PNPM

```sh
command -v brew &> /dev/null || /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)" # Setup brew
command -v nvm &> /dev/null || brew install nvm # Setup NVM
nvm install
nvm use
npm i -g pnpm
```

### Setup Node Deps

```sh
pnpm i
```

### Setup Golang Deps

```sh
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
```
