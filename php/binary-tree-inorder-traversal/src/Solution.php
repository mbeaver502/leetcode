<?php

// https://leetcode.com/problems/binary-tree-inorder-traversal/
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
    function inorderTraversal($root)
    {
        if ($root === null) {
            return [];
        }

        $parent = [$root->val];
        $left = $this->inorderTraversal($root->left);
        $right = $this->inorderTraversal($root->right);

        return array_merge($left, $parent, $right);
    }
}