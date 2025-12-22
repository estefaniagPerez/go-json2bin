
[go_ref]: https://go.dev/
[vscode_ref]: https://code.visualstudio.com/
[vscode_go_ref]:https://code.visualstudio.com/docs/languages/go
[go_install_ref]: https://go.dev/doc/install

# Golang - JSON To Binary - 

![Go/Golang](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

> _first version: estegp - 31/08/2023_ </br>
> _last edition: estefaniagPerez - 08/12/2025 _

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

## Build Project
After this, you can build the project using the Makefile by running the following command in the terminal from the root folder:
```Shell
make
```
This will generate an executable for the current OS.


## Run
The executable is run by passing 3 parameters:
1. The type of operation.
2. The path to the configuration JSON file.
3. The path of the binary file to be generated. 

There are two times of operations:
- gen_bin: for generating a binary file from the data of a JSON configuration file.
```Shell
  ./json2bin-parser --op=gen_bin --json=./assets/Conf_out.json --bin=./assets/Conf.dat
```
- parse_file: this operation gets a binary file, parses the data and writes it into a readable JSON file.
```Shell
  ./json2bin-parser --op=parse_file --json=./assets/Conf_out.json --bin=./assets/Conf_out.dat
```
