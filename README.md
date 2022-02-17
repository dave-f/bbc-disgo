# bbc-disgo
A simple 6502 disassembler written in Go

# Use
The program is intended to be a cheap and cheerful 6502 disassembler for getting output in a format I wanted, and also as an exercise in writing some Go.

Invoke the program by passing the filename of the control file:

````
go run disgo.go <filename>
````

The control file specifies which file to process as well as a base address and which parts of the file are code.

It supports the following commands:

````
file <string> - Specifies the input file (ie. a 6502 binary)
base <address> - Set the base address
code <address>,<length> - Mark the range at address,length as code
````

Numbers (ie `address` and `length`) can be decimal or hex (prefixed with `0x`).  See `example-control-file` for an example of these.

# Credits
I pulled out the 6502 definitions from https://github.com/hoglet67/AtomBusMon

6502 reference from https://www.masswerk.at/6502/6502_instruction_set.html
