# Arrietty

### 手順
- tokenize  -> 文字列をトークンに
- parse     -> トークンを構文解析しParseTreeに
- analyze   -> ParseTreeを意味解析しASTを作る(多分)
- interpret -> ...
- compile   -> ...

### Reserved Idents (keyword)
```text
return, if, else, while, for,
float, int, string, raw, array, dict, bool, true, false, null, void
print, len, type
```

### Grammar
```text
program    = toplevel*
toplevel   = types ident "(" funcParams? ")" block
           | comment
block      = "{" stmt* "}"
stmt       =  expr ";"
           | "return" expr? ";"
           | "if" "(" expr ")" block ("else" block)?
           | "while" "(" expr ")" block
           | "for" "(" expr? ";" expr? ";" expr? ")" block
           | comment
expr       = assign
assign     = "var" ident types ("=" andor)?   // varDecl (and assign)
           | ident ":=" andor                 // short varDecl
           | andor ("=" andor)?               // andor (or assign)
andor      = equality ("&&" equality | "||" equality)*
equality   = relational ("==" relational | "!=" relational)*
relational = add ("<" add | "<=" add | ">" add | ">=" add)*
add        = mul ("+" mul | "-" mul)*
mul        = unary ("*" unary | "/" unary | "%" unary)*
unary      = ("+" | "-" | "!")? primary
primary    = access
access     = literal ("[" expr "]")*
literal = "(" expr ")"
        | ident
        | ident "(" callArgs? ")"
        | float
        | int
        | string
        | raw
        | list
        | dict
        | bool
        | null


types      = "float" | "int" | "string" | "bool" | "void"
           | ident
           | "[" expr? "]" datType
           | "dict" "[" types "]" types


callArgs   = expr ("," expr)*
funcParams = ident types ("," ident types)*


list = "[" unary? "]"
      | "[" unary ("," unary)* "]"

dict  = "{" kv? "}"
      | "{" kv ("," kv)* "}"
kv    = string ":" unary

ident   = [a-zA-Z_][a-zA-Z0-9_]*
float   = [0-9]+[0-9.][0-9]+
int     = [0-9]+
string  = "\"" any-character* "\""
raw     = "`" any-character* "`"
bool    = "true" | "false"
null    = "null"


comment = "#" any-character*
white   = " " | "\t"
newline = "\n"
```

### Node implementation
- [x] Function
- [x] Block
- [x] Return 
- [x] If
- [x] IfElse
- [x] While
- [x] For
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