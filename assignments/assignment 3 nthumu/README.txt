NAME:
=====
	Neha Thumu

Programs Files:
===============
    Part 1:
        all program files related to part 1: helloworld.kt 
                                             helloworldtailrec.kt
                                             nontailrec.txt
                                             tailrec.txt
    Nested Functions:
        all program files related to part 2: nested.kt 
                                             nestedv2.kt
    Functions that return functions: 
        all program files related to part 3: returnedfunc.kt
    
    Filtering:
        all program files related to part 4: filtergeneric.kt
                                             filtering.kt
      
Answers to questions: 
===============
    Part 1:
        The largest number of times I can print "hello world" without getting a stack 
        overflow exception: 12698
            **Note! This number changes each time I run the program w/ JVM
        
        Adding the "tailrec" changes the stack overflow limit. In that there is no limit 
        in this case. 

    Nested Functions:
        I personally prefer the first version since it's easier for me to read/follow 
        along. 
    
     Functions that return functions: 
        The closure of the returned function is the paramter passed to nthpower (exponent). 
    
    Filtering: 
        (no questions?)
	
How to Compile:
===============
    Part 1:
       kotlinc helloworld.kt -include-runtime -d helloworld.jar   
       kotlinc helloworldtailrec.kt -include-runtime -d helloworldtailrec.jar
    Nested Functions:
        kotlinc nested.kt -include-runtime -d nested.jar
        kotlinc nestedv2.kt -include-runtime -d nestedv2.jar
    Functions that return functions:
        kotlinc returnfunc.kt -include-runtime -d returnfunc.jar
    Filtering: 
        kotlinc filtergeneric.kt -include-runtime -d filtergeneric.jar  
    
       
How to Run:
===========
      Part 1:
        Java -jar helloworld.jar 20000  >nontailrec.txt 27eout  
    	Java -jar helloworldtailrec.jar 20000  >tailrec.txt 27eout 
      Nested Functions:
        java -jar nested.jar 
        java -jar nestedv2.jar 
      Functions that return functions:
        java -jar returnfunc.jar
      Filtering:
        java -jar filtergeneric.jar
        
	
Reflection:
===========

     I really have started to appreciate Python. I keep accidentally using the 
     syntax and wondering what happened with my code. 
     I see the benefits of Kotlin but to be honest I'm not much of a fan (mostly
     due to the fact that vscode hates it and it's impossible to read my own code).


I Worked With:
==============
N/A

Approximate Hours worked:
=========================
11

Special Instructions to the grader:
===================================
N/A

Known Bugs or Limitations:
==========================
    The stack overflow is weird (giving different numbers)

Other comments:
===============
    N/A