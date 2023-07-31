<?php

// https://leetcode.com/problems/binary-tree-preorder-traversal/
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
    function preorderTraversal($root)
    {
        if ($root === null) {
            return [];
        }

        $parent = [$root->val];
        $left = $this->preorderTraversal($root->left);
        $right = $this->preorderTraversal($root->right);

        return array_merge($parent, $left, $right);
    }
}