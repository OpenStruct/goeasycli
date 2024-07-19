<div align="center">
  <img referrerpolicy="no-referrer-when-downgrade" src="https://avatars.githubusercontent.com/u/174039470?s=96&v=4" />
  <h1 align="center">Bootstrap your next Go project/library easily.
</h1>
</div>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

[![language][languages-shield]](https://golang.org)
![OS][os-shield]
[![Golang][releases]][repo-url]
![GitHub release date][release-date]
![GitHub last commit][last-commit]
[![Golang][project-download]][downloads-url]
![Contributors][contributors-shield]
[![License][license-shield]][license-url]
[![Free][free-for-dev]](#-license)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[releases]: https://img.shields.io/github/v/release/OpenStruct/goeasycli
[repo-url]: https://github.com/OpenStruct/goeasycli
[GitHub-Version-shield]: https://img.shields.io/github/v/release/OpenStruct/goeasycli
[project-download]: https://img.shields.io/github/downloads/OpenStruct/goeasycli/total
[downloads-url]: https://img.shields.io/github/downloads/OpenStruct/goeasycli/total
[contributors-shield]: https://img.shields.io/github/contributors/OpenStruct/goeasycli?color=7A3EF4
[license-shield]: https://img.shields.io/github/license/OpenStruct/goeasycli?color=9565F6
[codecov-shield]: https://img.shields.io/codecov/c/github/OpenStruct/goeasycli
[codecov-url]: https://codecov.io/gh/OpenStruct/goeasycli
[license-url]: https://opensource.org/licenses/MIT
[languages-shield]: https://img.shields.io/badge/language-go-blue.svg
[os-shield]: https://img.shields.io/badge/OS-linux%2C%20windows%2C%20macOS-0078D4
[release-date]: https://img.shields.io/github/release-date/OpenStruct/goeasycli
[last-commit]: https://img.shields.io/github/last-commit/OpenStruct/goeasycli
[free-for-dev]: https://img.shields.io/badge/free_for_non_commercial_use-brightgreen

---

## List of Contents

- [🔑 Prerequisites](#-prerequisites)
- [🤸 Quickstart](#-quickstart)
- [🪄 Simple and fast project setup](#-simple-and-fast-project-setup)
- [😺 API Documentation](#-api-documentation)
- [🏛️ Folder Structure](#-folder-structure)
- [🙁 Posible Issues](#-posible-issues)
- [📚 Creating a Library](#-creating-a-library)
- [🗺 Roadmap](#-roadmap)
- [🙌 Contributing and Community](#-contributing-and-community)
- [⭐️ Show Your Support](#-show-your-support)
- [📜 License](#-license)

---

## 🔑 Prerequisites

To use GoEasyCLI, you **need** to have the following tools installed on your system:

- [Go](https://golang.org/dl/)
- [Git](https://git-scm.com/downloads)

## 🤸 Quickstart

Install GoEasyCLI with the following command:

### Unix-like Systems (Linux, macOS)

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/OpenStruct/goeasycli/main/scripts/install.sh)"
```

### Windows

> **_NOTE:_** Please run Powershell in administrator mode.

```powershell
powershell -c "irm https://raw.githubusercontent.com/OpenStruct/goeasycli/main/scripts/install.ps1 | iex"
```

Check available options by running:

```bash
goeasycli --help
```

## 🪄 Simple and fast project setup

GoEasyCLI is a command-line interface that automates project creation and setup tasks for Go projects. It streamlines the process of setting up the project structure, installing dependencies, and configuring initial settings, making it easier to start with Go.

### Create a new project with specified options

```bash
goeasycli -p <project_name> -f {gin, fiber, echo}
```

Replace `<project_name>` with your desired project name.

**Supported frameworks:**

- [x] [Gin](https://gin-gonic.com)
- [x] [Fiber](https://gofiber.io)
- [x] [Echo](https://echo.labstack.com)

Here's an example:

```bash
goeasycli -p fafa_shop_api -f fiber
```

This command creates a new Go project named `fafa_shop_api` using the Fiber web framework.

The next step is to run your new project:

```bash
go run .
```

## 😺 API Documentation

After running the project, you can access the API documentation at `http://localhost:8080/swagger/index.html`.

More details on using the API documentation can be found in the [Swag documentation](https://github.com/swaggo/swag).

---

## 🗂️ Folder Structure

After using GoEasyCLI to create a project, the folder structure will be as follows:

```shell
└── 📁fafa_shop_api
    └── .env
    └── .gitignore
    └── 📁config
        └── config.go
    └── 📁controllers
        └── health.go
        └── user.go
    └── 📁database
        └── database.go
    └── go.mod
    └── go.sum
    └── 📁loggers
        └── sentry.go
        └── zap.go
    └── main.go
    └── 📁middlewares
        └── middlewares.go
    └── 📁models
        └── user.go
    └── 📁routers
        └── health.go
        └── routers.go
        └── user.go
    └── 📁seeds
        └── seed.go
    └── 📁structs
    └── 📁templates
        └── templates.go
    └── test_database.db
    └── 📁utils
        └── responses.go
        └── utils.go
```

This structure provides a solid foundation for your Go project, organized into common directories for configuration, controllers, database handling, logging, middleware, models, routes, templates, seeds, and utilities.

## 🏛️ Folder Structure Details

### config

The `config` folder contains the configuration settings for the project.

<details>

<summary>config.go</summary>

It initializes the configuration settings, loads the environment variables, and sets the default values for the configuration settings.

You can add more configuration settings to this file as needed.

</details>

### controllers

The `controllers` folder contains the controller files for the project. The controller files handle the project's business logic.

<details>

<summary>health.go</summary>

It contains the health check controller, which returns the status of the application.

</details>

<details>

<summary>user.go</summary>

It contains a sample user controller, which handles basic CRUD operations for users.

</details>

### database

The `database` folder contains the project's db configuration and connection settings.

<details>

<summary>database.go</summary>

It sets up the database connection and initializes the database.

It contains the database migration logic to create the required tables.
By default, the project uses a SQLite database. You can change the database settings in this file and the .env file to use a different database.

</details>

### loggers

The `loggers` folder contains the logger configuration settings for the project. The project uses two loggers: Zap and Sentry.

<details>

<summary>zap.go</summary>

It initializes the Zap logger, which logs messages to the console.

You can customize the logger to log messages to a file or a different output.

</details>

<details>

<summary>sentry.go</summary>

It initializes the Sentry logger, which sends error messages to the Sentry service.

You can configure the Sentry logger with your Sentry DSN to send error messages to your Sentry account.

</details>

### middlewares

The `middlewares` folder contains the middleware settings for the project.

### models

The `models` folder contains the model files for the project. The model files define the database schema for the project.

<details>
<summary>user.go</summary>

It contains the user model, which defines the sample user schema.

You can add more model files to define additional database schemas for the project.

</details>

### routes

The `routes` folder contains the route files for the project. The route files define the API routes for the project.

<details>
<summary>health.go</summary>

It contains the health check route, which returns the application status.

</details>

<details>
<summary>routers.go</summary>

It initializes the router and registers the API routes for the project.
Other route files are registered in this file.

</details>

<details>
<summary>user.go</summary>

It contains the sample user routes, which define the API routes for basic CRUD operations on users.

</details>

### structs

The `structs` folder contains the struct files for the project. The struct files define the data structures used in the project.

### seeds

The `seeds` folder contains the seed files for the project. The seed files populate the database with sample data.

<details>
<summary>seeds.go</summary>

It contains the seed logic to populate the database with sample users.
You can comment out the seed logic to avoid populating the database with sample data.

</details>

### templates

The `templates` folder contains the template files for the project. The template files define the HTML templates used in the project.
Sometimes, you may need to render HTML templates or send emails with HTML content. In such cases, you can use the template files in this folder.

### utils

The `utils` folder contains the utility files for the project. The utility files contain helper functions and utility functions used in the project.

<details>
<summary>responses.go</summary>

It contains the response utility functions to send JSON responses to the client.
You can customize the response functions to handle different response formats or error messages.

</details>

<details>
<summary>utils.go</summary>

It contains the project's utility functions.
Utility functions perform common tasks such as string manipulation and data validation.

</details>

## 🙁 Posible Issues

### Error: `Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub.`

```bash
2024/07/17 15:10:42 C:/WINDOWS/system32/justGo/database/database.go:66
[error] failed to initialize database, got error Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
failed to connect database
Database connection is nil. cannot run migration
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x28 pc=0x13e4992]
```

This error is caused by the project using `SQLite` as the default database.
To fix this issue, you need to:

1. Run this command in your `cmd` or `powershell`:

```bash
go env -w CGO_ENABLED =1
```

2.  Download a [tdm-gcc](https://jmeubank.github.io/tdm-gcc) for your architecture and install it.

> During installation, select `TDM-GCC Recommend, All Packages` from the Choose Components Screen. See Screenshot below
>
> ![TDM-GCC Setup](https://github.com/user-attachments/assets/44b06ab8-8f94-4b21-82bb-ba4b0c1f1cec)

## 📚 Creating a Library

GoEasyCLI can also be used to create a Go library. To create a new library, use the following command:

```bash
goeasycli -l <library_name> -r repo_url
```

Replace `<library_name>` with your desired library name and `repo_url` with the repository's URL.

Here's an example:

```bash
goeasycli -l my_awesome_fafa_lib -r https://github.com/heygoeasycli/my_awesome_fafa_lib
```

This command creates a new Go library named `my_awesome_fafa_lib` with the specified repository URL.

The folder structure for the library will be as follows:

```shell
└── 📁my_awesome_fafa_lib
    └── .github
        └── goeasycli_tag.yml
    └── config
        └── config.go
    └── database
        └── database.go
    └── email
        └── emails.go
    └── loggers
        └── zap.go
    └── go.mod
    └── go.sum
```

## 🗺 Roadmap

GoEasyCLI is being built in public. The [roadmap](https://github.com/OpenStruct/goeasycli/issues) is a regularly updated source of truth for the GoEasyCLI community to understand where the product is going in the short, medium, and long term.

GoEasyCLI is managed by [Open Struct](https://github.com/OpenStruct), a group with the aim of easing the burden of engineers. On our GitHub repo, you can directly influence the roadmap by [Creating an issue](https://github.com/OpenStruct/goeasycli/issues/new/choose).

## 🙌 Contributing and Community

We would love to develop GoEasyCLI together with our community! The best way to get started is to select any issue from the [repo](https://github.com/OpenStruct/goeasycli/issues) and open a Pull Request!

### Contributors

<a href="https://github.com/OpenStruct/goeasycli/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=OpenStruct/goeasycli" />
</a>

## ⭐️ Show Your Support

If you find GoEasyCLI helpful or interesting, please consider giving us a star on GitHub. Your support helps promote the project and lets others know it's worth checking out.

Thank you for being so supportive! 🌟

[![Star this project](https://img.shields.io/github/stars/OpenStruct/goeasycli?style=social)](github.com/OpenStruct/goeasycli/stargazers)

## 📜 License

GoEasyCLI is distributed under the terms of the MIT License.
A complete version of the license is available in the [LICENSE](LICENSE) file in
this repository. Any contribution made to this project will be licensed under the MIT License.
