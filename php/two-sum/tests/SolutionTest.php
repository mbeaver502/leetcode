<?php
declare(strict_types=1);


use PHPUnit\Framework\TestCase;
use PHPUnit\Framework\Attributes\DataProvider;

final class SolutionTest extends TestCase
{
    public static function twoSumProvider(): array
    {
        return [
            [[2, 7, 11, 15], 9, [0, 1]],
            [[3, 2, 4], 6, [1, 2]],
            [[3, 3], 6, [0, 1]],
            [[], 1, [-1, -1]],
            [[0, 1, 2, 3, 4, 5, 6, 7, 8, 9], 17, [8, 9]],
        ];
    }

    #[DataProvider('twoSumProvider')]
    public function testTwoSum(array $nums, int $target, array $expected): void
    {
        $solution = new Solution;

        $result = $solution->twoSum($nums, $target);

        $this->assertContains($expected[0], $result);
        $this->assertContains($expected[1], $result);

        if ($result[0] != -1) {
            $this->assertNotEquals($expected[0], $expected[1]);
        }
    }
}