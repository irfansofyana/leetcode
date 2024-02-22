class ProductOfNumbers {
public:
    int * arr;
    int last;
    
    ProductOfNumbers() {
        arr = new int[40001];
        last = -1;
    }
    
    void add(int num) {
        ++last;
        arr[last] = num;
    }
    
    int getProduct(int k) {
        int ret = 1;
        for (int i = 0; i < k; ++i){
            ret *= arr[last-i];
        }
        return ret;
    }
};

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * ProductOfNumbers* obj = new ProductOfNumbers();
 * obj->add(num);
 * int param_2 = obj->getProduct(k);
 */