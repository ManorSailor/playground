### References

I used the following resources to get started with `Go`. I did look at other resources when a resource couldn't help me understand a topic. Those are not included in here.

- [Learn Go By Alex Mux - Video](https://www.youtube.com/watch?v=8uiZC0l4Ajw&t=4s)
- [Official Go Tour - Text](https://go.dev/tour)

### Notes

Go is a multi-threaded, garbage collected language following the Procedural Paradigm like C. Unlike C tho, it has a Garbage Collector. It doesn't pursue OOPs or Functional Programming Paradigms strictly.

Go's OOP is limited to:
- Encapsulation
- Methods
- Polymorphism (via Interfaces)
- No classes or Inheritance. Composition all the way.

Go's Functional aspects are limited to:
- Functions as first-class citizens, i.e., pass as values to other func;
- Closures. Obviously.
- Higher Order Functions.

Everything is pass by value in Go. This is a design decision taken to ensure Go remains fast. `Pointers` themselves are passed by value, i.e., they are copied, but both copies point to the same object. You can mutate the value at the location but never the pointer - i.e., call-by-sharing or pass-by-value-reference.
An important distinction is with `map`, `slice` & `chan`. They are all pass-by-value-reference. `slice` is even more special, it passed by value for the slice, but the internal array is still pass-by-value-reference.

Functions in `go` do not support optional parameters. In fact, nothing in go can be actually optional. The Go's pattern suggest utilizing struct fields to simulate optional parameters because struct by itself resolves to the `undefined` flavors of its respective type. --- More on this later*

> [!NOTE]
> The reason behind pass by value is a clever trick to reduce GC overhead. Since everything is pass-by-value, the compiler attempts to allocate onto the stack via escape analysis. When the function exits, that memory is automatically reclaimed without GC involvement. However, values that escape their function scope get heap-allocated regardless i.e., closures capturing outer variables, returning a pointer to a local variable, or values too large for the stack. The compiler decides that tho. Still keeps Go fast while avoiding both manual memory management & most GC work. Clever!

---

Packages are directories of go files. Modules are group or collection of packages. Go files themselves have no distinction.
To go unconventional route, you can consider:
- Module as a single Go file.
- Namespace as collection of Module, i.e., Collection of Go files in same directory.
- Packages as Collection of Namespaces.

All data types in `go` have a default value, meaning you can access them without assigning a value.
For JS folks, you can consider this Go's version of `undefined`

Primitive Defaults:
- all `int, uint, float` have `0`.
- `bool` have `false`.
- `string` have an `""` empty string.
- `rune` have `0` - alias for `int32`
- `byte` have `0` - alias for `uint8`
- `interface, pointer, map, chan, slice` all have `nil`.

This doesn't apply to constants tho. They require a value upon definition.

---

While you can use `:=` walrus guy (or D*ck - coined by my Gf) to declare variables without writing `var`, you generally shouldn't (in case of func) as:
- you cannot define a data type
- type of the value is not immediately clear without hovering or going to func/variable definition

Updates:
- Walrus guy is your friend against the ugly ass syntax of this language. Don't alienate it.

---

Arrays are passed-by-value. Not reference. If you need to do that, pass a pointer of that array.

Slices are *view* into an array. They are basic wrappers over primitive arrays to give them vector-like feel from other languages.
Internally, slice can be imagined as a struct containing 3 fields:
- Pointer to the array
- Length of the array, i.e., the actual contents in it
- Capacity of the array, i.e., the room or space left for more `append`'s. Note: explicitly defined capacity are not extended without copying to create an entirely new slice.

Creating a slice from an existing slice does *not* create a copy of the internal array. Both slices point to the respective parts of the original array. A new slice is created tho.

---

Dealing with `strings` can be a PITA. Please be aware of the UTF-8 while looping over them, `index` will be skipped if a non-ascii char is encountered. Cast to `[]rune` if `index` is necessary.

---

`struct` are Go's answer to Objects. Just like `C`. However, `Go` allows struct types to have methods by a special *method* function syntax. An Empty Struct instance prints `{}`. A non-empty struct instance prints the default or `undefined` of its respective data types - nested struct fields print `{}`.

Some facts about `struct`:
- Only properties can be defined on them.
- Methods are attached via:
    - Pointer Receiver Functions
    - Value Receiver Functions

In general, you must use ONE of the following approaches, i.e., do not mix them together on types.


```go
type Counter struct {
    N int
}

func (c Counter) Inc() {
    c.N++
}

func (c *Counter) IncPtr() {
    c.N++
}

c := Counter{}

c.Inc()
c.IncPtr()
```


In the above example, we are defining methods on the `Counter` struct. `Inc` is a *Value-Receiver* method and will not mutate the `N` property of the original Counter instance as it receives a copy of the instance. `IncPtr` is a *Pointer-Receiver* method & will mutate the value of the original instance.

Note: You must think that `c` in `IncPtr` needs to be de-referenced before its properties can be accessed & you'd be right. Go does it implicitly for us so that we don't have to think about it.

---

Public functions or structs are exposed by capitalizing the first letter of their name. On the namespace level, everything is importable regardless of capitalization. In other words, private stuff can only be accessed within the same directory level.

Make fields and methods of a struct public by Capitalizing their first letter:
```go
type Example struct {
    private int
    Public int
}
```

---

`interface` in Go only allows methods. No `struct` properties are allowed. There is no `implements` keyword. Instead, you *implicitly* satisfy an interface by implementing its methods on your `type`. 

Pointers to `interface` are discouraged. Hence, any function accepting an argument with an interface type does **not** have a pointer behind its interface name. In other words, if a `func` accepts 2 different types implementing the same interface and mutates something in them, then the function signature must not define the parameter as a pointer to that interface.

---

Goroutine seem amazing. To *await* for the go-routines to finish, you must tell the parent go routine to wait for them via `WorkGroup.Wait`. This blocks until all go-routines have finished executing.

If a function internally makes use of `WaitGroup` & calls `wg.Wait`, it will block or await the parent function until all goroutines have been processed. The parent doesn't necessarily need to invoke anything. It's a bummer because you cannot know without looking at the definition of the function whether it will block your goroutine or not.

Race conditions, i.e., read or writes to the same memory location across multiple go-routines is unsafe without using Channels or Mutex. It results in *undefined* behavior. Sometimes you might even see it working perfectly across multiple runs only to break during a random run. In fact, the surrounding code can affect the possibility of a race condition, i.e., it may be visible or may not be depending on which code above it is executing.

---

Channels are useful when you don't require manual control via mutex. They are goro/thread-safe method of passing data around. If your goal is to pass data to goros without having to manually write `lock` or `unlock` each time a goro reads or write to memory, then you can use a channel.

Use `make` to create channels. There are:
- Unbuffered Channels, i.e., `chan` with 0 as their capacity.
- Buffered Channels, i.e., `chan` with a fixed capacity > 0.

Some facts about channels:

- A value once read from a channel is gone forever - store it somewhere if required later.
- Channel's default value is always `nil`. Reading from this channel blocks forever.
- Channel has two parties:
    - *Sender/Writer* - responsible for managing the lifecycle of the channel, i.e., write and close.
    - *Receiver/Reader* - responsible for reading from the channel. `close` on read side is compile-time error.
- Channels are *always* FIFO - preserving send/write order. However, with multiple concurrent writers, the order is non-deterministic because it depends on scheduling by the CPU, i.e., when X goro gets to run is not in our hands.
- Channels block the *write* side when their capacity is reached, i.e., they wait until the *read* side takes something out of the channel. 
    - Buffered channel lets the writer write without relying on reader. Write side can keep filling the channel until the `cap` is reached. In this case, the writer is independent of reader.
    - Unbuffered channel ensures that the write side is aware a read happened before the next value is written, i.e., delivery guarantee. They are *Slot-less*. In this case, the writer is blocked until the reader reads.
- The onus of closing the channel is on the writer. Go will throw a deadlock error if a channel wasn't closed as the invoking goro (or reader) will keep on waiting for more values to read.
- Reading from a closed channel returns `undefined` flavor value of the respective type. Writing to a closed channel causes a panic.

---

`select` is Go's equivalent of `switch` for channels. If you've multiple channels open and want to act accordingly when any of them resolves. You opt for `select`. It lets you colocate channel communication logic in a single select body. 

Some facts:

- It blocks until at least **one** channel resolves, i.e., is ready.
- If multiple channels have resolved at the same time, the order of invoking their respective `case` is random.
- Reading from `nil` channels block forever. However, `select` is smart enough to only pick ready channels.
- `nil` channels are never considered ready.

---

Types in Go are strict. Generic Types are basic. 

Some facts:

- Go supports the following ways to define your custom types:
    - Custom Type, i.e., `type A int` which is its own type inheriting all props of `int`, methods can be defined on it.
    - Struct Type, i.e., `type A struct { ... }` defining an object literal.
    - Type Alias, i.e., `type A = int` which is just a different name for same type. Both are interchangeable. No cast required.
    - Union Type, i.e., `type Union interface { ... }` allows defining TS like Union Types. Requires using `interface`.
    - Tagged/Discriminated Unions are not directly supported. We have to use `structs` with `switch` instead.
- Custom Types are branded by default, i.e., if something accepts Email as a type, then only Email will satisfy it, not string.
- Union Types which allow underlying values to satisfy their place can be used by declaring an interface and defining a union of all types starting with `~`. Only works with inbuilt types.
```go
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64...
}
```
- `struct` cannot be self-referential in Go. It raises a compiler error. However, *pointer* to the same types & slices of same type are allowed.
- `func & struct` can be generic in Go. Methods cannot be yet.
- `switch` in Go is used often. It can be used as a *type-switch*, i.e., choose a `case` based on the value's type and the type can be either inbuilt or custom.

