# bbc-disgo
A simple 6502 disassembler written in Go

# Use
The program is intended to be a cheap and cheerful 6502 disassembler for getting output in a format I wanted, and also as an exercise in writing some Go.

Invoke the program by passing a filename and a base address:

````
go run disgo.go <filename> <address>
````

The file is disassembled to standard output, taking `<address>` as its base address.

# Credits
I pulled out the 6502 definitions from https://github.com/hoglet67/AtomBusMon

6502 reference from https://www.masswerk.at/6502/6502_instruction_set.html
