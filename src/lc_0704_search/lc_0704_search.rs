
pub fn search(nums: Vec<i32>, target: i32) -> i32 {

    if nums.len() == 0 {
        return -1;
    }

    // 二分搜索 - [left, right]
    let mut left: usize = 0;
    let mut right: usize = nums.len();
    let mut middle: usize;

    while left <= right {

        middle = (left + right) / 2;
        if target < nums[middle] {
            right = middle - 1;
        } else if target > nums[middle] {
            left = middle + 1;
        } else {
            return middle as i32;
        }

    }

    return -1;

}

#[cfg(test)]
mod test {

    use super::*;
    #[test]
    fn test_search() {
        let nums = vec![-1, 0, 3, 5, 9, 12];
        println!("target is 9, result is: {}", search(nums.clone(), 9));
        println!("target is 2, result is: {}", search(nums, 2));
    }
}