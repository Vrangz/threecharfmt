# threecharfmt

A simple formatter making a string more readable, splitting the characters into blocks of three, separated by a separator of your choice

```go
Format("ABCDEF123456", threecharfmt.DefaultSeparator) // outputs: ABC DEF 123 456
Format("ABCDE123456", threecharfmt.DefaultSeparator)  // outputs: ABC DE1 234 56
Format("ABCDE12345", threecharfmt.DefaultSeparator)   // outputs: ABC DE1 23 45
```

Additionally, it allows replacing some characters before making readable form

```go
Format("ABCDE  F123456", threecharfmt.DefaultSeparator, threecharfmt.Replacement{Old: " ", New: ""}) // outputs: ABC DEF 123 456
Format("A BCDE123 456", threecharfmt.DefaultSeparator, threecharfmt.Replacement{Old: " ", New: ""})  // outputs: ABC DE1 234 56
Format("ABCD E1234 5", threecharfmt.DefaultSeparator, threecharfmt.Replacement{Old: " ", New: ""})   // outputs: ABC DE1 23 45
```

For example it can easily format phone numbers

```go
Format("7632935225", threecharfmt.DefaultSeparator})  // outputs: 763 293 52 25
Format("+1884113523", threecharfmt.DefaultSeparator, threecharfmt.Replacement{Old: "+1", New: ""})   // outputs: 884 113 523
```
