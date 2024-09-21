1. Need to show an error message for the line, when there are uneven number of quotation marks
2. Need to write a handler when firstByteIsQuote is open but lineIsTerminated, so that handles the cases like:
"asd", "has
kell", bmw
3. Need to write a converter to one quotation mark if there are even number of consecutive quotation marks