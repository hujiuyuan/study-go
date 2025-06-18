/*
    ✅ 反转字符串 (Reverse String)

题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
*/
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Reverse {

    function revertString(string memory str) public pure returns (string memory) {
        bytes memory byteStr = bytes(str);
        bytes memory result = new bytes(byteStr.length);

        for (uint i = 0; i < byteStr.length; i++) {
            result[i] = byteStr[byteStr.length - i - 1];
        }
        return string(result);
    }
}