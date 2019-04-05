grammar ProtonLang;

document
    :   includeStatement* 
        targetStatement? 
        namespaceStatement? 
        contextStatement* 
        aliasStatement* 
        componentStatement*
        entityIndexStatement*
    ;

includeStatement
    : INCLUDE StringLiteral SEMICOLON
    ;

targetStatement
    : TARGET identifier SEMICOLON
    ;

namespaceStatement
    : NAMESPACE namespace SEMICOLON
    ;

contextStatement
    : CONTEXT context (COMMA context)* SEMICOLON
    ;

context
    : identifier (COLON identifier)? (LPAREN keyValueList RPAREN)?
    ;

aliasStatement
    : ALIAS alias (COMMA alias)* SEMICOLON
    ;

alias
    : identifier EQUALS StringLiteral
    ;

componentStatement
    : COMPONENT component (COMMA component)* inStatement? (SEMICOLON || asStatement SEMICOLON || memberStatementList)
    ;

asStatement
    : AS (identifier (LESS_THAN keyValueList GREATER_THAN)? || StringLiteral)
    ;

inStatement
    : IN keyValueList
    ;

component
    :  identifier (LPAREN keyValueList RPAREN)?
    ;

entityIndexStatement
    : INDEX entityIndex
    ;

entityIndex
    :  identifier (LPAREN keyValueList RPAREN)? (IN identifier)? entityIndexMethodStatementList
    ;

entityIndexMethodStatementList
    : LCURLEY entityIndexMethodStatement+ RCURLEY
    ;

entityIndexMethodStatement
    : METHOD identifier memberStatementList
    ;

memberStatementList
    : LCURLEY memberStatement* RCURLEY
    ;

memberStatement
    : propertyList AS (identifier (LESS_THAN keyValueList GREATER_THAN)? | StringLiteral) SEMICOLON
    ;

propertyList
    : property (COMMA property)*
    ;
    
property
    : identifier (LPAREN keyValueList RPAREN)?
    ;

keyValueList
    : keyValue (COMMA keyValue)*
    ;

keyValue
    : identifier (COLON identifier)?
    ;

namespace
    : WORD+ (PERIOD WORD)*
    ;

identifier
   : UNDERSCORE? (UNDERSCORE | WORD | NUMBER)+
   ;

StringLiteral
	:	DOUBLE_QUOTE StringCharacters? DOUBLE_QUOTE
	;

fragment
StringCharacters
	:	StringCharacter+
	;

fragment
StringCharacter
	:	~["\\\r\n]
    ;

INCLUDE
    : '#include'
    ;
    
TARGET
    : 'target'
    ;

NAMESPACE
    : 'namespace'
    ;

CONTEXT
    : 'context'
    ;

ALIAS
    : 'alias'
    ;

INDEX
    : 'index'
    ;

METHOD
    : 'method'
    ;

IN 
    : 'in'
    ;

AS
    : 'as'
    ;

COMPONENT
    : 'component'
    ;

EQUALS
    : '='
    ;

SEMICOLON
    : ';'
    ;

UNDERSCORE
    : '_'
    ;

COLON 
    : ':'
    ;

DOUBLE_QUOTE
    : '"'
    ;

LCURLEY
    : '{'
    ;

RCURLEY
    : '}'
    ;

PERIOD
    : '.'
    ;

LESS_THAN
    : '<'
    ;

GREATER_THAN
    : '>'
    ;

LPAREN
    : '('
    ;

RPAREN
    : ')'
    ;

COMMA
    : ','
    ;

WORD
   : ('a' .. 'z' | 'A' .. 'Z')+ 
   ;

NUMBER
   : ('0' .. '9')+
   ;

COMMENT
    : '//' ~[\r\n]* -> skip;

WS
   : [ \r\n] -> skip
   ;