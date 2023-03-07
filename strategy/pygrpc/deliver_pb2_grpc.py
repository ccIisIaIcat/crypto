# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import deliver_pb2 as deliver__pb2


class CustomDataReceiverStub(object):
    """定义小时bar消息接收对象
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CustomDataReceiver = channel.unary_unary(
                '/CustomDataReceiver/CustomDataReceiver',
                request_serializer=deliver__pb2.BarData.SerializeToString,
                response_deserializer=deliver__pb2.Response.FromString,
                )


class CustomDataReceiverServicer(object):
    """定义小时bar消息接收对象
    """

    def CustomDataReceiver(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CustomDataReceiverServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CustomDataReceiver': grpc.unary_unary_rpc_method_handler(
                    servicer.CustomDataReceiver,
                    request_deserializer=deliver__pb2.BarData.FromString,
                    response_serializer=deliver__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'CustomDataReceiver', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CustomDataReceiver(object):
    """定义小时bar消息接收对象
    """

    @staticmethod
    def CustomDataReceiver(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/CustomDataReceiver/CustomDataReceiver',
            deliver__pb2.BarData.SerializeToString,
            deliver__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class SubmitServerReceiverStub(object):
    """定义历史bar消息接收对象
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SubmitServerReceiver = channel.unary_unary(
                '/SubmitServerReceiver/SubmitServerReceiver',
                request_serializer=deliver__pb2.LocalSubmit.SerializeToString,
                response_deserializer=deliver__pb2.Response.FromString,
                )


class SubmitServerReceiverServicer(object):
    """定义历史bar消息接收对象
    """

    def SubmitServerReceiver(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SubmitServerReceiverServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SubmitServerReceiver': grpc.unary_unary_rpc_method_handler(
                    servicer.SubmitServerReceiver,
                    request_deserializer=deliver__pb2.LocalSubmit.FromString,
                    response_serializer=deliver__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'SubmitServerReceiver', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SubmitServerReceiver(object):
    """定义历史bar消息接收对象
    """

    @staticmethod
    def SubmitServerReceiver(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/SubmitServerReceiver/SubmitServerReceiver',
            deliver__pb2.LocalSubmit.SerializeToString,
            deliver__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class TickDataReceiverStub(object):
    """定义tick消息接收对象
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.TickDataReceiver = channel.unary_unary(
                '/TickDataReceiver/TickDataReceiver',
                request_serializer=deliver__pb2.TickData.SerializeToString,
                response_deserializer=deliver__pb2.Response.FromString,
                )


class TickDataReceiverServicer(object):
    """定义tick消息接收对象
    """

    def TickDataReceiver(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TickDataReceiverServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'TickDataReceiver': grpc.unary_unary_rpc_method_handler(
                    servicer.TickDataReceiver,
                    request_deserializer=deliver__pb2.TickData.FromString,
                    response_serializer=deliver__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'TickDataReceiver', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TickDataReceiver(object):
    """定义tick消息接收对象
    """

    @staticmethod
    def TickDataReceiver(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/TickDataReceiver/TickDataReceiver',
            deliver__pb2.TickData.SerializeToString,
            deliver__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class OrerReceiverStub(object):
    """
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.OrerRReceiver = channel.unary_unary(
                '/OrerReceiver/OrerRReceiver',
                request_serializer=deliver__pb2.Order.SerializeToString,
                response_deserializer=deliver__pb2.Response.FromString,
                )


class OrerReceiverServicer(object):
    """
    """

    def OrerRReceiver(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OrerReceiverServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'OrerRReceiver': grpc.unary_unary_rpc_method_handler(
                    servicer.OrerRReceiver,
                    request_deserializer=deliver__pb2.Order.FromString,
                    response_serializer=deliver__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'OrerReceiver', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OrerReceiver(object):
    """
    """

    @staticmethod
    def OrerRReceiver(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OrerReceiver/OrerRReceiver',
            deliver__pb2.Order.SerializeToString,
            deliver__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class JsonReceiverStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.JsonReceiver = channel.unary_unary(
                '/JsonReceiver/JsonReceiver',
                request_serializer=deliver__pb2.JsonInfo.SerializeToString,
                response_deserializer=deliver__pb2.Response.FromString,
                )


class JsonReceiverServicer(object):
    """Missing associated documentation comment in .proto file."""

    def JsonReceiver(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_JsonReceiverServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'JsonReceiver': grpc.unary_unary_rpc_method_handler(
                    servicer.JsonReceiver,
                    request_deserializer=deliver__pb2.JsonInfo.FromString,
                    response_serializer=deliver__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'JsonReceiver', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class JsonReceiver(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def JsonReceiver(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/JsonReceiver/JsonReceiver',
            deliver__pb2.JsonInfo.SerializeToString,
            deliver__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
