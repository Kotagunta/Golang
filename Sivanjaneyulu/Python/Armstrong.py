x = input("Enter a number: ")
qubes_sum = 0
for i in x:
	qubes_sum += int(i) ** len(x)


if int(x) == qubes_sum:
	print(x + " is an Armstrong number")
else:
	print(x + " is not an Armstrong number")