# Solidity 全局变量 - 区块/交易上下文信息获取方法

Solidity 提供了一系列全局变量，用于获取当前区块和交易的相关信息。这些变量对于编写智能合约非常重要，特别是在需要与区块链状态交互时。

## 区块相关全局变量

1. **blockhash(uint blockNumber) returns (bytes32)**
    - 获取给定区块号的哈希值
    - 只能获取最近256个区块的哈希值

2. **block.basefee (uint)**
    - 当前区块的基础费用 (EIP-1559引入)

3. **block.chainid (uint)**
    - 当前区块链的链ID

4. **block.coinbase (address payable)**
    - 当前区块的矿工/验证者地址

5. **block.difficulty (uint)**
    - 当前区块的难度值 (PoW链)
    - 在PoS链(如以太坊2.0)上可能有不同含义

6. **block.gaslimit (uint)**
    - 当前区块的gas限制

7. **block.number (uint)**
    - 当前区块号

8. **block.timestamp (uint)**
    - 当前区块的时间戳(Unix时间)

## 交易相关全局变量

1. **gasleft() returns (uint256)**
    - 返回剩余的gas量

2. **msg.data (bytes calldata)**
    - 完整的调用数据

3. **msg.sender (address)**
    - 消息的发送者(当前调用者)
    - 最重要的全局变量之一

4. **msg.sig (bytes4)**
    - 调用数据的前4字节(函数选择器)

5. **msg.value (uint)**
    - 随消息发送的wei的数量

6. **tx.gasprice (uint)**
    - 交易的gas价格

7. **tx.origin (address)**
    - 交易的原始发送者(完整的调用链的起点)
    - 安全性考虑: 不应用于权限检查

## 示例用法

```solidity
pragma solidity ^0.8.0;

contract BlockInfo {
    function getBlockInfo() public view returns (
        uint number,
        uint timestamp,
        uint gaslimit,
        address miner
    ) {
        return (
            block.number,
            block.timestamp,
            block.gaslimit,
            block.coinbase
        );
    }
    
    function getTransactionInfo() public view returns (
        address sender,
        uint value,
        uint gasprice
    ) {
        return (
            msg.sender,
            msg.value,
            tx.gasprice
        );
    }
    
    function checkRemainingGas() public view returns (uint) {
        return gasleft();
    }
}
```

## 注意事项

1. `block.timestamp` 可以被矿工稍微操纵，不应作为关键随机源
2. `tx.origin` 不应用于权限检查，应使用 `msg.sender`
3. 区块变量在不同链上可能有不同行为
4. 从0.8.0版本开始，Solidity对区块变量进行了重新组织

这些全局变量为智能合约提供了与区块链环境交互的基本能力，是编写去中心化应用的基础工具。