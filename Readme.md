# GO-DSV
> A delimited separated values package for Golang that generate a single line with structs data

**Under development**

## How to use

- Create a struct and include flags to each field. Example:
```
type myTestStruct struct {
  FieldOne string `index:"1" length:"50" paddingChar:"0"`
  FieldTwo string `index:"3" length:"50"`
  FieldThree string `index:"2"`
}
```

### What tags are supported?

- `index` represents the position of each field in the generated line.
- `length` represents the length of the field in the line.
- `paddingChar` represents what char will be placed in the empty space until the length be achieved. Default is empty space.

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