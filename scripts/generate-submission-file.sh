#!/bin/bash

# Find all .go files matching the pattern
go_files=$(find ./pkg -type f -name "*.go" ! -name "*_test.go")

for input_file in $go_files; do
  # Extract the filename from the path
  filename=$(basename "$input_file")
  base_filename="${filename%.go}"

  # Define the destination path in /tmp
  mkdir "./tmp/${base_filename}"
  destination="./tmp/${base_filename}/submission-${filename}"

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

  # Change the package name to "main"
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' "s/Solve/Solve_${base_filename}/" "$destination"
  else
      sed -i "s/Solve/Solve_${base_filename}/" "$destination"
  fi

  echo "File copied and modified at: $destination"
done
