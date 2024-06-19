#!/bin/bash

# Step 1: Define the path to the JSON file
json_file="playlist_videos.json"

# Step 2: Read and parse the JSON file
# jq is a tool for processing JSON inputs
videos=$(jq -c '.[]' $json_file)

# Step 3: Iterate over each video
echo "$videos" | while read -r video; do
  title=$(echo "$video" | jq -r '.title')
  url=$(echo "$video" | jq -r '.url')
  
  # Extract day number from the title
  day=$(echo "$title" | sed -n 's/.*Day \([0-9]*\).*/\1/p')  
  
  # Construct the markdown file name
  md_file="day${day}.md"
  
  # Step 4: Execute the command for each video
  # Note: This is a placeholder command. You'll need to replace 'yt' and 'fabric' with the actual commands or scripts
  # that perform the operations you described.
  yt --transcript "$url" | fabric --pattern create_summary --model mistral:instruct >> "$md_file"
  
  echo "Processed $md_file"
done

echo "All videos processed."