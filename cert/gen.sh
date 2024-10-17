rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=DK/ST=Denmark/L=Copenhagen/O=ITU/OU=Education/CN=Max/emailAddress=mbko@itu.dk"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=DK/ST=Denmark/L=Copenhagen/O=ITU/OU=Education/CN=Server/emailAddress=mbko@itu.dk"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client0-key.pem -out client0-req.pem -subj "/C=DK/ST=Denmark/L=Copenhagen/O=PC Client/OU=Computer/CN=Client0/emailAddress=mbko@itu.dk"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in client0-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client0-cert.pem -extfile client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in client0-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client1-key.pem -out client1-req.pem -subj "/C=DK/ST=Denmark/L=Copenhagen/O=PC Client/OU=Computer/CN=Client1/emailAddress=mbko@itu.dk"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in client1-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client1-cert.pem -extfile client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in client1-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client2-key.pem -out client2-req.pem -subj "/C=DK/ST=Denmark/L=Copenhagen/O=PC Client/OU=Computer/CN=Client2/emailAddress=mbko@itu.dk"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in client2-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client2-cert.pem -extfile client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in client2-cert.pem -noout -text