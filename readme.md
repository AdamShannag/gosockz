# gosockz - A Simple WebSocket Library for Go

`gosockz` is a lightweight and minimal WebSocket library for Go that provides a flexible and event-driven approach to managing WebSocket clients. The library is designed with simplicity in mind, allowing users to easily integrate it into their applications.

## Features

- Event-driven approach for handling WebSocket requests.
- Built with Gorilla WebSocket for reliable WebSocket connections.
- Use of interfaces for the manager, client, and router for easy customization.

## Installation

To use `gosockz` in your Go project, you can use the `go get` command:

```bash
go get -u github.com/AdamShannag/gosockz
```

## Example Usage

Check out the [examples](./examples/) folder that demonstrates the usage of gosockz.

## Customization

gosockz is built from the ground up with interfaces, allowing users to provide their own implementations for the manager, client, and router. This flexibility enables seamless integration into various projects with different requirements.

To create custom implementations, you can implement the following interfaces:

- [Manager](./types/manager.go): Manages WebSocket clients and events.
- [Client](./types/client.go): Represents a WebSocket client.
- [EventRouter](./types/router.go): Routes incoming WebSocket events to appropriate handlers.

## Contributing

Feel free to contribute to the development of gosockz by opening issues or submitting pull requests on GitHub.

## Acknowledgments

This project was inspired and motivated by the blog post [Mastering WebSockets with Go](https://programmingpercy.tech/blog/mastering-websockets-with-go/) by Programming Percy.

## License

This project is licensed under the MIT License - see the [LICENSE](./license.md) file for details.
