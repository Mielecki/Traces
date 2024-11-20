# Traces

## Overview

This project, developed for a course on concurrency theory, analizes relationships between tasks in a concurrent system using trace theory. The implementation processes the following input:

1. Task list (a word) - a sequence of actions performed in a system.
2. Set Σ - a set of task names (alphabet).
3. Word - a string of symbols from Σ representing the sequence of executed tasks.

The project computes and returns:

1. The dependency set: pairs of tasks that are dependent on each other.
2. The independence set: pairs of tasks that can execute independently.
3. Graphs in `.gv` format in the `output` folder:
* Dependency graph,
* independence graph,
* Diekert graph,
* Hasse diagram.

## Usage

1. Create an input file `example.txt` with the following structure:
* Write the list of tasks, where each task is defined on a new line in the format:

```
variable := expression
```

* After one blank line, specify the sigma set in the format:

```
A = {a, b, c, ...}
```

* After another blank line, specify the word in the format:

```
w = ...
```

2. Run the program:

```
./Traces
```

3. Run the [Graphviz](https://graphviz.org/download/) for each `.gv` file in the `output` directory:

```
./graphs.sh
```