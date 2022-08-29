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
expr       = "!"? assign
assign     = andor ("=" andor)?
andor      = equality ("&&" equality | "||" equality)*
equality   = relational ("==" relational | "!=" relational)*
relational = add ("<" add | "<=" add | ">" add | ">=" add)*
add        = mul ("+" mul | "-" mul)*
mul        = unary ("*" unary | "/" unary | "%" unary)*
unary      = ("+" | "-")? primary
primary    = literal


callArgs   = literal ("," callArgs)?
funcParams = ident ("," funcParams)?


literal = ident
        | ident "(" callArgs? ")"
        | "(" expr ")"
        | float
        | int
        | string
        | raw
        | array
        | dict
        | bool
        | null

array = "[" literal? "]"
      | "[" literal ("," literal)* "]"

dict  = "{" kv? "}"
      | "{" kv ("," kv)* "}"
kv    = string ":" literal

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