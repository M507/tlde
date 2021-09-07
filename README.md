
I wrote this simple tool because I need a cross-platform and fast TLD parser to enhance my reconnaissance accuracy.

### Install
```
go install github.com/M507/tlde@latest
```

### Test list
```
$ cat /tmp/test.txt
http://www1.web.google.com.ru/
http://www2.yahoo.com.ru/
https://bwrc.am/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
https://syunikwrc.net/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
https://womeninblackarmenia.weebly.com/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
https://sose-ngo.am/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
https://womenofarmenia.org/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
https://oxygen.org.am/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
https://cyberhub.am/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
https://mediapoint.am/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
http://mediaadvocate.am/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
http://armhels.com/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
http://hahr.am/,HUMR,Human Rights Issues,2021-06-12,cyberhub-am,
```

### Get subdomain + domain + tld
```
$ cat /tmp/test.txt | cut -d',' -f1 | tlde -r -t -s 
www1.web.google.com.ru
www2.yahoo.com.ru
bwrc.am
syunikwrc.net
womeninblackarmenia.weebly.com
sose-ngo.am
womenofarmenia.org
oxygen.org.am
cyberhub.am
mediapoint.am
mediaadvocate.am
armhels.com
hahr.am
```

### Get hostname + tld
```
$ cat /tmp/test.txt | cut -d',' -f1 | tlde -r -t
google.com.ru
yahoo.com.ru
bwrc.am
syunikwrc.net
weebly.com
sose-ngo.am
womenofarmenia.org
oxygen.org.am
cyberhub.am
mediapoint.am
mediaadvocate.am
armhels.com
hahr.am
```

### Get hostname
```
$ cat /tmp/test.txt | cut -d',' -f1 | tlde -r
google
yahoo
bwrc
syunikwrc
weebly
sose-ngo
womenofarmenia
oxygen
cyberhub
mediapoint
mediaadvocate
armhels
hahr
```

### Get Subdomains
```
$ cat /tmp/test.txt | cut -d',' -f1 | tlde -s
www1.web
www2
womeninblackarmenia
```

### Help
```
Usage of /tmp/go-build672141740/b001/exe/tlde:
  -r    Print hostname
  -root
        Print hostname
  -s    Print subdomain
  -sub
        Print subdomain
  -t    Print Top-Level Domain
  -tld
        Print Top-Level Domain
  -u string
        URL (default "http://www1.web.google.com.sh/")
```
