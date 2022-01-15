NAME:
=====
	Neha Thumu

Programs Files:
===============
    Part 1:
        all program files related to part 1: structs.go 
    Part 2:
        all program files related to part 2: slices.go
    Part 3: [below]
        1.) 
        The original and the copy are not equal. Because the copy and the 
        original don't share any memory but have the same fields (when copied). 
        As such, once a field is changed in one of them, that change is not 
        reflected in the other. 

        On the other hand, in Java the original and the copy would point to 
        the same location in memory. 
        So if it's changed for the original or the copy then that change is 
        reflected in the other one. 

        2.) 
        The items in position 1 of the two slices would be equal.

        If the item was changed in the original it would not be equal to 
        the item in the same position of the new slice. Because the original 
        slice and the new slice are in two different locations in memory.

	
How to Compile:
===============
    Part 1:
       go build structs.go 
    Part 2:
       go build slices.go 
       
How to Run:
===========
      Part 1:
    	go run structs.go 
      Part 2:
        go run slices.go  
      Part 3:
        in this Readme file (shown above)
	
Reflection:
===========

     Writing this code made me appreciate Python more. I didn't realize how 
     kind the language was. Or maybe I'm just too used to stuff being in a
     certain order. The assignment itself was not bad. It took me a while to 
     understand the point of structs and how to use them properly (especially 
     that struct being used by another struct - that was so odd for some reason) 
     but now I (hopefully) understand. Did had to read some examples to get the 
     code to finally work though. 


I Worked With:
==============
I made plans to discuss the lab with Fatima (at the time of this submission we 
have not talked but maybe we will before she submits) 

[online resources]
I used https://gobyexample.com/ to understand how slices and structs worked. 
Also https://www.geeksforgeeks.org/was incredibly helpful for understanding where to 
add the go toString method (location in my code)

Approximate Hours worked:
=========================
7 (?)

Special Instructions to the grader:
===================================
N/A

Known Bugs or Limitations:
==========================
    The output isn't formatted well.

    VSC shows an error if the package names are the same for structs.go and 
    slices.go. Since they both have the same structs defined. 

Other comments:
===============
    Dragons are pretty neat. It's a pity they don't exist. 