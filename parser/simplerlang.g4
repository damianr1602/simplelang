grammar simplerlang;

program: ( stat? NEWLINE )* 
    ;

stat:	SHOW ID		#show
	| ID '=' expr0	#let
	| READINT ID   	#readint
	| READDOUBLE ID   	#readdouble
   ;
   
expr0:  expr1			#single0
      | expr1 ADD expr1		#add 
      | expr1 SUB expr1		#sub 
;

expr1:  expr2			#single1
      | expr2 MUL expr2	        #mul
      | expr2 DIV expr2	        #div

;

expr2:   INT			#int
       | REAL			#real
       | TOINT expr2		#toint
       | TOREAL expr2		#toreal
       | '(' expr0 ')'		#par
;	

SHOW:	'show' 
   ;

READINT:	'read_int' 
   ;
   
READDOUBLE:	'read_double' 
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
