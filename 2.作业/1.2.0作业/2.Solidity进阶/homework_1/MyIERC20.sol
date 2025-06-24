// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts (last updated v5.1.0) (token/ERC20/IERC20.sol)

pragma solidity ^0.8.0;

import "@openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";


/*

    合约包含以下标准 ERC20 功能：

balanceOf：查询账户余额。transfer：转账。approve 和 transferFrom：授权和代扣转账。使用 event 记录转账和授权操作。提供 mint 函数，允许合约所有者增发代币。提示：

    使用 mapping 存储账户余额和授权信息。

使用 event 定义 Transfer 和 Approval 事件。部署到sepolia 测试网，导入到自己的钱包
*/
contract MyIERC20 is IERC20 {
    // 合约所有者
    address private owner;
    // 账户
    mapping(address => uint256) private balances;
    // 授权额度
    mapping(address => mapping(address => uint256)) private allowances;

    // 修饰符：仅允许所有者调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // 构造函数：初始化所有者和初始供应量
    constructor(uint256 initialSupply) {
        owner = msg.sender; // 设置部署者为所有者
        _mint(msg.sender, initialSupply * 10**uint256(10)); // 初始铸造代币给部署者
    }

    /*
    查询账户余额
    */
    function balanceOf(address user_address) returns (uint256) {
        return balances[user_address];
    }


    // 内部铸造函数
    function _mint(address account, uint256 value) internal {
        _totalSupply += value;
        balances[account] += value;

        emit Transfer(address(0), account, value);
    }
}