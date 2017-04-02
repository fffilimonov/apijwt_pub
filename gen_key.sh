openssl genrsa -out settings/keys/private_key 2048;
openssl rsa -in settings/keys/private_key -pubout -out settings/keys/public_key
