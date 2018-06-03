from __future__ import print_function

import grpc

import tl_pb2
import tl_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:11011')
    stub = tl_pb2_grpc.MtprotoStub(channel)
    peer = tl_pb2.TypeInputPeer(InputPeerEmpty = tl_pb2.PredInputPeerEmpty())
    response = stub.MessagesGetDialogs(tl_pb2.ReqMessagesGetDialogs(OffsetPeer = peer))
    print("Greeter client received: " + str(response))


if __name__ == '__main__':
    run()
