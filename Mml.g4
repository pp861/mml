grammar Mml;

prog
    : ruleLine+
    ;

ruleLine
    : opType attribute ':' mmlList ';'
    ;

opType
    : 'ACT'
    | 'BLK'
    | 'CHECK'
    | 'CLB'
    | 'DEL'
    | 'DLD'
    | 'DSP'
    | 'HIS'
    | 'HISTORY'
    | 'LST'
    | 'MOD'
    | 'PING'
    | 'POWEROP'
    | 'QRY'
    | 'RBK'
    | 'RMV'
    | 'RST'
    | 'SCAN'
    | 'SET'
    | 'SHW'
    | 'SND'
    | 'STOP'
    | 'SWITCH'
    | 'SYNC'
    | 'UNACT'
    | 'UNBLK'
    | 'UPLD'
    | 'ADD'
    ;

attribute
    : identifier
    ;

identifier
    : Letter
    | Letter Digit
    | Letter (Letter | Digit | underScore)* (Letter | Digit)+
    ;

Letter
    : [a-z|A-Z]
    ;


underScore
    : '_'
    ;

mmlList
    : condition (',' condition)*
    ;

condition
    : attribute '=' value
    ;

value
    : numberInt
    | pattern
    ;

Digit
    : [0-9]
    ;

numberInt
    : IntegerLiteral
    | Digit
    | number
    | quotedString
    ;

IntegerLiteral
    : [1-9] Digit*
    ;

number
    : Digit+
    | Digit* '.' Digit*
    ;

quotedString
    : '"' (ESC|.)*? '"'
    ;

ESC
    : '\\"' | '\\\\'
    ;

pattern
    : '"' identifier '"'
    ;

WS 
    : [ \t\r\n]+ -> skip
    ;