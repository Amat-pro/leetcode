// 排序

// 快速排序
fn quick_sort(nums: &mut Vec<i32>) {
    let len = nums.len();
    if len == 0 {
        return;
    }

    _quick_sort(nums, 0, len - 1);
}

fn _quick_sort(nums: &mut Vec<i32>, left: usize, right: usize) {
    // 递归 - 确定终止条件
    if left >= right {
        return;
    }

    // 这里需要注意middle如果是usize, 则middle-1可能溢出导致panic
    let middle = partition(nums, left, right);
    // 左边 [left,middle-1]
    if middle - 1 > 0 {
        _quick_sort(nums, left, middle as usize - 1);
    }
    // 右边 [middle+1,right]
    _quick_sort(nums, middle as usize + 1, right);
}

// partition [left,right]区间排序，返回中间值(pivot)的index
fn partition(nums: &mut Vec<i32>, l: usize, r: usize) -> i32 {
    let pivot = nums[l];

    let mut left = l;
    let mut right = r;

    while left < right {
        while (left < right) && (nums[right] >= pivot) {
            right -= 1;
        }
        nums[left] = nums[right];

        while (left < right) && (nums[left] <= pivot) {
            left += 1;
        }
        nums[right] = nums[left];
    }

    nums[right] = pivot;
    return right as i32;
}

// 选择排序
fn selection_sort(nums: &mut Vec<i32>) {
    let len = nums.len();

    if len == 0 {
        return;
    }

    // 每次从[i,len]范围内选一个最小的，放在i的位置。默认nums[i]是最小的
    for i in 0..len {
        let mut min_index = i;
        // j范围:[i+1,len]
        for j in i + 1..len {
            if nums[j] < nums[min_index] {
                min_index = j;
            }
        }
        nums.swap(i, min_index);
    }
}

// 冒泡排序
fn bubble_sort(nums: &mut Vec<i32>) {
    let len = nums.len();
    if len == 0 {
        return;
    }

    // 两两对比，每次大泡泡向后移动
    let mut sorted = false;
    while !sorted {
        // 开始设置为true,当两两对比都是nums[i-1]>nums[i],说明排序完成了,否则设置sorted为false
        sorted = true;
        for i in 1..len {
            // i-1 i
            if nums[i - 1] > nums[i] {
                nums.swap(i, i - 1);
                sorted = false;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 选择排序
    #[test]
    fn test_selection_sort() {
        let mut nums: Vec<i32> = vec![2, 6, 8, 9, 4, 2, 5, 8];
        println!("nums is: {:?}", nums);
        selection_sort(&mut nums);
        println!("selection_sort, result is: {:?}", nums);
    }

    // 冒泡排序
    #[test]
    fn test_bubble_sort() {
        let mut nums: Vec<i32> = vec![2, 6, 8, 9, 4, 2, 5, 8];
        println!("nums is: {:?}", nums);
        bubble_sort(&mut nums);
        println!("bubble_sort, result is: {:?}", nums);
    }

    // 快速排序
    #[test]
    fn test_quick_sort() {
        let mut nums: Vec<i32> = vec![2, 6, 8, 9, 4, 2, 5, 8];
        println!("nums is: {:?}", nums);
        quick_sort(&mut nums);
        println!("quick_sort, result is: {:?}", nums);
    }
}
