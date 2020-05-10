grammar simplerlang;

program: ( stat? NEWLINE )* 
    ;

stat:	SHOW ID		#show
	| ID ASSIGN expr0	#let
	| READINT ID   	#readint
    | ID '['(INT)?']' (ASSIGN '{' array_items '}')?          #intarray
	| READDOUBLE ID   	#readdouble
	| SHOWARRAYELEM ID '['(INT)']'   	#showArrayElem
   ;
   
expr0:  expr1			#single0
      | expr1 ADD expr1		#add 
      | expr1 SUB expr1		#sub 
;

expr1:  expr2			        #single1
      | expr2 MUL expr2	        #mul
      | expr2 DIV expr2	        #div

;
//ID '['(INT)']'
expr2:   INT			    #int
       | ID '['(INT)']' 	#assignArrayElem
       | REAL			    #real
       | TOINT expr2		#toint
       | TOREAL expr2		#toreal
       | '(' expr0 ')'		#par
;	

array_items: INT (',' INT )*
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

WS: [ \n\t]+ -> skip  ;
