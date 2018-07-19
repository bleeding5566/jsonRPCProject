# jsonRPCProject
Parameter Args struct has two variables, X and Y <BR/>
Supporting methods and descriptions are listed below: <BR/>
Server has a map m to store value.

Create   : Set m[X] = Y. Key should not already exist in m. <BR/>
Delete   : Delete m[X] <BR/>
Set      : Set m[X] = Y. Key should already exist in m. <BR/>
Add      : Return m[X] + m[Y], both X and Y should already exist in m.<BR/>
Subtract : Return m[X] - m[Y], both X and Y should already exist in m. <BR/>
Multiply : Return m[X] * m[Y], both X and Y should already exist in m. <BR/>
Divide   : Return m[X] / m[Y], both X and Y should already exist in m. Resolution is 2^-16. <BR/>

Code Description: <BR/>
server.go : Open a json-rpc server <BR/>
arith.go : json-rpc method <BR/>
calculator.go : calculate addition/subtraction/mutiplication/division <BR/>
myMap.go : Define map stored in server <BR/>
