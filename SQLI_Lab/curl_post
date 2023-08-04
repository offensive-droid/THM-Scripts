#Task 3 Introduction to SQL Injection: Part 2 - SQL Injection 4: Update
#!/bin/bash

echo "
 _____            _  ______         _
/  __ \          | | | ___ \       |
| /  \/_   _ _ __| | | |_/ /__  ___| |_ ___ _ __
| |   | | | | '__| | |  __/ _ \/ __| __/ _ \ '__|
| \__/\ |_| | |  | | | | | (_) \__ \ ||  __/ |
 \____/\__,_|_|  |_| \_|  \___/|___/\__\___|_|

 | Made By Droid | Github: https://github.com/offensive-droid |
"

read -p "Full URL: " ip
read -p "How many data to post: " data
read -p "Cookies file: " cookie

arrayVal=()

for i in $(seq 1 $data); do
    read -p "Value Name $i: " name
    read -p "Value for $name: " value
    arrayVal[${#arrayVal[@]}]="$name=$value"
done

# Combine all array values into a single string
postData=$(IFS="&"; echo "${arrayVal[*]}")


# Perform the curl POST request including the cookie and following redirects
curl  -X GET "$ip" -H "Cookie: session=$(cat $cookie)" -d "$postData" -L
