# Arrietty

### Grammar
```text
program    = toplevel*
toplevel   = ident "(" funcParams? ")" block
block      = "{" stmt* "}"
stmt       =  expr ";"
           | "return" expr ";"
           | "if" "(" expr ")" block ("else" block)?
           | "while" "(" expr ")" block
           | "for" "(" expr? ";" expr? ";" expr? ")" block
expr       = assign
assign     = andor ("=" andor)?
andor      = equality ("&&" equality | "||" equality)*
equality   = relational ("==" relational | "!=" relational)*
relational = add ("<" add | "<=" add | ">" add | ">=" add)*
add        = mul ("+" mul | "-" mul)*
mul        = unary ("*" unary | "/" unary | "%" unary)*
unary      = ("+" | "-" | "!")? primary
primary    = access
access     = literal ("[" expr "]")*

callArgs   = expr ("," expr)*
funcParams = ident ("," ident)*


literal = "(" expr ")"
        | ident
        | ident "(" callArgs? ")"
        | float
        | int
        | string
        | raw
        | array
        | dict
        | bool
        | null

array = "[" primary? "]"
      | "[" primary ("," primary)* "]"

dict  = "{" kv? "}"
      | "{" kv ("," kv)* "}"
kv    = string ":" primary

ident   = [a-zA-Z_][a-zA-Z0-9_]*
float   = [0-9]+[0-9.][0-9]+
int     = [0-9]+
string  = "\"" any-character* "\""
raw     = "`" any-character* "`"
bool    = "true" | "false"
null    = "null"


comment = "//" any-character*
white   = " " | "\t"
newline = "\n"
```

### Node implementation
- [x] Function
- [x] Block
- [x] Return 
- [ ] If
- [ ] IfElse
- [ ] While
- [ ] For
- [x] Assign
- [x] Not
- [x] And
- [x] Or
- [x] Eq
- [x] Ne
- [x] Lt
- [x] Le
- [x] Gt
- [x] Ge
- [x] Add
- [x] Sub
- [x] Mul
- [x] Div
- [x] Mod
- [x] Ident
- [x] Call
- [x] Float, Int, String
- [x] Array, Dict, KV
- [ ] Raw
- [x] True, False, Null
- [x] Args
- [x] Params
- [x] Access(Dict, Array)
- [ ] Access(String)