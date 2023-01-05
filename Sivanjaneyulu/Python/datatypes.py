List = [1, "a", "b", 2, "z", "2", -1, 3+6j, 0.2]

result = []

get_integers = []
get_strings = []
get_floats = []

get_integers = [int(i) for i in List.split() if i.isdigit()]


print(get_integers)