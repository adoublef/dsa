use std::{
    collections::{BTreeMap, HashSet},
    hash::Hash,
};

struct Solution;

#[allow(dead_code)]
impl Solution {
    /// https://leetcode.com/problems/two-sum/
    /// Given an array of integers nums and an integer target,
    /// return indices of the two numbers such that they add up to target.
    fn problem_1(nums: Vec<i32>, target: i32) -> Vec<i32> {
        use std::collections::HashMap as hm;
        nums.iter()
            .enumerate()
            .fold((vec![], hm::new()), |(vec, mut map), (pos, val)| {
                map.insert(val, pos as i32);
                match map.get(&(target - val)) {
                    Some(&pos_cached) => (vec![pos_cached, pos as i32], map),
                    None => (vec, map),
                }
            })
            .0
    }

    /// https://leetcode.com/problems/maximum-subarray/
    fn problem_53(nums: Vec<i32>) -> i32 {
        nums.into_iter()
            .fold((i32::MIN, 0), |(best_sum, curr_sum), n| {
                let v = curr_sum + n;
                (best_sum.max(v), 0.max(v))
            })
            .0
    }

    /// https://leetcode.com/problems/climbing-stairs/
    fn problem_70(n: i32) -> i32 {
        let g = (1.0 + 5.0_f64.sqrt()) / 2.0;
        ((g.powi(n + 1) - (1.0 - g).powi(n + 1)) / 5.0_f64.sqrt()).round() as i32
    }

    /// You want to maximize your profit by choosing a single
    /// day to buy one stock and choosing a different day
    /// in the future to sell that stock.
    fn problem_121(prices: Vec<i32>) -> i32 {
        prices
            .iter()
            .fold((i32::MAX, 0), |(buy, profit), sell| {
                (buy.min(*sell), profit.max(sell - buy))
            })
            .1
    }

    /// https://leetcode.com/problems/contains-duplicate/
    fn problem_217<T>(nums: T) -> bool
    where
        T: IntoIterator,
        T::Item: Eq + Hash,
    {
        let mut m = HashSet::new();
        !nums.into_iter().all(|x| m.insert(x))
    }

    /// https://leetcode.com/problems/product-of-array-except-self/
    /// Given an integer array nums,
    /// return an array answer such that answer[i] is equal
    /// to the product of all the elements of nums except nums[i].
    fn problem_238(nums: Vec<i32>) -> Vec<i32> {
        // Why is this O(n) & not O(n*n)?
        //
        // One loop: O(n)  Another loop: O(n)  Total -> O(2n). We remove the constant 2, so we have O(n).
        // In this case, we only ever traverse a given array twice. An array of a thousand elements is only two traversals.
        // If loops were nested, then we would make one extra traversal for every element in the array i.e. O(n*n).

        // this would be an example of O(n*n) as I have nested looping
        let _slow = nums
            .iter()
            .enumerate()
            .map(|(k, _)| {
                nums[..k].iter().product::<i32>() * nums[(k + 1)..].iter().product::<i32>()
            })
            .collect::<Vec<_>>();

        use std::mem::replace;

        let prefixes = nums.iter().scan(1, |acc, n| Some(replace(acc, *acc * n)));

        let suffixes = nums
            .iter()
            .rev()
            .scan(1, |acc, n| Some(replace(acc, *acc * n)))
            .collect::<Vec<_>>()
            .into_iter() // the eq. of using a `let` binding to create a longer lived value
            .rev();

        prefixes
            .zip(suffixes)
            .map(|(a, b)| a * b)
            .collect::<Vec<_>>()
    }

    /// https://leetcode.com/problems/valid-anagram/
    /// Given two strings s and t,
    /// return true if t is an anagram of s,
    /// and false otherwise.
    fn problem_242(s: String, t: String) -> bool {
        use std::collections::HashMap as hm;

        if s.len() == t.len() {
            s.chars()
                .zip(t.chars())
                .fold(hm::<char, i32>::new(), |mut map, (x, y)| {
                    *map.entry(x).or_default() += 1;
                    *map.entry(y).or_default() -= 1;
                    map
                })
                .values()
                .all(|&x| x == 0)
        } else {
            false
        }
    }

    /// https://leetcode.com/explore/featured/card/the-leetcode-beginners-guide/
    fn problem_412(n: i32) -> Vec<String> {
        [1..=n].iter().enumerate().fold(Vec::new(), |res, (x, _)| {
            let mut val = String::new();

            if x % 3 == 0 {
                val += "Fizz"
            }

            if x % 5 == 0 {
                val += "Buzz"
            }

            if val.is_empty() {
                val = x.to_string();
            }

            [res, vec![val]].concat()
        })
    }

    // TODO: make generic
    /// https://leetcode.com/problems/fibonacci-number/
    fn problem_509(n: i32) -> i32 {
        (0..n).fold((0, 1), |(a, b), _| (b, a + b)).0
    }

