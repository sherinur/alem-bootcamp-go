#!/usr/bin/env bash
 
cat assets/star.txt
sleep 1
cat assets/welcome.txt
sleep 1

echo "Initializing the program..."
sleep 1

if [[ ! -d "rootkit" ]]; then
    echo "rootkit: No such file or directory" && exit 1
fi

cd "rootkit"

if [[ ! -d "modules" ]]; then
    echo "modules: No such file or directory" && exit 1
fi

cd "modules"
echo ""
echo "Loading modules..."
echo ""
sleep 1

files_total=0
for file in $(find . -name "*.ext1"); do
    wc -l "$file"
    ls -l "$file"
    cat "$file"

    files_total=$(( files_total + 1 ))
done

cd ..

if [[ ! -d "checksums" ]]; then
    echo "checksums: No such file or directory" && exit 1
fi

cd checksums
echo ""
echo "Checking modules..."
echo ""
sleep 1

for file in $(find . -name "*.md5"); do
    wc -l "$file"
    ls -l "$file"
    cat "$file"
    
    if md5sum -c "$file" ; then
        files_total=$(( files_total + 1 ))
    fi
done

cd ..

if [[ ! -d "configurations" ]]; then
    echo "configurations: No such file or directory" && exit 1
fi

cd configurations
if [[ -f main.conf ]]; then
    files_total=$(( files_total + 1 ))
fi

cd ../..

if [[ $files_total -eq 41 ]]; then
    bash ./run.sh
else
    echo ""
    echo "Program failed! Unable to initialize the modules..."
fi
