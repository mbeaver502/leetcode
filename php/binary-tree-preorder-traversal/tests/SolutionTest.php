<?php
declare(strict_types=1);


use PHPUnit\Framework\TestCase;
use PHPUnit\Framework\Attributes\DataProvider;

final class SolutionTest extends TestCase
{
    public static function preorderTraversalProvider(): array
    {
        return [
            [
                new TreeNode(1, null, new TreeNode(2, new TreeNode(3, null, null), null)),
                [1, 2, 3]
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

    #[DataProvider('preorderTraversalProvider')]
    public function testPreorderTraversal(?TreeNode $tree, array $expected): void
    {
        $solution = new Solution;
        $result = $solution->preorderTraversal($tree);
        $this->assertEquals($result, $expected);
    }
}