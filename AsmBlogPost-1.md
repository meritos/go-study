


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
```Assembly  
TEXT ·Add(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), BX
    MOVQ y+8(FP), BP
    ADDQ BP, BX
    MOVQ BX, ret+16(FP)
    RET
```

We can see that function body is defined in the regular *.go file, and the body is implemented in the *.s file which will be compiled and linked by the regular Go compiler, no need for new tools.

...

You can study more about Assembly basic in that [post](AsmBlogPost-1.md#1-github-repository-for-my-asm-games)  


## How to call back to go methods => `NO_LOCAL_POINTERS`
...
## Methods I wrote to trace go values => PrintVal, PrintFlags, PrintMemory
...
## Explain memory of array, function
...
## Finish with measuring cycles
There is very useful ASM instruction that can upgrade any performance benchmarking strategy. I am talking about: `RDTSC` -Read Time-Stamp Counter which gives you the actual CPU cycles counter. 
...
## To remember: 
  * To remove unnecessary directives 
  * ..
  
## Reference
###### [1]: [Github repository for my ASM games](https://github.com/meritos/go-study/tree/master/asm)

###### [2]: [Helpful post](https://blog.sgmansfield.com/2017/04/a-foray-into-go-assembly-programming/)

###### [3]: [RDTSC reference](http://x86.renejeschke.de/html/file_module_x86_id_278.html)


