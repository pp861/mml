grammar Mml;

prog
    : ruleLine+
    ;

ruleLine
    : opType Attribute ':' mmlList ';'
    ;

opType
    : ACT
    | BLK
    | CHECK
    | CLB
    | DEL
    | DLD
    | DSP
    | HIS
    | HISTORY
    | LST
    | MOD
    | PING
    | POWEROP
    | QRY
    | RBK
    | RMV
    | RST
    | SCAN
    | SET
    | SHW
    | SND
    | STOP
    | SWITCH
    | SYNC
    | UNACT
    | UNBLK
    | UPLD
    | ADD
    ;

ACT
    : [aA][cC][tT]
    ;

BLK
    : [bB][lL][kK]
    ;

CHECK
    : [cC][hH][eE][cC][kK]
    ;
CLB
    : [cC][lL][bB]
    ;

DEL
    : [dD][eE][lL]
    ;

DLD
    : [dD][lL][dD]
    ;

DSP
    : [dD][sS][pP]
    ;

HIS
    : [hH][iI][sS]
    ;

HISTORY
    : [hH][iI][sS][tT][oO][rR][yY]
    ;

LST
    : [lL][sS][tT]
    ;

MOD
    : [mM][oO][dD]
    ;
PING
    : [pP][iI][nN][gG]
    ;

POWEROP
    : [pP][oO][wW][eE][rR][oO][pP]
    ;
QRY
    : [qQ][rR][yY]
    ;

RBK
    : [rR][bB][kK]
    ;
RMV
    : [rR][mM][vV]
    ;
RST
    : [rR][sS][tT]
    ;
SCAN
    : [sS][cC][aA][nN]
    ;
SET
    : [sS][eE][tT]
    ;
SHW
    : [sS][hH][wW]
    ;
SND
    : [sS][nN][dD]
    ;
STOP
    : [sS][tT][oO][pP]
    ;

SWITCH
    : [sS][wW][iI][cC][hH]
    ;
SYNC
    : [sS][yY][nN][cC]
    ;
UNACT
    : [uU][nN][aA][cC][tT]
    ;
UNBLK
    : [uU][nN][bB][lL][kK]
    ;
UPLD
    : [uU][pP][lL][dD]
    ;
ADD
    : [aA][dD][dD]
    ;

Attribute
    : Identifier
    ;

// Identifier
//     : Letter
//     | Letter Digit
//     | Letter (Letter | Digit | underScore)* (Letter | Digit)+
//     ;

Identifier
    : Letter (Letter | Digit  | '_')*
    ;

// fragment
// ID_LETTER
//     : 'a'..'z'|'A'..'Z'|'_'  // [a-zA-Z_]
//     ;

Letter
    : [a-z|A-Z]
    ;


// underScore
//     : '_'
//     ;

mmlList
    : condition (',' condition)*
    ;

condition
    : Attribute '=' Value
    ;

Value
    : NumberInt
    | Pattern
    ;

Digit
    : [0-9]
    ;

NumberInt
    : IntegerLiteral
    | Digit
    | Number
    | QuotedString
    ;

IntegerLiteral
    : [1-9] Digit*
    ;

Number
    : Digit+
    | Digit* '.' Digit*
    ;

QuotedString
    : '"' (ESC|.)*? '"'
    ;

ESC
    : '\\"' | '\\\\'
    ;

Pattern
    : '"' Identifier '"'
    ;

WS 
    : [ \t\r\n]+ -> skip
    ;