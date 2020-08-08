# Cheapest flights within k stops

[Problem Link](https://leetcode.com/problems/cheapest-flights-within-k-stops/)

This is a node traversal problem which can be attacked with recursion and parallelism.

- Keep a common variable for the cheapest cost
- Traverse nodes recursively and accumulate cost.
- For each traversal, validate whether the destination is reached and if yes, push to cheapest cost store
- The cheapest cost store shall accept the path &amp; the cost only if the incoming cost is cheaper than the existing
- Once all the go routines complete, return the cheapest cost and the path