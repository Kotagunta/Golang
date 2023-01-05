List = [1,2,3,2,3,2,4,4]

Duplicate_numbers = []

for num in List:

	n = List.count(num)

	if n > 1:

		if Duplicate_numbers.count(num) == 0:

			Duplicate_numbers.append(num)

print(Duplicate_numbers)