# The Conf Script Language
```
package "v8.run/go/conf/script"

GoGLL: Statement_List;
```

# Lex specification
```
string_literal:
   '"'
   {
      (
         '\\' .
         | not "\"\\"
      )
   }
   '"';

integer_literal:
   (
      [any "-+"] any "0123456789" {any "0123456789_"}
      | '0' 'x' any "0123456789abcdefABCDEF" {any "0123456789abcdefABCDEF_"}
      | '0' 'b' any "01" {any "01_"}
      | '0' 'o' any "01234567" {any "01234567_"}
   );

float_literal:
   (
      [any "-+"] any "0123456789" '.' [any "0123456789"]
      |'.' <any "0123456789">
   );

operator: (
   ('*' | '/' | '%')
   | ('+' | '-')
   | ('<' '<' | '>' '>')
   | ('<' | '>' | '<' '=' | '>' '=')
   | ('=' '=' | '!' '=')
   | '&'
   | '^'
   | '|'
   | '&' '&'
   | '|' '|'
);

identifier: letter {letter | number | '_'};

!comment: (
      '/' '/' {not "\n"}
      | '/' '*' {not "*"| '*' not "/" } '*' '/'
      | '#' {not "\n"}
   );
```

# Syntax

```
Assignment_Statement:
   identifier "=" Expression ";";

Declaration_Statement:
   "let" identifier ":" identifier "=" Expression ";"
   | "let" "config" identifier ":" identifier "=" Expression ";";

Expression_Statement:
   Expression ";" | ";";

Function_Argument_List:
   Function_Argument
   | Function_Argument_List_Body;

Function_Argument_List_Body:
   Function_Argument ";"
   | Function_Argument_List_Body Function_Argument
   | Function_Argument_List_Body Function_Argument ";";

Function_Argument:
   identifier ":" identifier;

Function_Statement:
   "func" identifier "(" Function_Argument_List ")" identifier Statement
   | "func" identifier "(" ")" identifier Statement;

IF_Statement
   : "if" "(" Expression ")" Statement
   | "if" "(" Expression ")" Statement "else" Statement;

Return_Statement:
   "return" Expression ";";

Empty_Statement:
   ";";

Simple_Statement:
   Declaration_Statement
   | Assignment_Statement
   | Expression_Statement
   | Empty_Statement;

IncDec_Expression: identifier "++" | identifier "--" | "++" identifier | "--" identifier;

Literal_Expression:
   integer_literal
   | float_literal
   | string_literal;

Identifier_Expression: identifier;

Operator: operator;

Function_Call_Argument_List:
   Expression
   | Expression "," Function_Call_Argument_List;

Function_Call_Expression:
   identifier "(" Function_Call_Argument_List ")" | identifier "(" ")";

Expression:
   "(" Expression ")"
   | Literal_Expression
   | Identifier_Expression
   | Function_Call_Expression
   | IncDec_Expression
   | Expression Operator Expression;

For:
   "for" "(" Simple_Statement Expression_Statement Expression ")" Statement
   | "for" "(" Simple_Statement Expression_Statement ")" Statement;

Statement_List:
   Statement
   | Statement Statement_List;

Compound_Statement:
   "{" "}"
   | "{" Statement_List "}";

Statement:
   Empty_Statement
   | Simple_Statement
   | Compound_Statement
   | Function_Statement
   | IF_Statement
   | Return_Statement
   | For;
```
