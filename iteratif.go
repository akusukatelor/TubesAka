import random
import time
import matplotlib.pyplot as plt
import numpy as np

# Recursive Binary Search
def recursive_binary_search(arr, low, high, target):
    if low > high:
        return -1
    mid = (low + high) // 2
    if arr[mid] == target:
        return mid
    elif arr[mid] < target:
        return recursive_binary_search(arr, mid + 1, high, target)
    else:
        return recursive_binary_search(arr, low, mid - 1, target)

# Iterative Binary Search
def iterative_binary_search(arr, target):
    low, high = 0, len(arr) - 1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    return -1

# Generate Random Array
def generate_array(size, sorted_array):
    array = [random.randint(0, 1000) for _ in range(size)]
    if sorted_array:
        array.sort()
    return array

# Measure Execution Time for Recursive
def measure_recursive_time(array, target, sorted_array):
    if not sorted_array:
        array.sort()
    start = time.time()
    recursive_binary_search(array, 0, len(array) - 1, target)
    return time.time() - start

# Measure Execution Time for Iterative
def measure_iterative_time(array, target, sorted_array):
    if not sorted_array:
        array.sort()
    start = time.time()
    iterative_binary_search(array, target)
    return time.time() - start

# Plot Graph
def plot_graph(sizes, recursive_sorted, recursive_unsorted, iterative_sorted, iterative_unsorted):
    sizes = np.array(sizes)
    plt.figure(figsize=(10, 6))

    # Recursive Sorted
    plt.plot(sizes, recursive_sorted, label="Recursive Sorted", marker="o")
    # Recursive Unsorted
    plt.plot(sizes, recursive_unsorted, label="Recursive Unsorted", marker="o", linestyle="--")
    # Iterative Sorted
    plt.plot(sizes, iterative_sorted, label="Iterative Sorted", marker="s")
    # Iterative Unsorted
    plt.plot(sizes, iterative_unsorted, label="Iterative Unsorted", marker="s", linestyle="--")

    plt.xscale("log")
    plt.yscale("log")
    plt.title("Performance Comparison: Recursive vs Iterative (Sorted and Unsorted)")
    plt.xlabel("Size")
    plt.ylabel("Time (s)")
    plt.legend()
    plt.grid(True, which="both", linestyle="--", linewidth=0.5)
    plt.savefig("comparison.png")
    plt.show()

# Main
if _name_ == "_main_":
    sizes = [100, 500, 1000, 5000, 10000]
    print(f"{'Size':<10} {'Recursive Sorted':<20} {'Recursive Unsorted':<20} {'Iterative Sorted':<20} {'Iterative Unsorted':<20}")

    recursive_sorted = []
    recursive_unsorted = []
    iterative_sorted = []
    iterative_unsorted = []

    for size in sizes:
        array = generate_array(size, False)
        target = array[random.randint(0, len(array) - 1)]

        # Measure Recursive Times
        rs_time = measure_recursive_time(generate_array(size, True), target, True)
        ru_time = measure_recursive_time(array, target, False)
        recursive_sorted.append(rs_time)
        recursive_unsorted.append(ru_time)

        # Measure Iterative Times
        is_time = measure_iterative_time(generate_array(size, True), target, True)
        iu_time = measure_iterative_time(array, target, False)
        iterative_sorted.append(is_time)
        iterative_unsorted.append(iu_time)

        # Display Results
        print(f"{size:<10} {rs_time:<20.10f} {ru_time:<20.10f} {is_time:<20.10f} {iu_time:<20.10f}")

    # Generate Graph
    plot_graph(sizes, recursive_sorted, recursive_unsorted, iterative_sorted, iterative_unsorted)
