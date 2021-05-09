genrsa -des3 -out server.key 2048
req -new -key server.key -out server.csr
rsa -in server.key -out server_no_password.key
x509 -req -days 365 -in server.csr -signkey server_no_password.key -out server.cert
x509ignoreCN=0