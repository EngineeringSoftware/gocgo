## gocgo

gocgo (go-c-go), implemented in Go, includes lexer, parser, and
visitors for the C programming language.

This small project started as a demo of ANTLR and the visitor design
pattern.  It turned out that having a C parser available was valuable
for a few other cases on our side, thus, we decided to make it public
as it might be of interest to others as well.

gocgo is (currently) suitable for anlaysis of C code.  While some
transformation is possible, having a flexible transformation API could
be added later.


## A Quick Example

Below is an example of a visitor that collects C code into a buffer
and along the way removes function bodies.

```go
type FuncDeleteVisitor struct {
	*BaseVisitorImpl
	in   bool
	Buff bytes.Buffer
}

func NewFuncDeleteVisitor() *FuncDeleteVisitor {
	return &FuncDeleteVisitor{in: false}
}

func (v *FuncDeleteVisitor) VisitTerminal(n antlr.TerminalNode) error {
	v.Buff.WriteString(n.GetSymbol().GetText())
	v.Buff.WriteString(" ")
	return nil
}

func (v *FuncDeleteVisitor) VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) (bool, error) {
	v.in = true
	return true, nil
}

func (v *FuncDeleteVisitor) VisitFunctionDefinitionEnd(ctx *parsing.FunctionDefinitionContext) error {
	v.in = false
	return nil
}

func (v *FuncDeleteVisitor) VisitBlockItemList(ctx *parsing.BlockItemListContext) (bool, error) {
	if v.in {
		return false, nil
	} else {
		return true, nil
	}
}
```


## Usage

We provide a complete example of running the aforementioned visitor as
a command in this repo.  You can run the following:

```
go run ./cmd/gocgo/gocgo.go cprogs/for.c 
```

You can take a look at the command for details, but it is rather
simple code.


## Architecture and Design

Each visitor should embed a pointer to `BaseVisitorImpl`, which
provides the default implementation of `Visit` methods.

There is a `VisitX` method for each type of AST node (e.g.,
`VisitFunctionDefinition` for a function definition).  Consider this
example:

```go
func (v *FuncDeleteVisitor) VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) (bool, error) {
	v.in = true
	return true, nil
}
```

This `Visit` method is invoked for every function definition in parsed
C code.  There are two return values.  The first value (`bool`) says
if children should be visited.  The second value (`error`) is the
`error` Go interface.

Each `VisitX` method has a corresponding `VisitXEnd` method invoked
once all children of the node are visited.  From our earlier example:

```go
func (v *FuncDeleteVisitor) VisitFunctionDefinitionEnd(ctx *parsing.FunctionDefinitionContext) error {
	v.in = false
	return nil
}
```

`VisitXEnd` methods have a single return value (`error`).

This architecture is common in many tools, e.g., visitors in the
Eclipse implementation.


## Implementation

gocgo cuts a few steps for you if you were planning to use ANTLR to
obtain a parser for the C programming language.

We used the grammar from the [ANTLR
repo](https://github.com/antlr/grammars-v4/blob/19de2a44eaef0de599044b76ede8492fb8e07b4d/c/C.g4).

If you wish to reproduce the steps to get lexer, parser, and the
original visitors, you can run the following command in the root of
this repo:

```bash
go generate ./...
```

We made a few changes to simplify the usage of visitors and address
some issues (e.g., [Issue
4398](https://github.com/antlr/antlr4/issues/4398)).  Below is the
summary of key changes / additions:

* Addressed the issue 4398, by introducing (locally in this repo)
  `VisitChildren`.

* Removed the generated `c_base_visitor.go`, because it limits
  changing behavior of any `Visit` method.

* Introduced our own visitor interface (`BaseVisitor`) and its
  implementation (`BaseVisitorImpl`) that enable proper struct
  embedding.  Your visitor can now change behavior of (a subset of)
  `Visit` methods.

* Updated API design, such that each `Visit` method returns `bool` and
  `error` (rather than `interface{}`). We find the new approach more
  Go appropriate.


## License

[BSD-3-Clause license](LICENSE).


## Contact

Feel free to get in touch if you have any comments: Milos Gligoric
`<milos.gligoric@gmail.com>`.
