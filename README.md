# ExfilPro

![Go Version](https://img.shields.io/github/go-mod/go-version/golang/go) ![License](https://img.shields.io/github/license/yhk0/exfilpro-blue) ![Build Status](https://img.shields.io/badge/build-passing-brightgreen) ![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Linux-blue)


ExfilPro is a CLI tool designed for analyzing web content or files for potential sensitive information leaks. It supports filtering results by specific types of leaks and provides detailed output, including line numbers and file paths.
```bash
$ exfilpro -h

           __ _ _
  _____ __/ _(_) |_ __ _ _ ___
 / -_) \ /  _| | | '_ \ '_/ _ \
 \___/_\_\_| |_|_| .__/_| \___/
                 |_|

Version: dev (n/a) - Jabes Eduardo @yhk0 - 01/22/25

ExfilPro is a CLI tool designed for analyzing web content or files
for potential sensitive information leaks.
It supports filtering results by specific types of leaks and provides detailed output,
including line numbers and file paths.

Usage:
  exfilpro [flags]
  exfilpro [command]

Available Commands:
  file        Scans local files for sensitive data
  help        Help about any command
  web         Analyzes leaks in a URL

Flags:
  -h, --help   help for exfilpro

Use "exfilpro [command] --help" for more information about a command.
```
## Features

- Analyze web content for leaks (e.g., emails, API keys, passwords).
- Analyze files for metadata and potential security issues.
- Supports filtering results by type (e.g., emails, reCAPTCHA keys).
- Provides detailed output with line numbers and file paths.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/exfilpro.git
   cd exfilpro
   ```

2. Makefile (on Linux/Mac):
   ```bash
   make build-linux
   make build-macos
   ```
   Makefile (on Windows):
   ```bash
   make build-windows
   ```
3.  Enter the BIN directory
   ```bash
   cd bin
   ```

## Usage

### Web Mode
Analyze web content:

pattern: 
```bash
$ ./exfilpro web <url>
``` 
use flags
```bash
$ ./exfilpro web <url> -f [flags]
```

#### Flags
- `-f` : Filter results by data type (email, pass, API Key, reCAPTCHA)
- `-f email` 
- `-f pass` 
- `-f API Key` 
- `-f reCAPTCHA ` 
- `-h` : Display help for the `web` mode.

Example:
```bash
./exfilpro web https://example.com -f email
```

### File Mode
Analyze a file for metadata and leaks:
```bash
./exfilpro file -p <file_path>
```

#### Flags
- `-h` : Display help for the `file` mode.
- `-p` : Path of the file to be parsed (required)

Example:
```bash
./exfilpro file /path/to/file.txt
```

## Contributing

Feel free to contribute by submitting issues or pull requests. Contributions are welcome!

## License

This project is licensed under the Apache License. See the `LICENSE` file for details.

## Acknowledgments

Thanks to all contributors and the Go community for their support and resources.

