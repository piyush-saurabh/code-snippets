# code-snippets

This repository contains the code snippets which has been compiled from other git repositories or the codes which I have written during CTFs

**port-knocking.sh** : This is the script for port knocking. After sending the ping request to 'n' ports in a particular sequence, it will open up a particular port. e.g ssh
Usage :  **_./port-knocking.sh IP port1 port2 port3 && ssh username@ip_**

**ssrf_scan.py** : If we find SSRF vulnerability, then this script can be used to scan the internal network for the open web ports. This script can be modified for other scans related to internal network.
Usage :  **python ssrf_scan.py**
