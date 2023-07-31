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
        $l1Next = null;
        if ($l1 !== null) {
            $a = $l1->val;
            $l1Next = $l1->next;
        }

        $b = 0;
        $l2Next = null;
        if ($l2 !== null) {
            $b = $l2->val;
            $l2Next = $l2->next;
        }

        $c = $a + $b;

        if ($c >= 10) {
            $c %= 10;

            // Move the carry digit forward
            if ($l1Next !== null) {
                $l1Next->val += 1;
            } else if ($l2Next !== null) {
                $l2Next->val += 1;
            } else {
                $l1Next = new ListNode(1, null);
            }
        }

        return new ListNode(
            $c,
            $this->addTwoNumbers($l1Next, $l2Next)
        );
    }
}