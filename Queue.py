class Queue:

	def __init__(self):
		self.elements = []

	def enqueue(self, data):
		self.elements.append(data)
		return data

	def dequeue(self):
		return self.elements.pop(0)

	def rear(self):
		return self.elements[-1]

	def front(self):
		return self.elements[0]

	def is_empty(self):
		return len(self.elements) == 0

if __name__ == '__main__':
	queue = Queue()

	# checking is_empty method
	print(queue.is_empty())

	# adding the elements
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.enqueue(4)
	queue.enqueue(5)

	# again checking is_empty method
	print(queue.is_empty())

	# printing the front and rear elements
	print(queue.front(), end=' ')
	print(queue.rear())

	# removing the element -> 1
	queue.dequeue()

	# checking the front and rear elements
	print(queue.front(), end=' ')
	print(queue.rear())

	# removing all the elements
	queue.dequeue()
	queue.dequeue()
	queue.dequeue()
	queue.dequeue()

	# checking the is_empty
	print(queue.is_empty())