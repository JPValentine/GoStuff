#!/bin/bash
echo "wow"
cp -R /home/jp/snap/exercism/5/exercism/go/ /home/jp/Projects/GoStuff/exercism/ 
git -C /home/jp/Projects/GoStuff/exercism/go add *
git -C /home/jp/Projects/GoStuff/exercism/go commit -m "$*"
git -C /home/jp/Projects/GoStuff/exercism/go push origin main 

