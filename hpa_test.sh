echo Inicializando wget infinito em http://$1
while true;
do
    wget -q -O- http://$1;
done;


