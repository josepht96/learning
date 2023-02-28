# Stuff

## Create partition
parted > select /dev/sdc/ > mkpart
## Format partition
mkfs.ext4 (or whatever fs type) /dev/sdc2
## Mount
mkdir /var/lib/mount-target
mount /dev/sdc2 /var/lib/mount-target
## Umount
umount /var/lib/mount-target

## Get processes using system
fuser -v /var/lib/mount-target
## Kill process
kill <process id>