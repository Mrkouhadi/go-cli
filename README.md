# CLI Usage Details

This is a command-line interface (CLI) tool built in Golang. It provides users with various commands to retrieve system information such as CPU usage, RAM usage, disk usage, network details, and more.

## Installation

To install the CLI tool, you can use the following commands:

```bash
go get -u github.com/mrkouhadi/go-cli
```

Make sure your Go environment is properly set up.

## Usage

Once installed, you can use the CLI tool with the following commands:

### App Version

This application follows Semantic Versioning. The current version is `1.0.0`.

To check the version of your installed go-cli

```bash
go-cli --version
```

### IP Addresses

To retrieve IP addresses:

```bash
go-cli net ip -v ipv4
```

```bash
go-cli net ip -v ipv6
```

### Network Details

To retrieve network details:

```bash
go-cli net info
```

### Ping a URL

To ping a URL:

```bash
go-cli net ping -u bing.com
```

### CPU Usage

To retrieve CPU usage details:

```bash
go-cli info cpu
```

### RAM Usage

To retrieve RAM usage details:

```bash
go-cli info ram
```

### Disk Usage

To retrieve disk usage details:

```bash
go-cli info disk
```

### GPU Usage

To retrieve GPU usage details:

```bash
go-cli info gpu
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the [MIT License](LICENSE).
