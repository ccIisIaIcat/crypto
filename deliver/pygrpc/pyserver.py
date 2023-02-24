from concurrent import futures
import grpc
import deliver_pb2
import deliver_pb2_grpc
import  threading
import time

class InfoGather:
    def UpdataBar(self,bar_info):
        print("infogather bar called",bar_info)
    def UpdateTick(self,tick_info):
        print("infogather tick called",tick_info)
        

class BarServicer(deliver_pb2_grpc.BarDataReceiverServicer):
    MyGather = ""
    def __init__(self,newgather) :
        super().__init__()
        self.MyGather = newgather
    def BarDataReceiver(self, barinfo, context):   # Bar函数接收实现
        self.MyGather.UpdataBar(barinfo)
        return deliver_pb2.Response(response_me="bar delivered") 

class TickServicer(deliver_pb2_grpc.TickDataReceiverServicer):
    MyGather = ""
    def __init__(self,newgather) :
        super().__init__()
        self.MyGather = newgather
    def TickDataReceiver(self, tickinfo, context):   # Tick函数接收实现
        self.MyGather.UpdataBar(tickinfo)
        return deliver_pb2.Response(response_me="tick delivered") 


def serveBar(gather):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    deliver_pb2_grpc.add_BarDataReceiverServicer_to_server(BarServicer(gather),server)
    server.add_insecure_port("[::]:3902")
    server.start()
    print("pybar grpc server start...")
    server.wait_for_termination()
  
def serveTick(gather):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    deliver_pb2_grpc.add_TickDataReceiverServicer_to_server(TickServicer(gather),server)
    server.add_insecure_port("[::]:3903")
    server.start()
    print("pytick grpc server start...")
    server.wait_for_termination()
    
def NeverStop():
    while True:
        time.sleep(1000)
        
def start_gather(gather):
    # bar_thread = threading.Thread(target=serveBar,args={gather})
    # tick_thread = threading.Thread(target=serveTick,args={gather})
    # bar_thread.start()
    # tick_thread.start()
    serveBar(gather)

# if __name__ == '__main__':
#     lala = InfoGather()
#     start_gather(lala)
#     NeverStop()
    