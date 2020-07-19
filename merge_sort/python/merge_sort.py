import random
import timeit






def do_work():


	unsorted_list = []
	list_size = 1000000


	for i in range(1,int(list_size)):
                unsorted_list.append(random.randint(0, 100))
	print(f'------------------------> unsorted {unsorted_list}')

	start_time = timeit.default_timer()
	merge_sort(unsorted_list)
	end_time = timeit.default_timer()

	
	print(f'------------------------> sorted_list {unsorted_list}')
	print(f'------------------------> Execution list {end_time - start_time}')






def merge_sort(ulist):

	if len(ulist) > 1:

		#print(ulist)


		mid = len(ulist) // 2
		left = ulist[:mid]
		right = ulist[mid:]

		merge_sort(left)
		merge_sort(right)

		i,j,k = 0,0,0
		

		while (i < len(left)) and (j < len(right)):

			if left[i] < right[j]:
				ulist[k] = left[i]
				i += 1
				k += 1

			else :
				ulist[k] = right[j]
				j += 1
				k += 1



		while i< len(left):

			ulist[k] = left[i]
			i += 1
			k += 1

		while j <len(right):

			ulist[k] = right[j]
			j += 1
			k += 1



	#print(f'Merging  {ulist}')




if __name__ == '__main__':

	do_work()
