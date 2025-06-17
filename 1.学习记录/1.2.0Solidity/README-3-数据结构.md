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


