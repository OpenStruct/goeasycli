# Welcome to GoEasyCLI

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
goeasycli -p <project_name> -f {gin, fiber}
```

- Replace <project_name> with your desired project name.
- Use -f to select a web framework (gin or fiber). If not specified, the default framework is gin.

Example:

```bash
goeasycli -p fafa_shop_api -f fiber
```

This command creates a new Go project named fafa_shop_api using the Fiber web framework.

## Folder Structure

After using GoEasyCLI to create a project, the folder structure will be as follows:

```
fafa_shop_api
├── config
│   └── config.go
├── controllers
│   └── health.go
|   └── user.go
├── database
│   └── database.go
├── loggers
│   └── sentry.go
|   └── zap.go
├── middlewares
│   └── middlewares.go
├── models
│   └── user.go
├── routes
│   └── health.go
|   └── routers.go
|   └── user.go
├── structs
├── seeds
|   └── seeds.go
├── templates
|   └── templates.go
├── utils
│   └── responses.go
|   └── utils.go
```

This structure provides a solid foundation for your Go project, organized into common directories for configuration, controllers, database handling, logging, middleware, models, routes, templates, seeds and utilities.
