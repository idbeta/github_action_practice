#!/usr/bin/env python3
import os

def cur_dir():
    print(os.path.split(os.path.realpath(__file__))[0])

print(cur_dir())
