grammar simplelang;

prog: block;

block: ( stat? NEWLINE )* 
;

stat:	SHOW ID		#show
	| ID ASSIGN expr0	#let
	| READINT ID   	#readint
    | ID '['(INT)?']' (ASSIGN '{' array_items '}')?          #intarray
	| READDOUBLE ID   	#readdouble
	| SHOWARRAYELEM ID '['(INT)']'   	#showArrayElem
	| IF equal THEN block ENDIF	#if
	| REPEAT repetitions block ENDREPEAT	#repeat
   ;

expr0:  expr1			    #single0
      | expr1 ADD expr1		#add 
      | expr1 SUB expr1		#sub 
;

expr1:  expr2			        #single1
      | expr2 MUL expr2	        #mul
      | expr2 DIV expr2	        #div

;

expr2:   INT			    #int
       | ID             	#id
       | ID '['(INT)']' 	#assignArrayElem
       | REAL			    #real
       | STRING			    #string
       | TOINT expr2		#toint
       | TOREAL expr2		#toreal
       | '(' expr0 ')'		#par
;	


repetitions: expr2
;

REPEAT: 'repeat'
;

ENDREPEAT: 'endrepeat';



equal: ID '==' INT
    ;

array_items: INT (',' INT )*
    ;
    
IF: 'if'
;

THEN: 'then'
;

ENDIF: 'endif'
;

SHOW:	'show' 
    ;

SHOWARRAYELEM:	'show_array_elem' 
    ;

READINT:	'read_int' 
    ;
   
READDOUBLE:	'read_double' 
    ;
   
READARRAY:	'read_array' 
    ;
   
ID:   ('a'..'z'|'A'..'Z')+
    ;

STRING: '"' ( ESC | ~[\\"\r\n] )* '"'
    ;

fragment ESC : '\\"' | '\\\\' 
    ;

TOINT: '(int)'
    ;

TOREAL: '(real)'
    ;

INT:   '0'..'9'+
    ;
    
REAL: '0'..'9'+'.''0'..'9'+
    ;

ASSIGN: '='
    ;

ADD: '+'
    ;

MUL: '*'
    ;
    
SUB: '-'
    ;
    

DIV: '/'
    ;
    
NEWLINE:	'\r'? '\n'
    ;

WS: (' '|'\t')+ -> skip  ;
