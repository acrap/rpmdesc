# rpmdesc
Tool to get information about rpm packet by name (os and architecture are optional)

## Description
With this tool you can get information about files inside the rpm package by name.
This tool uses [rpmfind](https://rpmfind.net) to get information about packet.
It's just a web scrapper for rpmfind.net

## Dependencies 

[goquery](https://www.github.com/PuerkitoBio/goquery)

[urvafe/cli](https://www.github.com/urfave/cli)

## What kind of information can i get?
__rpmdesc__ tool gives you full package name, homepage and list of files in packet (most important feature).

# Usage 

### Usage:
   rpmdesc [global options] command [command options] [arguments...]

### Version:
   0.1

### Description:
   rpmdesc is a tool that gives you full package name, homepage and list of files in RPM packet

### commands:
     help, h  Shows a list of commands or help for one command

### Global options:
   _--nofilelist_   remove file list from the output
   
   _--nopname_      remove picked packet name from the output
   
   _--license_      add license info to output
   
   _--homepage_     add homepage info to output
   
   _--name value_   name of rpm packet
   
   _--os value_     OS of rpm packet
   
   _--arch value_   architecture of rpm packet
   
   _--help, -h_     show help
   
   --version, -v  print the version
