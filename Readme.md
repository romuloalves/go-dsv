# GO-DSV
> A delimited separated values package for Golang that generate a single line with structs data

**Under development**

## How to use

- Create a struct and include the `dsv` tag in each field. Example:
```
type myTestStruct struct {
  FieldOne string `dsv:"1,50,0"`
  FieldTwo string `dsv:"3,50"`
  FieldThree string `dsv:"2"`
  FieldFour string `dsv:"1,50,0,true"`
}
```

### What tags are supported?

*The sequence below is the same of the flags in example of `How to use`.*

*You NEED to include at least the `dsv` tag with the index or the field will not be written to the response.*

- `index` represents the position of each field in the generated line. (Integer)
- `length` represents the length of the field in the line. (Integer)
- `paddingChar` represents what char will be placed in the empty space until the length be achieved. Default is empty space. (Can't have more than one character) (String – default " ")
- `paddingRight` represents the side of the value that will receive the padding char. (Boolean – default False)

*If you do not want to replace same default value, use `-`. Example:* `dsv:"1,30,-,true"`.

## License

MIT License

Copyright (c) 2017 Rômulo Alves

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