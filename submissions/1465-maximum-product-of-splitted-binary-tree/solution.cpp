/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
public:
    long long sum_all;
    long long ans;
    long long mod = 1e9+7;
    
    long long find_sum(TreeNode * root){
        long long res = root->val;
        if (root->left != NULL){
            res = (res + find_sum(root->left))%mod;
        }
        if (root->right != NULL){
            res = (res + find_sum(root->right))%mod;
        }
        return res;
    }
    
    long long find_sum_below(TreeNode * root){
        long long t_sum = root->val;
        
        if (root->left != NULL){
            long long one = find_sum_below(root->left);
            t_sum = (t_sum + one) % mod;
            long long two = (sum_all - one + mod)%mod;
            if (one * two > ans){
                ans = one*two;
            }
        }
        if (root->right != NULL){
            long long one = find_sum_below(root->right);
            t_sum = (t_sum + one) % mod;
            long long two = (sum_all - one + mod)%mod;
            if ( one * two > ans){
                ans = (one * two);
            }
        }
        return t_sum;
    }
    
    long long maxProduct(TreeNode* root) {
        sum_all = find_sum(root);
        ans = 0;
        long long t = find_sum_below(root);
        return ans%mod;
    }
};
