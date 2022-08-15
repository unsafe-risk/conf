
// Package symbols is generated by gogll. Do not edit.
package symbols

type Symbol interface{
	isSymbol()
	IsNonTerminal() bool
	String() string
}

func (NT) isSymbol() {}
func (T) isSymbol() {}

// NT is the type of non-terminals symbols
type NT int
const( 
	NT_Assignment_Statement NT = iota
	NT_Compound_Statement 
	NT_Declaration_Statement 
	NT_Empty_Statement 
	NT_Expression 
	NT_Expression_Statement 
	NT_For 
	NT_Function_Argument 
	NT_Function_Argument_List 
	NT_Function_Argument_List_Body 
	NT_Function_Statement 
	NT_GoGLL 
	NT_Identifier_Expression 
	NT_IncDec_Expression 
	NT_Literal_Expression 
	NT_Return_Statement 
	NT_Simple_Statement 
	NT_Statement 
	NT_Statement_List 
)

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // assign 
	T_1  // break 
	T_2  // colon 
	T_3  // comma 
	T_4  // comment 
	T_5  // config 
	T_6  // continue 
	T_7  // dec 
	T_8  // else 
	T_9  // float_literal 
	T_10  // for 
	T_11  // function 
	T_12  // identifier 
	T_13  // if 
	T_14  // in 
	T_15  // inc 
	T_16  // integer_literal 
	T_17  // lbrace 
	T_18  // let 
	T_19  // lparen 
	T_20  // rbrace 
	T_21  // return 
	T_22  // rparen 
	T_23  // semicolon 
	T_24  // string_literal 
	T_25  // while 
	T_26  // whitespace 
)

type Symbols []Symbol

func (ss Symbols) Strings() []string {
	strs := make([]string, len(ss))
	for i, s := range ss {
		strs[i] = s.String()
	}
	return strs
}

func (NT) IsNonTerminal() bool {
	return true
}

func (T) IsNonTerminal() bool {
	return false
}

func (nt NT) String() string {
	return ntToString[nt]
}

func (t T) String() string {
	return tToString[t]
}

var ntToString = []string { 
	"Assignment_Statement", /* NT_Assignment_Statement */
	"Compound_Statement", /* NT_Compound_Statement */
	"Declaration_Statement", /* NT_Declaration_Statement */
	"Empty_Statement", /* NT_Empty_Statement */
	"Expression", /* NT_Expression */
	"Expression_Statement", /* NT_Expression_Statement */
	"For", /* NT_For */
	"Function_Argument", /* NT_Function_Argument */
	"Function_Argument_List", /* NT_Function_Argument_List */
	"Function_Argument_List_Body", /* NT_Function_Argument_List_Body */
	"Function_Statement", /* NT_Function_Statement */
	"GoGLL", /* NT_GoGLL */
	"Identifier_Expression", /* NT_Identifier_Expression */
	"IncDec_Expression", /* NT_IncDec_Expression */
	"Literal_Expression", /* NT_Literal_Expression */
	"Return_Statement", /* NT_Return_Statement */
	"Simple_Statement", /* NT_Simple_Statement */
	"Statement", /* NT_Statement */
	"Statement_List", /* NT_Statement_List */ 
}

var tToString = []string { 
	"assign", /* T_0 */
	"break", /* T_1 */
	"colon", /* T_2 */
	"comma", /* T_3 */
	"comment", /* T_4 */
	"config", /* T_5 */
	"continue", /* T_6 */
	"dec", /* T_7 */
	"else", /* T_8 */
	"float_literal", /* T_9 */
	"for", /* T_10 */
	"function", /* T_11 */
	"identifier", /* T_12 */
	"if", /* T_13 */
	"in", /* T_14 */
	"inc", /* T_15 */
	"integer_literal", /* T_16 */
	"lbrace", /* T_17 */
	"let", /* T_18 */
	"lparen", /* T_19 */
	"rbrace", /* T_20 */
	"return", /* T_21 */
	"rparen", /* T_22 */
	"semicolon", /* T_23 */
	"string_literal", /* T_24 */
	"while", /* T_25 */
	"whitespace", /* T_26 */ 
}

var stringNT = map[string]NT{ 
	"Assignment_Statement":NT_Assignment_Statement,
	"Compound_Statement":NT_Compound_Statement,
	"Declaration_Statement":NT_Declaration_Statement,
	"Empty_Statement":NT_Empty_Statement,
	"Expression":NT_Expression,
	"Expression_Statement":NT_Expression_Statement,
	"For":NT_For,
	"Function_Argument":NT_Function_Argument,
	"Function_Argument_List":NT_Function_Argument_List,
	"Function_Argument_List_Body":NT_Function_Argument_List_Body,
	"Function_Statement":NT_Function_Statement,
	"GoGLL":NT_GoGLL,
	"Identifier_Expression":NT_Identifier_Expression,
	"IncDec_Expression":NT_IncDec_Expression,
	"Literal_Expression":NT_Literal_Expression,
	"Return_Statement":NT_Return_Statement,
	"Simple_Statement":NT_Simple_Statement,
	"Statement":NT_Statement,
	"Statement_List":NT_Statement_List,
}