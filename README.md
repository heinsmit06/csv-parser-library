DONE 1. Need to show an error message for the line, when there are uneven number of quotation marks

DONE 2. Need to write a handler when firstByteIsQuote is open but lineIsTerminated, so that handles the cases like:
"asd", "has
kell", bmw

DONE 3. Need to write a converter to one quotation mark if there are even number of consecutive quotation marks

DONE 4. Added handling of the cases when a field starts with a quotation mark and when it does not
