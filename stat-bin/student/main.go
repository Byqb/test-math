package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func readData(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func calculateAverage(data []float64) int {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return int(math.Round(sum / float64(len(data))))
}

func calculateMedian(data []float64) int {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return int(math.Round((data[n/2-1] + data[n/2]) / 2))
	}
	return int(math.Round(data[n/2]))
}

func calculateVariance(data []float64, mean float64) int {
	sum := 0.0
	for _, value := range data {
		sum += (value - mean) * (value - mean)
	}
	return int(math.Round(sum / float64(len(data))))
}

func calculateStandardDeviation(variance float64) int {
	return int(math.Round(math.Sqrt(variance)))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}

	filename := os.Args[1]

	data, err := readData(filename)
	if err != nil {
		log.Fatal("Error reading data:", err)
	}

	average := calculateAverage(data)
	median := calculateMedian(data)
	mean := float64(average)
	variance := calculateVariance(data, mean)
	standardDeviation := calculateStandardDeviation(float64(variance))

	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", standardDeviation)
}

/* package main

import (
	"bufio" // Package for buffered I/O operations
	"fmt"
	"math"
	"os"      // Package for operating system functionality
	"sort"    // Package for sorting functionality
	"strconv" // Package for string conversions
)

// stats function calculates the average, median, variance, and standard deviation of a slice of float64 values

func stats(data []float64) (avg, med, variance, std float64) {
	sum := 0.0
	// Iterate over the slice and calculate the sum
	for _, v := range data {
		sum += v
	}
	// Calculate the average
	avg = sum / float64(len(data))

	// Sort the slice in ascending order
	sort.Float64s(data)
	n := len(data)
	// If the length is even
	if n%2 == 0 {
		// Calculate the median as the average of the two middle values
		med = (data[n/2-1] + data[n/2]) / 2
	} else {
		// If the length is odd, the median is the middle value
		med = data[n/2]
	}

	variance = 0.0

	// Iterate over the slice to calculate the variance
	for _, v := range data {
		// Square the difference between each value and the average
		variance += math.Pow(v-avg, 2)
	}
	// Divide the sum of squared differences by the length to get the variance
	variance /= float64(n)
	// Calculate the standard deviation as the square root of the variance
	std = math.Sqrt(variance)

	return // Return the calculated values
}

func main() {
	// Check if a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run program.go <data_file>") // Print usage instructions
		return
	}

	// Open the file specified in the command-line argument
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function returns

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Slice to store the data values
	var data []float64
	// Iterate over each line in the file
	for scanner.Scan() {
		// Convert the line to a float64 value
		v, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			// Print the error if the conversion fails
			fmt.Println(err)
			// Skip to the next line
			continue
		}
		// Append the value to the data slice
		data = append(data, v)
	}

	avg, med, variance, std := stats(data)                                                                         // Call the stats function with the data slice
	fmt.Printf("Average: %.0f\nMedian: %.0f\nVariance: %.0f\nStandard Deviation: %.0f\n", avg, med, variance, std) // Print the calculated statistics
}

*/
