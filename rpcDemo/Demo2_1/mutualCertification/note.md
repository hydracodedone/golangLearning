genrsa -out ca.key 2048
req -new -x509 -days 365 -key ca.key -out ca.pem

genrsa -out server.key 2048
req -new -key server.key -out server.csr
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem

ecparam -genkey -name secp384r1 -out client.key
req -new -key client.key -out client.csr
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem


ps:common name:localhost
ps:if error happened,renter the openssl and carry on
