/*
    ✅  二分查找 (Binary Search)

    题目描述：在一个有序数组中查找目标值。
*/
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract Homework_6 {
    function binarySearch(uint[] memory arr1, uint num) public pure returns (int) {
        uint left = 0;
        uint right = arr1.length - 1;

        while (left <= right) {
            uint index = left + (right - left) / 2; // 确保 index 是 uint 类型
            if (arr1[index] == num) {
                return int(index); // 找到目标值，返回索引
            } else if (arr1[index] > num) {
                // 目标在左侧
                right = index - 1;
            } else {
                // 目标在右侧
                left = index + 1;
            }
        }
        return -1; // 目标值不在数组中
    }
}
