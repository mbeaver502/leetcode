<?php

// https://leetcode.com/problems/add-two-numbers/
// Definition for a singly-linked list.
class ListNode
{
    public $val = 0;
    public $next = null;
    function __construct($val = 0, $next = null)
    {
        $this->val = $val;
        $this->next = $next;
    }
}

class Solution
{
    /**
     * @param ?ListNode $l1
     * @param ?ListNode $l2
     * @return ?ListNode
     */
    function addTwoNumbers($l1, $l2)
    {
        if ($l1 === null && $l2 === null) {
            return null;
        }

        $a = 0;
        if ($l1 !== null) {
            $a = $l1->val;
            $l1 = $l1->next;
        }

        $b = 0;
        if ($l2 !== null) {
            $b = $l2->val;
            $l2 = $l2->next;
        }

        $c = $a + $b;

        if ($c >= 10) {
            $c %= 10;

            // Move the carry digit forward
            if ($l1 !== null) {
                $l1->val += 1;
            } else if ($l2 !== null) {
                $l2->val += 1;
            } else {
                $l1 = new ListNode(1, null);
            }
        }

        return new ListNode(
            $c,
            $this->addTwoNumbers($l1, $l2)
        );
    }
}