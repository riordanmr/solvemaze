# solvemaze
Simple program to solve a maze.

The input file contains a description of the multiway intersections in the maze.

That is, the file describes only those nodes in the maze where you have a choice of which way to go. Each such node has a number, and each node has letters describing the choices.  In the original maze for which this program was written - from the 2024-05-26 issue of the New York Times - each node in the maze has only two choices, so the directions to go are "a" and "b". 

The input file contains lines, each of which describes a path leading out of a given node.

- "1a 2"  means that if you start at node 1 and take direction a, you get to node 2.
- "1b x"  means that if you start at node 1 and take direction b, you eventually reach a dead end.
- "2a e"  means that if you start at node 2 and take direction a, you will reach the end.

The maze must be annotated (outside of this program) as to which nodes have numbers, what those numbers are, 
and which direction out of a node is "a" and "b".
