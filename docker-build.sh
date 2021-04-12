if [[ "$1" = "NO-CACHE" ]]
then
   docker build --no-cache --tag atlas-fec:latest .
else
   docker build --tag atlas-fec:latest .
fi
