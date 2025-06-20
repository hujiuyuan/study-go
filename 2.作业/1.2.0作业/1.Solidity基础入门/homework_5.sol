/*
    ✅  合并两个有序数组 (Merge Sorted Array)

    题目描述：将两个有序数组合并为一个有序数组。
*/
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MergeSortedArray {
    function mergeArray(uint[] memory arr1, uint[] memory arr2) public pure returns(uint[] memory) {
        uint[] memory merged = new uint[](arr1.length + arr2.length);
        uint i = 0; // a的指针
        uint j = 0; // b的指针
        uint k = 0; // merged的指针

        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] < arr2[j]) {
                merged[k++] = arr1[i++];
            } else {
                merged[k++] = arr2[j++];
            }
        }
        // 处理剩余元素
        while (i < arr1.length) {
            merged[k++] = arr1[i++];
        }

        while (j < arr2.length) {
            merged[k++] = arr2[j++];
        }
        return merged;
    }
}