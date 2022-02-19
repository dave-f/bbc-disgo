# bbc-disgo
A simple 6502 disassembler written in Go

# Use
The program is intended to be a cheap and cheerful 6502 disassembler for getting output in a format I wanted, and also as an exercise in writing some Go.

Invoke the program by passing the filename of the control file:

````
go run disgo.go <filename>
````

The program uses a "control file" which specifies which file to process as well as a base address and which parts of the file are data.  It also specifies the output file and by default the program will look at the output file first to parse comments out in order to re-apply them.

# Command line arguments

The tool has one required argument, the filename of the control file (see below).

However there are a few options too:

````
--wipe     Don't parse comments out of the target file, ie. wipe all comments (default off)
--column n Look for comment indicator (currently ';') at specified column (default 28)
--dry      Write disassembly to stdout rather than the target file (default off)
````

# Control file commands

The control file currently supports 4 commands:

````
load <string>           Specifies the input file (ie. a 6502 binary)
base <address>          Set the base address
data <address>,<length> Mark the range at address,length as data
save <string>           Set the output file
````

Numbers (ie `address` and `length`) can be decimal or hex (prefixed with `0x`).  See `example-control-file` for an example of a control file.

# Credits
I pulled out the 6502 definitions from https://github.com/hoglet67/AtomBusMon

6502 reference from https://www.masswerk.at/6502/6502_instruction_set.html
