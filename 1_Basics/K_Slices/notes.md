### More important points
- Slicing is a way to access a range of elements in a sequence (like a list or string) using the syntax `sequence[start:stop:step]`.
- The `start` index is inclusive, while the `stop` index is exclusive. This means that the slice will include elements from the `start` index up to, but not including, the `stop` index.
- The `step` parameter allows you to specify the interval between elements in the slice. For example, a step of 2 will include every second element in the specified range.
- If the `start` index is omitted, it defaults to the beginning of the sequence. If the `stop` index is omitted, it defaults to the end of the sequence. If the `step` is omitted, it defaults to 1.
- Negative indices can be used to slice from the end of the sequence. For example, `sequence[-1]` refers to the last element, `sequence[-2]` refers to the second-to-last element, and so on.
- Slicing can also be used to create a new list or string that contains only the elements specified in the slice. This is useful for creating sublists or substrings without modifying the original sequence.
- Slicing can be used with any sequence type in Python, including lists, strings, tuples, and more. The syntax is the same for all sequence types.
- Slicing can also be used to reverse a sequence by using a negative step. For example, `sequence[::-1]` will return the sequence in reverse order.
- Slicing can be combined with other operations, such as concatenation and repetition, to create new sequences. For example, `sequence[start:stop] + sequence[start:stop]` will concatenate two slices of the same sequence.

## More links
- [Go Slice](https://golang.org/doc/effective_go.html#slices)
- [more](https://go.dev/doc/effective_go#slices)