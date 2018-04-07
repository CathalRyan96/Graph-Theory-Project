# Graph Theory Project

This is a repository for my project for the Graph Theory Module.


<h2>Problem</h2>
You must write a program in the Go programming language [2] that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “∗
” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission.

<h3>Structure of Project</h3>
I divived the project up to 4 different classes:

- Shunt.go
- nfa.go
- rega.go
- runner.go

# Shunt Algorithm
In this class I displayed that the shunting yard algorithm is a method for parsing mathematical expressions specified in infix notation.

<h5>Thompsons Algorithm<h5>
In the nfa.go class I displayed the thomspsons construction alogorithm, which is a algorithm for transforming a regular expression into an equivalent nondeterministic finite automaton.

<h6>Matching Function<h6>
In the rega.go class I display how to use a match function on a string and a postfix regular expression.

