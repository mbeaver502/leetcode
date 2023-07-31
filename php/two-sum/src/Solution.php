<?php

// https://leetcode.com/problems/two-sum/

class Solution
{

    /**
     * @param int[] $nums
     * @param int $target
     * @return int[]
     */
    function twoSum($nums, $target)
    {
        $map = [];
        for ($i = 0; $i < count($nums); $i++) {
            $num = $nums[$i];
            $diff = $target - $num;
            if (array_key_exists($diff, $map)) {
                return [$map[$diff], $i];
            }
            $map[$num] = $i;
        }
        return [-1, -1];
    }
}