import random
import timeit




def do_work():

	unsorted_list  = []

	for i in range(1,10000):
		unsorted_list.append(random.randint(0, 100))

	print(f'\n Unsorted list ------> {unsorted_list}')

	get_sorted(unsorted_list)


def get_sorted(uns_list):

	start_time = timeit.default_timer()

	len_uns = len(uns_list)

	for i in range(0,len_uns-1):

		j = i + 1

		if uns_list[j] < uns_list[i]:

			cp_item = uns_list[j]

			while (j> 0 and cp_item < uns_list[j-1]):

				#print(n)
				uns_list[j] = uns_list[j-1]
				j= j -1

			uns_list[j] = cp_item

	end_time = timeit.default_timer()

	print(f'\n Sorted List ------->{uns_list}')
	print(f'\n Execution Time is ----> {end_time - start_time} seconds')


if __name__ == '__main__':
	do_work()
