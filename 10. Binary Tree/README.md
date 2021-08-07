# Binary Tree
1. Level Order Traversal Problem.
   <br /> [Question](/10.%20Binary%20Tree/docs/1.level-order-traversal-question.jpg)
   <br /> [Solution](/10.%20Binary%20Tree/1.level-order-traversal-solution.go)
   <br /> Command:-
   ```shell
   $ go run 1.level-order-traversal-solution.go < 0.input-file.txt
   ```

   | Method | TimeComplexity | SpaceComplexity | Source |
   |---|---|---|---|
   | Using Queue | O(N) | O(N) | [Link](/10.%20Binary%20Tree/1.level-order-traversal-solution.go) |
   
2. Level Order Build Tree Problem.
      <br /> [Question](/10.%20Binary%20Tree/docs/2.level-order-build-tree-question.jpg)
      <br /> [Solution](/10.%20Binary%20Tree/2.level-order-build-tree-solution.go)
      <br /> Command:-
      ```shell
      $ go run 2.level-order-build-tree-solution.go < 0.input-file.txt
      ```
   
      | Method | TimeComplexity | SpaceComplexity | Source |
      |---|---|---|---|
      | Using Queue | O(N) | O(N) | [Link](/10.%20Binary%20Tree/2.level-order-build-tree-solution.go) |
 
3. Recursive Pre Order Traversal Problem.
    <br /> [Question](/10.%20Binary%20Tree/docs/3.preorder-traversal-recursive-question.jpg)
    <br /> [Solution](/10.%20Binary%20Tree/3.preorder-traversal-recursive-solution.go)
    <br /> Command:-
    ```shell
    $ go run 3.preorder-traversal-recursive-solution.go < 0.input-file.txt
    ```
 
    | Method | TimeComplexity | SpaceComplexity | Source |
    |---|---|---|---|
    | Recursive | O(N) | O(N) | [Link](/10.%20Binary%20Tree/3.preorder-traversal-recursive-solution.go) |  
    
4. Iterative Pre Order Traversal Problem.
     <br /> [Question](/10.%20Binary%20Tree/docs/4.preorder-traversal-iterative-question.jpg)
     <br /> [Solution](/10.%20Binary%20Tree/4.preorder-traversal-iterative-solution.go)
     <br /> Command:-
     ```shell
     $ go run 4.preorder-traversal-iterative-solution.go < 0.input-file.txt
     ```
  
     | Method | TimeComplexity | SpaceComplexity | Source |
     |---|---|---|---|
     | Iterative using stack | O(N) | O(N) | [Link](/10.%20Binary%20Tree/4.preorder-traversal-iterative-solution.go) |  
     
5. Recursive In Order Traversal Problem.
     <br /> [Question](/10.%20Binary%20Tree/docs/5.inorder-traversal-recursive-question.jpg)
     <br /> [Solution](/10.%20Binary%20Tree/5.inorder-traversal-recursive-solution.go)
     <br /> Command:-
     ```shell
     $ go run 5.inorder-traversal-recursive-solution.go < 0.input-file.txt
     ```
  
     | Method | TimeComplexity | SpaceComplexity | Source |
     |---|---|---|---|
     | Recursive | O(N) | O(N) | [Link](/10.%20Binary%20Tree/5.inorder-traversal-recursive-solution.go) | 
     
6. Iterative In Order Traversal Problem.
     <br /> [Question](/10.%20Binary%20Tree/docs/6.inorder-traversal-iterative-question.jpg)
     <br /> [Solution](/10.%20Binary%20Tree/6.inorder-traversal-iterative-solution.go)
     <br /> Command:-
     ```shell
     $ go run 6.inorder-traversal-iterative-solution.go < 0.input-file.txt
     ```
  
     | Method | TimeComplexity | SpaceComplexity | Source |
     |---|---|---|---|
     | Iterative using stack | O(N) | O(N) | [Link](/10.%20Binary%20Tree/6.inorder-traversal-iterative-solution.go) |  
     
7. Recursive Post Order Traversal Problem.
     <br /> [Question](/10.%20Binary%20Tree/docs/7.postorder-traversal-recursive-question.jpg)
     <br /> [Solution](/10.%20Binary%20Tree/7.postorder-traversal-recursive-solution.go)
     <br /> Command:-
     ```shell
     $ go run 7.postorder-traversal-recursive-solution.go < 0.input-file.txt
     ```
  
     | Method | TimeComplexity | SpaceComplexity | Source |
     |---|---|---|---|
     | Recursive | O(N) | O(N) | [Link](/10.%20Binary%20Tree/7.postorder-traversal-recursive-solution.go) | 
     
