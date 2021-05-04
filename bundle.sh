mkdir -p res
for folder in lists/*/; do tar -czf "res/$(echo $folder | sed -e "s/^lists\///" -e "s/\/$//").tar.gz" "$folder"; done
