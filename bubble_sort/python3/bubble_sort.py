import random
import timeit




def do_work():

	unsorted_list = []
	list_size = 100


	for i in range(1,int(list_size)):
	        unsorted_list.append(random.randint(0, 100))
	print(f'------------------------> unsorted {unsorted_list}')

	start_time = timeit.default_timer()
	bubble_sort(unsorted_list, list_size)
	end_time = timeit.default_timer()


	print(f'------------------------> sorted_list {unsorted_list}')
	print(f'------------------------> Execution list {end_time - start_time}')	



def bubble_sort(ulist, lsize):

	swapd = False

	for i in range(lsize-1):
		swapd = False
		for j in range(lsize-2):

			if ulist[j] > ulist[j+1]:
				ulist[j], ulist[j+1] = ulist[j+1], ulist[j]
				swapd = True



		if swapd == False:
			return


if __name__ == '__main__':

	do_work()
