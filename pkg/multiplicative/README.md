# Introduction

Multiply all the digits of a number n by each other, repeating with the product until a single digit is obtained.
The number of steps required is known as the multiplicative persistence, and the final digit obtained is called the multiplicative digital root of n.

For example, the sequence obtained from the starting number 9876 is (9876, 3024, 0),
so 9876 has an multiplicative persistence of two and a multiplicative digital root of 0.
The multiplicative persistences of the first few positive integers are 
0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 1, 1, ... (OEIS A031346).
The smallest numbers having multiplicative persistences of 1, 2, ... are `10, 25,
39, 77, 679, 6788, 68889, 2677889, 26888999, 3778888999, 277777788888899`, ... (OEIS A003001; Wells 1986, p. 78).

There is no number <10^(233) with multiplicative persistence >11 (Carmody 2001; updating Wells 1986, p. 78).
It is conjectured that the maximum number lacking the digit 1 with persistence 11 is `77777733332222222222222222222`

There is a stronger conjecture that there is a maximum number lacking the digit 1 for each persistence >2.

# Goals

- We want to programitically find the multiplicative persistence for all the `int` numbers as fast as possible.
- We also want to find the minimum & maximum value value of `n` with highest persistence of F(n) for `int`.