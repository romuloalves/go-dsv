# GO-DSV
> A delimited separated values package for Golang that generate or read a single line

## How to use

- Create a struct and include the `dsv` tag in each field. For example:
```
type MyTestStruct struct {
  FieldOne string `dsv:"1,50,0"`
  FieldTwo string `dsv:"3,50"`
  FieldThree string `dsv:"2"`
  FieldFour string `dsv:"1,50,0,true"`
}
```

### What tags are supported?

*In order to write or read, you need to add the `dsv` tag to the fields containing at least their indexes.*

Tag values order:

1. `index`: represents the position of the field in the line. (Integer)
2. `length`: represents the length of the field in the line. (Integer)
3. `paddingChar`: represents what char that is/will be placed in the empty spaces until the length be achieved. The default value is an empty space. (Can't have more than one character) (String – default " ")
4. `paddingRight`: represents the side of the value that contains/will receive the padding char. (Boolean – default False)

*If you do not want to replace some default value, use `-` instead. Example:* `dsv:"1,30,-,true"`.

## License

MIT License

Copyright (c) 2020 Rômulo Alves

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.