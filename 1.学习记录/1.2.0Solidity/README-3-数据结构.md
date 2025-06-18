| 数据类型         	 | 描述                         |
| ------------------ | ---------------------------- |
| 整数（`int`/`uint`） |	 有符号和无符号整数，大小从 8 位到 256 位。  |                             	
|布尔（`bool`）       	 |只有两个值：`true` 和 `false`。             |                   	
|地址（`address`）    	 |用于存储以太坊地址。                         |                  	
|字符串（`string`）   	 |可变长度的 UTF-8 编码字符串。                |                  	
|字节数组（`bytes`）  	 |动态长度的字节数组。                         |                  	
|固定字节数组（`bytesN`）| 	 长度固定的字节数组，最大长度为 32。         |                  	
|数组（`array`）      	 |静态数组和动态数组。                         |                  	
|结构体（`struct`）   	 |自定义复杂数据类型。                         |                  	
|映射（`mapping`）    	 |键值对映射，键可以是任意类型，值可以是任意类型。 |             	
|枚举（`enum`）       	 |定义一组命名的常量。                         |                  	
|函数指针（`function`）|	 存储函数的引用。                         |                     	
|用户定义类型（`type`）	 |自定义类型。                               |

---

1.整数（Integers）


无符号整数（`uint`）

```solidity
// 定义一个无符号整数
uint256 myUint = 100;

// 使用规则：无符号整数不能为负数
function incrementUint(uint256 _value) public pure returns (uint256) {
    return _value + 1; // 简单的加法操作
}
```



有符号整数（`int`）

```solidity
// 定义一个有符号整数
int256 myInt = -50;

// 使用规则：有符号整数可以为正数或负数
function negateInt(int256 _value) public pure returns (int256) {
    return -_value; // 取反操作
}
```



2.布尔类型（Boolean）

```solidity
// 定义一个布尔变量
bool myBool = true;

// 使用规则：布尔值只有 `true` 和 `false`
function invertBool(bool _value) public pure returns (bool) {
    return !_value; // 取反操作
}
```



3.地址类型（Address）

```solidity
// 定义一个地址变量
address myAddress = 0x1234567890abcdef1234567890abcdef12345678;

// 使用规则：地址通常用于存储以太坊钱包地址
function getAddressBalance(address _address) public view returns (uint) {
    return _address.balance; // 获取地址的余额
}
```



4.字符串（String）

```solidity
// 定义一个字符串变量
string myString = "Hello, Solidity!";

// 使用规则：字符串用于存储文本数据
function concatenateStrings(string memory _str1, string memory _str2) public pure returns (string memory) {
    return string(abi.encodePacked(_str1, " ", _str2)); // 拼接字符串
}
```



5.字节数组（Bytes）


固定长度字节数组（`bytesN`）

```solidity
// 定义一个固定长度字节数组
bytes32 myBytes32 = 0x1234567890abcdef1234567890abcdef1234567890abcdef;

// 使用规则：固定长度字节数组的长度在定义时确定
function getBytes32Length(bytes32 _bytes) public pure returns (uint) {
    return 32; // 固定长度为 32 字节
}
```



动态长度字节数组（`bytes`）

```solidity
// 定义一个动态长度字节数组
bytes myBytes = new bytes(10);

// 使用规则：动态长度字节数组的长度可以变化
function appendToBytes(bytes memory _bytes, uint8 _value) public pure returns (bytes memory) {
    bytes memory result = new bytes(_bytes.length + 1);
    for (uint i = 0; i < _bytes.length; i++) {
        result[i] = _bytes[i];
    }
    result[_bytes.length] = _value;
    return result; // 在字节数组末尾追加一个字节
}
```



6.数组（Arrays）


静态数组

```solidity
// 定义一个静态数组
uint[3] myStaticArray = [1, 2, 3];

// 使用规则：静态数组的长度固定
function getStaticArrayElement(uint[3] memory _array, uint _index) public pure returns (uint) {
    require(_index < 3, "Index out of bounds");
    return _array[_index]; // 获取数组元素
}
```



动态数组

```solidity
// 定义一个动态数组
uint[] myDynamicArray;

// 使用规则：动态数组的长度可以变化
function pushToDynamicArray(uint _value) public {
    myDynamicArray.push(_value); // 向动态数组中添加元素
}

function getDynamicArrayLength() public view returns (uint) {
    return myDynamicArray.length; // 获取动态数组的长度
}
```



7.结构体（Structs）

```solidity
// 定义一个结构体
struct Person {
    string name;
    uint age;
    address wallet;
}

// 使用规则：结构体用于组合多个字段
Person myPerson = Person("Alice", 30, 0x1234567890abcdef1234567890abcdef12345678);

function getPersonName(Person memory _person) public pure returns (string memory) {
    return _person.name; // 获取结构体中的字段
}
```



8.映射（Mappings）

```solidity
// 定义一个映射
mapping(address => uint) public balances;

// 使用规则：映射用于存储键值对
function setBalance(address _address, uint _value) public {
    balances[_address] = _value; // 设置映射中的值
}

function getBalance(address _address) public view returns (uint) {
    return balances[_address]; // 获取映射中的值
}
```



9.枚举（Enums）

