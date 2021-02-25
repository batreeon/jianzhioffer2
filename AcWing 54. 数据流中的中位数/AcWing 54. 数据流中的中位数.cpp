class Solution {
public:
    priority_queue<int> maxHeap;
    priority_queue<int,vector<int>,greater<int>> minHeap;
    void insert(int num){
        if ((maxHeap.size()+minHeap.size())&1==1) {
            maxHeap.push(num);
            minHeap.push(maxHeap.top());
            maxHeap.pop();
        }else{
            minHeap.push(num);
            maxHeap.push(minHeap.top());
            minHeap.pop();
        }
    }

    double getMedian(){
        if ((maxHeap.size()+minHeap.size())&1==1) {
            return (double)maxHeap.top();
        }else{
            return ((double)maxHeap.top()+(double)minHeap.top())/2.0;
        }
    }
};