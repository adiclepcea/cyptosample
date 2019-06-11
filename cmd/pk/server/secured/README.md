# PK

This is a Spring server intended to serve web requests over https and also to check the client certificate

## Build

To make it work, you must run:

```
make all
```

If your aliases are already present you can remove them:

```
keytool.exe -delete -alias localhost -keystore keystore.jks
```

and/or

```
keytool.exe -delete -alias localhost -keystore keystore.jks
```

then copy the keystore.jks and truststore.jks in the resources folder and

```
gradle bootJar
java -jar build/libs/yourJar.jsr
```
