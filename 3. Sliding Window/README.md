# Sliding Window
1. Sub Array Sum Problem.
   <br /> [Question](/3.%20Sliding%20Window/docs/1.sub-array-sum-question.jpg)
   <br /> [Solution](/3.%20Sliding%20Window/1.sub-array-sum-solution.go)
   <br /> Command:-
   ```shell
   $ go run 1.sub-array-sum-solution.go
   ```

   | Method | TimeComplexity | SpaceComplexity | Source |
   |---|---|---|---|
   | Brute Force  | O(N^3) | O(1) |
   | Prefix SumArray | O(N^2)  | O(N) |
   | Prefix SumArray + Binary Search | O(Nlog(N))  | O(N) |
   | Sliding Window | O(N) | O(1) | [Link](/3.%20Sliding%20Window/1.sub-array-sum-solution.go) |