```solidity
// 定义一个枚举
enum State { Pending, Active, Closed }

// 使用规则：枚举用于定义一组命名的常量
State myState = State.Pending;

function nextState(State _state) public pure returns (State) {
    if (_state == State.Pending) {
        return State.Active;
    } else if (_state == State.Active) {
        return State.Closed;
    } else {
        return _state; // 如果已经是 Closed，则保持不变
    }
}
```



10.函数指针（Function Pointers）

```solidity
// 定义一个函数指针
function(uint) external returns (uint) myFunctionPointer;

// 使用规则：函数指针用于存储函数的引用
function setFunctionPointer(function(uint) external returns (uint) _func) public {
    myFunctionPointer = _func; // 设置函数指针
}

function callFunctionPointer(uint _value) public returns (uint) {
    return myFunctionPointer(_value); // 调用函数指针指向的函数
}
```



11.用户定义类型（User-Defined Types）

```solidity
// 定义一个用户定义类型
type MyInt is int;

// 使用规则：用户定义类型是对现有类型的封装
MyInt myValue = MyInt.wrap(10);

function incrementMyInt(MyInt _value) public pure returns (MyInt) {
    return MyInt.wrap(MyInt.unwrap(_value) + 1); // 对用户定义类型进行操作
}
```



12.特殊类型：固定小数点数（Fixed Point Numbers）

```solidity
// 定义一个固定小数点数（需要 Solidity 0.8.0 及以上版本）
int128x128 myFixedPoint = int128x128.wrap(123456);

// 使用规则：固定小数点数用于精确的数学运算
function multiplyFixedPoint(int128x128 _value, int128x128 _multiplier) public pure returns (int128x128) {
    return int128x128.wrap(int128x128.unwrap(_value) * int128x128.unwrap(_multiplier)); // 乘法操作
}
```



13.特殊类型：元组（Tuples）

```solidity
// 定义一个元组
(uint, string) myTuple = (100, "Solidity");

// 使用规则：元组用于组合多个不同类型的值
function getTupleValues((uint, string) memory _tuple) public pure returns (uint, string memory) {
    return (_tuple[0], _tuple[1]); // 获取元组中的值
}
```
---


# Solidity 基础数据类型转换示例

Solidity 是一种静态类型语言，不同类型之间的转换需要显式进行。以下是 Solidity 中常见基础数据类型之间的转换示例：

## 1. 整数类型之间的转换

```solidity
// 整数显式转换
uint8 a = 10;
uint16 b = uint16(a); // 小类型转大类型

uint256 c = 256;
uint8 d = uint8(c); // 大类型转小类型（会截断，256会变成0）

int256 e = -10;
uint256 f = uint256(e); // 有符号转无符号（负数会变成很大的正数）
```

## 2. 地址类型转换

```solidity
// 地址与整数转换
address addr = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;
uint160 addrInt = uint160(addr); // 地址转整数

address newAddr = address(addrInt); // 整数转地址

// payable地址转换
address payable payableAddr = payable(addr);
```

## 3. bytes 类型转换

```solidity
// bytes 类型转换
bytes2 b2 = 0x1234;
bytes4 b4 = bytes4(b2); // 短bytes转长bytes（前面补0）

bytes4 b4full = 0x12345678;
bytes2 b2short = bytes2(b4full); // 长bytes转短bytes（截断后面）

// bytes 与 uint 转换
uint256 num = 0x123456;
bytes32 b32 = bytes32(num); // uint转bytes32

bytes32 fullBytes = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
uint256 numFromBytes = uint256(fullBytes); // bytes32转uint256
```

## 4. 字符串转换

```solidity
// 字符串与bytes转换
string memory str = "hello";
bytes memory byteStr = bytes(str); // string转bytes

bytes memory byteData = new bytes(5);
byteData[0] = 'h';
byteData[1] = 'e';
byteData[2] = 'l';
byteData[3] = 'l';
byteData[4] = 'o';
string memory strFromBytes = string(byteData); // bytes转string
```

## 5. 固定大小字节数组与动态字节数组转换

```solidity
// bytes32 与 bytes 转换
bytes32 fixedBytes = "fixed";
bytes memory dynamicBytes = bytes(fixedBytes); // bytes32转动态bytes

bytes memory dynBytes = new bytes(32);
dynBytes[0] = 'd';
dynBytes[1] = 'y';
dynBytes[2] = 'n';
bytes32 fixedFromDyn = bytes32(dynBytes); // 动态bytes转bytes32（必须长度匹配）
```

## 6. 布尔类型转换

```solidity
// 布尔类型转换
uint256 zero = 0;
bool isZero = zero != 0; // 整数转布尔（0为false，非0为true）

bool flag = true;
uint256 flagNum = flag ? 1 : 0; // 布尔转整数
```

## 7. 枚举类型转换

```solidity
// 枚举类型转换
enum Color { Red, Green, Blue }
Color myColor = Color.Green;
uint colorIndex = uint(myColor); // 枚举转uint（Green对应1）

uint index = 2;
Color colorFromIndex = Color(index); // uint转枚举
```

## 注意事项

1. **显式转换**：Solidity 要求大多数转换必须显式进行
2. **截断风险**：大类型转小类型时可能丢失数据
3. **符号处理**：有符号和无符号整数转换时要注意符号处理
4. **长度匹配**：bytes 类型转换时要注意长度匹配
5. **Gas 消耗**：复杂类型转换可能消耗较多 gas

