from random import randint

player = input("Player: ")

rand_num = randint(0,2)

if rand_num == 0:
	computer = "rock"
elif rand_num == 1:
	computer = "paper"
if rand_num == 2:
	computer = "scissors"

print(f"compuetr plays:{computer}")


if player == computer:
	print("Its a Tie")

elif player == "rock":
	if computer == "scissors":
		print("Player Wins!")
	else:
		print("computer wins!")

elif player == "paper":
	if computer == "rock":
		print("Player Wins!")
	else:
		print("computer Wins!")

elif player == "scissors":
	if computer == "paper":
		print("Player Wins!")
	else:
		print("computer Wins!")

else:
	print("Something went wrong")