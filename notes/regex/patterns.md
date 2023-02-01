# regex

https://regexone.com/lesson/matching_characters?

```bash
# match digit
\d
# match any string
.
# match period
\.
# match chars in array
[c, m, f]
# match chars not in array
[^test]
# range
[0-9]
[a-z]
[A-Z]
# repetitions
a{3} # match 3 a's
[0-9]{3} #match 3 nums 0-9
a* # match any number of a's
a+ # match at least one a

# optional blocks
# match question mark
\?
# previous block is optional
a?
\d? files? found\?

# whitespace
\t # tab
\n # new line
\r # carriage return
   # space
\s # any whitespace
\d+\.\s+abc # 1.    abc # space and a tab, + is required to catch multiple whitespaces

# beginning and end
^start with this
end with this$

# group match - return only a portion of the matched string
^(test.).pdf # returns only test.

# nested group, return multiple groups
^(IMG(\d+))\.png$
^(\w+ (\d+))$

# or operator
^I love (cats|dogs)$

# opposites
# not...
\D
\W
\S
```
