# awk command example

```shell
awk '{print $1}' temp.txt 
awk '{print $1,$2}' temp.txt
for i in $(awk '{print}' temp.txt); do echo $i; done
for i in $(awk '{print $1}' temp.txt); do echo $i; done
for i in $(awk '{print $1}' urls.txt); do curl -L $i; done
```
