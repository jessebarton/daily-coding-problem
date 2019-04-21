


# Problem --- Given an array of integers, return a new array such that each element at
# index i of the new array is the product of all the numbers
# in the original array excpet the one at i.
# For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]
nums = [1, 2, 3, 4, 5]
def products(nums):
    prefix_numbers = []
    for num in nums:
        if prefix_numbers:
            prefix_numbers.append(prefix_numbers[-1] * num)
        else:
            prefix_numbers.append(num)
    
    suffix_numbers = []
    for num in reversed(nums):
        if suffix_numbers:
            suffix_numbers.append(suffix_numbers[-1] * num)
            print(suffix_numbers[-1])
        else:
            suffix_numbers.append(num)
    suffix_numbers = list(reversed(suffix_numbers))

    result = []
    for i in range(len(nums)):
        if i == 0:
            result.append(suffix_numbers[i + 1])
        elif i == len(nums) - 1:
            result.append(prefix_numbers[i - 1])
        else:
            result.append(
                prefix_numbers[i - 1] * suffix_numbers[i + 1]
            )
    print(result)
    return result

products(nums)