在 Solidity 中，可以通过 **全局变量** 和 **内置函数** 获取当前区块和交易的上下文信息。这些信息对于智能合约的逻辑控制、安全验证和 Gas 优化非常重要。以下是详细的分类和示例：

---

## **1. 区块信息（Block Context）**
这些变量提供当前区块的元数据：

| 变量/函数 | 类型 | 描述 | 示例 |
|-----------|------|------|------|
| `block.number` | `uint` | 当前区块高度 | `uint currentBlock = block.number;` |
| `block.timestamp` | `uint` | 区块时间戳（Unix 时间，秒级） | `uint now = block.timestamp;` |
| `block.coinbase` | `address payable` | 当前区块矿工地址 | `address miner = block.coinbase;` |
| `block.difficulty` | `uint` | 当前区块的 PoW 难度 | `uint diff = block.difficulty;` |
| `block.gaslimit` | `uint` | 当前区块的 Gas 上限 | `uint limit = block.gaslimit;` |
| `block.basefee` | `uint` | 当前区块的基础费用（EIP-1559） | `uint baseFee = block.basefee;` |
| `block.chainid` | `uint` | 当前链的 ID（如以太坊主网=1） | `uint chainId = block.chainid;` |
| `blockhash(uint blockNumber)` | `bytes32` | 获取指定区块的哈希（仅最近 256 个区块） | `bytes32 hash = blockhash(block.number - 1);` |

---

## **2. 交易信息（Transaction Context）**
这些变量提供当前交易的相关数据：

| 变量/函数 | 类型 | 描述 | 示例 |
|-----------|------|------|------|
| `msg.sender` | `address` | **当前函数调用者**（最常用！） | `address caller = msg.sender;` |
| `msg.value` | `uint` | 随交易发送的 ETH（单位：wei） | `uint sentEth = msg.value;` |
| `msg.data` | `bytes calldata` | 完整的调用数据（calldata） | `bytes memory data = msg.data;` |
| `msg.sig` | `bytes4` | 函数选择器（如 `transfer(address,uint256)` 的 4 字节哈希） | `bytes4 selector = msg.sig;` |
| `tx.origin` | `address` | **交易的原始发起者**（可能是 EOA，慎用！） | `address origin = tx.origin;` |
| `tx.gasprice` | `uint` | 交易的 Gas 价格（单位：wei/gas） | `uint gasPrice = tx.gasprice;` |
| `gasleft()` | `uint256` | 剩余可用 Gas | `uint remainingGas = gasleft();` |

---

## **3. 合约信息（Contract Context）**
这些变量提供当前合约的元数据：

| 变量/函数 | 类型 | 描述 | 示例 |
|-----------|------|------|------|
| `address(this)` | `address` | 当前合约地址 | `address myContract = address(this);` |
| `this.balance` | `uint256` | 当前合约的 ETH 余额（wei） | `uint balance = address(this).balance;` |
| `selfdestruct(address recipient)` | - | 销毁合约并发送余额 | `selfdestruct(payable(owner));` |

---

## **4. 常见使用场景**
### **(1) 权限控制（`msg.sender`）**
```solidity
address owner;

function withdraw() public {
    require(msg.sender == owner, "Only owner can withdraw");
    payable(msg.sender).transfer(address(this).balance);
}
```

### **(2) 时间锁（`block.timestamp`）**
```solidity
uint public unlockTime = 1650000000;

function withdraw() public {
    require(block.timestamp >= unlockTime, "Funds are locked");
    // ...
}
```

### **(3) 防止重放攻击（`block.number`）**
```solidity
mapping(address => uint) public lastWithdrawBlock;

function withdraw() public {
    require(block.number > lastWithdrawBlock[msg.sender] + 5, "Wait 5 blocks");
    lastWithdrawBlock[msg.sender] = block.number;
    // ...
}
```

### **(4) 动态 Gas 调整（`gasleft()`）**
```solidity
function complexCalculation() public {
    uint startGas = gasleft();
    // ... 复杂计算 ...
    uint gasUsed = startGas - gasleft();
    emit GasUsed(gasUsed);
}
```

### **(5) 链识别（`block.chainid`）**
```solidity
function getChainId() public view returns (uint) {
    return block.chainid; // 主网=1, Ropsten=3, 等
}
```

---

## **5. 注意事项**
1. **`block.timestamp` 不完全可靠**
   - 矿工可轻微调整（±15秒），**不要用于精确计时**。

2. **`tx.origin` vs `msg.sender`**
   - `tx.origin` 是交易的**最初发起者**（可能是用户钱包）。
   - `msg.sender` 是**直接调用者**（可能是另一个合约）。
   - **避免用 `tx.origin` 做权限检查**（易受钓鱼攻击）。

3. **`blockhash` 仅支持最近 256 个区块**
   - 超出范围返回 `0`。

4. **`gasleft()` 用于优化复杂逻辑**
   - 可在循环中检查剩余 Gas，避免 Out-of-Gas 错误。

---

## **总结**
Solidity 的区块/交易上下文信息广泛用于：
- **权限管理**（`msg.sender`）
- **时间控制**（`block.timestamp`）
- **链识别**（`block.chainid`）
- **Gas 优化**（`gasleft()`）
- **安全验证**（避免 `tx.origin`）

掌握这些全局变量能帮助你编写更安全、高效的智能合约！ 🚀