
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"v8.run/go/conf/script/parser/symbols"
)

type Label int

const(
	Assignment_Statement0R0 Label = iota
	Assignment_Statement0R1
	Assignment_Statement0R2
	Assignment_Statement0R3
	Assignment_Statement0R4
	Compound_Statement0R0
	Compound_Statement0R1
	Compound_Statement0R2
	Compound_Statement1R0
	Compound_Statement1R1
	Compound_Statement1R2
	Compound_Statement1R3
	Declaration_Statement0R0
	Declaration_Statement0R1
	Declaration_Statement0R2
	Declaration_Statement0R3
	Declaration_Statement0R4
	Declaration_Statement0R5
	Declaration_Statement0R6
	Declaration_Statement0R7
	Declaration_Statement1R0
	Declaration_Statement1R1
	Declaration_Statement1R2
	Declaration_Statement1R3
	Declaration_Statement1R4
	Declaration_Statement1R5
	Declaration_Statement1R6
	Declaration_Statement1R7
	Declaration_Statement1R8
	Empty_Statement0R0
	Empty_Statement0R1
	Expression0R0
	Expression0R1
	Expression1R0
	Expression1R1
	Expression_Statement0R0
	Expression_Statement0R1
	Expression_Statement0R2
	Expression_Statement1R0
	Expression_Statement1R1
	For0R0
	For0R1
	For0R2
	For0R3
	For0R4
	For0R5
	For0R6
	For0R7
	For1R0
	For1R1
	For1R2
	For1R3
	For1R4
	For1R5
	For1R6
	Function_Argument0R0
	Function_Argument0R1
	Function_Argument0R2
	Function_Argument0R3
	Function_Argument_List0R0
	Function_Argument_List0R1
	Function_Argument_List1R0
	Function_Argument_List1R1
	Function_Argument_List_Body0R0
	Function_Argument_List_Body0R1
	Function_Argument_List_Body0R2
	Function_Argument_List_Body1R0
	Function_Argument_List_Body1R1
	Function_Argument_List_Body1R2
	Function_Argument_List_Body2R0
	Function_Argument_List_Body2R1
	Function_Argument_List_Body2R2
	Function_Argument_List_Body2R3
	Function_Statement0R0
	Function_Statement0R1
	Function_Statement0R2
	Function_Statement0R3
	Function_Statement0R4
	Function_Statement0R5
	Function_Statement0R6
	Function_Statement0R7
	Function_Statement1R0
	Function_Statement1R1
	Function_Statement1R2
	Function_Statement1R3
	Function_Statement1R4
	Function_Statement1R5
	Function_Statement1R6
	GoGLL0R0
	GoGLL0R1
	IncDec_Expression0R0
	IncDec_Expression0R1
	IncDec_Expression0R2
	IncDec_Expression1R0
	IncDec_Expression1R1
	IncDec_Expression1R2
	Literal_Expression0R0
	Literal_Expression0R1
	Literal_Expression1R0
	Literal_Expression1R1
	Literal_Expression2R0
	Literal_Expression2R1
	Simple_Statement0R0
	Simple_Statement0R1
	Simple_Statement1R0
	Simple_Statement1R1
	Simple_Statement2R0
	Simple_Statement2R1
	Simple_Statement3R0
	Simple_Statement3R1
	Statement0R0
	Statement0R1
	Statement1R0
	Statement1R1
	Statement2R0
	Statement2R1
	Statement3R0
	Statement3R1
	Statement4R0
	Statement4R1
	Statement_List0R0
	Statement_List0R1
	Statement_List1R0
	Statement_List1R1
	Statement_List1R2
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
	Label 	Label
}

