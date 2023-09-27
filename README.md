# CS3339

Description: In this project you will create a simplified ARM v8 disassembler. Your
disassembler will read in a binary file containing an ARM program in machine code
and output the assembly code that generated the machine code. The instructions
will follow simplifications outlined in Zybook chapter X. You will use this code in
project 2 do a good job and test it completely. I will test your code completely and
make sure nothing has been missed and edge cases have been accounted for!
This program is not exceptionally complicated and I will provide you a pseudocode
outline and several specific things I want do in a certain way so that you practice
several skills. Please follow my instructions and do not do your own thing or you
grade will be harshly affected. I am the customer and what I want is what you will
create.
Change Notices: As in real life I reserve the right to add to this project whatever
and whenever I want to. In real life you would charge more but this is only a partial
simulation:) Point is, you should build your code to make it easy to add additional
instructions if needed!
Implementation: You MUST implement this project in Go.
Instruction Set: The instruction set your disassembler must deal with are
contained in the following table: (see table in lecture 7 slides)

You will be given an input file containing a sequence of 32 bit words. Assume that the first instruction is always
at memory address 96 (0x1100000). The final "instruction" is always a BREAK instruction. All words following
the BREAK instruction are program data. This data is a sequence of 32 bit 2's complement signed
integers.These continue to the end of file. ARM does not have a BREAK instruction so I have made on up and it
is described in Figure 13.1.1
Notes on Instructions:
ADDI, SUBI - the immediate value will be treated as a signed integer that needs to be added or subtracted
SHIFTS - The shamt amount will be an integer describing how many places to shift. The real ARM instruction is
more complicated.
CNBZ, CBZ - Actual assembly is written with tags. Instead of tags the instruction will contain the number of
words offset from the current Program Counter location. So if the branch is to the next instruction, there will be a
1 in the appropriate CB instruction field.
Your Program Execution: Your program must accept command line arguments for execution. The following
and only the following arguments must be supported in exactly the way that I am showing i assuming a single
file to executet:
$go run . team#_project1.go -i test1_bin.txt -o team#_out
The test1_bin.txt input file can be any valid input file that I want to run on your code. I am using test1_bin just an
example! The output file name should be able to be any path I want to aim the file to. Obviously in this case the
output will go into the current folder. PLEASE show some thought and do not enter "team#"! I will not be happy
if I have to mess around when my automation attempts to run your code. DO Not Hardcode file names or
paths!!!! EVER!!!!
MY output file is team0_out_dis.txt
Your code will produce one output file named team#_out_dis.txt which contains the disassembled program.


with input that is not provided to you. It is recommended you create your own input
files since the ones I give you do not test everything. In fact you should create lots of
small test files to test every instruction possible. Figure out the machine code.
Output: The output will contain one line for every line of input. It will be essentially 4
tab separated columns. The columns will contain the following:
1. The binary representation of the instruction word. If the word is an instruction as
opposed to program data after the BREAK, the instruction will be broken into
space separated groups representing the instructions parts.
2. The address of the instruction as an int
3. The name of the instruction operation in CAPS , ie ADDI
4. Each argument in the appropriate place for the actual assembly instruction
separated by a comma and space. ie ADD R2, R1, R3
As an example: R instruction
10001011000 10111 000000 01100 10110 96 ADD R22, R12, R23Your program will be graded with both the sample input files I will give to you and
with input that is not provided to you. It is recommended you create your own input
files since the ones I give you do not test everything. In fact you should create lots of
small test files to test every instruction possible. Figure out the machine code.
Output: The output will contain one line for every line of input. It will be essentially 4
tab separated columns. The columns will contain the following:
1. The binary representation of the instruction word. If the word is an instruction as
opposed to program data after the BREAK, the instruction will be broken into
space separated groups representing the instructions parts.
2. The address of the instruction as an int
3. The name of the instruction operation in CAPS , ie ADDI
4. Each argument in the appropriate place for the actual assembly instruction
separated by a comma and space. ie ADD R2, R1, R3
As an example: R instruction
10001011000 10111 000000 01100 10110 96 ADD R22, R12, R23
If the spacing of the binary instruction causes visual misalignment, this is ok. Please
note the CAPS!!! All integer values should be in decimal, Immediate values should
be preceded by a "#". Be careful and consider which instructions take signed values
and which take unsigned values. Use the correct form depending on the context.
Negative numbers must have a leading "-" sign.
You will be graded on exactness of following these rules. Any deviation will be
punished severely. I will grade using diff to compare your output to mine.

