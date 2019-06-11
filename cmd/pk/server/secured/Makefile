.PHONY: createKeystore addHost exportAuthority createTruststore addClient clean

all:	clean createKeystore addHost createTruststore addClient

createKeystore:
	keytool -genkey -alias ca -ext san=dns:localhost,ip:127.0.0.1  \
	    -ext BC=ca:true -keyalg RSA -keysize 4096 -sigalg SHA512withRSA  \
	    -keypass supersafe -validity 3650 -dname "CN=ITArad,OU=itarad.ro,O=ITArad,L=Arad,ST=Arad,C=RO"  \
	    -keystore keystore.jks -storepass supersafe

addHost:
	# Generate a host certificate
	keytool -genkey -alias localhost -ext san=dns:localhost,ip:127.0.0.1  \
	    -keyalg RSA -keysize 4096 -sigalg SHA512withRSA -keypass supersafe \
	    -validity 3650 -dname "CN=localhost,OU=itarad.com,O=ITArad,L=Arad,ST=Arad,C=RO" \
	    -keystore "keystore.jks" -storepass supersafe
	# Generate a host certificate signing request
	keytool -certreq -alias localhost -ext san=dns:localhost,ip:127.0.0.1 -ext BC=ca:true \
	    -keyalg RSA -keysize 4096 -sigalg SHA512withRSA \
	    -validity 3650 -file "localhost.csr" \
	    -keystore keystore.jks -storepass supersafe
	# Generate signed certificate with the certificate authority
	keytool -gencert -alias ca -ext san=dns:localhost,ip:127.0.0.1 \
	    -validity 3650 -sigalg SHA512withRSA \
	    -infile "localhost.csr" -outfile "localhost.crt" -rfc \
	    -keystore keystore.jks -storepass supersafe
	# Import signed certificate into the keystore
	keytool -import -trustcacerts -alias localhost -ext san=dns:localhost,ip:127.0.0.1 \
	    -file "localhost.crt" \
	    -keystore keystore.jks -storepass supersafe

exportAuthority:
	# Export certificate authority
	keytool -export -alias ca -ext san=dns:localhost,ip:127.0.0.1 -file ca.crt -rfc \
	    -keystore keystore.jks -storepass supersafe


createTruststore: exportAuthority
	# Import certificate authority into a new truststore
	keytool -import -trustcacerts -noprompt -alias ca -ext san=dns:localhost,ip:127.0.0.1 -file ca.crt \
	    -keystore truststore.jks -storepass supersafe

addClient:
	# Generate client certificate
	keytool -genkey -alias cid_pk -ext san=dns:localhost,ip:127.0.0.1 \
	    -keyalg RSA -keysize 4096 -sigalg SHA512withRSA -keypass supersafe \
	    -validity 3650 -dname "CN=cid,OU=itarad.com,O=ITArad,L=Arad,ST=Arad,C=RO" \
	    -keystore truststore.jks -storepass supersafe
	# Generate a host certificate signing request
	keytool -certreq -alias cid_pk -ext san=dns:localhost,ip:127.0.0.1 -ext  BC=ca:true \
	    -keyalg RSA -keysize 4096 -sigalg SHA512withRSA \
	    -validity 3650 -file "cid.csr" \
	    -keystore truststore.jks -storepass supersafe
	# Generate signed certificate with the certificate authority
	keytool -gencert -alias ca -ext san=dns:localhost,ip:127.0.0.1 \
	    -validity 3650 -sigalg SHA512withRSA \
	    -infile "cid.csr" -outfile "cid.crt" -rfc \
	    -keystore keystore.jks -storepass supersafe
	# Import signed certificate into the truststore
	keytool -import -trustcacerts -alias cid -ext san=dns:localhost,ip:127.0.0.1 \
	    -file "cid.crt" \
	    -keystore truststore.jks -storepass supersafe
	# Export private certificate for importing into a browser
	keytool -importkeystore -srcalias cid_pk -ext san=dns:localhost,ip:127.0.0.1 \
	    -srckeystore truststore.jks -srcstorepass supersafe \
	    -destkeystore "cid.p12" -deststorepass supersafe \
	    -deststoretype PKCS12
	# Delete client private key as truststore should not contain any private keys
	keytool -delete -alias cid_pk -keystore truststore.jks -storepass supersafe

clean:
	# Remove generated artifacts
	find . \( -name "cid*" -o -name "localhost*" -o -name "keystore,jks" -o -name "truststore.jks" -o -name ca.crt \) -type f -exec rm -f {} \;
