import pyserver
import threading

class strategy:
    def UpdataBar(self,bar_info):
        print("infogather bar called",bar_info)
    def UpdateTick(self,tick_info):
        print("infogather tick called",tick_info)
        
    def Start(self):
        bar_thread = threading.Thread(target=pyserver.serveBar,args={self})
        tick_thread = threading.Thread(target=pyserver.serveTick,args={self})
        bar_thread.start()
        tick_thread.start()


if __name__ == '__main__':
    datagather = strategy()
    datagather.Start()
    pyserver.NeverStop()