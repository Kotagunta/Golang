import random

random_number = random.randint(1,100)

guess_number = None

while guess_number != random_number:

    guess_number = int(input("Enete a number: "))

    if guess_number == random_number:
        print("You Won!")
    else:
        print("You Lose!")