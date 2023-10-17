# performance tools

uptime - CPU usage over time, 1m 5m 15m
13:29  up 75 days,  5:59, 2 users, load averages: 2.53 2.70 2.91

top
Processes: 410 total, 3 running, 407 sleeping, 2515 threads                                                                                                        13:31:13
Load Avg: 2.70, 2.76, 2.90  CPU usage: 7.48% user, 10.10% sys, 82.41% idle  SharedLibs: 534M resident, 102M data, 24M linkedit.
MemRegions: 377318 total, 2165M resident, 87M private, 5322M shared. PhysMem: 15G used (1532M wired), 86M unused.
VM: 185T vsize, 3831M framework vsize, 1022048(0) swapins, 1520209(0) swapouts. Networks: packets: 88293278/96G in, 28435665/10G out.
Disks: 180791313/2719G read, 119287278/2115G written.

ps


iostat
              disk0       cpu    load average
    KB/t  tps  MB/s  us sy id   1m   5m   15m
   16.89   46  0.76   2  7 91  3.00 3.06 3.00

strace
tcpdump

for i in $(docker container ls --format "{{.ID}}"); do docker inspect -f '{{.State.Pid}} {{.Name}}' $i; done

vmstat
iostat
top -o %MEM
ps
lsof

ps aux | grep 14573
ps aux

find . -name "test*"
lsof ~/work/test.swp
sudo lsof returns more results
sudo lsof +D

