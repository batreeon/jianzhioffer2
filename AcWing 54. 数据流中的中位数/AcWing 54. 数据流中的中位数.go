type minHeap []int
func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i,j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i,j int) { h[i],h[j] = h[j],h[i] }
func (h *minHeap) Push(x interface{}) {
    *h = append(*h,x.(int))
}
func (h *minHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type maxHeap []int
func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i,j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i,j int) { h[i],h[j] = h[j],h[i] }
func (h *maxHeap) Push(x interface{}) {
    *h = append(*h,x.(int))
}
func (h *maxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

max := &maxHeap{}
min := &minHeap{}
heap.Init(max)
heap.Init(min)

func insert(num int) {
    if (heap.Len(min)+heap.Len(max))&1 == 1 {
        heap.Push(max,num)
        heap.Push(min,heap.Pop(max))
    }else{
        heap.Push(min,num)
        heap.Push(max,heap.Pop(min))
    }
}
func getMedian() float64{
    if (heap.Len(min)+heap.Len(max))&1 == 1 {
        return float64(heap.Pop(max))
    }else{
        heap.Push(min,num)
        heap.Push(max,heap.Pop(min))
        return (float64(heap.Pop(max))+float64(heap.Pop(min)))/2.0
    }
}