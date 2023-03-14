import datetime
import json
import sys
import os
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
from tool import ToolUtil as TU
import numpy as np
# print(datetime.datetime.now())

# dict_ = {"lala":{"a":"b"},"ccc":["kkk","ppp"]}

# print(json.dumps(dict_))


# def get_list_of_dicts(name: str, surname: str) -> map[str,TU.position]:
#     return {}

# print(str(datetime.datetime.now()))

# lala = ["a","b"]

# for k in lala:
#     print(k)

print(np.linspace(start = 1, stop = 2, num = 6)[1:])