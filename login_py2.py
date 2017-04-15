#!/usr/bin/python
# Date: 17-July-2016
#author : Muhammad Falak R Wani (mfrw)
# No f***ing copyright here.. take it and do whatever you want with the code.

# random script to login to iiitd network
# A very hacky script, expect refinments over time
# 
# The most irritating way, but ofcourse, the most showoff way is to use it like an interacitive script, but if you are lazy,
# then its better you comment line the raw_input('Username:') and the raw_input('Password:') file and just punch in your 
# password and username in the commented fields

# I hope I rewrite this in a much saner way.. but till then this is what it is...
# btw i hope fidu has gotten the script

# If the target machine does not have requests upload the request.zip file from here, unzip it in the same folder as the script :| hopefully it should run


import requests
import getpass
s = requests.Session()
r = s.get('http://www.google.com')
#uname = Your_username
#passd = Your_passwd

if 'IIIT-D' in str(r.content):
    uname = raw_input('Username:')
    passd = getpass.getpass()
    login_data = dict(username=uname, magic=r.content[6354:6370],password=passd)
    # Instead of a re search, extract the magic cookie^^^^^^^^^(should be faster) [may break] 
    r = s.post(r.url, data=login_data)
else:
    print ('dude you are already connected')
