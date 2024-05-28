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
2) Then, ensure the remaining numbers is only equal to one type of number ( since 0 is majority, we flip away 1, located at bottom left, to become 0)

1  0  0  0  2          1  0  0  0  2 
0  1  0  2  0          0  1  0  2  0
0  0  1  0  0   --->   0  0  1  0  0
0  0  2  0  0          0  0  2  0  0
1  0  2  0  0          0  0  2  0  0
```

