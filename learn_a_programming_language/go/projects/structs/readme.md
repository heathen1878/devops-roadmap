# Structs

A struct is a collection of properties that are related but can be different types.

## Pointers

### TL;DR

If you have a variable value and you want the memory address us `&var`, if you have a memory address and want the value use `*var`.

Go is a Pass By Value language, so when for example a receiver function for a particular type is defined, the function will receive a copy of the original value not the actual value. If the type is an `int`
`float`
`string`
`bool`
or `struct` as these are value types.

__Note__
`slices`
`maps`
`channels`
`pointers`
and `functions` are reference types therefore have a pointer back to the data.

Pointers are used to get the memory address of variable of whatever type; this way the receiver function can update the variable of that type directly rather than updating the copy passed into the function.

```go
pointer := &var

pointer.function()
```

To use the address reference place an asterisk in front of the variable that holds the address reference for example...

### Type

Note with this function you can simply pass in the variable of a matching type you don't need to reference the pointer using `&var`

```go
func (pointer *type) function_name () {
    // then use the operator as below
}
```

### Operator

```go
(*.pointer).field = "something"
```

