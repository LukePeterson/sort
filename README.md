# Sorting algorithms in Go

This repository contains common sorting algorithms, implemented in Go:

Algorithm         | Best case                | Average case             | Worst case              | Space                  | Stability*
------------------|--------------------------|--------------------------|-------------------------|------------------------|-------------------
[Bubble sort](https://github.com/LukePeterson/sort/tree/main/bubble)       | O(n)                     | O(n^2)                   | O(n^2)                  | O(1)                   | Stable
[Insertion sort](https://github.com/LukePeterson/sort/tree/main/insertion)    | O(n)                     | O(n^2)                   | O(n^2)                  | O(1)                   | Stable
[Selection sort](https://github.com/LukePeterson/sort/tree/main/selection)    | O(n^2)                   | O(n^2)                   | O(n^2)                  | O(1)                   | Not stable
Merge sort        | O(n log n)               | O(n log n)               | O(n log n)              | O(n)                   | Stable
[Quick sort](https://github.com/LukePeterson/sort/tree/main/quick)        | O(n log n)               | O(n log n)               | O(n^2)                  | O(log n)               | Not stable
Heap sort         | O(n log n)               | O(n log n)               | O(n log n)              | O(1)                   | Not stable

_* "Stability" in sorting algorithms refers to the preservation of the original order of equal elements. A stable sorting algorithm ensures that if two items are equal, they will appear in the same order in the sorted output as they did in the input._