8. Recursive Post Order Traversal Problem.
     <br /> [Question](/10.%20Binary%20Tree/docs/8.postorder-traversal-iterative-question.jpg)
     <br /> [Solution](/10.%20Binary%20Tree/8.postorder-traversal-iterative-solution.go)
     <br /> Command:-
     ```shell
     $ go run 8.postorder-traversal-iterative-solution.go < 0.input-file.txt
     ```
  
     | Method | TimeComplexity | SpaceComplexity | Source |
     |---|---|---|---|
     | Iterative using stack | O(N) | O(N) | [Link](/10.%20Binary%20Tree/8.postorder-traversal-iterative-solution.go) |            

9. Height of Binary Tree Problem.
   <br /> [Question](/10.%20Binary%20Tree/docs/9.height-of-tree-question.jpg)
   <br /> [Solution](/10.%20Binary%20Tree/9.height-of-tree-solution.go)
   <br /> Command:-
   ```shell
   $ go run 9.height-of-tree-solution.go < 0.input-file.txt
   ```

   | Method | TimeComplexity | SpaceComplexity | Source |
   |---|---|---|---|
   | Recursive | O(N) | O(N) | [Link](/10.%20Binary%20Tree/9.height-of-tree-solution.go) |
   
10. Diameter of Binary Tree Problem.
       <br /> [Question](/10.%20Binary%20Tree/docs/10.diameter-of-tree-question.jpg)
       <br /> [Solution O(N^2)](/10.%20Binary%20Tree/10.diameter-of-tree-o(n^2)-solution.go)
       <br /> [Solution O(N)](/10.%20Binary%20Tree/10.diameter-of-tree-o(n)-solution.go)
       <br /> Command:-
       ```shell
       $ go run 10.diameter-of-tree-o(n^2)-solution.go < 0.input-file.txt
    
       $ go run 10.diameter-of-tree-o(n)-solution.go < 0.input-file.txt
       ```
    
       | Method | TimeComplexity | SpaceComplexity | Source |
       |---|---|---|---|
       | Calculating height and diameter in different call | O(N^2) | O(N) | [Link](/10.%20Binary%20Tree/10.diameter-of-tree-o(n^2)-solution.go) |
       | Calculating height and diameter in same call | O(N) | O(N) | [Link](/10.%20Binary%20Tree/10.diameter-of-tree-o(n)-solution.go) |

11. Replace Descendants Sum Tree Problem.
       <br /> [Question](/10.%20Binary%20Tree/docs/11.replace-with-descendants-sum-question.jpg)
       <br /> [Solution](/10.%20Binary%20Tree/11.replace-with-descendants-sum-solution.go)
       <br /> Command:-
       ```shell
       $ go run 11.replace-with-descendants-sum-solution.go
       ```
    
       | Method | TimeComplexity | SpaceComplexity | Source |
       |---|---|---|---|
       | Recursive | O(N) | O(N) | [Link](/10.%20Binary%20Tree/11.replace-with-descendants-sum-solution.go) |
   
12. Height Balanced Tree Problem.
        <br /> [Question](/10.%20Binary%20Tree/docs/12.is-balanced-tree-question.jpg)
        <br /> [Solution O(N^2)](/10.%20Binary%20Tree/12.is-balanced-tree-o(n^2)-solution.go)
        <br /> [Solution O(N)](/10.%20Binary%20Tree/12.is-balanced-tree-o(n)-solution.go)
        <br /> Command:-
       ```shell
       $ go run 12.is-balanced-tree-o(n^2)-solution.go < 0.input-file.txt
    
       $ go run 12.is-balanced-tree-o(n)-solution.go < 0.input-file.txt
       ```
    
       | Method | TimeComplexity | SpaceComplexity | Source |
       |---|---|---|---|
       | Calculating height and balance check in different call | O(N^2) | O(N) | [Link](/10.%20Binary%20Tree/12.is-balanced-tree-o(n^2)-solution.go) |
       | Calculating height and balance check in same call | O(N) | O(N) | [Link](/10.%20Binary%20Tree/12.is-balanced-tree-o(n)-solution.go) |
   
   13. Max Sub Set Sum Tree Problem.
        <br /> [Question](/10.%20Binary%20Tree/docs/13.max-sub-set-sum-question.jpg)
        <br /> [Solution](/10.%20Binary%20Tree/13.max-sub-set-sum-solution.go)
        <br /> Command:-
      
        ```shell
        $ go run 13.max-sub-set-sum-solution.go  < 0.input-file.txt
        ```
       
       | Method | TimeComplexity | SpaceComplexity | Source |
       |---|---|---|---|
       | Recursive | O(N) | O(N) | [Link](/10.%20Binary%20Tree/13.max-sub-set-sum-solution.go) |
                         