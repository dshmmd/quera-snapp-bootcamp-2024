#!/bin/bash

# Check if an argument was provided
if [ $# -eq 0 ]; then
    echo "No .go file provided. Usage: $0 <path-to-go-file>"
    exit 1
fi

# Get the path of the .go file
input_file="$1"

# Extract the filename from the path
filename=$(basename "$input_file")

# Define the destination path in /tmp
destination="./tmp/submission-${filename}"

# Copy the original .go file to /tmp with the new name
cp "$input_file" "$destination"

# Change the package name to "main"
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/^package .*/package main/' "$destination"
else
    sed -i 's/^package .*/package main/' "$destination"
fi

# Append the specified code to the end of the file
cat <<EOF >> "$destination"
func main() {
	ans, err := Solve(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Print(ans)
}
EOF

echo "File copied and modified at: $destination"
