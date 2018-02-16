#rpmdesc
Tool to get information about rpm packet by name

##Description
With this tool you can get information about files inside the rpm package by name.
This tool uses https://rpmfind.net to get information about packet.

##What kind of information can i get?
rpmdesc gives you full package name, homepage and list of files in packet (most important feature).

##Usage
Use:
_rpmdesc -name packagename_
or
_rpmdesc -name part_of_packagename_

For example:
_rpmdesc -name cron_

Also, you can start tool without any arguments and input name after greeting.

You can use -arch option to input custom architecture. By default tool gives you first architecture from the result result.
But be sure that you input full architecture name. Use