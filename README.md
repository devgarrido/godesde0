# GO Lenguaje

## Tipo de datos.

`uint`

* uint8     unsigned    8-bit   integers (0 to 255)
* uint16    unsigned    16-bit  integers (0 to 65535)
* uint32    unsigned    32-bit  integers (0 to 4294967295)

`int`

* int8      signed       8-bit  integers (-128 to 127)
* int16     signed      16-bit  integers (-32768 to 32767)
* int32     signed      32-bit  integers (-2147483648 to 2147483647)
* int64     signed      64-bit  integers (-9223372036854775808 to 9223372036854775807)

`byte` // alias for uint8.

`rune` // alias for int32
       // represents a Unicode code point.

`float32`
`float62` 


## Verbs 

`General`

* %v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
* %#v	a Go-syntax representation of the value
* %T	a Go-syntax representation of the type of the value
* %%	a literal percent sign; consumes no value

`Boolean`

* %t	the word true or false

`Integers`

* %b	base 2
* %c	the character represented by the corresponding Unicode code point
* %d	base 10
* %o	base 8
* %O	base 8 with 0o prefix
* %q	a single-quoted character literal safely escaped with Go syntax.
* %x	base 16, with lower-case letters for a-f
* %X	base 16, with upper-case letters for A-F
* %U	Unicode format: U+1234; same as "U+%04X"

`Floating-point and complex constituents:`

* %b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
* %e	scientific notation, e.g. -1.234456e+78
* %E	scientific notation, e.g. -1.234456E+78
* %f	decimal point but no exponent, e.g. 123.456
* %F	synonym for %f
* %g	%e for large exponents, %f otherwise. Precision is discussed below.
* %G	%E for large exponents, %F otherwise
* %x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
* %X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
 
`String and slice of bytes (treated equivalently with these verbs):` 

* %s	the uninterpreted bytes of the string or slice
* %q	a double-quoted string safely escaped with Go syntax
* %x	base 16, lower-case, two characters per byte
* %X	base 16, upper-case, two characters per byte

`Slice:`

* %p	address of 0th element in base 16 notation, with leading 0x

`Pointer:` 

* %p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.
The default format for %v is:

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p

For compound objects, the elements are printed using these rules, recursively, laid out like this:

struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]

Width is specified by an optional decimal number immediately preceding the verb. If absent, the width is whatever is necessary to represent the value. Precision is specified after the (optional) width by a period followed by a decimal number. If no period is present, a default precision is used. A period with no following number specifies a precision of zero. Examples:

%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
Width and precision are measured in units of Unicode code points, that is, runes. (This differs from C's printf where the units are always measured in bytes.) Either or both of the flags may be replaced with the character '*', causing their values to be obtained from the next operand (preceding the one to format), which must be of type int.

For most values, width is the minimum number of runes to output, padding the formatted form with spaces if necessary.

For strings, byte slices and byte arrays, however, precision limits the length of the input to be formatted (not the size of the output), truncating if necessary. Normally it is measured in runes, but for these types when formatted with the %x or %X format it is measured in bytes.

For floating-point values, width sets the minimum width of the field and precision sets the number of places after the decimal, if appropriate, except that for %g/%G precision sets the maximum number of significant digits (trailing zeros are removed). For example, given 12.345 the format %6.3f prints 12.345 while %.3g prints 12.3. The default precision for %e, %f and %#g is 6; for %g it is the smallest number of digits necessary to identify the value uniquely.

For complex numbers, the width and precision apply to the two components independently and the result is parenthesized, so %f applied to 1.2+3.4i produces (1.200000+3.400000i).

When formatting a single integer code point or a rune string (type []rune) with %q, invalid Unicode code points are changed to the Unicode replacement character, U+FFFD, as in strconv.QuoteRune.

Other flags:

'+'	always print a sign for numeric values;
	guarantee ASCII-only output for %q (%+q)
'-'	pad with spaces on the right rather than the left (left-justify the field)
'#'	alternate format: add leading 0b for binary (%#b), 0 for octal (%#o),
	0x or 0X for hex (%#x or %#X); suppress 0x for %p (%#p);
	for %q, print a raw (backquoted) string if strconv.CanBackquote
	returns true;
	always print a decimal point for %e, %E, %f, %F, %g and %G;
	do not remove trailing zeros for %g and %G;
	write e.g. U+0078 'x' if the character is printable for %U (%#U).
' '	(space) leave a space for elided sign in numbers (% d);
	put spaces between bytes printing strings or slices in hex (% x, % X)
'0'	pad with leading zeros rather than spaces;
	for numbers, this moves the padding after the sign;
	ignored for strings, byte slices and byte arrays