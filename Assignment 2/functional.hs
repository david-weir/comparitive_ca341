data BinTree t = Empty | Root t (BinTree t)(BinTree t)  -- BST data type
    deriving (Eq, Ord, Show)

leaf x = Root x Empty Empty  -- leaf is a node with no child nodes

insert :: Ord t => t -> BinTree t -> BinTree t  -- Takes an integer and a BST and inserts the int into the bst

insert a Empty = Root a Empty Empty  -- if bst is empty create a bst with root a
insert x (Root a left right)  -- else there is a left and/or left and right child node
    | x < a = Root a (insert x left) right -- if x is less than root a, insert x in the left subtree using recursion
    | otherwise = Root a left (insert x right)  -- otherwise insert the value x recursively in the right subtree 

search :: Ord t => t -> BinTree t -> Bool -- takes an integer and bst and returns True/False depending on if the value is in the bst

search x Empty = False -- if the bst is empty return false
search x (Root a left right) -- else if it is non-empty
    | x == a = True  -- if value x equals root node a, return true
    | x < a = search x left  -- if x is < a, search the left subtree
    | x > a = search x right -- if x > a, search the right subtree

inorder :: BinTree t -> [t] -- takes a bst and returns a list of nodes from inorder traversal 

inorder Empty = []  -- if the tree is empty return an empty list
inorder (Root x Empty Empty) = [x] -- if the tree only has a root node, return the root
inorder (Root x left right) = (inorder left) ++ [x] ++ (inorder right) -- otherwise, recursively traverse the left subtree, then the root and finally the right subtree

preorder :: BinTree t -> [t] -- takes a bst and returns a list of nodes from preorder traversal 

preorder Empty = [] -- if the tree is empty return an empty list
preorder (Root x Empty Empty) = [x] -- if the tree only has a root node, return the root
preorder (Root x left right) = [x] ++ (preorder left) ++ (preorder right) -- otherwise, recursively traverse the tree from root node, left subtree and then the right subtree

postorder :: BinTree t -> [t] -- takes a bst and returns a list of nodes from postorder traversal 

postorder Empty = [] -- if the tree is empty return an empty list
postorder (Root x Empty Empty) = [x] -- if the tree only has a root node, return the root
postorder (Root x left right) = (postorder left) ++ (postorder right) ++ [x] -- otherwise, recursively traverse the left subtree, then the right subtree, then the root

-- for testing, creates a tree from a list
makeTree :: Ord t => [t] -> BinTree t
makeTree [] = Empty -- the list empty create an empty tree
makeTree [x] = Root x Empty Empty -- the list has 1 element, return a root node with 2 empty child nodes
makeTree (x:xs) = insert x (makeTree xs) -- else insert the head of the list to the bst and recursivly call makeTree on the tail until the list is empty