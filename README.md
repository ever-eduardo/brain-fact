# brain-fact
A brainf**k interpreter written in Go.

## Project Milestones.
| Component      	| Status 	|
|----------------	|--------	|
| Lexer          	| ✅      	|
| Compiler       	| ✅      	|
| Interpreter    	| 🔲      	|
| Code Examples  	| 🔲      	|
| Configurations 	| 🔲      	|

## Implementation Details.
According to [The epistle to the Implementors](http://brainfuck.org/epistle.html), there are some specific topics to consider before doing an implementation of the language. In our implementation we follow some of them. Here is the list:

- When the value of a cell is 10, it means end-of-line.
- When found an EOF, it wont affect the value of a cell. It is like a no-op instruction.
- The supported character set of operators is { '+', ',', '-', '.', '<', '>', '[', ']' }. Any other character is a comment.
- The input must be a valid utf-8 encoded file.
- The default size of the array is 32767 cells.
- The pointer starts at position 0.
- Output is not necessarily done with a monospace font. It will depend on the local context.
- I/O operations are done during execution, so all programs are interactive.
- The input operation has a default configurable prompt.
- The compiler matches [] before execution, and it will report an error with unbalanced [].
- The interpreter will have some configurable features like: prompt, wrapping on/off, cell and array size.
- Peace be with us.


The interpreter is intended to be used as a library to embed in your Go projects. Nobody knows when you need a concurrent safe brainf**k interpreter at hand. 


Contributions are welcome!


Happy brain-fact coding! ❤️