    fn problem_740(nums: Vec<i32>) -> i32 {
        // as I need it to be sorted better to use this
        // than a HashMap
        let (a, b) = nums
            .iter()
            .fold(BTreeMap::new(), |mut m, n| {
                *m.entry(n).or_insert(0) += 1;
                m
            })
            .iter()
            .fold(
                (0, 0),
                |(a, b), (k, v)| {
                    if a == *k - 1 {
                        (a, b)
                    } else {
                        (b, b + v)
                    }
                },
            );

        a.max(b)
    }

    /// https://leetcode.com/problems/min-cost-climbing-stairs/
    fn problem_746(cost: Vec<i32>) -> i32 {
        cost.iter()
            .chain(std::iter::once(&0))
            .fold((0, 0), |(x, y), n| (y, n + x.min(y)))
            .1
    }

    // TODO: make generic
    /// https://leetcode.com/problems/n-th-tribonacci-number/
    fn problem_1137(n: i32) -> i32 {
        (0..n).fold((0, 1, 1), |(a, b, c), _| (b, c, a + b + c)).0
    }

    /// https://leetcode.com/problems/running-sum-of-1d-array/
    fn problem_1480(nums: Vec<i32>) -> Vec<i32> {
        nums.iter()
            .fold((Vec::new(), 0), |(res, mut acc), n| {
                acc += n;
                ([res, vec![acc]].concat(), acc)
            })
            .0
    }

    /// https://leetcode.com/problems/richest-customer-wealth/
    fn problem_1672(nums: Vec<i32>) -> Vec<i32> {
        nums.iter()
            .fold((Vec::new(), 0), |(res, mut acc), n| {
                acc += n;
                ([res, vec![acc]].concat(), acc)
            })
            .0
    }
}

#[cfg(test)]
pub mod test {
    use super::*;

    #[test]
    fn test_1() {
        let run = |(nums, target), expected| {
            let res = Solution::problem_1(nums, target);
            assert_eq!(res, expected)
        };

        run((vec![2, 7, 11, 15], 9), vec![0, 1]);
        run((vec![3, 2, 4], 6), vec![1, 2]);
        run((vec![3, 3], 6), vec![0, 1]);
    }

    #[test]
    fn test_53() {
        assert_eq!(Solution::problem_53(vec![-2, 1, -3, 4, -1, 2, 1, -5, 4]), 6);
        assert_eq!(Solution::problem_53(vec![1]), 1);
        assert_eq!(Solution::problem_53(vec![-1]), -1);
        assert_eq!(Solution::problem_53(vec![5, 4, -1, 7, 8]), 23)
    }

    #[test]
    fn test_70() {
        assert_eq!(Solution::problem_70(1), 1);
        assert_eq!(Solution::problem_70(6), 13);
        assert_eq!(Solution::problem_70(9), 55);
        assert_eq!(Solution::problem_70(13), 377);
        assert_eq!(Solution::problem_70(18), 4181);
    }

    #[test]
    fn problem_121() {
        let input = vec![7, 6, 4, 3, 1];
        let result = Solution::problem_121(input);
        assert_eq!(result, 0)
    }

    #[test]
    fn test_217() {
        assert!(Solution::problem_217(vec![1, 2, 3, 1]));
        assert!(!Solution::problem_217(vec![1, 2, 3, 4]));
    }

    #[test]
    fn problem_238() {
        let input = vec![1, 2, 3, 4];
        let result = Solution::problem_238(input);
        assert_eq!(result, vec![24, 12, 8, 6])
    }

    #[test]
    fn test_242() {
        let (s, t) = ("cat", "tac");
        let result = Solution::problem_242(s.to_string(), t.to_string());
        assert!(result)
    }

    #[test]
    fn test_412() {
        let run = |nums, expected| {
            let res = Solution::problem_412(nums);
            assert_eq!(res, expected)
        };

        run(3, vec!["1", "2", "Fizz"]);
        run(5, vec!["1", "2", "Fizz", "4", "Buzz"]);
        run(
            15,
            vec![
                "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz",
                "13", "14", "FizzBuzz",
            ],
        );
    }

    #[test]
    fn test_509() {
        assert_eq!(Solution::problem_509(0), 0);
        assert_eq!(Solution::problem_509(2), 1);
        assert_eq!(Solution::problem_509(4), 3);
    }

    #[test]
    fn test_746() {
        assert_eq!(Solution::problem_746(vec![10, 15, 20]), 15);
        assert_eq!(
            Solution::problem_746(vec![1, 100, 1, 1, 1, 100, 1, 1, 100, 1]),
            6
        );
    }

    #[test]
    fn test_1137() {
        assert_eq!(Solution::problem_1137(4), 4);
        assert_eq!(Solution::problem_1137(25), 1389537);
    }

    #[test]
    fn test_1672() {
        let run = |nums, expected| {
            let res = Solution::problem_1672(nums);
            assert_eq!(res, expected)
        };

        run(vec![1, 2, 3, 4], vec![1, 3, 6, 10]);
        run(vec![1, 1, 1, 1, 1], vec![1, 2, 3, 4, 5]);
        run(vec![3, 1, 2, 10, 1], vec![3, 4, 6, 16, 17]);
    }
}
