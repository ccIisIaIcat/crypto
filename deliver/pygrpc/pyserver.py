from concurrent import futures
import grpc
import deliver_pb2
import deliver_pb2_grpc
import time

        
class BarCustomServicer(deliver_pb2_grpc.CustomDataReceiverServicer):
    MyGather = ""
    def __init__(self,newgather) :
        super().__init__()
        self.MyGather = newgather
    def CustomDataReceiver(self, barinfo, context):
        self.MyGather.UpdateBarCustom(barinfo)
        return deliver_pb2.Response(response_me="Custom bar delivered") 

def serveBarCustom(gather):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    deliver_pb2_grpc.add_CustomDataReceiverServicer_to_server(BarCustomServicer(gather),server)
    server.add_insecure_port("[::]:"+gather.portbarcustom)
    server.start()
    print("pybarhour grpc server start...")
    server.wait_for_termination()
        

class TickServicer(deliver_pb2_grpc.TickDataReceiverServicer):
    MyGather = ""
    def __init__(self,newgather) :
        super().__init__()
        self.MyGather = newgather
    def TickDataReceiver(self, tickinfo, context):   # Tick函数接收实现
        self.MyGather.UpdateTick(tickinfo)
        return deliver_pb2.Response(response_me="tick delivered") 

def serveTick(gather):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    deliver_pb2_grpc.add_TickDataReceiverServicer_to_server(TickServicer(gather),server)
    server.add_insecure_port("[::]:"+gather.porttick)
    server.start()
    print("pytick grpc server start...")
    server.wait_for_termination()

class AccountServicer(deliver_pb2_grpc.JsonReceiverServicer):
    MyGather = ""
    def __init__(self,newgather) :
        super().__init__()
        self.MyGather = newgather
    def JsonReceiver(self, accountinfo, context):   # account函数接收实现
        self.MyGather.UpdateAccount(accountinfo.jsoninfo)
        return deliver_pb2.Response(response_me="account info delivered") 
    
def serveAccount(gather):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    deliver_pb2_grpc.add_JsonReceiverServicer_to_server(AccountServicer(gather),server)
    server.add_insecure_port("[::]:"+gather.portaccount)
    server.start()
    print("pyaccount grpc server start...")
    server.wait_for_termination()
    
  



    
def NeverStop():
    while True:
        time.sleep(1000)
        


# if __name__ == '__main__':
#     lala = InfoGather()
#     start_gather(lala)
#     NeverStop()
    