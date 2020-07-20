import random
import timeit









def do_work():

        unsorted_list = []
        list_size = 10000


        for i in range(1,int(list_size)):
                unsorted_list.append(random.randint(0, 100))
        print(f'------------------------> unsorted {unsorted_list}')

        start_time = timeit.default_timer()
        odd_even_sort(unsorted_list, list_size)
        end_time = timeit.default_timer()


        print(f'------------------------> sorted_list {unsorted_list}')
        print(f'------------------------> Execution list {end_time - start_time} seconds')      






def odd_even_sort(ulist, usize):

	sorted = False

	while not sorted:

		sorted = True

		for i in range(1,usize-2,2):
			if ulist[i] > ulist[i+1]:
				ulist[i], ulist[i+1] = ulist[i+1], ulist[i]
				sorted = False

		for i  in range(0,usize-2, 2):
			if ulist[i] > ulist[i+1]:
				ulist[i], ulist[i+1] = ulist[i+1], ulist[i]
				sorted = False







if __name__ == '__main__':

	do_work()
