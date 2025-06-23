在 Solidity 中，**错误处理**和**自定义修饰符（Modifiers）** 是编写健壮、安全智能合约的重要工具。以下是详细解析和最佳实践：

---

## **1. 错误处理（Error Handling）**
Solidity 提供了三种主要的错误处理机制：

### **(1) `require(condition, message)`**
- **用途**：验证输入或状态（如权限检查、参数有效性）。
- **效果**：如果 `condition` 为 `false`，回滚所有状态变更，并返回 `message`。
- **Gas**：消耗剩余 Gas，不退款。
- **示例**：
  ```solidity
  function withdraw(uint amount) public {
      require(amount <= balance[msg.sender], "Insufficient balance");
      balance[msg.sender] -= amount;
  }
  ```

### **(2) `revert(message)`**
- **用途**：主动触发回滚，通常用于复杂条件判断。
- **效果**：直接终止执行并回滚，返回 `message`。
- **Gas**：同 `require`。
- **示例**：
  ```solidity
  function transfer(address to, uint amount) public {
      if (to == address(0) || amount == 0) {
          revert("Invalid recipient or amount");
      }
      // ...
  }
  ```

### **(3) `assert(condition)`**
- **用途**：检查内部错误（如合约逻辑漏洞）。
- **效果**：如果 `condition` 为 `false`，回滚并消耗所有 Gas。
- **Gas**：不退款（设计用于严重错误）。
- **示例**：
  ```solidity
  function divide(uint a, uint b) public pure returns (uint) {
      assert(b != 0); // 除零错误应永不发生
      return a / b;
  }
  ```

### **对比表**
| 机制         | 适用场景               | Gas 处理       | 是否退款 |
|--------------|------------------------|----------------|----------|
| `require`    | 输入/状态验证          | 消耗剩余 Gas   | ❌        |
| `revert`     | 复杂条件回滚           | 消耗剩余 Gas   | ❌        |
| `assert`     | 内部逻辑错误（应永不发生） | 消耗所有 Gas   | ❌        |

---

## **2. 自定义修饰符（Custom Modifiers）**
修饰符（Modifiers）用于**复用代码逻辑**（如权限检查、状态验证），可附加到函数上。

### **(1) 基本语法**
```solidity
modifier onlyOwner() {
    require(msg.sender == owner, "Not owner");
    _; // 继续执行函数体
}
```

### **(2) 使用示例**
```solidity
address owner;

constructor() {
    owner = msg.sender;
}

modifier onlyOwner() {
    require(msg.sender == owner, "Not owner");
    _;
}

function changeOwner(address newOwner) public onlyOwner {
    owner = newOwner;
}
```

### **(3) 带参数的修饰符**
```solidity
modifier minimumAmount(uint amount) {
    require(msg.value >= amount, "Insufficient ETH");
    _;
}

function buyToken() public payable minimumAmount(1 ether) {
    // 仅当 msg.value >= 1 ether 时执行
}
```

### **(4) 修饰符执行顺序**
修饰符按声明顺序执行：
```solidity
function example() public mod1 mod2 {
    // 先执行 mod1，再 mod2，最后函数体
}
```

---

## **3. 高级错误处理（Custom Errors）**
Solidity 0.8.4+ 支持**自定义错误类型**，比 `revert(string)` 更省 Gas。

### **(1) 定义错误**
```solidity
error InsufficientBalance(uint available, uint required);
```

### **(2) 触发错误**
```solidity
function withdraw(uint amount) public {
    if (amount > balance[msg.sender]) {
        revert InsufficientBalance(balance[msg.sender], amount);
    }
    // ...
}
```

### **(3) 优势**
- **Gas 优化**：比字符串消息更便宜。
- **结构化数据**：可返回错误详细信息（如数值）。

---

