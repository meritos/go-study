


# Golang: How to trace embedded Assembly modules?


Plan: 

It was great discovery for me to find out that Golang can have direct calls to 
Assembly modules. Assembly programming gives you that feeling that you can make a program to be ultra optimized and every programmer like to have that power on tips of his fingers

## Intro: how ASM embedded in Golang works 
To introduce briefly how the Assembly modules are defined let's see very simple example: 

### AsmSamples.go

```Golang 
func Add(x, y int64) int64
```

### AsmSamples.s
```c  
TEXT ·Add(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), BX
    MOVQ y+8(FP), BP
    ADDQ BP, BX
    MOVQ BX, ret+16(FP)
    RET
```

We can see that function body is defined in the regular *.go file, and the body is implemented in the *.s file which will be compiled and linked by the regular Go compiler, no need for new tools. The strange `·` - middle dot in front of the function name is actually a package delimiter, so you can add package path to the signature. `NOSPLIT` - is directive on memory management which I will not explain here but you can study more about that here: [...](...). `$0-24` in the function is `$stack_size-arguments_size` in bytes. Here you we can see that no stack allocation being used and 24 bytes is used for params: x,y: int64 = 8 bytes and ret_val: int64 = 8 bytes. You can do absolutely anything you want in assembly but I will stop here just to set the stage for the tracing model .You can study more about Assembly basic in that useful blog post: [[2]](blog-1.md#2-helpful-post)  

## How to call back to go methods  
So far we had studied how to make methods totally implemented in Asm language. But Asm is as cool as it is, also is a very close to machine language and that
requires very good tools to develop any algorithm. Fortunately we have option to call regular go code form Asm and that will give us the required setup to make very strong tracing for the development use.

So how I call from Asm to go? first of all I will need stack: if in the `Add` method no stack was defined it is only cause it was not required, for any call back to go we will define stack - at least temporary until the code will not require anymore tracing. There is also GC that has to be informed about nature of our method that is why we should use that directive `NO_LOCAL_POINTERS`, it will set the map of pointers the GC is asking from any Go function. You can try to make from Asm to Go call without that directive, the best chance you will have an ugly runtime exception. So here is how looks like code that is ready to be traced by our methods: 

```c
TEXT ·Add(SB), NOSPLIT, $8-24
    
    NO_LOCAL_POINTERS

    MOVQ x+0(FP), BX
    MOVQ y+8(FP), BP
    ADDQ BP, BX
    MOVQ BX, ret+16(FP)
    RET
```

## Functions I sketched to trace Assembly values 
I had developed 3 Go functions that helps me to trace what is going on with the code: `PrintVal(val)`, `PrintFlags(FlagsReg)`, `PrintMemory(address, size)`. Lets see how this three will help us. 

  * `PrintVal(val)` - Any value can be dumped with that function, and because in Asm we usually want to make bitwise operations - also add some clear representation of the value in hex and in binary. Here is how it works: 

```c
    
    // Send argument through the stack
    MOVQ    $20, 0(SP)
    CALL ·PrintVal(SB)
```

 The output will look like that: 

 (!) todo: crop screen for PrintVal 
  

  * `PrintFlags(FlagsReg)` - 
The flags register is the fundamental method to control program flow in Asm. We can use `PUSHFQ` to access all of it and later send it to `·PrintFlags` to
print the translation of it. Here is the simple syntax for that: 
 
```c
    
    // Access flag reg
    PUSHFQ

    // Set glag reg as argument 
    POPQ    R9
    MOVQ    R9,    0(SP)

    // Make the function call 
    CALL    ·PrintFlags(SB)
```

(!) todo: insert the terminal output

  * `PrintMemory(address, size)` - The most tricky function, it will print for us full blocks of memory and ensure that our memory manipulation is correct. 
  to try it lets define a local array of `int64`: that is how we do it in Asm: 

```c
    DATA intArray<>+0x00(SB)/8, $0
    DATA intArray<>+0x08(SB)/8, $1
    DATA intArray<>+0x10(SB)/8, $2
    DATA intArray<>+0x18(SB)/8, $3
    DATA intArray<>+0x20(SB)/8, $4
    DATA intArray<>+0x28(SB)/8, $5
    DATA intArray<>+0x30(SB)/8, $6
    DATA intArray<>+0x38(SB)/8, $7
    DATA intArray<>+0x40(SB)/8, $8
    GLOBL intArray<>(SB), (RODATA | NOPTR), $0x48
    ```

  Here is how to print the real memory that is allocated for that array: 

```c    
    // Load Address of the array
    LEAQ intArray<>(SB), DX
    
    // Init args for the CALL
    // make CALL with the args
    // address and $10 for size 
    // we want to print 
    MOVQ    DX,    0(SP)
    MOVQ    $10,    8(SP)

    // Make the call
    CALL    ·PrintMem(SB)
```

The output will be something like that: 

![intArrya Memory Dump](http://i.imgur.com/By9OSLKg.png)

We can clearly see the values 01..08 that we set in the init part of  the 
`intArray`


(!) todo: Print memory of the function address...


The interesting thing I studied is that you can turn any address saved as `int` into a pointer using that syntax: 

```go 
    ptr := (*uint64)(unsafe.Pointer(uintptr(addr + int64(i * 8) )))
```

Obviously you don't want to have anything like this in a production code , but for tracing at developing it should be fine.


## Finish with measuring cycles
There is very useful ASM instruction that can upgrade any performance benchmarking strategy. I am talking about: `RDTSC` -Read Time-Stamp Counter which gives you the actual CPU cycles counter. 
...
## To remember: 
  * After the call to a function the registers can be reset
  * To remove unnecessary directives 
  * To run my sample you can use: (!) todo: ...
  
## Reference
###### [1]: [Github repository for my ASM games](https://github.com/meritos/go-study/tree/master/asm)
###### [2]: [Helpful post](https://blog.sgmansfield.com/2017/04/a-foray-into-go-assembly-programming/)
###### [3]: [Helpful post 2](https://goroutines.com/asm)
###### [4]: [RDTSC reference](http://x86.renejeschke.de/html/file_module_x86_id_278.html)

