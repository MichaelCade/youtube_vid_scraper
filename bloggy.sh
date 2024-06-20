#!/bin/bash

# Define the path to the JSON file
json_file="playlist_videos.json"

# Read and parse the JSON file using jq
videos=$(jq -c '.[]' $json_file)

# Iterate over each video
echo "$videos" | while read -r video; do
  title=$(echo "$video" | jq -r '.title')
  url=$(echo "$video" | jq -r '.url')
  
  # Extract day number from the title
  day=$(echo "$title" | sed -n 's/.*Day \([0-9]*\).*/\1/p')  
  
  # Construct the markdown file name
  md_file="day${day}.md"
  
  # Insert video title, thumbnail, and clickable URL at the top of the markdown file
  echo "# $title" > "$md_file"
  echo "[![Watch the video](thumbnails/day${day}.png)]($url)" >> "$md_file"
  echo "" >> "$md_file" # Adds an extra newline for spacing

  # Execute the command for each video (Placeholder commands)
  yt --transcript "$url" | fabric --pattern create_summary --model mistral:instruct >> "$md_file"
  yt --transcript "$url" | fabric --pattern create_summary --model llama3:latest >> "$md_file"
  
  echo "Processed $md_file"
done

echo "All videos processed."