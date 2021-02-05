func verifySequenceOfBST(sequence []int) bool{
    if len(sequence) == 0 {
        return true
    }
	more := 0	//找左右子树分界点
	for more < len(sequence)-1 && sequence[more] < sequence[len(sequence)-1] {
	    more++
	}
	left := sequence[:more]
	right := sequence[more:len(sequence)-1]
	for i := 0 ; i < len(left) ; i++ {
	    if left[i] > sequence[len(sequence)-1] {
	        return false
	    }
	}
	for i := 0 ; i < len(right) ; i++ {
	    if right[i] < sequence[len(sequence)-1] {
	        return false
	    }
	}
	return verifySequenceOfBST(left) && verifySequenceOfBST(right)
}
//后序遍历，根节点在末尾，左子树（小于等于根节点）在前部，右子树（大于等于根节点）在后部，根据这个特点进行判断
//在递归地对左子树和右子树进行判断