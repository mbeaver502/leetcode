<?php
declare(strict_types=1);


use PHPUnit\Framework\TestCase;
use PHPUnit\Framework\Attributes\DataProvider;

final class SolutionTest extends TestCase
{
    public static function sortedArrayToBSTProvider(): array
    {
        return [
            [[-10, -3, 0, 5, 9], [0, -3, 9, -10, 5]],
            [[1, 3], [3, 1]],
        ];
    }

    // Create an array using breadth-first traversal, ignore "null" leaves
    private function toArray(?TreeNode $tree): array
    {
        $queue = [$tree];
        $out = [];

        while (count($queue) !== 0) {
            $current = array_pop($queue);
            $out[] = $current->val;

            if ($current->left !== null) {
                array_unshift($queue, $current->left);
            }

            if ($current->right !== null) {
                array_unshift($queue, $current->right);
            }
        }

        return $out;
    }

    #[DataProvider('sortedArrayToBSTProvider')]
    public function testSortedArrayToBST(array $nums, array $expected): void
    {
        $solution = new Solution;

        $result = $solution->sortedArrayToBST($nums);

        $this->assertEquals($expected, $this->toArray($result));
    }
}