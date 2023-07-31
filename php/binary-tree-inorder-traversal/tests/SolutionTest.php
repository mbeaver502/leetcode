<?php
declare(strict_types=1);


use PHPUnit\Framework\TestCase;
use PHPUnit\Framework\Attributes\DataProvider;

final class SolutionTest extends TestCase
{
    public static function inorderTraversalProvider(): array
    {
        return [
            [
                new TreeNode(1, null, new TreeNode(2, new TreeNode(3, null, null), null)),
                [1, 3, 2]
            ],
            [
                null,
                []
            ],
            [
                new TreeNode(1, null, null),
                [1]
            ],
        ];
    }

    #[DataProvider('inorderTraversalProvider')]
    public function testInorderTraversal(?TreeNode $tree, array $expected): void
    {
        $solution = new Solution;
        $result = $solution->inorderTraversal($tree);
        $this->assertEquals($result, $expected);
    }
}