from concurrent import futures
import grpc
import deliver_pb2
import deliver_pb2_grpc

class CalServicer(deliver_pb2_grpc.BarDataRevicerServicer):
  def BarDataRevicer(self, barinfo, context):   # Add函数的实现逻辑
    print(barinfo.Insid)
    return deliver_pb2.Response(response_me="lalala") 


def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
  deliver_pb2_grpc.add_BarDataRevicerServicer_to_server(CalServicer(),server)
  server.add_insecure_port("[::]:4352")
  server.start()
  print("grpc server start...")
  server.wait_for_termination()

if __name__ == '__main__':
  serve()