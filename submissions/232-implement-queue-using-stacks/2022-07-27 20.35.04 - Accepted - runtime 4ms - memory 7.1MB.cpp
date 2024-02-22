class MyQueue {
public:
    bool isStack; // true if it's stack mode, false if it's queue mode
    stack<int> st[2]; // st[0] for stack, st[1] for queue
    
    MyQueue() {
        this->isStack = true;
    }
    
    void push(int x) {
        if (this->isStack) {
            st[0].push(x);
        } else {
            while (!st[1].empty()) {
                int curr = st[1].top();
                st[0].push(curr);
                st[1].pop();
            }
            st[0].push(x);
            this->isStack = true;
        }
    }
    
    int pop() {
        int val;
        if (this->isStack) {
            while (!st[0].empty()) {
                int curr = st[0].top();
                st[1].push(curr);
                st[0].pop();
            }
            val = st[1].top();
            st[1].pop();
            this->isStack = false;
        } else {
            val = st[1].top();
            st[1].pop();
        }
        
        return val;
    }
    
    int peek() {
        int val;
        if (this->isStack) {
            while (!st[0].empty()) {
                int curr = st[0].top();
                st[1].push(curr);
                st[0].pop();
            }
            val = st[1].top();
            this->isStack = false;
        } else {
            val = st[1].top();
        }
        
        return val;
    }
    
    bool empty() {
        if (this->isStack) {
            return st[0].size() == 0;
        } 
        return st[1].size() == 0;
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */