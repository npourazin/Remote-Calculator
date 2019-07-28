# Remote-Calculator

A simple server-client calculator in Golang!

How it works:

client sents a string of serveral numbers with + or - characters and without any white spaces in the middle.

For example: 3+234-232+2676-44

(UPDATE: can take several space sepated strings and calculate each of them

Example: 23+5 345-1+3

Result:28

Result:347
)

Then, server sends back the result of the calculation.
