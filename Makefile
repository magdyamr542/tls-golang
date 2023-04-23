ca:
	# Generate a private key for the CA:
	openssl genpkey -algorithm RSA -out ca.key
	# Create a self-signed certificate for the CA:
	openssl req -new -x509 -key ca.key -out ca.crt

serverkey:
	# Generate a private key for the server:
	openssl genpkey -algorithm RSA -out server.key
	# Create a certificate signing request (CSR) for the server:
	openssl req -new -key server.key -out server.csr -config san.cnf
	# Create a server certificate by signing the CSR with the CA:
	openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -extensions reqexts -extfile san.cnf

clientkey:
	# Generate a private key for the client:
	openssl genpkey -algorithm RSA -out client.key
	# Create a certificate signing request (CSR) for the client:
	openssl req -new -key client.key -out client.csr
	# Create a client certificate by signing the CSR with the CA:
	openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt

server:
	go run ./cmd/server/server.go

client:
	go run ./cmd/client/client.go
