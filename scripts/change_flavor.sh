#!/bin/bash

function helper {
    echo '''
    Usage: change_flavor.sh [OPTIONS]

    -h or --help: show this help message and exit
    -d or --debug: debug mode
    -f or --flavor: flavor to change to (default: dev) (dev/uat/prod) 
    -r or --release: release to change to
    -t or --test: test mode
    '''
}

# find argument contains -d or --debug
if [[ $@ == *"-h"* || $@ == *"--help"* ]]
then
    helper
else 
    GOOD_TO_GO=false
    if [[ $@ == *"--debug "* || $@ == *"--release "* || $@ == *"--test "* || $@ == *"--flavor"* ]]
    then
        GOOD_TO_GO=true
    fi

    if [[ $@ == *"-f "* || $@ == *"-d "* || $@ == *"-r "* || $@ == *"-t "* ]]
    then
        GOOD_TO_GO=true
    fi

    if [ $GOOD_TO_GO == false ]
    then
        echo "Invalid arguments"
        helper
        exit 1
    fi

    # split arguments into an array with spaces as delimiter
    IFS=' ' read -r -a args <<< "$@"

    if [ "${#args[@]}" != "3" ]
    then
        echo "Not enough arguments"
        helper
        exit 1
    fi

    # for loop with index and value of each argument
    for i in "${!args[@]}"
    do
        if [ "${args[$i]}" == "-d" ] || [ "${args[$i]}" == "--debug" ]
        then
            MODE=debug
        fi
        if [ "${args[$i]}" == "-f" ] || [ "${args[$i]}" == "--flavor" ]
        then
            FLAVOR=${args[$i+1]}
        fi
        if [ "${args[$i]}" == "-r" ] || [ "${args[$i]}" == "--release" ]
        then
            MODE=release
        fi
        if [ "${args[$i]}" == "-t" ] || [ "${args[$i]}" == "--test" ]
        then
            MODE=test
        fi
    done

    if [ "$FLAVOR" == "dev" ]
    then
    	cp ./config/envs/dev.env ./config/envs/config.env
    else
        cp ./config/envs/prod.env ./config/envs/config.env
    fi

    sed -re  "s/(FLAVOR=)[^^=]*$/\1${FLAVOR}/g" "$PWD/config/envs/config.env" > "$PWD/config/envs/config.env " && mv "$PWD/config/envs/config.env " "$PWD/config/envs/config.env"
    sed -re  "s/(GIN_MODE=)[^^=]*$/\1${MODE}/g" "$PWD/config/envs/config.env" > "$PWD/config/envs/config.env " && mv "$PWD/config/envs/config.env " "$PWD/config/envs/config.env"
fi
