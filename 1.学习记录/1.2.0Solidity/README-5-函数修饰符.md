# Solidity 函数修饰符(Modifiers)相关知识点

函数修饰符(Modifiers)是Solidity中用于修改函数行为的强大工具，可以增强代码的可重用性和安全性。以下是关于函数修饰符的全面知识点：

## 1. 修饰符基础语法

```solidity
modifier 修饰符名称 {
    // 修饰符逻辑
    _; // 表示原函数体执行的位置
}

function 函数名() 修饰符名称 {
    // 函数体
}
```

## 2. 常用内置修饰符

### visibility修饰符(可见性修饰符)
- `public` - 任意账户/合约可调用
- `private` - 仅当前合约内部可调用
- `internal` - 当前合约及继承合约可调用
- `external` - 仅外部账户/合约可调用

### state mutability修饰符(状态可变性)
- `view` - 承诺不修改状态
- `pure` - 承诺不读取也不修改状态
- `payable` - 函数可以接收ETH

## 3. 自定义修饰符

### 基本自定义修饰符
```solidity
modifier onlyOwner {
    require(msg.sender == owner, "Not owner");
    _;
}

function withdraw() public onlyOwner {
    // 只有owner可以调用
}
```

### 带参数的修饰符
```solidity
modifier costs(uint price) {
    require(msg.value >= price, "Insufficient funds");
    _;
}

function buyItem(uint itemId) public payable costs(1 ether) {
    // 需要支付1ETH才能调用
}
```

## 4. 修饰符执行顺序

```solidity
function example() mod1 mod2 mod3 {
    // 执行顺序: mod1 → mod2 → mod3 → 函数体 → mod3剩余部分 → mod2剩余部分 → mod1剩余部分
}
```

## 5. 特殊修饰符技巧

### 多个修饰符组合
```solidity
function sensitiveAction() 
    public 
    onlyOwner 
    whenNotPaused 
    costs(0.1 ether) 
    payable 
{
    // 组合多个修饰条件
}
```

### 修饰符继承
子合约会继承父合约的修饰符

### 修饰符中的return
```solidity
modifier earlyReturn(bool condition) {
    if (condition) {
        return; // 提前返回
    }
    _;
}
```

## 6. 常见修饰符模式

### 权限控制
```solidity
modifier onlyAdmin {
    require(hasRole(ADMIN_ROLE, msg.sender), "Requires admin role");
    _;
}
```

### 状态检查
```solidity
modifier whenNotPaused {
    require(!paused, "Contract is paused");
    _;
}
```

### 重入保护
```solidity
modifier nonReentrant {
    require(!locked, "Reentrant call");
    locked = true;
    _;
    locked = false;
}
```

### 时间限制
```solidity
modifier duringSale {
    require(block.timestamp >= saleStart && block.timestamp <= saleEnd, "Sale not active");
    _;
}
```

## 7. 修饰符最佳实践

1. **保持简短**：修饰符应该简单且专注于单一职责
2. **明确命名**：如`onlyOwner`、`whenNotPaused`等
3. **避免副作用**：修饰符不应执行复杂逻辑
4. **合理排序**：将最可能失败的检查放在前面
5. **考虑gas成本**：复杂的修饰符会增加gas消耗

## 8. 修饰符与require的比较

| 特性 | 修饰符 | require |
|------|--------|---------|
| 可重用性 | 高 | 低 |
| 代码清晰度 | 高 | 中 |
| Gas成本 | 可能更高 | 通常更低 |
| 错误信息 | 可自定义 | 可自定义 |
| 适用场景 | 多函数共享条件 | 单一函数特定条件 |

## 示例合约

```solidity
pragma solidity ^0.8.0;

contract ModifierExample {
    address public owner;
    bool public paused;
    uint public price;
    
    constructor() {
        owner = msg.sender;
    }
    
    modifier onlyOwner {
        require(msg.sender == owner, "Not owner");
        _;
    }
    
    modifier whenNotPaused {
        require(!paused, "Contract paused");
        _;
    }
    
    modifier costs(uint _amount) {
        require(msg.value >= _amount, "Insufficient funds");
        _;
    }
    
    function setPause(bool _paused) public onlyOwner {
        paused = _paused;
    }
    
    function setPrice(uint _price) public onlyOwner whenNotPaused {
        price = _price;
    }
    
    function buy() public payable whenNotPaused costs(price) {
        // 购买逻辑
    }
    
    // 多个修饰符组合
    function sensitiveAction() 
        public 
        onlyOwner 
        whenNotPaused 
        costs(0.1 ether) 
        payable 
    {
        // 敏感操作
    }
}
```

修饰符是Solidity中实现DRY(Don't Repeat Yourself)原则的重要工具，合理使用可以大大提高合约的安全性和可维护性。