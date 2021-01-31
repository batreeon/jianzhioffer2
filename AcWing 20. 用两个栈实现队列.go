class MyQueue {
    stack<int> s1,s2;
public:
    /** Initialize your data structure here. */
    MyQueue() {
        
    }
    
    /** Push element x to the back of queue. */
    void push(int x) {
        s1.push(x);
    }
    
    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        while(s1.size()){
            s2.push(s1.top());
            s1.pop();
        }
        int pop = s2.top();
        s2.pop();
        while(s2.size()){
            s1.push(s2.top());
            s2.pop();
        }
        return pop;
    }
    
    /** Get the front element. */
    int peek() {
        while(s1.size()){
            s2.push(s1.top());
            s1.pop();
        }
        int pop = s2.top();
        while(s2.size()){
            s1.push(s2.top());
            s2.pop();
        }
        return pop;
    }
    
    /** Returns whether the queue is empty. */
    bool empty() {
        return s1.empty();
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue obj = MyQueue();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.peek();
 * bool param_4 = obj.empty();
 */