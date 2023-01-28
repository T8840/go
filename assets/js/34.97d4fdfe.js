(window.webpackJsonp=window.webpackJsonp||[]).push([[34],{444:function(t,a,n){"use strict";n.r(a);var s=n(2),e=Object(s.a)({},(function(){var t=this,a=t._self._c;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("p",[t._v("Go在语法层面将并发原语channel作为一等公民对待。这意味着我们可以像使用普通变量那样使用channel，比如，定义channel类型变量、给channel变量赋值、将channel作为参数传递给函数/方法、将channel作为返回值从函数/方法中返回，甚至将channel发送到其他channel中。")]),t._v(" "),a("h2",{attrs:{id:"channel的操作"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#channel的操作"}},[t._v("#")]),t._v(" channel的操作")]),t._v(" "),a("h3",{attrs:{id:"创建channel"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#创建channel"}},[t._v("#")]),t._v(" 创建channel")]),t._v(" "),a("ul",[a("li",[t._v("方式一：声明一个元素为int类型的channel类型变量ch")])]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" ch "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),t._v("\n")])])]),a("p",[t._v("如果channel类型变量在声明时没有被赋予初值，那么它的默认值为nil。")]),t._v(" "),a("ul",[a("li",[t._v("方式二：使用make这个Go预定义的函数为channel类型变量赋初值")])]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[t._v("ch1 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 无缓冲")]),t._v("\nch2 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("5")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 带缓冲")]),t._v("\n")])])]),a("p",[t._v("声明了两个元素类型为int的channel类型变量ch1和ch2，并给这两个变量赋了初值。\n第一行我们通过make(chan T)创建的、元素类型为T的channel类型，是无缓冲channel"),a("br"),t._v("\n第二行中通过带有capacity参数的make(chan T, capacity)创建的元素类型为T、缓冲区长度为capacity的channel类型，是带缓冲channel\n这两种类型的变量关于发送（send）与接收（receive）的特性是不同的")]),t._v(" "),a("h3",{attrs:{id:"发送与接收"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#发送与接收"}},[t._v("#")]),t._v(" 发送与接收")]),t._v(" "),a("p",[t._v("Go提供了<-操作符用于对channel类型变量进行发送与接收操作：")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[t._v("ch1 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("13")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 将整型字面值13发送到无缓冲channel类型变量ch1中")]),t._v("\nn "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" ch1 "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 从无缓冲channel类型变量ch1中接收一个整型值存储到整型变量n中")]),t._v("\nch2 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("17")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 将整型字面值17发送到带缓冲channel类型变量ch2中")]),t._v("\nm "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" ch2 "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 从带缓冲channel类型变量ch2中接收一个整型值存储到整型变量m中")]),t._v("\n\n")])])]),a("p",[t._v("在理解channel的发送与接收操作时，你一定要始终牢记：channel是用于Goroutine间通信的，所以绝大多数对channel的读写都被分别放在了不同的Goroutine中。")]),t._v(" "),a("p",[t._v("使用操作符<-，我们还可以声明只发送channel类型（send-only）和只接收channel类型。\n试图从一个只发送channel类型变量中接收数据，或者向一个只接\n收channel类型发送数据，都会导致编译错误。。")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[t._v("ch1 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 只发送channel类型")]),t._v("\nch2 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 只接收channel类型")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v("ch1 "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// invalid operation: <-ch1 (receive from send-only type chan<- int)")]),t._v("\nch2 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("13")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// invalid operation: ch2 <- 13 (send to receive-only type <-chan int)")]),t._v("\n")])])]),a("p",[t._v("通常只发送channel类型和只接收channel类型，会被用作函数的参数类型或返回值，用于限制对channel内的操作，或者是明确可对channel进行的操作的类型，如下：")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("produce")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("ch "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("for")]),t._v(" i "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("0")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(";")]),t._v(" i "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("10")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(";")]),t._v(" i"),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("++")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    ch "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" i "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("+")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),t._v("\n    time"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Sleep")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("time"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Second"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("close")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("ch"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("consume")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("ch "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("for")]),t._v(" n "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("range")]),t._v(" ch "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("println")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("main")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n ch "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("5")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" wg sync"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("WaitGroup\n wg"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Add")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("2")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("go")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("produce")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("ch"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n wg"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Done")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("go")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("consume")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("ch"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n  wg"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Done")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n wg"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Wait")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("在这个例子中，我们启动了两个Goroutine，分别代表生产者（produce）与消费者（consume）。生产者只能向channel中发送数据，我们使用chan<- int作为produce函\n数的参数类型；消费者只能从channel中接收数据，我们使用<-chan int作为consume函\n数的参数类型。\n在消费者函数consume中，我们使用了for range循环语句来从channel中接收数据，for\nrange会阻塞在对channel的接收操作上，直到channel中有数据可接收或channel被关闭循环，才会继续向下执行。channel被关闭后，for range循环也就结束了")]),t._v(" "),a("h4",{attrs:{id:"无缓冲发送与接收"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#无缓冲发送与接收"}},[t._v("#")]),t._v(" 无缓冲发送与接收")]),t._v(" "),a("p",[t._v("对无缓冲channel类型的发送与接收操作，一定要放在两个不同的Goroutine中进行，否则会导致deadlock。")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("main")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n ch1 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("go")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    ch1 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("13")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 将发送操作放入一个新goroutine中执行")]),t._v("\n    "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n n "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v("ch1\n "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("println")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("下面这个报错：")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("main")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n ch1 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n ch1 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("13")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// fatal error: all goroutines are asleep - deadlock!")]),t._v("\n n "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v("ch1\n "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("println")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("h4",{attrs:{id:"带缓冲发送与接收"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#带缓冲发送与接收"}},[t._v("#")]),t._v(" 带缓冲发送与接收")]),t._v(" "),a("p",[t._v("和无缓冲channel相反，带缓冲channel的运行时层实现带有缓冲区，因此，对带缓冲\nchannel的发送操作在缓冲区未满、接收操作在缓冲区非空的情况下是异步的（发送或接收不需要阻塞等待）。"),a("br"),t._v("\n也就是说，对一个带缓冲channel来说，在缓冲区未满的情况下，对它进行发送操作的\nGoroutine并不会阻塞挂起；在缓冲区有数据的情况下，对它进行接收操作的Goroutine也不会阻塞挂起。但当缓冲区满了的情况下，对它进行发送操作的Goroutine就会阻塞挂起；当缓冲区为空的情况下，对它进行接收操作的Goroutine也会阻塞挂起。")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[t._v("ch2 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\nn "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v("ch2 "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起")]),t._v("\nch3 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("make")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\nch3 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("17")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 向ch3发送一个整型数17")]),t._v("\nch3 "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("27")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起")]),t._v("\n")])])]),a("h3",{attrs:{id:"关闭channel"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#关闭channel"}},[t._v("#")]),t._v(" 关闭channel")]),t._v(" "),a("p",[t._v("采用不同接收语法形式的语句，在channel被关闭后的返回值的情况")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[t._v("n "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v(" ch "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 当ch被关闭后，n将被赋值为ch元素类型的零值")]),t._v("\nm"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" ok "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),t._v("ch "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 当ch被关闭后，m将被赋值为ch元素类型的零值, ok值为false")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("for")]),t._v(" v "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("range")]),t._v(" ch "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 当ch被关闭后，for range循环结束")]),t._v("\n "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("...")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("...")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("注意：")]),t._v(" "),a("ul",[a("li",[t._v("一般，发送端负责关闭channel。为什么要在发送端关闭channel呢？这是因为发送端没有像接受端那样的、可以安全判断channel是否被关闭了的方法")]),t._v(" "),a("li",[t._v("一旦向一个已经关闭的channel执行发送操作，这个操作就会引发panic")])]),t._v(" "),a("h2",{attrs:{id:"无缓冲channel的惯用法"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#无缓冲channel的惯用法"}},[t._v("#")]),t._v(" 无缓冲channel的惯用法")]),t._v(" "),a("p",[t._v("无缓冲channel兼具通信和同步特性，在并发程序中应用颇为广泛。以下应用示例：")]),t._v(" "),a("h3",{attrs:{id:"第一种用法-用作信号传递"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#第一种用法-用作信号传递"}},[t._v("#")]),t._v(" 第一种用法：用作信号传递")]),t._v(" "),a("h3",{attrs:{id:"第二种用法-用于替代锁机制"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#第二种用法-用于替代锁机制"}},[t._v("#")]),t._v(" 第二种用法：用于替代锁机制")]),t._v(" "),a("h2",{attrs:{id:"带缓冲channel的惯用法"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#带缓冲channel的惯用法"}},[t._v("#")]),t._v(" 带缓冲channel的惯用法")]),t._v(" "),a("h3",{attrs:{id:"第一种用法-用作消息队列"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#第一种用法-用作消息队列"}},[t._v("#")]),t._v(" 第一种用法：用作消息队列")]),t._v(" "),a("h3",{attrs:{id:"第二种用法-用作计数信号量-counting-semaphore"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#第二种用法-用作计数信号量-counting-semaphore"}},[t._v("#")]),t._v(" 第二种用法：用作计数信号量（counting semaphore）")]),t._v(" "),a("h2",{attrs:{id:"len-channel-的应用"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#len-channel-的应用"}},[t._v("#")]),t._v(" len(channel)的应用")]),t._v(" "),a("p",[t._v("len是Go语言的一个内置函数，它支持接收数组、切片、map、字符串和channel类型的参\n数，并返回对应类型的“长度”，也就是一个整型值。\n针对channel ch的类型不同，len(ch)有如下两种语义：\n这样一来，针对带缓冲channel的len调用似乎才是有意义的。那我们是否可以使用len函数来\n实现带缓冲channel的“判满”、“判有”和“判空”逻辑呢？\n当ch为无缓冲channel时，len(ch)总是返回0；\n当ch为带缓冲channel时，len(ch)返回当前channel ch中尚未被读取的元素个数")]),t._v(" "),a("h2",{attrs:{id:"nil-channel的应用"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#nil-channel的应用"}},[t._v("#")]),t._v(" nil channel的应用")]),t._v(" "),a("p",[t._v("如果一个channel类型变量的值为nil，我们称它为nil channel。nil channel有一个特性，那就是对nil channel的读写都会发生阻塞。")])])}),[],!1,null,null,null);a.default=e.exports}}]);