if [ ! -f "conf/app.ini" ]; then
    echo "conf/app.ini not found"
    if [ -f "conf/app.ini.sample" ]; then
        cp conf/app.ini.sample conf/app.ini
    else
        echo "conf/app.ini.sample not  founded"
    fi
fi
nohup ./app -g "daemon off;" > /dev/null &