type Index struct {
	NT      symbols.NT
	Alt     int
	Pos     int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
	Assignment_Statement0R0: {
		symbols.NT_Assignment_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Assignment_Statement0R0, 
	},
	Assignment_Statement0R1: {
		symbols.NT_Assignment_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Assignment_Statement0R1, 
	},
	Assignment_Statement0R2: {
		symbols.NT_Assignment_Statement, 0, 2, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Assignment_Statement0R2, 
	},
	Assignment_Statement0R3: {
		symbols.NT_Assignment_Statement, 0, 3, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Assignment_Statement0R3, 
	},
	Assignment_Statement0R4: {
		symbols.NT_Assignment_Statement, 0, 4, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Assignment_Statement0R4, 
	},
	Compound_Statement0R0: {
		symbols.NT_Compound_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.T_20,
		}, 
		Compound_Statement0R0, 
	},
	Compound_Statement0R1: {
		symbols.NT_Compound_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.T_20,
		}, 
		Compound_Statement0R1, 
	},
	Compound_Statement0R2: {
		symbols.NT_Compound_Statement, 0, 2, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.T_20,
		}, 
		Compound_Statement0R2, 
	},
	Compound_Statement1R0: {
		symbols.NT_Compound_Statement, 1, 0, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.NT_Statement_List, 
			symbols.T_20,
		}, 
		Compound_Statement1R0, 
	},
	Compound_Statement1R1: {
		symbols.NT_Compound_Statement, 1, 1, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.NT_Statement_List, 
			symbols.T_20,
		}, 
		Compound_Statement1R1, 
	},
	Compound_Statement1R2: {
		symbols.NT_Compound_Statement, 1, 2, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.NT_Statement_List, 
			symbols.T_20,
		}, 
		Compound_Statement1R2, 
	},
	Compound_Statement1R3: {
		symbols.NT_Compound_Statement, 1, 3, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.NT_Statement_List, 
			symbols.T_20,
		}, 
		Compound_Statement1R3, 
	},
	Declaration_Statement0R0: {
		symbols.NT_Declaration_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement0R0, 
	},
	Declaration_Statement0R1: {
		symbols.NT_Declaration_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement0R1, 
	},
	Declaration_Statement0R2: {
		symbols.NT_Declaration_Statement, 0, 2, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement0R2, 
	},
	Declaration_Statement0R3: {
		symbols.NT_Declaration_Statement, 0, 3, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement0R3, 
	},
	Declaration_Statement0R4: {
		symbols.NT_Declaration_Statement, 0, 4, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement0R4, 
	},
	Declaration_Statement0R5: {
		symbols.NT_Declaration_Statement, 0, 5, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement0R5, 
	},
	Declaration_Statement0R6: {
		symbols.NT_Declaration_Statement, 0, 6, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement0R6, 
	},
	Declaration_Statement0R7: {
		symbols.NT_Declaration_Statement, 0, 7, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement0R7, 
	},
	Declaration_Statement1R0: {
		symbols.NT_Declaration_Statement, 1, 0, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R0, 
	},
	Declaration_Statement1R1: {
		symbols.NT_Declaration_Statement, 1, 1, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R1, 
	},
	Declaration_Statement1R2: {
		symbols.NT_Declaration_Statement, 1, 2, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R2, 
	},
	Declaration_Statement1R3: {
		symbols.NT_Declaration_Statement, 1, 3, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R3, 
	},
	Declaration_Statement1R4: {
		symbols.NT_Declaration_Statement, 1, 4, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R4, 
	},
	Declaration_Statement1R5: {
		symbols.NT_Declaration_Statement, 1, 5, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R5, 
	},
	Declaration_Statement1R6: {
		symbols.NT_Declaration_Statement, 1, 6, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R6, 
	},
	Declaration_Statement1R7: {
		symbols.NT_Declaration_Statement, 1, 7, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R7, 
	},
	Declaration_Statement1R8: {
		symbols.NT_Declaration_Statement, 1, 8, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.T_5, 
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12, 
			symbols.T_0, 
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Declaration_Statement1R8, 
	},
	Empty_Statement0R0: {
		symbols.NT_Empty_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.T_23,
		}, 
		Empty_Statement0R0, 
	},
	Empty_Statement0R1: {
		symbols.NT_Empty_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.T_23,
		}, 
		Empty_Statement0R1, 
	},
	Expression0R0: {
		symbols.NT_Expression, 0, 0, 
		symbols.Symbols{  
			symbols.NT_IncDec_Expression,
		}, 
		Expression0R0, 
	},
	Expression0R1: {
		symbols.NT_Expression, 0, 1, 
		symbols.Symbols{  
			symbols.NT_IncDec_Expression,
		}, 
		Expression0R1, 
	},
	Expression1R0: {
		symbols.NT_Expression, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Literal_Expression,
		}, 
		Expression1R0, 
	},
	Expression1R1: {
		symbols.NT_Expression, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Literal_Expression,
		}, 
		Expression1R1, 
	},
	Expression_Statement0R0: {
		symbols.NT_Expression_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Expression_Statement0R0, 
	},
	Expression_Statement0R1: {
		symbols.NT_Expression_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Expression_Statement0R1, 
	},
	Expression_Statement0R2: {
		symbols.NT_Expression_Statement, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Expression, 
			symbols.T_23,
		}, 
		Expression_Statement0R2, 
	},
	Expression_Statement1R0: {
		symbols.NT_Expression_Statement, 1, 0, 
		symbols.Symbols{  
			symbols.T_23,
		}, 
		Expression_Statement1R0, 
	},
	Expression_Statement1R1: {
		symbols.NT_Expression_Statement, 1, 1, 
		symbols.Symbols{  
			symbols.T_23,
		}, 
		Expression_Statement1R1, 
	},
	For0R0: {
		symbols.NT_For, 0, 0, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.NT_Expression, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For0R0, 
	},
	For0R1: {
		symbols.NT_For, 0, 1, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.NT_Expression, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For0R1, 
	},
	For0R2: {
		symbols.NT_For, 0, 2, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.NT_Expression, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For0R2, 
	},
	For0R3: {
		symbols.NT_For, 0, 3, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.NT_Expression, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For0R3, 
	},
	For0R4: {
		symbols.NT_For, 0, 4, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.NT_Expression, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For0R4, 
	},
	For0R5: {
		symbols.NT_For, 0, 5, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.NT_Expression, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For0R5, 
	},
	For0R6: {
		symbols.NT_For, 0, 6, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.NT_Expression, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For0R6, 
	},
	For0R7: {
		symbols.NT_For, 0, 7, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.NT_Expression, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For0R7, 
	},
	For1R0: {
		symbols.NT_For, 1, 0, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For1R0, 
	},
	For1R1: {
		symbols.NT_For, 1, 1, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For1R1, 
	},
	For1R2: {
		symbols.NT_For, 1, 2, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For1R2, 
	},
	For1R3: {
		symbols.NT_For, 1, 3, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For1R3, 
	},
	For1R4: {
		symbols.NT_For, 1, 4, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For1R4, 
	},
	For1R5: {
		symbols.NT_For, 1, 5, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For1R5, 
	},
	For1R6: {
		symbols.NT_For, 1, 6, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_19, 
			symbols.NT_Simple_Statement, 
			symbols.NT_Expression_Statement, 
			symbols.T_22, 
			symbols.NT_Statement,
		}, 
		For1R6, 
	},
	Function_Argument0R0: {
		symbols.NT_Function_Argument, 0, 0, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12,
		}, 
		Function_Argument0R0, 
	},
	Function_Argument0R1: {
		symbols.NT_Function_Argument, 0, 1, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12,
		}, 
		Function_Argument0R1, 
	},
	Function_Argument0R2: {
		symbols.NT_Function_Argument, 0, 2, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12,
		}, 
		Function_Argument0R2, 
	},
	Function_Argument0R3: {
		symbols.NT_Function_Argument, 0, 3, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_2, 
			symbols.T_12,
		}, 
		Function_Argument0R3, 
	},
	Function_Argument_List0R0: {
		symbols.NT_Function_Argument_List, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Function_Argument,
		}, 
		Function_Argument_List0R0, 
	},
	Function_Argument_List0R1: {
		symbols.NT_Function_Argument_List, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Function_Argument,
		}, 
		Function_Argument_List0R1, 
	},
	Function_Argument_List1R0: {
		symbols.NT_Function_Argument_List, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body,
		}, 
		Function_Argument_List1R0, 
	},
	Function_Argument_List1R1: {
		symbols.NT_Function_Argument_List, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body,
		}, 
		Function_Argument_List1R1, 
	},
	Function_Argument_List_Body0R0: {
		symbols.NT_Function_Argument_List_Body, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Function_Argument, 
			symbols.T_3,
		}, 
		Function_Argument_List_Body0R0, 
	},
	Function_Argument_List_Body0R1: {
		symbols.NT_Function_Argument_List_Body, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Function_Argument, 
			symbols.T_3,
		}, 
		Function_Argument_List_Body0R1, 
	},
	Function_Argument_List_Body0R2: {
		symbols.NT_Function_Argument_List_Body, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Function_Argument, 
			symbols.T_3,
		}, 
		Function_Argument_List_Body0R2, 
	},
	Function_Argument_List_Body1R0: {
		symbols.NT_Function_Argument_List_Body, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body, 
			symbols.NT_Function_Argument,
		}, 
		Function_Argument_List_Body1R0, 
	},
	Function_Argument_List_Body1R1: {
		symbols.NT_Function_Argument_List_Body, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body, 
			symbols.NT_Function_Argument,
		}, 
		Function_Argument_List_Body1R1, 
	},
	Function_Argument_List_Body1R2: {
		symbols.NT_Function_Argument_List_Body, 1, 2, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body, 
			symbols.NT_Function_Argument,
		}, 
		Function_Argument_List_Body1R2, 
	},
	Function_Argument_List_Body2R0: {
		symbols.NT_Function_Argument_List_Body, 2, 0, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body, 
			symbols.NT_Function_Argument, 
			symbols.T_3,
		}, 
		Function_Argument_List_Body2R0, 
	},
	Function_Argument_List_Body2R1: {
		symbols.NT_Function_Argument_List_Body, 2, 1, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body, 
			symbols.NT_Function_Argument, 
			symbols.T_3,
		}, 
		Function_Argument_List_Body2R1, 
	},
	Function_Argument_List_Body2R2: {
		symbols.NT_Function_Argument_List_Body, 2, 2, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body, 
			symbols.NT_Function_Argument, 
			symbols.T_3,
		}, 
		Function_Argument_List_Body2R2, 
	},
	Function_Argument_List_Body2R3: {
		symbols.NT_Function_Argument_List_Body, 2, 3, 
		symbols.Symbols{  
			symbols.NT_Function_Argument_List_Body, 
			symbols.NT_Function_Argument, 
			symbols.T_3,
		}, 
		Function_Argument_List_Body2R3, 
	},
	Function_Statement0R0: {
		symbols.NT_Function_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.NT_Function_Argument_List, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement0R0, 
	},
	Function_Statement0R1: {
		symbols.NT_Function_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.NT_Function_Argument_List, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement0R1, 
	},
	Function_Statement0R2: {
		symbols.NT_Function_Statement, 0, 2, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.NT_Function_Argument_List, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement0R2, 
	},
	Function_Statement0R3: {
		symbols.NT_Function_Statement, 0, 3, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.NT_Function_Argument_List, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement0R3, 
	},
	Function_Statement0R4: {
		symbols.NT_Function_Statement, 0, 4, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.NT_Function_Argument_List, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement0R4, 
	},
	Function_Statement0R5: {
		symbols.NT_Function_Statement, 0, 5, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.NT_Function_Argument_List, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement0R5, 
	},
	Function_Statement0R6: {
		symbols.NT_Function_Statement, 0, 6, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.NT_Function_Argument_List, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement0R6, 
	},
	Function_Statement0R7: {
		symbols.NT_Function_Statement, 0, 7, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.NT_Function_Argument_List, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement0R7, 
	},
	Function_Statement1R0: {
		symbols.NT_Function_Statement, 1, 0, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement1R0, 
	},
	Function_Statement1R1: {
		symbols.NT_Function_Statement, 1, 1, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement1R1, 
	},
	Function_Statement1R2: {
		symbols.NT_Function_Statement, 1, 2, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement1R2, 
	},
	Function_Statement1R3: {
		symbols.NT_Function_Statement, 1, 3, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement1R3, 
	},
	Function_Statement1R4: {
		symbols.NT_Function_Statement, 1, 4, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement1R4, 
	},
	Function_Statement1R5: {
		symbols.NT_Function_Statement, 1, 5, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement1R5, 
	},
	Function_Statement1R6: {
		symbols.NT_Function_Statement, 1, 6, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.T_12, 
			symbols.T_19, 
			symbols.T_22, 
			symbols.T_12, 
			symbols.NT_Statement,
		}, 
		Function_Statement1R6, 
	},
	GoGLL0R0: {
		symbols.NT_GoGLL, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Statement,
		}, 
		GoGLL0R0, 
	},
	GoGLL0R1: {
		symbols.NT_GoGLL, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Statement,
		}, 
		GoGLL0R1, 
	},
	IncDec_Expression0R0: {
		symbols.NT_IncDec_Expression, 0, 0, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_15,
		}, 
		IncDec_Expression0R0, 
	},
	IncDec_Expression0R1: {
		symbols.NT_IncDec_Expression, 0, 1, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_15,
		}, 
		IncDec_Expression0R1, 
	},
	IncDec_Expression0R2: {
		symbols.NT_IncDec_Expression, 0, 2, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_15,
		}, 
		IncDec_Expression0R2, 
	},
	IncDec_Expression1R0: {
		symbols.NT_IncDec_Expression, 1, 0, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_7,
		}, 
		IncDec_Expression1R0, 
	},
	IncDec_Expression1R1: {
		symbols.NT_IncDec_Expression, 1, 1, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_7,
		}, 
		IncDec_Expression1R1, 
	},
	IncDec_Expression1R2: {
		symbols.NT_IncDec_Expression, 1, 2, 
		symbols.Symbols{  
			symbols.T_12, 
			symbols.T_7,
		}, 
		IncDec_Expression1R2, 
	},
	Literal_Expression0R0: {
		symbols.NT_Literal_Expression, 0, 0, 
		symbols.Symbols{  
			symbols.T_16,
		}, 
		Literal_Expression0R0, 
	},
	Literal_Expression0R1: {
		symbols.NT_Literal_Expression, 0, 1, 
		symbols.Symbols{  
			symbols.T_16,
		}, 
		Literal_Expression0R1, 
	},
	Literal_Expression1R0: {
		symbols.NT_Literal_Expression, 1, 0, 
		symbols.Symbols{  
			symbols.T_9,
		}, 
		Literal_Expression1R0, 
	},
	Literal_Expression1R1: {
		symbols.NT_Literal_Expression, 1, 1, 
		symbols.Symbols{  
			symbols.T_9,
		}, 
		Literal_Expression1R1, 
	},
	Literal_Expression2R0: {
		symbols.NT_Literal_Expression, 2, 0, 
		symbols.Symbols{  
			symbols.T_24,
		}, 
		Literal_Expression2R0, 
	},
	Literal_Expression2R1: {
		symbols.NT_Literal_Expression, 2, 1, 
		symbols.Symbols{  
			symbols.T_24,
		}, 
		Literal_Expression2R1, 
	},
	Simple_Statement0R0: {
		symbols.NT_Simple_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Declaration_Statement,
		}, 
		Simple_Statement0R0, 
	},
	Simple_Statement0R1: {
		symbols.NT_Simple_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Declaration_Statement,
		}, 
		Simple_Statement0R1, 
	},
	Simple_Statement1R0: {
		symbols.NT_Simple_Statement, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Assignment_Statement,
		}, 
		Simple_Statement1R0, 
	},
	Simple_Statement1R1: {
		symbols.NT_Simple_Statement, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Assignment_Statement,
		}, 
		Simple_Statement1R1, 
	},
	Simple_Statement2R0: {
		symbols.NT_Simple_Statement, 2, 0, 
		symbols.Symbols{  
			symbols.NT_Expression_Statement,
		}, 
		Simple_Statement2R0, 
	},
	Simple_Statement2R1: {
		symbols.NT_Simple_Statement, 2, 1, 
		symbols.Symbols{  
			symbols.NT_Expression_Statement,
		}, 
		Simple_Statement2R1, 
	},
	Simple_Statement3R0: {
		symbols.NT_Simple_Statement, 3, 0, 
		symbols.Symbols{  
			symbols.NT_Empty_Statement,
		}, 
		Simple_Statement3R0, 
	},
	Simple_Statement3R1: {
		symbols.NT_Simple_Statement, 3, 1, 
		symbols.Symbols{  
			symbols.NT_Empty_Statement,
		}, 
		Simple_Statement3R1, 
	},
	Statement0R0: {
		symbols.NT_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Empty_Statement,
		}, 
		Statement0R0, 
	},
	Statement0R1: {
		symbols.NT_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Empty_Statement,
		}, 
		Statement0R1, 
	},
	Statement1R0: {
		symbols.NT_Statement, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Simple_Statement,
		}, 
		Statement1R0, 
	},
	Statement1R1: {
		symbols.NT_Statement, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Simple_Statement,
		}, 
		Statement1R1, 
	},
	Statement2R0: {
		symbols.NT_Statement, 2, 0, 
		symbols.Symbols{  
			symbols.NT_Compound_Statement,
		}, 
		Statement2R0, 
	},
	Statement2R1: {
		symbols.NT_Statement, 2, 1, 
		symbols.Symbols{  
			symbols.NT_Compound_Statement,
		}, 
		Statement2R1, 
	},
	Statement3R0: {
		symbols.NT_Statement, 3, 0, 
		symbols.Symbols{  
			symbols.NT_Function_Statement,
		}, 
		Statement3R0, 
	},
	Statement3R1: {
		symbols.NT_Statement, 3, 1, 
		symbols.Symbols{  
			symbols.NT_Function_Statement,
		}, 
		Statement3R1, 
	},
	Statement4R0: {
		symbols.NT_Statement, 4, 0, 
		symbols.Symbols{  
			symbols.NT_For,
		}, 
		Statement4R0, 
	},
	Statement4R1: {
		symbols.NT_Statement, 4, 1, 
		symbols.Symbols{  
			symbols.NT_For,
		}, 
		Statement4R1, 
	},
	Statement_List0R0: {
		symbols.NT_Statement_List, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Statement,
		}, 
		Statement_List0R0, 
	},
	Statement_List0R1: {
		symbols.NT_Statement_List, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Statement,
		}, 
		Statement_List0R1, 
	},
	Statement_List1R0: {
		symbols.NT_Statement_List, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Statement_List, 
			symbols.NT_Statement,
		}, 
		Statement_List1R0, 
	},
	Statement_List1R1: {
		symbols.NT_Statement_List, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Statement_List, 
			symbols.NT_Statement,
		}, 
		Statement_List1R1, 
	},
	Statement_List1R2: {
		symbols.NT_Statement_List, 1, 2, 
		symbols.Symbols{  
			symbols.NT_Statement_List, 
			symbols.NT_Statement,
		}, 
		Statement_List1R2, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_Assignment_Statement,0,0 }: Assignment_Statement0R0,
	Index{ symbols.NT_Assignment_Statement,0,1 }: Assignment_Statement0R1,
	Index{ symbols.NT_Assignment_Statement,0,2 }: Assignment_Statement0R2,
	Index{ symbols.NT_Assignment_Statement,0,3 }: Assignment_Statement0R3,
	Index{ symbols.NT_Assignment_Statement,0,4 }: Assignment_Statement0R4,
	Index{ symbols.NT_Compound_Statement,0,0 }: Compound_Statement0R0,
	Index{ symbols.NT_Compound_Statement,0,1 }: Compound_Statement0R1,
	Index{ symbols.NT_Compound_Statement,0,2 }: Compound_Statement0R2,
	Index{ symbols.NT_Compound_Statement,1,0 }: Compound_Statement1R0,
	Index{ symbols.NT_Compound_Statement,1,1 }: Compound_Statement1R1,
	Index{ symbols.NT_Compound_Statement,1,2 }: Compound_Statement1R2,
	Index{ symbols.NT_Compound_Statement,1,3 }: Compound_Statement1R3,
	Index{ symbols.NT_Declaration_Statement,0,0 }: Declaration_Statement0R0,
	Index{ symbols.NT_Declaration_Statement,0,1 }: Declaration_Statement0R1,
	Index{ symbols.NT_Declaration_Statement,0,2 }: Declaration_Statement0R2,
	Index{ symbols.NT_Declaration_Statement,0,3 }: Declaration_Statement0R3,
	Index{ symbols.NT_Declaration_Statement,0,4 }: Declaration_Statement0R4,
	Index{ symbols.NT_Declaration_Statement,0,5 }: Declaration_Statement0R5,
	Index{ symbols.NT_Declaration_Statement,0,6 }: Declaration_Statement0R6,
	Index{ symbols.NT_Declaration_Statement,0,7 }: Declaration_Statement0R7,
	Index{ symbols.NT_Declaration_Statement,1,0 }: Declaration_Statement1R0,
	Index{ symbols.NT_Declaration_Statement,1,1 }: Declaration_Statement1R1,
	Index{ symbols.NT_Declaration_Statement,1,2 }: Declaration_Statement1R2,
	Index{ symbols.NT_Declaration_Statement,1,3 }: Declaration_Statement1R3,
	Index{ symbols.NT_Declaration_Statement,1,4 }: Declaration_Statement1R4,
	Index{ symbols.NT_Declaration_Statement,1,5 }: Declaration_Statement1R5,
	Index{ symbols.NT_Declaration_Statement,1,6 }: Declaration_Statement1R6,
	Index{ symbols.NT_Declaration_Statement,1,7 }: Declaration_Statement1R7,
	Index{ symbols.NT_Declaration_Statement,1,8 }: Declaration_Statement1R8,
	Index{ symbols.NT_Empty_Statement,0,0 }: Empty_Statement0R0,
	Index{ symbols.NT_Empty_Statement,0,1 }: Empty_Statement0R1,
	Index{ symbols.NT_Expression,0,0 }: Expression0R0,
	Index{ symbols.NT_Expression,0,1 }: Expression0R1,
	Index{ symbols.NT_Expression,1,0 }: Expression1R0,
	Index{ symbols.NT_Expression,1,1 }: Expression1R1,
	Index{ symbols.NT_Expression_Statement,0,0 }: Expression_Statement0R0,
	Index{ symbols.NT_Expression_Statement,0,1 }: Expression_Statement0R1,
	Index{ symbols.NT_Expression_Statement,0,2 }: Expression_Statement0R2,
	Index{ symbols.NT_Expression_Statement,1,0 }: Expression_Statement1R0,
	Index{ symbols.NT_Expression_Statement,1,1 }: Expression_Statement1R1,
	Index{ symbols.NT_For,0,0 }: For0R0,
	Index{ symbols.NT_For,0,1 }: For0R1,
	Index{ symbols.NT_For,0,2 }: For0R2,
	Index{ symbols.NT_For,0,3 }: For0R3,
	Index{ symbols.NT_For,0,4 }: For0R4,
	Index{ symbols.NT_For,0,5 }: For0R5,
	Index{ symbols.NT_For,0,6 }: For0R6,
	Index{ symbols.NT_For,0,7 }: For0R7,
	Index{ symbols.NT_For,1,0 }: For1R0,
	Index{ symbols.NT_For,1,1 }: For1R1,
	Index{ symbols.NT_For,1,2 }: For1R2,
	Index{ symbols.NT_For,1,3 }: For1R3,
	Index{ symbols.NT_For,1,4 }: For1R4,
	Index{ symbols.NT_For,1,5 }: For1R5,
	Index{ symbols.NT_For,1,6 }: For1R6,
	Index{ symbols.NT_Function_Argument,0,0 }: Function_Argument0R0,
	Index{ symbols.NT_Function_Argument,0,1 }: Function_Argument0R1,
	Index{ symbols.NT_Function_Argument,0,2 }: Function_Argument0R2,
	Index{ symbols.NT_Function_Argument,0,3 }: Function_Argument0R3,
	Index{ symbols.NT_Function_Argument_List,0,0 }: Function_Argument_List0R0,
	Index{ symbols.NT_Function_Argument_List,0,1 }: Function_Argument_List0R1,
	Index{ symbols.NT_Function_Argument_List,1,0 }: Function_Argument_List1R0,
	Index{ symbols.NT_Function_Argument_List,1,1 }: Function_Argument_List1R1,
	Index{ symbols.NT_Function_Argument_List_Body,0,0 }: Function_Argument_List_Body0R0,
	Index{ symbols.NT_Function_Argument_List_Body,0,1 }: Function_Argument_List_Body0R1,
	Index{ symbols.NT_Function_Argument_List_Body,0,2 }: Function_Argument_List_Body0R2,
	Index{ symbols.NT_Function_Argument_List_Body,1,0 }: Function_Argument_List_Body1R0,
	Index{ symbols.NT_Function_Argument_List_Body,1,1 }: Function_Argument_List_Body1R1,
	Index{ symbols.NT_Function_Argument_List_Body,1,2 }: Function_Argument_List_Body1R2,
	Index{ symbols.NT_Function_Argument_List_Body,2,0 }: Function_Argument_List_Body2R0,
	Index{ symbols.NT_Function_Argument_List_Body,2,1 }: Function_Argument_List_Body2R1,
	Index{ symbols.NT_Function_Argument_List_Body,2,2 }: Function_Argument_List_Body2R2,
	Index{ symbols.NT_Function_Argument_List_Body,2,3 }: Function_Argument_List_Body2R3,
	Index{ symbols.NT_Function_Statement,0,0 }: Function_Statement0R0,
	Index{ symbols.NT_Function_Statement,0,1 }: Function_Statement0R1,
	Index{ symbols.NT_Function_Statement,0,2 }: Function_Statement0R2,
	Index{ symbols.NT_Function_Statement,0,3 }: Function_Statement0R3,
	Index{ symbols.NT_Function_Statement,0,4 }: Function_Statement0R4,
	Index{ symbols.NT_Function_Statement,0,5 }: Function_Statement0R5,
	Index{ symbols.NT_Function_Statement,0,6 }: Function_Statement0R6,
	Index{ symbols.NT_Function_Statement,0,7 }: Function_Statement0R7,
	Index{ symbols.NT_Function_Statement,1,0 }: Function_Statement1R0,
	Index{ symbols.NT_Function_Statement,1,1 }: Function_Statement1R1,
	Index{ symbols.NT_Function_Statement,1,2 }: Function_Statement1R2,
	Index{ symbols.NT_Function_Statement,1,3 }: Function_Statement1R3,
	Index{ symbols.NT_Function_Statement,1,4 }: Function_Statement1R4,
	Index{ symbols.NT_Function_Statement,1,5 }: Function_Statement1R5,
	Index{ symbols.NT_Function_Statement,1,6 }: Function_Statement1R6,
	Index{ symbols.NT_GoGLL,0,0 }: GoGLL0R0,
	Index{ symbols.NT_GoGLL,0,1 }: GoGLL0R1,
	Index{ symbols.NT_IncDec_Expression,0,0 }: IncDec_Expression0R0,
	Index{ symbols.NT_IncDec_Expression,0,1 }: IncDec_Expression0R1,
	Index{ symbols.NT_IncDec_Expression,0,2 }: IncDec_Expression0R2,
	Index{ symbols.NT_IncDec_Expression,1,0 }: IncDec_Expression1R0,
	Index{ symbols.NT_IncDec_Expression,1,1 }: IncDec_Expression1R1,
	Index{ symbols.NT_IncDec_Expression,1,2 }: IncDec_Expression1R2,
	Index{ symbols.NT_Literal_Expression,0,0 }: Literal_Expression0R0,
	Index{ symbols.NT_Literal_Expression,0,1 }: Literal_Expression0R1,
	Index{ symbols.NT_Literal_Expression,1,0 }: Literal_Expression1R0,
	Index{ symbols.NT_Literal_Expression,1,1 }: Literal_Expression1R1,
	Index{ symbols.NT_Literal_Expression,2,0 }: Literal_Expression2R0,
	Index{ symbols.NT_Literal_Expression,2,1 }: Literal_Expression2R1,
	Index{ symbols.NT_Simple_Statement,0,0 }: Simple_Statement0R0,
	Index{ symbols.NT_Simple_Statement,0,1 }: Simple_Statement0R1,
	Index{ symbols.NT_Simple_Statement,1,0 }: Simple_Statement1R0,
	Index{ symbols.NT_Simple_Statement,1,1 }: Simple_Statement1R1,
	Index{ symbols.NT_Simple_Statement,2,0 }: Simple_Statement2R0,
	Index{ symbols.NT_Simple_Statement,2,1 }: Simple_Statement2R1,
	Index{ symbols.NT_Simple_Statement,3,0 }: Simple_Statement3R0,
	Index{ symbols.NT_Simple_Statement,3,1 }: Simple_Statement3R1,
	Index{ symbols.NT_Statement,0,0 }: Statement0R0,
	Index{ symbols.NT_Statement,0,1 }: Statement0R1,
	Index{ symbols.NT_Statement,1,0 }: Statement1R0,
	Index{ symbols.NT_Statement,1,1 }: Statement1R1,
	Index{ symbols.NT_Statement,2,0 }: Statement2R0,
	Index{ symbols.NT_Statement,2,1 }: Statement2R1,
	Index{ symbols.NT_Statement,3,0 }: Statement3R0,
	Index{ symbols.NT_Statement,3,1 }: Statement3R1,
	Index{ symbols.NT_Statement,4,0 }: Statement4R0,
	Index{ symbols.NT_Statement,4,1 }: Statement4R1,
	Index{ symbols.NT_Statement_List,0,0 }: Statement_List0R0,
	Index{ symbols.NT_Statement_List,0,1 }: Statement_List0R1,
	Index{ symbols.NT_Statement_List,1,0 }: Statement_List1R0,
	Index{ symbols.NT_Statement_List,1,1 }: Statement_List1R1,
	Index{ symbols.NT_Statement_List,1,2 }: Statement_List1R2,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_GoGLL:[]Label{ GoGLL0R0 },
	symbols.NT_Assignment_Statement:[]Label{ Assignment_Statement0R0 },
	symbols.NT_Declaration_Statement:[]Label{ Declaration_Statement0R0,Declaration_Statement1R0 },
	symbols.NT_Expression_Statement:[]Label{ Expression_Statement0R0,Expression_Statement1R0 },
	symbols.NT_Function_Argument_List:[]Label{ Function_Argument_List0R0,Function_Argument_List1R0 },
	symbols.NT_Function_Argument_List_Body:[]Label{ Function_Argument_List_Body0R0,Function_Argument_List_Body1R0,Function_Argument_List_Body2R0 },
	symbols.NT_Function_Argument:[]Label{ Function_Argument0R0 },
	symbols.NT_Function_Statement:[]Label{ Function_Statement0R0,Function_Statement1R0 },
	symbols.NT_Empty_Statement:[]Label{ Empty_Statement0R0 },
	symbols.NT_Simple_Statement:[]Label{ Simple_Statement0R0,Simple_Statement1R0,Simple_Statement2R0,Simple_Statement3R0 },
	symbols.NT_IncDec_Expression:[]Label{ IncDec_Expression0R0,IncDec_Expression1R0 },
	symbols.NT_Literal_Expression:[]Label{ Literal_Expression0R0,Literal_Expression1R0,Literal_Expression2R0 },
	symbols.NT_Expression:[]Label{ Expression0R0,Expression1R0 },
	symbols.NT_For:[]Label{ For0R0,For1R0 },
	symbols.NT_Statement_List:[]Label{ Statement_List0R0,Statement_List1R0 },
	symbols.NT_Compound_Statement:[]Label{ Compound_Statement0R0,Compound_Statement1R0 },
	symbols.NT_Statement:[]Label{ Statement0R0,Statement1R0,Statement2R0,Statement3R0,Statement4R0 },
}

