// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
    ✅ 创建一个名为Voting的合约，包含以下功能：
    一个mapping来存储候选人的得票数
    一个vote函数，允许用户投票给某个候选人一个getVotes函数，返回某个候选人的得票数一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    // 存储候选人得票数的mapping
    mapping(address => uint256) public votes;

    // 投票函数
    function vote(address candidate) public {
        require(candidate != address(0), "Invalid candidate address");
        votes[candidate] += 1;
    }

    // 获取某个候选人的得票数
    function getVotes(address candidate) public view returns (uint256) {
        return votes[candidate];
    }

    // 重置所有候选人的得票数
    function resetVotes() public {
        // 遍历mapping并重置得票数为0
        // 注意：Solidity中没有直接遍历mapping的方法，但可以通过其他方式实现
        // 这里假设我们有一个已知的候选人列表（candidateList）来遍历
        address[] memory candidateList = getCandidateList();
        for (uint256 i = 0; i < candidateList.length; i++) {
            votes[candidateList[i]] = 0;
        }
    }

    // 辅助函数：获取候选人列表（假设已知）
    function getCandidateList() private pure returns (address[] memory) {
        // 这里返回一个示例候选人列表，实际应用中可以根据需要修改
        address[] memory candidates = new address[](3);
        candidates[0] = 0x1234567890123456789012345678901234567890;
        candidates[1] = 0x0987654321098765432109876543210987654321;
        candidates[2] = 0x1111111111111111111111111111111111111111;
        return candidates;
    }
}
