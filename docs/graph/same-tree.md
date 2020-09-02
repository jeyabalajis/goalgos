# Same Tree

## Problem Statement:
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

### Example:
![Same Tree](https://jeyabalajis.github.io/goalgos/images/same-tree.png)

Difficulty: :moneybag:

[Leet Code Link](https://leetcode.com/problems/same-tree/)

## Solution:

This is a simple tree traversal problem. If we simply get a hash out of the node values, we will miss out on the structure of the tree.

> To solve this problem, we need a way to get a hash value of a tree that respects tree's structure.

To incorporate a tree's structure also into it's hash, we can implement the following approach:

- For a node, include node's depth, it's qualifer (_ROOT_, _LEFT_ or _RIGHT_) and it's value together and get a hash value.
- The sum of this hash value across all nodes is the hash value of a tree.


### Algorithm:

- Implement each node traversal as a go routine. 
- At each node, compute Hash Value as a combination of it's depth, qualifier (left, right or root) and it's value.
- Collect all such hash values for the tree and compare the same with that of the other tree