/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int findSum(TreeNode * root, int level, queue<int> state) {
        if (root == NULL) {
            return 0;
        } else {
            int isIncluded = 0;
            if (!state.empty()) {
                if (level == state.front()) {
                    state.pop();
                    isIncluded += root -> val;
                }
            }
            if ((root->val) % 2 == 0) state.push(level + 2);
            return findSum(root->left, level + 1, state) + findSum(root->right, level + 1, state) + isIncluded; 
        }
    }
    
    int sumEvenGrandparent(TreeNode* root) {
        queue<int> st;
        return findSum(root, 0, st);
    }
};
