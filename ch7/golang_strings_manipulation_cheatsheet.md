## I. String Manipulation cheatsheet

Print this out, tape it by your desk, or put it in your favorite three-ring binder for school. It contains almost every string manipulation with an expression as an example for when you're working.

### String Literals


| Expression | Result   | Comment                                                |
|------------|----------|--------------------------------------------------------|
| " "        |          | Empty string is default zero value for the string type |
| Japan 日本   | Japan 日本 | Go has UTF-8 built in                                  |
| \          | \        | Backslash                                              |
| \n         |          | newline                                                |
| \t         |          | tab                                                    |



### Concatenate

| **Expression,** | **Result,** | **Comment**   |
|-----------------|-------------|---------------|
| "Go"\+"Lang",   | Golang,     | concatenation |


### Equal and Compare
| **Expression,**         | **Result,** | **Comment**               |
|-------------------------|---------|---------------------------|
| "Go" == "Go"            | true    | Equality                  |
| strings\.EqualFold\("GoLang"\) | GOLANG  | true unicode case folding |
| "Go" < "go"             | true    | Lexicographic order       |


### Length in bytes or runes
| **Expression,**              | **Result,** | **Comment**     |
|------------------------------|---------|-----------------|
| len\("行く"\)                  | 6       | Length in Bytes |
| utf8\.RuneCountInString\("行く"\) | 2       | in runes        |
| utf8\.ValidString\("行く"\)    | true    | UTF\-8          |


### Index, substring, iterate
| **Expression,**  | **Result,** | **Comment**         |
|------------------|-------------|---------------------|
| "GO"\[1\],       | O           | ,byte position at 1 |
| "GoLang"\[1:3\], | oL          | ,byte indexing      |
| "GO"\[:2\],      | GO          | ,byte indexing      |
| "GO"\[1:\],      | O           | ,byte indexing      |


### Search
| **Expression**                  | **Result** | **Comment** | ****                                |
|---------------------------------|------------|-------------|-------------------------------------|
| strings\.Contains\("GoLang"     | "abc"\)    | false       | Is "abc" in GoLang                  |
| strings\.Contains\("GoLang"     | "abc"\)    | true        | Is 'a' or 'b' or 'c' in Golang      |
| strings\.Count\("Orange"        | "an"\)     | 1           | non\-overlapping instances of range |
| strings\.HasPrefix\("GoLang"    | "Go"\)     | true        | Does GoLang start with 'Go'         |
| strings\.HasSuffix\("GoLang"    | "Lang"\)   | true        | Does GoLang end with 'Lang'         |
| strings\.Index\("GoLang"        | "abc"\)    | \-3         | Index of first abc                  |
| strings\.IndexAny\("GoLang"     | "abc"\)    | 3           | a or b or c                         |
| strings\.LastIndex\("GoLang"    | "abc"\)    | \-3         | Index of last abc                   |
| strings\.LastIndexAny\("GoLang" | "abc"\)    | 3           | a or b or c                         |


### Replace
| **Expression**                  | **Result** | **Comment**                            | ****                             | ****  | ****                            |
|---------------------------------|------------|----------------------------------------|----------------------------------|-------|---------------------------------|
| strings\.Replace\("foo"         | "o"        | "\."                                   | 2\)                              | f\.\. | Replace first two “o” with “\.” |
| strings\.ToUpper\("golang"\)    | GOLANG     | Uppercase                              |                                  |       |                                 |
| strings\.ToLower\("GoLang"\)    | golang     | Lowercase                              |                                  |       |                                 |
| strings\.Title\("go lang"\)     | Go Lang    | Initial letters to uppercase           |                                  |       |                                 |
| strings\.TrimSpace\(" foo\\n"\) | foo        | Strip leading and trailing white space |                                  |       |                                 |
| strings\.Trim\("foo"            | "fo"\)     | Strip                                  | leading and trailing f:s and o:s |       |                                 |
| strings\.TrimLeft\("foo"        | "f"\)      | oo                                     | only leading                     |       |                                 |
| strings\.TrimRight\("foo"       | "o"\)      | f                                      | only trailing                    |       |                                 |
| strings\.TrimPrefix\("foo"      | "fo"\)     | o                                      | only                             |       |                                 |
| strings\.TrimSuffix\("foo"      | "o"\)      | fo                                     | only                             |       |                                 |


### Split by space or comma
| **Expression**                | **Result**  | **Comment**        | **** | ****        | ****             | ****           |
|-------------------------------|-------------|--------------------|------|-------------|------------------|----------------|
| strings\.Fields\(" a	 b\\n"\) | \["a" "b"\] | Remove white space |      |             |                  |                |
| strings\.Split\("a            | b"          | "                  | "\)  | \["a" "b"\] | Remove separator |                |
| strings\.SplitAfter\("a       | b"          | "                  | "\)  | \["a        | " "b"\]          | Keep separator |


### Join strings with a Separator
| **Expression**                 | **Result** | **Comment** | ****             | ****                 |
|--------------------------------|------------|-------------|------------------|----------------------|
| strings\.Join\(\[\]string\{"a" | "b"\}      | ":"\)       | a:b              | Add ':' as separator |
| strings\.Repeat\("da"          | 2\)        | dada        | 2 copies of “da” |                      |

### Format and convert
| **Expression**              | **Result** | **Comment**   |
|-----------------------------|------------|---------------|
| strconv\.Itoa\(\-42\)       | \-42       | Int to string |
| strconv\.FormatInt\(255,16) | ff         | Base 16       |


### Regular Expressions
| **Expression,** | **Meaning**                                     |
|-----------------|-------------------------------------------------|
| \.,             | any character                                   |
| \[ab\],         | the character a or b                            |
| \[^ab\],        | any character except a or b                     |
| \[a\-z\],       | any character from a to z                       |
| \[a\-z0\-9\],   | any character from a to z or 0 to 9             |
| \\d,            | a digit: \[0\-9\]                               |
| \\D,            | a non\-digit: \[^0\-9\]                         |
| \\s,            | a whitespace character: \[                      |
| \\S,            | a non\-whitespace character: \[^\\t\\n\\f\\r \] |
| \\w,            | a word character: \[0\-9A\-Za\-z\_\]            |
| \\W,            | a non\-word character: \[^0\-9A\-Za\-z\_\]      |
| \\p\{Greek\},   | Unicode character class\*                       |
| \\pN,           | one\-letter name                                |
| \\P\{Greek\},   | negated Unicode character class\*               |
| \\PN,           | one\-letter name                                |
| Regexp,         | Meaning                                         |
| x\*,            | zero or more x prefer more                      |
| x\*?,           | prefer fewer \(non\-greedy\)                    |
| x\+,            | one or more x prefer more                       |
| x\+?,           | prefer fewer \(non\-greedy\)                    |
| x?,             | zero or one x prefer one                        |
| x??,            | prefer zero                                     |
| x\{n\},         | exactly n x                                     |
| Regexp,         | Meaning                                         |
| xy,             | x followed by y                                 |
| x\|y,           | x or y prefer x                                 |
| xy\|z,          | same as \(xy\)\|z                               |
| xy\*,           | same as x\(y\*\)                                |
| Symbol,         | Matches                                         |
| \\A,            | at beginning of text                            |
| ^,              | at beginning of text or line                    |
| $,              | at end of text                                  |
| \\b,            | at ASCII word boundary                          |
| \\B,            | not at ASCII word boundary                      |




