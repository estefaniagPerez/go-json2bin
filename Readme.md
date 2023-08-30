
[go_ref]: https://go.dev/
[vscode_ref]: https://code.visualstudio.com/
[go_install_ref]: https://go.dev/doc/install

# Golang - JSON To Binary - 

![Go/Golang](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

> _first version: estegp - 31/08/2023_ </br>

## Introduction
Welcome to Json2Binary! This project, developed in [Go][go_ref] (or Golang), is dedicated to transforming a demo JSON configuration file into binary files seamlessly. The primary motivation behind this utility is to demostrate the process of binary file generation from human-readable JSON files, eliminating the complexities associated with manual conversions and potential errors.


<!--
![maintenance-status](https://img.shields.io/badge/maintenance-actively--developed-brightgreen.svg)-->

## Features

- Efficient Transformation: Convert JSON files to binary rapidly without sacrificing accuracy.
- Ease of Use: With a straightforward command-line interface, you can generate binary files in seconds.
- Golang Optimized: Benefit from the speed and performance optimizations that Golang offers.
- Extensible: As open-source software, you can contribute and tailor the tool to fit your exact needs.

## Getting Started
To manage the [Rust][rust_ref] installation use _rustup_ toolchain manager. To install this tool the installation script needs to be downloaded and run from the terminal.

```properties
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Besides the toolchain manager, this script will also install the [Rust][rust_ref] compiler _rustc_ and a build system named _cargo_. 


## Using Rust Compiler
The simplest way to create a rust application is to create a file named _main.rs_. Inside the file create a _main_ function and code the functionality that you want the program to do.


For example, a simple application that prints a _Hello World_ message.
```rust
fn main() {
    println!("Hello, world!");
}
```

This program can be compiled by running the [Rust][rust_ref] compiler
```Shell
rustc main.rs 
```


## Using Cargo Build System
Rust provides a build system manager named [Cargo][cargo_ref] that can be used to initializate proyects, manage dependencies, and compile and package rust applications.

### Create a Project
To create a new proyect with cargo run the following comand inside the folder where the project will be located.

```Shell
cargo new hello_cargo
cd hello_cargo
```

### Build a Project
The project can be build and then run by using the following command.
```Shell
cargo run main.rsc
```

### Build For Release
The project can be build and then run by using the following command.
```Shell
cargo build --release main.rsc
```


## Visual Strudio Code and Debugging
Rust can be developed using [VSCode][vscode_ref] editor. It is recomendable to use the following extensión  &#91; [_ref_][vscode_rust_ref] &#93;:

* rust-analyzer
* CodeLLDB

In order to run the debugger a launch.json file needs to be created on the .vscode folder on the project
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "type": "lldb",
            "request": "launch",
            "name": "Debug",
            "program": "${workspaceRoot}/project_name/target/debug/${workspaceRootFolderName}",
            "args": [],
            "cwd": "${workspaceRoot}/project_name/target/debug/",
            "sourceLanguages": ["rust"]
        }
    ]
}
```
<br/>

> **_NOTE:_**  _Any rust developed application can be debugged using gdb debugger._