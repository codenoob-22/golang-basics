/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var pointer *ListNode

func helper(x *ListNode) bool {
    if( x.Next == nil) {
        return pointer.Val == x.Val
    }
    
    if (helper(x.Next)) {
        pointer = pointer.Next
        return pointer.Val == x.Val
    }
    
    return false
}
func isPalindrome(head *ListNode) bool {
    if(head == nil) {
        return true
    }
    pointer = head
    return helper(head)
}
