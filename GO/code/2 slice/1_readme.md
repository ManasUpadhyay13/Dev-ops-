slice is as same as array but there is a quite differece berween them.

the first one begin is that their declaratioin is different as compared to array.

their length is not fixed as in case of arrays.

if a = [ 1, 2, 3]
   b := a
if we change b[1] = 4 and the print the result the result wil be
a = [ 1, 2, 3]
b = [1 , 4 , 3]
but in case of slices the output will be
a = [ 1, 4 , 3]
b = [1 , 4 , 3]

This was because in arrays the duplicate of a variable is created but  in slices the reference of the variable is passed and any changes in the second variable will lead to the changes in the first variable also.