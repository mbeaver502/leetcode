<?php

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
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
     * @param int[] $nums
     * @return ?TreeNode
     */
    function sortedArrayToBST($nums)
    {
        $len = count($nums);

        // Probably unnecessary to check this, but you never know...
        if ($len < 1) {
            return null;
        } else if ($len == 1) {
            return new TreeNode($nums[0], null, null);
        }

        $halfway = floor($len / 2);
        return new TreeNode(
            $nums[$halfway],
            $this->sortedArrayToBST(array_slice($nums, 0, $halfway)),
            $this->sortedArrayToBST(array_slice($nums, $halfway + 1))
        );
    }
}