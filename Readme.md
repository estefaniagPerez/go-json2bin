
[go_ref]: https://go.dev/
[vscode_ref]: https://code.visualstudio.com/
[vscode_go_ref]:https://code.visualstudio.com/docs/languages/go
[go_install_ref]: https://go.dev/doc/install

# Golang - JSON To Binary - 

![Go/Golang](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

> _first version: estegp - 31/08/2023_ </br>
> _last edition: estefaniagPerez - 02/02/2025 (This repository was moved from my old repository)_

## Introduction
Welcome to Json2Binary! This project, developed in [Go/Golang][go_ref], converts demo JSON configuration files into binary files seamlessly. The primary motivation behind this utility is to demonstrate the process of binary file generation from human-readable JSON files, eliminating the complexities associated with manual conversions and potential errors.
 

<!--
![maintenance-status](https://img.shields.io/badge/maintenance-actively--developed-brightgreen.svg)-->

## Features

- Efficient Transformation: Convert JSON files to binary rapidly without sacrificing accuracy.
- Ease of Use: With a straightforward command-line interface, you can generate binary files in seconds.
- Golang Optimized: Benefit from the speed and performance optimizations that Golang offers.
- Extensible: As open-source software, you can contribute and tailor the tool to fit your exact needs.

## Installation
This application has been developed in [Go][go_ref], so the first thing you need to do is install Go language tools. The best way to do it is by following the installation guide at the [Go][go_ref] documentation site.
<pre>
<a href="https://go.dev/doc/install">https://go.dev/doc/install</a>
</pre>

## Build
### Initialize Module
To build the project, it is necessary to initialize the project. Open a terminal and go to the folder that contains the project, then run the following command.

```Shell
go mod init Json2Bin.go
```
### Build Project
After this, you can build the project by running the go build command
```Shell
go build .
```
This will generate an executable for the current OS.

### Build Project for Other OS
If you want to generate executables for other platforms you can use the GOOS parameter. For example, if you want to build for Windows 64 bits from Linux you can run the following commad:

```Shell
env GOOS=windows GOARCH=amd64 go build .
```
## Run
The executable is run by passing 3 parameters:
1. The type of operation.
2. The path to the configuration JSON file.
3. The path of the binary file to be generated. 

There are two times of operations:
- --gen_bin: for generating a binary file from the data of a JSON configuration file.
```Shell
  ./Json2Bin --gen_bin Conf.json Conf.dat
```
- --parse_file: this operation gets a binary file, parses the data and writes it into a readable JSON file.
```Shell
  ./Json2Bin --parse_file Conf.json Conf.dat
```
## Debbug with VSCode
This project can be used in [VSCode][vscode_ref]. If you haven't set up [Go][go_ref] support in your IDE, you can follow the [VSCode][vscode_go_ref] documentation to install the [Go][go_ref] plugin.
<pre>
<a href="https://code.visualstudio.com/docs/languages/go">https://code.visualstudio.com/docs/languages/go</a>
</pre>

Once the plugin is configured, you will need to configure the GOROOT variable in [VSCode][vscode_ref], go to File -> Preferences -> Settings and search for GOROOT. On the GOROOT setting, click Edit Settings and add the following line to the settings.json:

```json
  "go.goroot": "/home/userno/.go/"
```

If the project is not initialized with the "mod init" you can configure the GO111MODULE parameter to avoid an error when running the application.
```json
    "go.toolsEnvVars": {
        "GO111MODULE":"auto"
    }
```

Finally, to start debugging the application, press F5 on the keyboard, and the project will start debugging.

Also, if you want to change the arguments used to launch the project, go to the .vscode folder in the project explorer, open the launch.json file, and modify the parameters in the "args" property.

```json
    "args": ["--gen_bin", "Conf.json", "Conf.dat"]
```


