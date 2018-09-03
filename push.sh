#!/bin/bash
git add .
comment=""
while [ "$comment" = "" ]; do
    echo -n "Введи комментарий для пуша: "
    read comment
done
echo "${comment}"
git commit -m "${comment}"
git push