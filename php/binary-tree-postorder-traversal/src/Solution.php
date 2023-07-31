<?php

// https://leetcode.com/problems/binary-tree-postorder-traversal/
// Definition for a binary tree node.
class TreeNode
{
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($val = 0, $left = null, $right = null)
    {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
    }
}

class Solution
{

    /**
     * @param ?TreeNode $root
     * @return int[]
     */
    function postorderTraversal($root)
    {
        if ($root === null) {
            return [];
        }

        $parent = [$root->val];
        $left = $this->postorderTraversal($root->left);
        $right = $this->postorderTraversal($root->right);

        return array_merge($left, $right, $parent);
    }
}