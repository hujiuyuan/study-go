# Solidity 地址类型与接口详解

## 地址类型(Address)

Solidity 中的地址类型(address)是专门用于存储以太坊地址的数据类型，长度为20字节(160位)。

### 地址类型分类

1. **普通地址(address)**
    - 基本地址类型
    - 示例：`address user = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;`

2. **可支付地址(address payable)**
    - 可以接收以太币的地址
    - 具有`transfer()`和`send()`方法
    - 示例：`address payable recipient = payable(0x...);`

### 地址类型成员

地址类型有以下内置成员和方法：

| 成员/方法 | 描述 | 示例 |
|-----------|------|------|
| `.balance` | 地址余额(wei) | `address(this).balance` |
| `.transfer(uint amount)` | 发送以太币(失败时revert) | `recipient.transfer(1 ether)` |
| `.send(uint amount)` | 发送以太币(返回bool) | `bool success = recipient.send(1 ether)` |
| `.call(bytes memory)` | 低级调用 | `(bool success, ) = addr.call{value: 1 ether}("")` |
| `.delegatecall(bytes memory)` | 委托调用 | `addr.delegatecall(abi.encodeWithSignature("func()"))` |
| `.staticcall(bytes memory)` | 静态调用 | `addr.staticcall(abi.encodeWithSignature("get()"))` |

### 地址类型转换

1. **普通地址 → 可支付地址**
   ```solidity
   address payable recipient = payable(someAddress);
   ```

2. **合约地址 → 合约实例**
   ```solidity
   MyContract contract = MyContract(contractAddress);
   ```

## 接口(Interface)

接口是定义合约外部可见函数原型的抽象类型，不包含实现。

### 接口特性

1. 只能声明函数(不能有状态变量)
2. 函数不能有实现(不能有函数体)
3. 所有函数默认为`external`和`virtual`
4. 不能继承其他合约，但可以继承其他接口

### 接口定义

```solidity
interface IERC20 {
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address recipient, uint256 amount) external returns (bool);
    
    event Transfer(address indexed from, address indexed to, uint256 value);
}
```

### 接口使用场景

1. **与其他合约交互**
   ```solidity
   IERC20 token = IERC20(tokenAddress);
   token.transfer(msg.sender, amount);
   ```

2. **定义标准规范** (如ERC20、ERC721等)

3. **合约升级模式** (通过接口指向不同实现)

### 接口与合约的区别

| 特性 | 接口(Interface) | 抽象合约(Abstract Contract) |
|------|----------------|----------------------------|
| 函数实现 | 不能有 | 可以有部分实现 |
| 状态变量 | 不能有 | 可以有 |
| 继承 | 只能继承接口 | 可以继承合约 |
| 构造函数 | 不能有 | 可以有 |
| 修饰符 | 不能使用 | 可以使用 |

## 地址与接口结合使用

```solidity
interface IWallet {
    function deposit() external payable;
    function withdraw(uint amount) external;
}

contract WalletUser {
    function sendToWallet(address walletAddr) public payable {
        IWallet wallet = IWallet(walletAddr);
        wallet.deposit{value: msg.value}();
    }
}
```

## 最佳实践

1. 优先使用`transfer()`而非`send()`(因为会自动处理失败)
2. 使用`call()`时总是检查返回值
3. 与外部合约交互时使用接口而非直接调用
4. 对用户输入地址进行零地址检查
5. 使用OpenZeppelin的`Address`库进行安全操作

## 安全注意事项

1. 调用外部合约前验证合约是否存在(检查代码长度)
   ```solidity
   require(addr.code.length > 0, "No contract at address");
   ```
2. 处理`send()`的返回值
3. 使用`call()`时要防范重入攻击
4. 避免使用`tx.origin`做权限检查

## 问题

1. 怎么实现一个接口
2. 怎么实现抽象合约
3. event 是什么，有什么用，怎么用