# 实现细节 

[English version](./implementation_en.md)

## 流程

```
          lexer           parser                              parser and linker
xxx.fbs --------> tokens --------> abstract syntax tree (AST) ================> fbs descriptor
```

lexer 是直接实现的

parser 则使用 [goyacc](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc) 来生成

linker 也是直接实现的 

## lexer 的实现

见 `fbsLex.Lex`，这个函数每次被调用时会读取文件内容，返回一个 token 对应的编号

## parser 的实现

### 第零步：阅读官方网站的语法并实现 `fbs.y`

* [flatbuffers 语法](https://google.github.io/flatbuffers/flatbuffers_grammar.html)

* 实现 `fbs.y`

__Notes:__ 

1. 官网提供的 [flatbuffers 语法](https://google.github.io/flatbuffers/flatbuffers_grammar.html) 并不准确，需要修改，以现在实现的 [`fbs.y
`](../fbs.y) 为准
2. `fbs.y` 这个文件定义了
	1. terminals (AST 的叶节点) 以及 non-terminals (AST 的内部节点)
	2. 生成规则：terminals 和 non-terminals 之间如何组合最后得到其他的 non-terminals，直到最后组合成根节点

### 第一步：使用 `fbs.y` 来生成 `fbs.y.go`

首先安装 `goyacc` 工具

```shell
$ go get golang.org/x/tools/cmd/goyacc
```

然后使用以下命令来生成 `fbs.y.go`

```shell
$ go generate ./...
```

该命令写在 `parser.go` 的文件起始处：

```go
//go:generate goyacc -o fbs.y.go -p fbs fbs.y
```

可以直接使用如下的原始命令来生成：

```shell
$ goyacc -o fbs.y.go -p fbs fbs.y
```

除了 `fbs.y` 中已经提供了的信息之外, `fbs.y.go` 还定义了： 

1. `fbsParser` interface 以及其具体实现 `fbsParserImpl`
2. `fbsLexer` interface:
```go
type fbsLexer interface {
	Lex(lval *fbsSymType) int
	Error(s string)
}
```
我们需要为 `fbsLexer` 提供一个具体实现然后将这个 lexer 传给 `fbsParser.Parse`， `fbsParser.Parse` 会读取我们 lexer 生成的 token (terminal) 流，然后根据 `fbs.y` 中定义的生成规则来构造一个完整的 AST

### 第二步：用 Go 写 terminal 以及 non-terminal 的节点定义以及相应的构造方法

Package `ast` 提供了 terminal 以及 non-terminal 的节点定义以及相应的构造方法，这些内容会被 `fbs.y` 及其生成文件用到

举例如下：

1. 节点定义
	1. Terminal 节点定义:
	```go
	type IdentNode struct { terminalNode .. }
	```
	2. Non-terminal 节点定义:
	```go
	type SchemaNode struct { compositeNode .. }
	```
2. 构造方法
	1. Terminal 节点构造方法, 这些方法会被 lexer 用来生成 terminal 节点:
	```go
	func NewIdentNode(val string, info TokenInfo) *IdentNode {..}
	```
	2. Non-terminal 节点构造方法, 这些方法会被 `fbs.y` 使用:
	```go
	func NewSchemaNode(includes []*IncludeNode, decls []DeclElement) *SchemaNode {..}
	```
	`fbs.y` 中的使用例子:
	```go
	// schema is the root node of AST, it is of type *ast.SchemaNode
	// $$ represents left hand side of colon, in this case is `schema`
	// $1 represents the first node on the right hand side of colon, `includes`
	// $2 refers to the second node `decls`
	// $3 $4 ..
	schema: includes decls {
			$$ = ast.NewSchemaNode($1, $2)
			fbslex.(*fbsLex).res = $$  // store result into lexer
		}
	```

### 第三步：找到不合理的语法，修改语法并重复以上步骤

`.fbs` 文件的具体使用例子可以参考： [monster_test.fbs](https://github.com/google/flatbuffers/blob/master/tests/monster_test.fbs)

我们提供了一个带有语法注释的 [test file](../fbsfiles/annotated_test.fbs), 该文件基本包含了所有常见用法

现在写好的 `fbs.y` 文件和 [原始语法](https://google.github.io/flatbuffers/flatbuffers_guide_writing_schema.html) 已经差别较大，一些改动可以参考 `fbs.y` 文件的开头注释部分，这些变动正是本步骤的结果

### 第四步：基于 AST 来生成描述符

实现部分见 `Parser.ParseFiles`，该方法会为每一个输入的文件产生对应描述符的输出，如果设置了 `Recursive` 标志（默认设置），那么会递归解析所有 include 的 flatbuffers 文件

对于每一个文件的操作步骤如下：

1. 创建一个 lexer
2. 调用生成文件 `fbs.y.go` 中定义的 `fbsParse`，该函数会每次调用 lexer 的 `Lex` 函数来产生 token，然后根据 token 序列以及语法信息来构造抽象语法树 (AST)
3. 调用 `parseResult.createSchemaDescriptor` 来为 AST 生成对应的描述符，方法是遍历树上的结点，依次根据结点信息来构造相应的描述符

## linker 的实现

实现部分见 `linker.linkFiles`，其主要作用是对一个文件中的类型引用进行解析，因为这些类型可能定义在该文件所 include 的文件中

分为两步：

1. 记录所有已经定义好的类型（type definition），将这些类型定义的完全限定名（fully qualified name, 即前面带有自己所在 namespace
 的完整类型名）以及其对应的描述符放到 descriptor pool 中，在这个过程中可以检查出是否出现重复定义的错误

2. 对文件中所有的类型引用（type reference）进行解析，将原始类型引用名依次加上 namespace 的各种前缀形成完全限定名后在 descriptor pool
 中查找是否存在该符号的定义，如果存在，则解析成功，保存其成功解析的完全限定名

以上步骤都结束后，我们便可以得到输入文件列表对应的描述符列表 