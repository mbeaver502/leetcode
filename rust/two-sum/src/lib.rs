// https://leetcode.com/problems/two-sum
#[allow(dead_code)]
struct Solution;

impl Solution {
    #[allow(dead_code)]
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        use std::collections::HashMap;

        let mut map: HashMap<i32, i32> = HashMap::new();
        for (i, &n) in nums.iter().enumerate() {
            if let Some(&v) = map.get(&(target - n)) {
                return vec![v, i as i32];
            }
            map.insert(n, i as i32);
        }

        vec![-1, -1]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum() {
        struct Test {
            nums: Vec<i32>,
            target: i32,
            want: Vec<i32>,
        }

        let tests = vec![
            Test {
                nums: vec![2, 7, 11, 15],
                target: 9,
                want: vec![0, 1],
            },
            Test {
                nums: vec![3, 2, 4],
                target: 6,
                want: vec![1, 2],
            },
            Test {
                nums: vec![3, 3],
                target: 6,
                want: vec![0, 1],
            },
            Test {
                nums: vec![],
                target: 100,
                want: vec![-1, -1],
            },
            Test {
                nums: vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
                target: 19,
                want: vec![8, 9],
            },
        ];

        for test in tests {
            let got = Solution::two_sum(test.nums, test.target);
            assert!(
                (got[0] == test.want[0] || got[0] == test.want[1])
                    && (got[1] == test.want[0] || got[1] == test.want[1])
            );
        }
    }
}
