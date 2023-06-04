package main

/*

### Problem Statement ###

Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
You may not alter the values in the list's nodes, only nodes themselves may be changed.


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	// Trivial case, just return the linked list
	if k == 1 {
		return head
	}

	var array_of_nodes = make([]*ListNode, k)
	var new_head *ListNode

	var list_length int = 0
	temp_head := head

	for temp_head != nil {
		list_length++

		if list_length == k {
			new_head = temp_head
		}
		temp_head = temp_head.Next
	}

	temp_head = head
	iterations := list_length / k

	// Iterate through the linked list
	for iter := 0; iter < iterations; iter++ {

		// iterate until we hit the k node
		for i := 0; i < k; i++ {

			// Check if we need to point to the next k nodes
			if iter > 0 && i == 0 {

				// Point to the beginning of the next k nodes
				for next_node := 1; next_node < k; next_node++ {
					temp_head = temp_head.Next
				}
				array_of_nodes[k-1].Next = temp_head.Next
				temp_head = head
			}
			array_of_nodes[(k-1)-i] = head
			head = head.Next
		}
		array_of_nodes[k-1].Next = array_of_nodes[0].Next

		for j := 0; j < k-1; j++ {
			array_of_nodes[j].Next = array_of_nodes[j+1]
		}
	}
	return new_head
}
