class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<int> st;
        for (int i = 0; i < tokens.size(); i++) {
            string token = tokens[i];
            if (!isOperator(token)) {
                int intToken = convertToInt(token);
                st.push(intToken);
            } else {
                int t1 = st.top(); st.pop();
                int t2 = st.top(); st.pop();
                int result = eval(t2, t1, token);
                st.push(result);
            }
        }
        
        return st.top();
    }

    bool isOperator(string s) {
        return (s == "+" || s == "-" || s == "*" || s == "/");
    }

    int eval(int a, int b, string op) {
        if (op == "+") {
            return a + b;
        }
        if (op == "-") {
            return a - b;
        }
        if (op == "*") {
            return a * b;
        }

        return a / b;
    }

    int convertToInt(string s) {
        bool isPositive = true;
        if (s[0] == '-') {
            isPositive = false;
        }

        int num = 0;
        for (int i = (isPositive ? 0 : 1); i < s.size(); i++) {
            num = (10 * num) + (s[i] - '0');
        }

        return (isPositive ? num : -1 * num);
    }
};
