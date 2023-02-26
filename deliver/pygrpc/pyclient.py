import grpc
import deliver_pb2
import deliver_pb2_grpc

def run():
    # 本次不使用SSL，所以channel是不安全的
    channel = grpc.insecure_channel('localhost:4352')
    # 客户端实例
    stub = deliver_pb2_grpc.OrerReceiverStub(channel)
    # 调用服务端方法
    response = stub.OrerRReceiver(deliver_pb2.Order(insId = "BTC-USDT"))
    print("Greeter client received: " + response.response_me)

if __name__ == '__main__':
    run()