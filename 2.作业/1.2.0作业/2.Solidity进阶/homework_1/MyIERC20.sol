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
    address private _owner;
    // 账户
    mapping(address => uint256) private _balances;
    // 授权额度
    mapping(address => mapping(address => uint256)) private _allowances;
    uint256 private _totalSupply; // 总供应量

    // 定义转账事件
    event Transfer(address owner, address account, uint256 balance);
    // 定义授权事件
    event Approval(address owner, address account, uint256 balance);

    // 修饰符：仅允许所有者调用
    modifier onlyOwner() {
        require(msg.sender == _owner, "Only owner can call this function");
        _;
    }

    // 构造函数：初始化所有者和初始供应量
    constructor(uint256 initialSupply) {
        _owner = msg.sender; // 设置部署者为所有者
        _mint(msg.sender, initialSupply * 10**uint256(10)); // 初始铸造代币给部署者
    }

    /*
    address(0) 代表一个 空账户，
    emit Transfer(address(0), account, value); // 空账号转 给 目标账户转代币 代表新增发
    emit Transfer(account, address(0), value); // 目标账户 给 空账号转代币 代表销毁代币
    */
    // 铸造代币（仅所有者可调用）
    function mint(address to, uint256 value) external onlyOwner {
        require(to != address(0), "Invalid address");

        _mint(to, value);
    }

    // 内部铸造函数
    function _mint(address account, uint256 value) internal {
        _totalSupply += value;
        _balances[account] += value;

        emit Transfer(address(0), account, value);
    }

    /*
    查询账户余额
    */
    function balanceOf(address account) external view  returns (uint256) {
        return _balances[account];
    }

    /*
    转账
    */
    function transfer(address to_account, uint256 value) external view  returns (uint256) {
        //  校验 账户&余额
        require(address(0) == msg.sender, "当前账号发起转账失败！");
        require(_balances[msg.sender] < value, "当前账户的余额不支持本次转账！");

        _balances[msg.sender] -= value;
        _balances[to_account] += value;
        // 发送转账消息
        emit Transfer(msg.sender, to_account, value);
        return _balances[account];
    }

    /*
    授权
    */
    function approve(address from_account, uint256 value)  {
        require(address(0) == msg.sender, "当前账号发起转账失败！");
    }
}