## **4. 最佳实践**
### **(1) 错误处理选择**
- 使用 `require` 验证**外部输入**（如用户参数）。
- 使用 `assert` 检查**内部一致性**（如数学不变量）。
- 用 `revert` 处理**复杂条件**（如业务逻辑错误）。

### **(2) 修饰符设计**
- **单一职责**：一个修饰符只做一件事（如 `onlyOwner`、`whenNotPaused`）。
- **避免嵌套**：复杂逻辑应放在函数内，而非修饰符。

### **(3) 安全注意事项**
- **慎用 `tx.origin`**：修饰符中避免使用 `tx.origin` 做权限检查（易受钓鱼攻击）。
- **Gas 限制**：修饰符中的操作应尽量轻量。

---

## **5. 完整示例**
```solidity
pragma solidity ^0.8.0;

contract ErrorHandlingExample {
    address public owner;
    bool public paused;
    mapping(address => uint) public balance;

    error ContractPaused();
    error Unauthorized();

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        if (msg.sender != owner) revert Unauthorized();
        _;
    }

    modifier whenNotPaused() {
        if (paused) revert ContractPaused();
        _;
    }

    function pause() public onlyOwner {
        paused = true;
    }

    function deposit() public payable whenNotPaused {
        balance[msg.sender] += msg.value;
    }

    function withdraw(uint amount) public whenNotPaused {
        require(amount <= balance[msg.sender], "Insufficient balance");
        balance[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
    }
}
```

---
## **6. Solidity 0.8.0 及之后版本中的自定义错误**
    在 Solidity 0.8.0 之后，Solidity 引入了自定义错误机制（custom errors），提供了一种更加 Gas 高效的错误处理方式。自定义错误比 require 或 revert 的字符串消息消耗更少的 Gas，因为自定义错误只传递函数选择器和参数。

定义和使用自定义错误

    自定义错误定义：
        自定义错误通过 error 关键字声明。

    示例：
```solidity
pragma solidity ^0.8.0;
contract CustomErrorExample {
    error Unauthorized(address caller);  // 自定义错误
    address public owner;
    constructor() {
        owner = msg.sender;
    }
    function restrictedFunction() public {
        if (msg.sender != owner) {
        revert Unauthorized(msg.sender);  // 使用自定义错误
        }
    }
}
```


自定义错误的优势：

    自定义错误不会在错误消息中传递冗长的字符串，因此相比传统的 require 和 revert，节省了更多的 Gas。
---

## **7. try/catch 错误捕获** 

Solidity 0.6.0 版本后引入了 try/catch 结构，用于捕获外部合约调用中的异常。此功能允许开发者捕获和处理外部调用中的错误，增强了智能合约编写的灵活性。
try/catch 的使用场景：

    捕获外部合约调用失败时的错误，而不让整个交易失败。
    在同一个交易中可以对失败的调用进行处理或重试。

示例代码
```solidity

pragma solidity >=0.6.0;
contract ExternalContract {
    function getValue() public pure returns (uint) {
        return 42;
    }
    function willRevert() public pure {
        revert("This function always fails");
    }
}
contract TryCatchExample {
    ExternalContract externalContract;
    constructor() {
        externalContract = new ExternalContract();
    }
    function tryCatchTest() public returns (uint, string memory) {
        try externalContract.getValue() returns (uint value) {
            return (value, "Success");
        } catch {
            return (0, "Failed");
        }
    }
    function tryCatchWithRevert() public returns (string memory) {
        try externalContract.willRevert() {
            return "This will not execute";
        } catch Error(string memory reason) {
            return reason;  // 捕获错误信息
        } catch {
            return "Unknown error";
        }
    }
}
```


## **总结**
- **错误处理**：`require`、`revert`、`assert` 分别用于不同场景。
- **修饰符**：复用权限和状态检查逻辑。
- **自定义错误**：节省 Gas 并提供结构化错误信息。
- **安全第一**：合理设计修饰符，避免常见漏洞（如重入攻击）。