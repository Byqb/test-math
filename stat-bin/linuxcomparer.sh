#!/bin/bash

# Execute the binary and store its output
binary_output=$(bin/math-skills)

# Execute the Go program and store its output
go_output=$(go run ../../main.go data.txt)

# if you would to see the output of the binary and Go program, uncomment the following lines


# echo -e "*******************************EXPECTED OUTPUT*******************************\n$binary_output"
# echo -e "*******************************ACTUAL OUTPUT*******************************\n$go_output"

# Extract the average, median, variance, and standard deviation from the binary output
binary_average=$(echo "$binary_output" | grep -oP 'Average: \K\d+')
binary_median=$(echo "$binary_output" | grep -oP 'Median: \K\d+')
binary_variance=$(echo "$binary_output" | grep -oP 'Variance: \K\d+')
binary_std_dev=$(echo "$binary_output" | grep -oP 'Standard Deviation: \K\d+')

# Extract the average, median, variance, and standard deviation from the Go output
go_average=$(echo "$go_output" | grep -oP 'Average: \K\d+')
go_median=$(echo "$go_output" | grep -oP 'Median: \K\d+')
go_variance=$(echo "$go_output" | grep -oP 'Variance: \K\d+')
go_std_dev=$(echo "$go_output" | grep -oP 'Standard Deviation: \K\d+')


# initialize a counter and if it is 4 then we passed
counter=0


# Compare the average
if [[ "$binary_average" == "$go_average" ]]; then
    echo -e "\033[38;2;0;255;0mAverage PASSED \033[0m"
    counter=$((counter+1))
else
    echo -e "\033[38;2;255;0;0mAverage FAILED \033[0m => Expected Average: $binary_average but got $go_average instead"
fi

# Compare the median
if [[ "$binary_median" == "$go_median" ]]; then
    echo -e "\033[38;2;0;255;0mMedian PASSED \033[0m"
    counter=$((counter+1))
else
    echo -e "\033[38;2;255;0;0mMedian FAILED \033[0m => Expected Median: $binary_median but got $go_median instead"
fi

# Compare the variance
if [[ "$binary_variance" == "$go_variance" ]]; then
    echo -e "\033[38;2;0;255;0mVariance PASSED \033[0m"
    counter=$((counter+1))
else
    echo -e "\033[38;2;255;0;0mVariance FAILED \033[0m => Expected Variance: $binary_variance but got $go_variance instead"
fi

# Compare the standard deviation
if [[ "$binary_std_dev" == "$go_std_dev" ]]; then
    echo -e "\033[38;2;0;255;0mStandard Deviation PASSED \033[0m"
    counter=$((counter+1))
else
    echo -e "\033[38;2;255;0;0mStandard Deviation FAILED \033[0m => Expected Standard Deviation: $binary_std_dev but got $go_std_dev instead"
fi

# Get the terminal width
term_width=$(tput cols)


# Create a string of spaces for padding
padding=$(printf '%*s' $(( (term_width - ${#message}) / 2 )))

# If all the tests passed, print a success message
if [[ $counter -eq 4 ]]; then
    message="\033[1m\033[38;2;0;255;0mAll tests PASSED \033[0m"
    echo -e "$padding$message"
else
    message="\033[38;2;255;0;0mSome tests FAILED \033[0m"
    echo -e "$padding$message"
fi