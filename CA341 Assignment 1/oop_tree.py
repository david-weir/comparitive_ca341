class Node:  # Node class with name phone and address strings and Left and Right child nodes
    def __init__(self, name, phone, address):
        self.name = name
        self.phone = phone
        self.addr = address
        self.left = None
        self.right = None

    def __str__(self):
        message = ""
        message += f"Name: {self.name}\n"
        message += f"Phone Number: {self.phone}\n"
        message += f"Address: {self.addr}\n"

        return message


# phonebook bst based on contact name
class NameTree:

    def __init__(self):
        self.root = None

    # insert a node
    def insert(self, name, phone, address):
        # if the tree is empty set a root node
        if self.root is None:
            self.root = Node(name, phone, address)

        # else add a new node to the tree
        else:
            self.binary_insert(self.root, name, phone, address)

    def binary_insert(self, curr, name, phone, address):

        # if the name is greater than the current node, continue down the right subtree
        if name > curr.name:
            if curr.right is None:
                curr.right = Node(name, phone, address)
            else:
                self.binary_insert(curr.right, name, phone, address)

        # if the name is less than the current node, continue down the left subtree
        else:
            if curr.left is None:
                curr.left = Node(name, phone, address)
            else:
                self.binary_insert(curr.left, name, phone, address)

    # search for a node based on name
    def search(self, name, node=None):

        # if we are starting at the root we want to set node to the root
        # but if we are starting at the child node we do not want to restart at the node
        if node is None:
            node = self.root

        # if we find the correct node return it
        if node.name == name:
            return node

        # if the name is greater than the current node name, search the right subtree
        elif name > node.name:

            if node.right is None:
                return None
            else:
                return self.search(name, node.right)

        # else search the left subtree
        else:
            if node.left is None:
                return None
            else:
                return self.search(name, node.left)

    def remove(self, name):
        if self.root is None:
            return False

        if name < self.root.name:
            self.root.left = self.remove(name)
        elif name > self.root.name:
            self.root.right = self.remove(name)
        else:
            if self.root.left is None or self.root.right is None:
                tmp = Node(None, None, None)

                if self.root.name is None:
                    tmp = self.root.left

                else:
                    tmp = self.root.right

                if tmp is None:
                    return None
                else:
                    return self.root

            else:
                child = Node(None, None, None)
                child = self.getchild()
                self.root.name = child.name
                self.root.right = self.remove(child.name)
                return name

        return name

    def getchild(self):
        if self.root is None:
            return None

        tmp = self.root.right

        while tmp is not None:
            tmp = tmp.left

        return tmp

    def traverse(self, curr):
        # inorder traversal of a tree
        if curr:
            self.traverse(curr.left)
            print(curr)
            self.traverse(curr.right)


# phonebook bst based on contact phone number
class PhoneTree:

    def __init__(self):
        self.root = None

    # insert a node based on phone number
    def insert(self, name, phone, address):
        # if the tree is empty set a root node
        if self.root is None:
            self.root = Node(name, phone, address)

        # else add a new node to the tree
        else:
            self.binary_insert(self.root, name, phone, address)

    def binary_insert(self, curr, name, phone, address):

        # if the phone is greater than the current node, continue down the right subtree
        if phone > curr.phone:
            if curr.right is None:
                curr.right = Node(name, phone, address)
            else:
                self.binary_insert(curr.right, name, phone, address)

        # if the name is less than the current node, continue down the left subtree
        else:
            if curr.left is None:
                curr.left = Node(name, phone, address)
            else:
                self.binary_insert(curr.left, name, phone, address)

    def search(self, phone, node=None):

        # if we are starting at the root we want to set node to the root
        # but if we are starting at the child node we do not want to restart at the node
        if node is None:
            node = self.root

        # if we find the correct node return it
        if node.phone == phone:
            return node

        # if the phone is greater than the current node name, search the right subtree
        elif phone > node.phone:

            if node.right is None:
                return None
            else:
                return self.search(phone, node.right)

        # else search the left subtree
        else:
            if node.left is None:
                return None
            else:
                return self.search(phone, node.left)

    def remove(self, phone):

        if self.root is None:
            return False

        if phone < self.root.phone:
            self.root.left = self.remove(phone)
        elif phone > self.root.phone:
            self.root.right = self.remove(phone)
        else:
            if self.root.left is None or self.root.right is None:
                tmp = Node(None, None, None)

                if self.root.name is None:
                    tmp = self.root.left

                else:
                    tmp = self.root.right

                if tmp is None:
                    return None
                else:
                    return tmp

            else:
                child = self.getchild()
                self.root.name = child.name
                self.root.right = self.remove(child.name)
                return phone

        return phone

    def getchild(self):
        if self.root is None:
            return None

        tmp = self.root.right

        while tmp is not None:
            tmp = tmp.left

        return tmp

    def traverse(self, curr):
        if curr:
            self.traverse(curr.left)
            print(curr)
            self.traverse(curr.right)


def main():
    phone_tree = PhoneTree()
    name_tree = NameTree()

    while True:

        command = input("$ ")
        args = command.split()

        if command == "exit":
            break

        elif command == "help":
            print("CA341 Procedural BST Manual\n")
            print("Supported Commands:\n")

            print("insertname: inserts a node by name")
            print("insertphone: inserts a node by phone number\n")

            print("findname: searches the tree for a node by name")
            print("findphone: searches the tree for a node by phone\n")

            print("Deletion is WIP")
            print("deletename: deletes the node with the given name")
            print("deletephone: deletes the node with the given name\n")

            print("exit: Exits program.")

        elif args[0] == "insertname":
            name_tree.insert(args[1], args[2], " ".join(args[3:]))
            print(name_tree.traverse(name_tree.root))

        elif args[0] == "insertphone":
            phone_tree.insert(args[1], args[2], " ".join(args[3:]))
            print(phone_tree.traverse(phone_tree.root))

        elif args[0] == "findname":
            print(name_tree.search(args[1]))

        elif args[0] == "findphone":
            print(name_tree.search(args[1]))

        elif args[0] == "deletename":
            name_tree.remove(args[1])
            print(name_tree.traverse(name_tree.root))

        elif args[0] == "deletename":
            phone_tree.remove(args[1])
            print(phone_tree.traverse(phone_tree.root))
            
        else:
            print("Invalid command")

if __name__ == "__main__":
    main()
