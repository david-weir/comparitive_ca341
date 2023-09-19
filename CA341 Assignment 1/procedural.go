package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Tree struct {
	Root *Node
}

// Node structure with name phone and address strings and Left and Right child nodes
type Node struct {
	Name    string
	Phone   string
	Address string
	Left    *Node
	Right   *Node
}

// insert a node to the tree using the name
func (t *Tree) insertName(name, phone string, address string) error {

	if t.Root == nil { // if the tree is empty create a new node
		t.Root = &Node{Name: name, Phone: phone, Address: address}
		return nil
	}

	// else insert a node
	return t.Root.insertName(name, phone, address)
}

// insert a node to the tree using the phone number
func (t *Tree) insertPhone(name, phone string, address string) error {

	if t.Root == nil { // if the tree is empty create a new node
		t.Root = &Node{Name: name, Phone: phone, Address: address}
		return nil
	}

	// else insert a node
	return t.Root.insertPhone(name, phone, address)
}

func (n *Node) insertName(name, phone string, address string) error {
	// Function to insert node based on name

	// if it is empty we cannot add it
	if n == nil {
		return errors.New("Cannot insert an entry into a nil tree")
	}

	switch {

	// if the node exists already
	case name == n.Name:
		return nil

	// if the name is less than the current node and the left child is empty insert a node on the left
	// else call insert on the left subtree
	case name < n.Name:

		if n.Left == nil {
			n.Left = &Node{Name: name, Phone: phone, Address: address}
			return nil
		}

		return n.Left.insertName(name, phone, address)

	// if the name is greater than the current node and the right child is empty insert a node on the right
	// else call insert on the right subtree
	case name > n.Name:

		if n.Right == nil {
			n.Right = &Node{Name: name, Phone: phone, Address: address}
			return nil
		}

		return n.Right.insertName(name, phone, address)
	}

	return nil
}

func (n *Node) insertPhone(name, phone string, address string) error {
	// Function to insert based on phone number

	// error check
	if n == nil {
		return errors.New("Cannot insert an entry into a nil tree")
	}

	switch {

	// if the node already exists
	case phone == n.Phone:
		return nil

	// if the name is less than the current node and the left child is empty insert a node on the left
	// else call insert on the left subtree
	case phone < n.Phone:

		if n.Left == nil {
			n.Left = &Node{Name: name, Phone: phone, Address: address}
			return nil
		}

		return n.Left.insertPhone(name, phone, address)

	// if the name is greater than the current node and the right child is empty insert a node on the right
	// else call insert on the right subtree
	case phone > n.Phone:

		if n.Right == nil {
			n.Right = &Node{Name: name, Phone: phone, Address: address}
			return nil
		}

		return n.Right.insertPhone(name, phone, address)
	}

	return nil
}

// search for a node by name
func (t *Tree) findName(name string) []string {

	// search tree unless the root node is empty
	if t.Root == nil {
		return []string{"", "", ""}
	}

	return t.Root.findName(name)
}

// search for a node by phone number
func (t *Tree) findPhone(phone string) []string {

	// search tree unless the root node is empty
	if t.Root == nil {
		return []string{"j", "", ""}
	}

	return t.Root.findPhone(phone)
}

// search for a node with the given name
func (n *Node) findName(name string) []string {

	if n == nil {
		return []string{"", "", ""}
	}

	switch {

	// if we find the correct node
	case name == n.Name:
		return []string{n.Name, n.Phone, n.Address}

	// if the given name is less than the current node search the left subtree
	case name < n.Name:
		return n.Left.findName(name)

	// else search the right subtree
	default:
		return n.Right.findName(name)
	}

}

// search for a node with the given phone number
func (n *Node) findPhone(phone string) []string {

	if n == nil {
		return []string{"", "", ""}
	}

	switch {

	// if we find the correct node
	case phone == n.Phone:
		return []string{n.Name, n.Phone, n.Address}

	// if the given name is less than the current node search the left subtree
	case phone < n.Phone:
		return n.Left.findPhone(phone)

	// else search the right subtree
	default:
		return n.Right.findPhone(phone)
	}

}

// finds the largest node in a tree / subtree
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}

	if n.Right == nil {
		return n, parent
	}

	return n.Right.findMax((n))
}

// replaces the parent's child pointer with a pointer to the replacement node
func (n *Node) replaceNode(parent, replace *Node) error {
	if n == nil {
		return errors.New("Cannot replace node in empty tree")
	}

	if n == parent.Left {
		parent.Left = replace
		return nil
	}

	parent.Right = replace
	return nil
}

// delete a node using a given name
func (t *Tree) deleteName(name string) error {

	// error if tree is empty
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	// avoids treating the root node as a special case
	fakeparent := &Node{Right: t.Root}
	err := t.Root.deleteName(name, fakeparent)

	if err != nil {
		return err
	}

	// if the tree consists of only a root node, set the root to nil
	if fakeparent.Right == nil {
		t.Root = nil
	}

	return nil
}

