# The Conf Script Language
```
package "v8.run/go/conf/script"

GoGLL: Statement;
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


whitespace: any " \t\r\n";

function: 'f' 'u' 'n' 'c';
let: 'l' 'e' 't';
for: 'f' 'o' 'r';
in: 'i' 'n';
if: 'i' 'f';
else: 'e' 'l' 's' 'e';
while: 'w' 'h' 'i' 'l' 'e';
return: 'r' 'e' 't' 'u' 'r' 'n';
break: 'b' 'r' 'e' 'a' 'k';
continue: 'c' 'o' 'n' 't' 'i' 'n' 'u' 'e';
config: 'c' 'o' 'n' 'f' 'i' 'g';

inc: '+' '+';
dec: '-' '-';

colon: ':';
comma: ',';
assign: '=';
semicolon: ';';

lparen: '(';
rparen: ')';

lbrace: '{';
rbrace: '}';

identifier: upcase {letter | number | '_'};

!comment: (
      '/' '/' {not "\n"}
      | '/' '*' {not "*"| '*' not "/" } '*' '/'
      | '#' {not "\n"}
   );
```

# Syntax

```
Assignment_Statement:
   identifier assign Expression semicolon;

Declaration_Statement:
   let identifier colon identifier assign Expression semicolon
   | let config identifier colon identifier assign Expression semicolon;

Expression_Statement:
   Expression semicolon | semicolon;

Function_Argument_List:
   Function_Argument
   | Function_Argument_List_Body;

Function_Argument_List_Body:
   Function_Argument comma
   | Function_Argument_List_Body Function_Argument
   | Function_Argument_List_Body Function_Argument comma;

Function_Argument:
   identifier colon identifier;

Function_Statement:
   function identifier lparen Function_Argument_List rparen identifier Statement
   | function identifier lparen rparen identifier Statement;

Empty_Statement:
   semicolon;

Simple_Statement:
   Declaration_Statement
   | Assignment_Statement
   | Expression_Statement
   | Empty_Statement;

IncDec_Expression: identifier inc | identifier dec;

Literal_Expression:
   integer_literal
   | float_literal
   | string_literal;

Expression:
   IncDec_Expression
   | Literal_Expression;

For:
   for lparen Simple_Statement Expression_Statement Expression rparen Statement
   | for lparen Simple_Statement Expression_Statement rparen Statement;

Statement_List:
   Statement
   | Statement_List Statement;

Compound_Statement:
   lbrace rbrace
   | lbrace Statement_List rbrace;

Statement:
   Empty_Statement
   | Simple_Statement
   | Compound_Statement
   | Function_Statement
   | For;
```
