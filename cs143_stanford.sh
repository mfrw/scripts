#!/bin/sh
#This shell script exposes jlex to the user

echo -e '#!/bin/sh \njava -classpath /usr/class/cs143/cool/lib/jlex.jar JLex.Main $1' > jlex
chmod a+x jlex
sudo mv jlex /bin/jlex
