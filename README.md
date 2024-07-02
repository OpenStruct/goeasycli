# Welcome to GoEasyCLI

![Release](https://img.shields.io/github/v/release/OpenStruct/goeasycli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Github forks](https://img.shields.io/github/forks/OpenStruct/goeasycli)

Starting a new Go project involves setting up the project structure, installing dependencies, and configuring initial settings, which can be time-consuming. **GoEasyCLI** streamlines this process with a command-line interface, automating project creation and setup tasks. Whether you're new to Go or looking to speed up your workflow, GoEasyCLI helps you get started quickly.

## Installation

Install GoEasyCLI with the following command:

### Unix-like Systems (Linux, macOS)

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/OpenStruct/goeasycli/main/scripts/install.sh)"
```

### windows

```powershell
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/OpenStruct/goeasycli/main/scripts/install.ps1" -OutFile "$env:TEMP\install.ps1"; & "$env:TEMP\install.ps1"
```

## Usage

To begin, use the following command to explore available options:

```bash
goeasycli --help
```

## Commands

### Create a New Project

Create a new Go project with specified options:

```bash
goeasycli -p <project_name> -f {gin, fiber, echo}
```

> Replace `<project_name>` with your desired project name.

**The supported frameworks are:**
- [x] [Gin](https://gin-gonic.com)
- [x] [Fiber](https://gofiber.io)
- [x] [Echo](https://echo.labstack.com)

Example:

```bash
goeasycli -p fafa_shop_api -f fiber
```

This command creates a new Go project named `fafa_shop_api` using the Fiber web framework.

---

## Folder Structure

After using GoEasyCLI to create a project, the folder structure will be as follows:

```
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

This structure provides a solid foundation for your Go project, organized into common directories for configuration, controllers, database handling, logging, middleware, models, routes, templates, seeds and utilities.

### config
The `config` folder contains the configuration settings for the project.

<details>

<summary>config.go</summary>

- <kbd> It initializes the configuration settings.</kbd>
- <kbd> It loads the environment variables and sets the default values for the configuration settings.</kbd>
- <kbd> You can add more configuration settings to this file as needed.</kbd>
</details>

### controllers
The `controllers` folder contains the controller files for the project. The controller files handle the business logic for the project.

<details>

<summary>health.go</summary>

- <kbd> It contains the health check controller, which returns the status of the application.</kbd>

</details>

<details>

<summary>user.go</summary>

- <kbd> It contains a sample user controller, which handles basic CRUD operations for users.</kbd>

</details>

### database
The `database` folder contains the database configuration and connection settings for the project.

<details>

<summary>database.go</summary>

- <kbd> It sets up the database connection and initializes the database.</kbd>
- <kbd> It contains the database migration logic to create the required tables.</kbd>
- <kbd> By default, the project uses a SQLite database. You can change the database settings in this file, as well as the .env file to use a different database.</kbd>

</details>

### loggers
The `loggers` folder contains the logger configuration settings for the project. The project uses two loggers: Zap and Sentry.

<details>

<summary>zap.go</summary>

- <kbd> It initializes the Zap logger, which logs messages to the console.</kbd>
- <kbd> You can customize the logger to log messages to a file or a different output.</kbd>

</details>

<details>

<summary>sentry.go</summary>

- <kbd> It initializes the Sentry logger, which sends error messages to the Sentry service.</kbd>
- <kbd> You can configure the Sentry logger with your Sentry DSN to send error messages to your Sentry account.</kbd>

</details>

### middlewares
The `middlewares` folder contains the middleware settings for the project.

### models
The `models` folder contains the model files for the project. The model files define the database schema for the project.

<details>
<summary>user.go</summary>

- <kbd> It contains the user model, which defines the sample user schema.</kbd>
- <kbd> You can add more model files to define additional database schemas for the project.</kbd>
</details>

### routes
The `routes` folder contains the route files for the project. The route files define the API routes for the project.

<details>
<summary>health.go</summary>

- <kbd> It contains the health check route, which returns the status of the application.</kbd>
</details>

<details>
<summary>routers.go</summary>

- <kbd> It initializes the router and registers the API routes for the project.</kbd>
- <kbd> Other route files are registered in this file.</kbd>
</details>

<details>
<summary>user.go</summary>

- <kbd> It contains the sample user routes, which define the API routes for basic CRUD operations on users.</kbd>
</details>

### structs
The `structs` folder contains the struct files for the project. The struct files define the data structures used in the project.

### seeds
The `seeds` folder contains the seed files for the project. The seed files populate the database with sample data.

<details>
<summary>seeds.go</summary>

- <kbd> It contains the seed logic to populate the database with sample user.</kbd>
- <kbd> You can comment out the seed logic if you don't want to populate the database with sample data.</kbd>
</details>

### templates
The `templates` folder contains the template files for the project. The template files define the HTML templates used in the project.
Sometimes, you may need to render HTML templates, or send emails with HTML content. In such cases, you can use the template files in this folder.

### utils
The `utils` folder contains the utility files for the project. The utility files contain helper functions and utility functions used in the project.

<details>
<summary>responses.go</summary>

- <kbd> It contains the response utility functions to send JSON responses to the client.</kbd>
- <kbd> You can customize the response functions to handle different response formats or error messages.</kbd>
</details>

<details>
<summary>utils.go</summary>

- <kbd> It contains the utility functions for the project.</kbd>
- <kbd> Utility functions are used to perform common tasks such as string manipulation, and data validation.</kbd>
</details>

---

## Contributors

<a href="https://github.com/OpenStruct/goeasycli/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=OpenStruct/goeasycli" />
</a>

---

## License

GoEasyCLI is released under the [MIT license](LICENSE).
