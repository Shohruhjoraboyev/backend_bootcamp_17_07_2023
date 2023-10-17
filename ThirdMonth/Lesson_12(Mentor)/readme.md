| No. | Topic                                                                   |
| --- | ----------------------------------------------------------------------- |
| 1   | [**File and Directory Management**](#File and Directory Management)     |
| 2   | [**File Viewing and Editing**](#File Viewing and Editing)               |
| 3   | [**System Information:**](#System Information:)                         |
| 4   | [**Process Management:**](#Process Management:)                         |
| 5   | [**Networking:**](#Networking:)                                         |
| 6   | [**Package Management (package managers may vary):**](#Package Management (package managers may vary):)     |



## File and Directory Management:
**- `ls`:** List files and directories.
```bash
ls
ls -l
ls -a
```


- `cd`: Change directory.

```bash
cd /path/to/directory
cd ..
cd ~
```
- `pwd`: Print working directory.


- `mkdir`: Create a new directory.

```bash
mkdir new_directory
```
- `rmdir`: Remove an empty directory.

```bash

```
- `cp`: Copy files and directories.

```bash
cp file.txt /path/to/destination
cp -r directory /path/to/destination
```
- `mv`: Move or rename files and directories.

```bash
mv file.txt new_file.txt
mv directory /path/to/destination
```
- `rm`: Remove files and directories.

```bash
rm file.txt
rm -r directory
```
- `find`: Search for files and directories.

```bash
nano file.txt
```


Certainly! Let's continue with the sections you mentioned:

## File Viewing and Editing:
- `cat`: Concatenate and display file contents.
```bash
cat file.txt
```
- `less`: View file contents interactively.
```bash
less file.txt
```
- `head`: Display the first lines of a file.
```bash
head file.txt
```
- `tail`: Display the last lines of a file.
```bash
tail file.txt
```
- `nano`: Text editor for the command line.
```bash
nano file.txt
```
- `vi` or `vim`: Powerful text editor for the command line.
```bash
vi file.txt
vim file.txt
```

## System Information:
- `uname`: Print system information.
```bash
uname
```
- `whoami`: Display the current user.
```bash
whoami
```
- `hostname`: Show or set the system's hostname.
```bash
hostname
```
- `df`: Display disk space usage.
```bash
df
```
- `du`: Estimate file and directory space usage.
```bash
du -sh /path/to/directory
```
- `top`: Display system resource usage.
```bash
top
```

## Process Management:
- `ps`: Display running processes.
```bash
ps
```
- `kill`: Terminate processes.
```bash
kill process_id
```
- `killall`: Kill processes by name.
```bash
killall process_name
```
- `pgrep`: Find processes by name.
```bash
pgrep process_name
```
- `pstree`: Display process hierarchy.
```bash
pstree
```

## Networking:
- `ifconfig`: Configure or display network interface parameters.
```bash
ifconfig
```
- `ping`: Send ICMP echo requests to a host.
```bash
ping host
```
- `netstat`: Network statistics.
```bash
netstat
```
- `ssh`: Secure shell client for remote login.
```bash
ssh user@hostname
```
- `scp`: Securely copy files between hosts.
```bash
scp file.txt user@hostname:/path/to/destination
```
- `wget`: Download files from the web.
```bash
wget URL
```
- `curl`: Transfer data from or to a server.
```bash
curl URL
```

## Package Management (package managers may vary):
- `apt`: Advanced Package Tool (Debian/Ubuntu).
```bash
apt install package_name
apt remove package_name
apt update
apt upgrade
```
- `yum`: Yellowdog Updater Modified (RHEL/CentOS).
```bash
yum install package_name
yum remove package_name
yum update
yum upgrade
```
- `dnf`: Dandified YUM (Fedora).
```bash
dnf install package_name
dnf remove package_name
dnf update
dnf upgrade
```
- `pacman`: Package Manager (Arch Linux).
```bash
pacman -S package_name
pacman -R package_name
pacman -Syu
```
- `zypper`: Package Manager (openSUSE).
```bash
zypper install package_name
zypper remove package_name
zypper update
zypper upgrade
```
- `snap`: Package manager for snaps (universal Linux packages).
```bash
snap install package_name
snap remove package_name
snap refresh
```

