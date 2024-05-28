# CodeSignal Technical Assessment 

**Flipping to get letter Y** 

given an odd matrix, find the minimum number of flips to achieve letter Y. Ensure only two numbers remains on the board after flipping.
Possible remaining combination = (0,1), (0,2), (1,2)

```
Example, matrix of len 3, only one flip (change 1 => 2). Result is 1.
1  0  2            2  0  2
0  2  0    ---->   0  2  0
0  2  0            0  2  0
```

```
Example, matrix of len 5, three flips (change 1 to 2). Result is 3.

1  0  0  0  2          1  0  0  0  2 
0  1  0  2  0          0  1  0  2  0
0  0  1  0  0   --->   0  0  1  0  0
0  0  2  0  0          0  0  2  0  0
0  0  2  0  0          0  0  2  0  0

Reason for flipping 1 is because, there is four '2's and only three '1's in the letter Y.
Since we need to return the minimum flips, we should flip 1 to become 2.
```

```
Example, matrix of len 5, result = four flips. Reason:
1) First, change 1 to 2 for those in letter Y = 3 flips
2) Then, ensure the remaining numbers is only equal to one type of number
( since 0 is majority, we flip away 1, located at bottom left, to become 0)

1  0  0  0  2          1  0  0  0  2 
0  1  0  2  0          0  1  0  2  0
0  0  1  0  0   --->   0  0  1  0  0
0  0  2  0  0          0  0  2  0  0
1  0  2  0  0          0  0  2  0  0

Total space in Y =       5           +   2      = 7
                    (len of matrix)      (mid)

Total outside space =       5*5             -     7
                     (total len of matrix)      (total Y spaces)

                    = 25 - 7 = 18
```


**Approach**

Determine:
1) What's the majority inside letter Y - is it 0, 1 or 2?
2) What's the majority outside of letter Y - is it 0, 1 or 2?

Then, flips required = (total Y space - majority inside Y) + (total outside space - majority outside Y)

```
Total spaces inside Y = len(matrix) + mid of matrix 

Total spaces outside Y = total length n*n - total spaces inside Y
```

```
To determine majority of Y, can try using a recursive loop, starting from mid of matrix, where mid = len(matrix) // 2

use an array to store the count of each number [0, 0, 0]

Starting from mid, recursive count the frequency by using => directions [ (-1,-1), (1,1), (1,0) ]
  upperleft of letter Y = (-1,-1)
  upperright of letter Y = (1, 1)
  downwards of letter Y  = (1,0)

From here, i can know the number of highest freq. All i have to do is subtract the majority from total Y space, which leaves me the remaining that I must flip.
```

```
For space outside Y:
Iterate through matrix to get total frequency of each number in the grid.

Then subtract the freq in Y to get the outside number's frequencies
From here, I can't have the outside space to be the same number as the number inside Y.

How to determine which to flip if matrix is all 1s? in this case we shld flip the inside of Y instead for minimum flips.
```




