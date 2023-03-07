import pymongo

mongo_client = pymongo.MongoClient('mongodb://127.0.0.1:27017')

# print(mongo_client.server_info()) #判断是否连接成功

mongo_db = mongo_client['test']
mongo_collection = mongo_db['demo']

info = {"name":"justin","age":16}

find_condition = {
    'name' : 'justin'
}
find_result = mongo_collection.find_one(find_condition)

print(find_result["age"])