// removes a node from the tree with the given name
func (n *Node) deleteName(name string, parent *Node) error {

	// if the node does not exist
	if n == nil {
		return errors.New("Node does not exist")
	}

	// search the tree to find the given node
	switch {
	case name < n.Name:
		return n.Left.deleteName(name, n)

	case name > n.Name:
		return n.Right.deleteName(name, n)

	// after finding the node remove it from the parent node
	default:

		// if the node has no child nodes
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		// if the node only has one child, replace the node with the child
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}

		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		// if the node has two child nodes find the maximum node in the left subtree
		replace, replacepar := n.Left.findMax(n)

		// replace the nodes data
		n.Name, n.Phone, n.Address = replace.Name, replace.Phone, replace.Address

		// remove the replacement node
		return replace.deleteName(replace.Name, replacepar)
	}
}

// delete a node using a given phone number
func (t *Tree) deletePhone(name string) error {

	// error if tree is empty
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	// avoids treating the root node as a special case
	fakeparent := &Node{Right: t.Root}
	err := t.Root.deletePhone(name, fakeparent)

	if err != nil {
		return err
	}

	// if the tree consists of only a root node, set the root to nil
	if fakeparent.Right == nil {
		t.Root = nil
	}

	return nil
}

// removes a node from the tree with the given number
func (n *Node) deletePhone(name string, parent *Node) error {

	// if the node does not exist
	if n == nil {
		return errors.New("Node does not exist")
	}

	// search the tree to find the given node
	switch {
	case name < n.Phone:
		return n.Left.deletePhone(name, n)

	case name > n.Phone:
		return n.Right.deletePhone(name, n)

	// after finding the node remove it from the parent node
	default:

		// if the node has no child nodes
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		// if the node only has one child, replace the node with the child
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}

		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		// if the node has two child nodes find the maximum node in the left subtree
		replace, replacepar := n.Left.findMax(n)

		// replace the nodes data
		n.Name, n.Phone, n.Address = replace.Name, replace.Phone, replace.Address

		// remove the replacement node
		return replace.deletePhone(replace.Name, replacepar)
	}
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {

	// Traverses the tree using inorder traversal (for testing)
	if n == nil {
		return
	}

	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func main() {

	// read input
	reader := bufio.NewReader(os.Stdin)

	// initialze trees
	nameTree := &Tree{}
	phoneTree := &Tree{}

	// loop command shell
	for {

		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// get args
		var commandStr = strings.TrimSuffix(cmdString, "\n")
		arrCommandStr := strings.Fields(commandStr)

		switch arrCommandStr[0] {

		// commands
		case "exit":
			os.Exit(0)

		case "help":
			fmt.Println("CA341 Procedural BST Manual")
			fmt.Println("Supported Commands:")
			fmt.Println()

			fmt.Println("insertname: inserts a node by name")
			fmt.Println("insertphone: inserts a node by phone number")
			fmt.Println()

			fmt.Println("findname: searches the tree for a node by name")
			fmt.Println("findphone: searches the tree for a node by phone")

			fmt.Println()
			fmt.Println("deletename: deletes the node with the given name")
			fmt.Println("deletephone: deletes the node with the given name")

		case "insertname":

			nameTree.insertName(arrCommandStr[1], arrCommandStr[2], strings.Join(arrCommandStr[3:], " "))

			nameTree.Traverse(nameTree.Root, func(n *Node) {
				fmt.Print("Name: ", n.Name, "\nPhone Number: ", n.Phone, "\nAddress: ", n.Address, "\n\n")
			})

		case "insertphone":

			phoneTree.insertPhone(arrCommandStr[1], arrCommandStr[2], strings.Join(arrCommandStr[3:], " "))

			phoneTree.Traverse(phoneTree.Root, func(n *Node) {
				fmt.Print("Name: ", n.Name, "\nPhone Number: ", n.Phone, "\nAddress: ", n.Address, "\n\n")
			})

		case "findname":

			result := nameTree.findName(arrCommandStr[1])

			fmt.Println("Name: ", result[0], "\nPhone Number: ", result[1], "\nAddress: ", result[2])

		case "findphone":
			result := phoneTree.findPhone(arrCommandStr[1])

			fmt.Println("Name: ", result[0], "\nPhone Number: ", result[1], "\nAddress: ", result[2])

		case "deletename":
			nameTree.deleteName(arrCommandStr[1])

			nameTree.Traverse(nameTree.Root, func(n *Node) {
				fmt.Print("Name: ", n.Name, "\nPhone Number: ", n.Phone, "\nAddress: ", n.Address, "\n\n")
			})

		case "deletephone":
			phoneTree.deletePhone(arrCommandStr[1])

			phoneTree.Traverse(phoneTree.Root, func(n *Node) {
				fmt.Print("Name: ", n.Name, "\nPhone Number: ", n.Phone, "\nAddress: ", n.Address, "\n\n")
			})
		}

		cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

	}
}
