<?php
declare(strict_types=1);


use PHPUnit\Framework\TestCase;
use PHPUnit\Framework\Attributes\DataProvider;

final class SolutionTest extends TestCase
{
    public static function addTwoNumbersProvider(): array
    {
        return [
            [[0], [0], [0]],
            [[1], [0], [1]],
            [[1], [1], [2]],
            [[2], [2], [4]],
            [[9], [1], [0, 1]],
            [[0], [0, 1], [0, 1]],
            [[0, 1], [0], [0, 1]],
            [[1], [1, 1], [2, 1]],
            [[1, 1], [1], [2, 1]],
            [[2, 4, 3], [5, 6, 4], [7, 0, 8]],
            [[9, 9, 9, 9, 9, 9, 9], [9, 9, 9, 9], [8, 9, 9, 9, 0, 0, 0, 1]],
        ];
    }

    private function toList(array $arr): ?ListNode
    {
        if (count($arr) == 0) {
            return new ListNode; // defaults to 0 with no next
        }

        return new ListNode($arr[0], $this->toList(array_slice($arr, 1)));
    }

    private function toArray(ListNode $list): array
    {
        $out = [];
        while ($list !== null) {
            if ($list->val === 0 && $list->next === null) {
                break;
            }

            $out[] = $list->val;
            $list = $list->next;
        }
        return $out;
    }

    #[DataProvider('addTwoNumbersProvider')]
    public function testAddTwoNumbers(array $l1, array $l2, array $expected): void
    {
        $solution = new Solution;

        $list1 = $this->toList($l1);
        $list2 = $this->toList($l2);

        $result = $solution->addTwoNumbers($list1, $list2);

        $this->assertEquals($this->toArray($result), $expected);
    }
}