package desafios

/*
Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored.
*/

func Parse(data string) []int {
	value := 0
	values := make([]int, 0)

	for i := 0; i < len(data); i++ {
		switch data[i] {
		case 'i':
			value++
		case 'd':
			value--
		case 's':
			value *= value
		case 'o':
			values = append(values, value)
		default:
			continue
		}
	}
	return values
}