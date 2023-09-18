bt(X, Left, Right). /* a tree with a node value, left and right subtrees */

% insert a node into a tree
insert(X, nil, bt(X, nil, nil)). /* insert a node into an empty tree*/
insert(X, bt(Y, Left, Right), bt(Y, Left1, Right)) :- X =< Y, insert(X, Left, Left1). /* if X value is less than the root, insert in the left subtree using recursion*/
insert(X, bt(Y, Left, Right), bt(Y, Left, Right1)) :- insert(X, Right, Right1). /* else recursively insert in the right subtree*/

% search for a node in the tree
search(X, bt(X, _, _)). /* Base case: the tree is empty*/
search(X, bt(Y, Left, _)) :- X =< Y, search(X, Left). /* if X value is less than the root, search the left subtree using recursion */
search(X, bt(_, _, Right)) :- search(X, Right). /* else recursively search in the right subtree*/

% inorder traversal of a tree, returns a list
inorder(nil, []). /* Base Case: the tree is empty, return an empty list */
inorder(bt(X, Left, Right), R) :- % if there is subtrees
    inorder(Left, R1), % recursively traverse the left subtree
    inorder(Right, R2), % recursively traverse the right subtree
    append(R1, [X|R2], R). % append the left subtree, root node and right subtree

% preorder traversal of a tree, returns a list
preorder(nil, []). /* Base Case: the tree is empty, return an empty list */
preorder(bt(X, Left, Right), R) :- % if there is subtrees
    preorder(Left, R1), % recursively traverse the left subtree
    preorder(Right, R2), % recursively traverse the right subtree
    append([X|R1], R2, R).

% postorder traversal of a tree, returns a list
postorder(nil, []). /* Base Case: the tree is empty, return an empty list */
postorder(bt(X, Left, Right), List) :- % if there is subtrees
    postorder(Left, LL), % recursively traverse the left subtree
    postorder(Right, RL), % recursively traverse the right subtree
    append(LL, RL, NewList), % append the left and right subtree to a NewList
    append(NewList, [X], List). % append the root node X to the NewList to create a final traversed list
