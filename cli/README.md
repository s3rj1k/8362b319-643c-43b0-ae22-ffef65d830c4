# Homework Application

NordVPN public API parser, a Go-based project.

## Building the Application

To build the application, you can use the provided Makefile. Here are the available commands:

- `make build`: Compiles the application and generates the binary `homework`.
- `make test`: Runs all the tests in the application.
- `make clean`: Cleans up any build artifacts.
- `make build-linux-amd64`: Builds the application specifically for Linux AMD64 architecture.

### Prerequisites

Ensure you have Go installed on your system. You can download and install Go from [the official Go website](https://golang.org/dl/).

### Building the Application

To build the application, run the following command in the root directory of the project:

```bash
make build
```

This will compile the source code and generate an executable file named `homework`.

### Application Flags

The application supports several command-line flags to customize its behavior:

    --config: Path to the configuration file. Default is config.ini.
        Usage: ./homework --config=path/to/config.ini

    --log_level: Set the log verbosity level. Default is debug.
        Usage: ./homework --log_level=info

    --log: Path to the log file. Logs to stdout if this is empty.
        Usage: ./homework --log=path/to/logfile

### Examples

Running the application with custom configuration and log level:

```bash
./homework --config=config.ini --log_level=info
```

Running the application with a specified log file:

```bash
./homework --log=/var/log/homework.log
```

For any additional information or help, refer to the application's help:

```bash
./homework --help
```
