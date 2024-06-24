#!/bin/bash

lower=$1
upper=$2

cp splash "$lower" -r
sed -e "s/Splash/$upper/g" -e "s/splash/$lower/g" $lower/* -i
