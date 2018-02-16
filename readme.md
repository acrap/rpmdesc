# rpmdesc
Tool to get information about rpm packet by name (os and architecture are optional)

## Description
With this tool you can get information about files inside the rpm package by name.
This tool uses [rpmfind](https://rpmfind.net) to get information about packet.
It's just a web scrapper for rpmfind.net

## What kind of information can i get?
__rpmdesc__ tool gives you full package name, homepage and list of files in packet (most important feature).

## Usage

rpmdesc [global options] command [command options] [arguments...]

#### Commands:
     help, h  Shows a list of commands or help for one command

#### Global options:
   --name value   name of rpm packet
   --os value     OS of rpm packet
   --arch value   architecture of rpm packet
   --help, -h     show help
   --version, -v  print the version
