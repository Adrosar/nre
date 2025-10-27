# Node Runtime Environment

An extension for:

* [NVM for Windows](https://github.com/coreybutler/nvm-windows)
* [NVM for Linux](https://github.com/nvm-sh/nvm)
* [NVM for MacOS](https://github.com/nvm-sh/nvm) (work in progress)

This utility ensures that any CMD command is executed within an environment configured for the selected **NodeJS** version.



## Requirements (Windows)

Before installation, ensure the following software is installed:

* [NVM for Windows](https://github.com/coreybutler/nvm-windows)



## Installation (Windows)

If you have GoLang version 1.20.4 or higher, run the command:

```
go install github.com/Adrosar/nre@latest
```

... or download the compiled version:

1. Download and extract the archive `nre-0.3.0-win64.zip`.
2. Move the `nre.exe` binary to the `%NVM_HOME%` directory.
3. Open a terminal (console) and run the `nre` command to confirm that the application is functioning correctly.



## Requirements (Linux)

Before installation, ensure the following software is installed:

* [NVM for Linux](https://github.com/nvm-sh/nvm)



## Installation (Linux)

If you have GoLang version 1.20.4 or higher, run the command:

```
go install github.com/Adrosar/nre@latest
```

... or download the compiled version:

1. Download and extract the archive `nre-0.3.0-linux64.zip`.
2. Move the `nre` binary to the `$NVM_DIR` directory.
3. Set execute permissions: `chmod +x nre`
4. Open a terminal (console) and run the `nre` command to confirm that the application is functioning correctly.



## Usage

```bash
nre <NODE_VERSION> <COMMAND>
```

* `NODE_VERSION` - The NodeJS version number managed by **NVM**.
* `COMMAND` - Any command valid for your operating system or terminal.

### Examples

* `nre 8 npm run start`: Executes `npm run start` using NodeJS version 8.
* `nre 10 node -v`: Executes `node -v` using NodeJS version 10.
* `nre 12 git commit -m "Update"`: Executes `git commit -m "Update"` using NodeJS version 12.



## Development (Windows)

To build and run NRE from source, ensure the following prerequisites are installed:

* [Go 1.20.4 for Windows](https://go.dev/dl/go1.20.4.windows-amd64.msi)
* [NVM for Windows](https://github.com/coreybutler/nvm-windows)

### Running from Source

```
go run main.go
```

### Building the Binary

To generate the `nre.exe` executable for Windows:

```
 go build -o nre.exe main.go
```



## Development (Linux)

To build and run NRE from source, ensure the following prerequisites are installed:

* Go 1.20.4 for Linux
* [NVM for Linux](https://github.com/nvm-sh/nvm)

### Running from Source

```
go run main.go
```

### Building the Binary

To generate the `nre` executable for Linux:

```
 go build -o nre main.go
```



## Contributing

This project was initially created for personal use, but it’s open to the community.  
Contributions are welcome — feel free to:

- Open an issue to report bugs or suggest new features  
- Submit pull requests with improvements  
- Share ideas for future development  

Your feedback and participation are highly appreciated!



## Author

Adrian Gargula
