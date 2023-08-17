# fbs 

一个用 Go 实现的 flatbuffers 解析器

[English version](./docs/README_en.md)

## 快速上手

使用方法：

1. 首先导入 `package fbs`：

```go
import "trpc.group/trpc-go/fbs"
```

2. 然后传入要解析的 flatbuffers 文件列表，创建 `Parser` 并调用 `ParseFiles` 

```go
filenames := []string{
    "./file1.fbs", // 给出一系列需要解析的 flatbuffers 文件名
    "./file2.fbs",
}
p := fbs.NewParser()
fds, err := p.ParseFiles(filenames...)
fd1 := fds[0] // file1.fbs 的解析结果（描述符）
fd2 := fds[1] // file2.fbs 的解析结果（描述符）
```

通过 `fds` 即可访问到描述符，从而可以使用 flatbuffers 文件中定义的 rpc service 里各个 method 的名字、输入输出类型、是否为流式等信息，描述符定义见 `desc.go`

其中每个文件的描述符 `SchemaDesc` 各字段如下：

```go
// SchemaDesc 描述了一个完整的 flatbuffers 文件
type SchemaDesc struct {
	Schema *ast.SchemaNode // Schema 存储抽象语法树(AST)的根节点
	Name string // Name 存储 .fbs 文件的文件名，如 "./file1.fbs" 
	// Namespaces 存储该文件中遇到的所有 namespace，其作用类似于 protobuf 中的 package 
	// 不同之处在于 flatbuffers 允许一个文件中出现多次 namespace 
	// 不同 namespace 下定义的 table/struct 等类型拥有各自不同的 namespace 
	Namespaces []string 
	Root string // Root 存储 flatbuffers 文件中的 root_type 声明
	FileExt string // FileExt 存储 flatbuffers 文件中的 file_extension 声明
	FileIdent string // FileIdent 存储 flatbuffers 文件中的 file_identifier 声明
	Attrs []string // Attrs 存储 flatbuffers 文件中的 attribute 声明（可能有多个）
	Includes []string // Includes 存储所有 include 语句中的文件名
	Dependencies []*SchemaDesc // Dependencies 存储所有 include 语句对应的解析出的描述符 
	Tables []*TableDesc // Tables 存储所有的 table 描述符
	Structs []*StructDesc // Structs 存储所有的 struct 描述符
	Enums []*EnumDesc // Enums 存储所有的 enumeration 描述符
	Unions []*UnionDesc // Unions 存储所有的 union 描述符 
	RPCs []*RPCDesc // RPCs 存储所有 rpc service 描述符
}
```

通过访问这些字段，如 `RPCs`，就可以得到 flatbuffers 文件中定义的信息，从而完成一系列相关的工作（如 rpc 桩代码的生成）

## 工程目录结构

```
.
├── desc.go         # 各种节点对应描述符的定义
├── desc_test.go    
├── doc.go          
├── docs            # 实现文档
├── errors.go       # 错误处理机制 
├── errors_test.go  
├── fbsfiles        # 存放用于测试的 .fbs 文件
├── fbs.y           # 根据 flatbuffers 的语法写成
├── fbs.y.go        # 由 fbs.y 生成, 用于将 token 流解析为抽象语法树
├── go.mod          
├── go.sum          
├── internal        
│   └── ast         # 存放抽象语法树中各节点的定义以及构造方法
├── lexer.go        # 实现 lexer
├── lexer_test.go   
├── linker.go       # 实现 linker 
├── linker_test.go  
├── parse_result.go # 存放解析结果的定义 
├── parser.go       # 实现 parser 
├── parser_test.go  
├── reader.go       # 实现 reader 
├── README.md       
└── scope.go        # 实现 scope 
```

更多内容见 [实现细节](./docs/implementation_cn.md)

## 应用实例

本项目为 trpc-go 提供了支持生成 flatbuffers 相关桩代码的功能