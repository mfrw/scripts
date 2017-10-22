# make your own (large) PDF from slideshare url
# image resolutions available: 320, 728 or 1024 pixels
# example: $0 http://de.slideshare.net/olvemaudal/deep-c 728 deepc.pdf

test $# -eq 3||exec echo usage: url $0 {320,728,1024} file.pdf
browser='curl -s -O';
#browser='ftp -4';
test ${#TMPDIR} -eq 0||cd $TMPDIR||exec echo set TMPDIR 
$browser -o 1.htm $1;
echo downloading images... >&2;
tr '\40' '\12' < 1.htm \
	|sed '/-'"$2"'\.jpg/!d;s/\.jpg.*/.jpg/;
s/.*https:/https:/;s/\".*//' \
	|while read a;do $browser $a;done #slow step;
test -f  *-1-$2.jpg ||exec echo download failed 
pdfimage -o $3 *-$2.jpg;
exec rm  *-$2.jpg 1.htm;

# pdfimage is sample program included with pdflib
# http://www.pdflib.com/
