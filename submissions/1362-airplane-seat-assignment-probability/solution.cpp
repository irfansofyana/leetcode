class Solution {
public:
    double nthPersonGetsNthSeat(int n) {
        if (n == 1) {
            return 1.00000;
        }

        return 0.50000;
    }
};

// n-1 * (d(n-1) + d(n-2))

// d(1) = 0
// d(2) = 1
// d(3) = 2
// d(4) = 9
// d(5) = 44
// d(6) = 265

// (n-1)! * (n-1) = 3 * 6 = 18
